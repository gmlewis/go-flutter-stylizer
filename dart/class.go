/*
Copyright © 2021 Glenn M. Lewis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dart

import (
	"fmt"
	"strings"
)

// Class represents a Dart class.
type Class struct {
	e         *Editor
	classBody string
	lines     []*Line

	className                 string
	openCurlyOffset           int
	closeCurlyOffset          int
	groupAndSortGetterMethods bool

	theConstructor         *Entity
	namedConstructors      []*Entity
	staticVariables        []*Entity
	instanceVariables      []*Entity
	overrideVariables      []*Entity
	staticPrivateVariables []*Entity
	privateVariables       []*Entity
	overrideMethods        []*Entity
	otherMethods           []*Entity
	buildMethod            *Entity
	getterMethods          []*Entity
}

// NewClass returns a new Dart Class.
//
// className is the name of the class.
// lineOffset is the offset of the starting line of the class (0-based).
// openCurlyOffset is the position of the "{" for that class.
// closeCurlyOffset is the position of the "}" for that class.
// groupAndSortGetterMethods determines how getter methods are processed.
func NewClass(editor *Editor, className string,
	lineOffset, openCurlyOffset, closeCurlyOffset int,
	groupAndSortGetterMethods bool) *Class {
	lessThanOffset := strings.Index(className, "<")
	if lessThanOffset >= 0 { // Strip off <T>.
		className = className[0:lessThanOffset]
	}
	classBody := editor.fullBuf[openCurlyOffset:closeCurlyOffset]
	p := strings.Split(classBody, "\n")
	numClassLines := len(p)
	classLines := editor.lines[lineOffset : lineOffset+numClassLines]
	// log.Printf("p[%v]=%q(%v), closeCurlyOffset=%v, len(classBody)=%v, last classLine=%#v", len(p)-1, p[len(p)-1], len(p[len(p)-1]), closeCurlyOffset, len(classBody), classLines[len(classLines)-1])

	// Replace last line for truncated classBody:
	ll := classLines[len(classLines)-1]
	classLines[len(classLines)-1] = &Line{
		line:           p[len(p)-1],
		stripped:       strings.TrimSpace(p[len(p)-1]),
		strippedOffset: ll.strippedOffset,
		startOffset:    ll.startOffset,
		endOffset:      ll.endOffset,
		entityType:     BlankLine,
	}

	return &Class{
		e:                editor,
		classBody:        classBody,
		lines:            classLines,
		className:        className,
		openCurlyOffset:  openCurlyOffset,
		closeCurlyOffset: closeCurlyOffset,

		groupAndSortGetterMethods: groupAndSortGetterMethods,
	}
}

func (c *Class) FindFeatures() error {
	for i, line := range c.lines {
		// Change a blank line following a comment to a SingleLineComment in
		// order to keep it with the following entity.
		if line.entityType == Unknown && len(line.stripped) == 0 {
			line.entityType = BlankLine
		}
		if i > 1 && c.lines[i-1].entityType == BlankLine && isComment(c.lines[i-2]) {
			c.lines[i-1].entityType = SingleLineComment
		}
	}

	// c.identifyMultiLineComments()  // This step is not needed, as the editor marked these already.
	if err := c.identifyMainConstructor(); err != nil {
		return err
	}
	if err := c.identifyNamedConstructors(); err != nil {
		return err
	}
	if err := c.identifyOverrideMethodsAndVars(); err != nil {
		return err
	}
	if err := c.identifyOthers(); err != nil {
		return err
	}

	if c.e.Verbose {
		for i := 0; i < len(c.lines); i++ {
			line := c.lines[i]
			c.e.logf("line #%v type=%v: %v", i, line.entityType, line.line)
		}
	}

	return nil
}

// func (c *Class) genStripped(startLine int) string {
// 	var strippedLines []string
// 	for i := startLine; i < len(c.lines); i++ {
// 		strippedLines = append(strippedLines, c.lines[i].stripped)
// 	}
// 	return strings.Join(strippedLines, "\n")
// }

func (c *Class) identifyMainConstructor() error {
	className := c.className + "("
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown {
			continue
		}
		offset := strings.Index(line.stripped, className)
		if offset >= 0 {
			if offset > 0 {
				char := line.stripped[offset-1 : offset]
				if char != " " && char != "\t" {
					continue
				}
			}
			if c.lines[i].entityType > MainConstructor {
				if err := c.repairIncorrectlyLabeledLine(i); err != nil {
					return err
				}
			}
			c.lines[i].entityType = MainConstructor
			var err error
			c.theConstructor, err = c.markMethod(i, className, MainConstructor)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func (c *Class) identifyNamedConstructors() error {
	className := c.className + "."
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown {
			continue
		}
		offset := strings.Index(line.stripped, className)
		if offset >= 0 {
			if offset > 0 {
				char := line.stripped[offset-1 : offset]
				if line.stripped[0] == '?' || line.stripped[0] == ':' || (char != " " && char != "\t") {
					continue
				}
			}
			openParenOffset := offset + strings.Index(line.stripped[offset:], "(")
			namedConstructor := line.stripped[offset : openParenOffset+1] // Include open parenthesis.
			if c.lines[i].entityType >= MainConstructor && c.lines[i].entityType != NamedConstructor {
				if err := c.repairIncorrectlyLabeledLine(i); err != nil {
					return err
				}
			}
			c.lines[i].entityType = NamedConstructor
			entity, err := c.markMethod(i, namedConstructor, NamedConstructor)
			if err != nil {
				return err
			}
			c.namedConstructors = append(c.namedConstructors, entity)
		}
	}

	return nil
}

func (c *Class) identifyOverrideMethodsAndVars() error {
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown {
			continue
		}

		if strings.HasPrefix(line.stripped, "@override") && i < len(c.lines)-1 {
			offset := strings.Index(c.lines[i+1].stripped, "(")
			if offset >= 0 {
				// Include open paren in name.
				ss := c.lines[i+1].stripped[0 : offset+1]
				// Search for beginning of method name.
				nameOffset := strings.LastIndex(ss, " ") + 1
				name := ss[nameOffset:]
				entityType := OverrideMethod
				if name == "build(" {
					entityType = BuildMethod
				}
				if c.lines[i].entityType >= MainConstructor && c.lines[i].entityType != entityType {
					if err := c.repairIncorrectlyLabeledLine(i); err != nil {
						return err
					}
				}
				c.lines[i].entityType = entityType
				entity, err := c.markMethod(i+1, name, entityType)
				if err != nil {
					return err
				}
				if name == "build(" {
					c.buildMethod = entity
				} else {
					c.overrideMethods = append(c.overrideMethods, entity)
				}
			} else {
				entity := &Entity{
					entityType: OverrideMethod,
				}
				lineNum := i + 1
				// No open paren - could be a getter. See if it has a body.
				if strings.Index(c.lines[i+1].stripped, "{") >= 0 {
					lineOffset := strings.Index(c.classBody, c.lines[i+1].line)
					inLineOffset := strings.Index(c.lines[i+1].line, "{")
					relOpenCurlyOffset := lineOffset + inLineOffset

					if c.classBody[relOpenCurlyOffset] != '{' {
						return fmt.Errorf("expected open curly bracket at relative offset %v but got %q", relOpenCurlyOffset, c.classBody[relOpenCurlyOffset:])
					}

					absOpenCurlyOffset := c.openCurlyOffset + relOpenCurlyOffset
					absCloseCurlyOffset, err := c.e.findMatchingBracket(absOpenCurlyOffset)
					if err != nil {
						return err
					}

					relCloseCurlyOffset := absCloseCurlyOffset - c.openCurlyOffset

					if c.classBody[relOpenCurlyOffset] != '}' {
						return fmt.Errorf("expected close curly bracket at relative offset %v but got %q", relCloseCurlyOffset, c.classBody[relCloseCurlyOffset:])
					}

					nextOffset := absCloseCurlyOffset - c.openCurlyOffset
					bodyBuf := c.classBody[lineOffset : nextOffset+1]
					numLines := len(strings.Split(bodyBuf, "\n"))
					for j := 0; j < numLines; j++ {
						if c.lines[lineNum+j].entityType >= MainConstructor && c.lines[lineNum+j].entityType != entity.entityType {
							if err := c.repairIncorrectlyLabeledLine(lineNum + j); err != nil {
								return err
							}
						}
						c.lines[lineNum+j].entityType = entity.entityType
						entity.lines = append(entity.lines, c.lines[lineNum+j])
					}
				} else {
					// Does not have a body - if it has no fat arrow, it is a variable.
					if strings.Index(c.lines[i+1].stripped, "=>") < 0 {
						entity.entityType = OverrideVariable
					}
					// Find next ";", marking entityType forward.
					for j := i + 1; j < len(c.lines); j++ {
						if c.lines[j].entityType >= MainConstructor && c.lines[j].entityType != entity.entityType {
							if err := c.repairIncorrectlyLabeledLine(j); err != nil {
								return err
							}
						}
						c.lines[j].entityType = entity.entityType
						entity.lines = append(entity.lines, c.lines[j])
						semicolonOffset := strings.Index(c.lines[j].stripped, ";")
						if semicolonOffset >= 0 {
							break
						}
					}
				}

				// Preserve the comment lines leading up to the method.
				for lineNum--; lineNum > 0; lineNum-- {
					if isComment(c.lines[lineNum]) || strings.HasPrefix(c.lines[lineNum].stripped, "@") {
						c.lines[lineNum].entityType = entity.entityType
						entity.lines = append([]*Line{c.lines[lineNum]}, entity.lines...)
						continue
					}
					break
				}

				if entity.entityType == OverrideVariable {
					c.overrideVariables = append(c.overrideVariables, entity)
				} else {
					c.overrideMethods = append(c.overrideMethods, entity)
				}
			}
		}
	}

	return nil
}

func (c *Class) identifyOthers() error {
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown {
			continue
		}

		entity, err := c.scanMethod(i)
		if err != nil {
			return err
		}

		if entity.entityType == Unknown {
			continue
		}

		// Preserve the comment lines leading up to the entity.
		for lineNum := i - 1; lineNum > 0; lineNum-- {
			if isComment(c.lines[lineNum]) {
				c.lines[lineNum].entityType = entity.entityType
				entity.lines = append([]*Line{c.lines[lineNum]}, entity.lines...)
				continue
			}
			break
		}

		switch entity.entityType {
		case OtherMethod:
			c.otherMethods = append(c.otherMethods, entity)
		case GetterMethod:
			c.getterMethods = append(c.getterMethods, entity)
		case StaticVariable:
			c.staticVariables = append(c.staticVariables, entity)
		case StaticPrivateVariable:
			c.staticPrivateVariables = append(c.staticPrivateVariables, entity)
		case InstanceVariable:
			c.instanceVariables = append(c.instanceVariables, entity)
		case OverrideVariable:
			c.overrideVariables = append(c.overrideVariables, entity)
		case PrivateInstanceVariable:
			c.privateVariables = append(c.privateVariables, entity)
		default:
			return fmt.Errorf("unexpected EntityType=%v", entity.entityType)
		}
	}

	return nil
}

func (c *Class) scanMethod(lineNum int) (*Entity, error) {
	entity := &Entity{}

	sequence, lineCount, leadingText, err := c.findSequence(lineNum)
	if err != nil {
		return nil, err
	}
	c.e.logf("scanMethod(line=#%v), sequence=%v, lineCount=%v, leadingText=%q", lineNum+1, sequence, lineCount, leadingText)

	nameParts := strings.Split(leadingText, " ")
	var staticKeyword bool
	var privateVar bool
	if len(nameParts) > 0 {
		entity.name = nameParts[len(nameParts)-1]
		if strings.HasPrefix(entity.name, "_") {
			privateVar = true
		}
		if nameParts[0] == "static" {
			staticKeyword = true
		}
	}

	entity.entityType = InstanceVariable
	switch true {
	case privateVar && staticKeyword:
		entity.entityType = StaticPrivateVariable
		break
	case staticKeyword:
		entity.entityType = StaticVariable
		break
	case privateVar:
		entity.entityType = PrivateInstanceVariable
		break
	}

	switch sequence {
	case "(){}":
		entity.entityType = OtherMethod
		break

	case "();": // instance variable or abstract method.
		if !strings.HasSuffix(leadingText, " Function") {
			entity.entityType = OtherMethod
		}
		break

	case "=(){}":
		entity.entityType = OtherMethod
		break

	default:
		if strings.Index(sequence, "=>") >= 0 {
			entity.entityType = OtherMethod
		}
		break
	}

	// Force getters to be methods.
	if strings.Index(leadingText, " get ") >= 0 {
		if c.groupAndSortGetterMethods {
			entity.entityType = GetterMethod
		} else {
			entity.entityType = OtherMethod
		}
	}

	for i := 0; i <= lineCount; i++ {
		if lineNum+i >= len(c.lines) {
			break
		}

		if c.lines[lineNum+i].entityType >= MainConstructor && c.lines[lineNum+i].entityType != entity.entityType {
			c.e.logf("scanMethod: Changing line #%v from type %v to type %v", lineNum+i+1, c.lines[lineNum+i].entityType, entity.entityType)
			if err := c.repairIncorrectlyLabeledLine(lineNum + i); err != nil {
				return nil, err
			}
		}
		c.lines[lineNum+i].entityType = entity.entityType
		entity.lines = append(entity.lines, c.lines[lineNum+i])
	}

	return entity, nil
}

func (c *Class) repairIncorrectlyLabeledLine(lineNum int) error {
	incorrectLabel := c.lines[lineNum].entityType
	switch incorrectLabel {
	case NamedConstructor:
		for i := 0; i < len(c.namedConstructors); i++ {
			el := c.namedConstructors[i]
			for j := 0; j < len(el.lines); j++ {
				line := el.lines[j]
				if line != c.lines[lineNum] {
					continue
				}
				c.namedConstructors[i].lines = append(c.namedConstructors[i].lines[:j], c.namedConstructors[i].lines[j+1:]...)
				if len(c.namedConstructors[i].lines) == 0 {
					c.namedConstructors = append(c.namedConstructors[:i], c.namedConstructors[i+1:]...)
				}
				return nil
			}
		}
	default:
		return fmt.Errorf("repairIncorrectlyLabeledLine: line #%v, unhandled case %v. Please report on GitHub Issue Tracker with example test case.", lineNum+1, incorrectLabel)
	}

	return nil
}

func (c *Class) findSequence(lineNum int) (string, int, string, error) {
	var result []string

	startLine := c.lines[lineNum]
	cursor := &Cursor{
		e:         c.e,
		absOffset: startLine.startOffset + startLine.strippedOffset,
		lineIndex: startLine.originalIndex,

		relStrippedOffset: 0,

		reader: strings.NewReader(c.e.lines[startLine.originalIndex].stripped),
	}
	c.e.logf("\n\nfindSequence(line=%v): %#v, cursor=%v", lineNum+1, startLine, cursor)
	features, err := cursor.advanceUntil(";", "}")
	if err != nil {
		return "", 0, "", fmt.Errorf("advanceUntil(';', '}'): %v", err)
	}
	c.e.logf("cursor=%v", cursor)

	buildLeadingText := true
	var buildStr []string
	for _, f := range features {
		if strings.ContainsAny(f, "(){}=>;") {
			result = append(result, f)
			if f != ">" {
				buildLeadingText = false
			}
		}
		if buildLeadingText {
			buildStr = append(buildStr, f)
		}
	}
	leadingText := strings.TrimSpace(strings.Join(buildStr, ""))
	lineCount := cursor.lineIndex - startLine.originalIndex

	// var openParenCount int
	// var openBraceCount int
	// var openCurlyCount int
	// for i := 0; i < len(buf); i++ {
	// 	if openParenCount > 0 {
	// 		for ; i < len(buf); i++ {
	// 			switch buf[i] {
	// 			case '(':
	// 				openParenCount++
	// 				break
	// 			case ')':
	// 				openParenCount--
	// 				break
	// 			case '\n':
	// 				lineCount++
	// 				break
	// 			}
	// 			if openParenCount == 0 {
	// 				result = append(result, string(buf[i]))
	// 				break
	// 			}
	// 		}
	// 	} else if openBraceCount > 0 {
	// 		for ; i < len(buf); i++ {
	// 			switch buf[i] {
	// 			case '[':
	// 				openBraceCount++
	// 				break
	// 			case ']':
	// 				openBraceCount--
	// 				break
	// 			case '\n':
	// 				lineCount++
	// 				break
	// 			}
	// 			if openBraceCount == 0 {
	// 				result = append(result, string(buf[i]))
	// 				return strings.Join(result, ""), lineCount, leadingText
	// 			}
	// 		}
	// 	} else if openCurlyCount > 0 {
	// 		for ; i < len(buf); i++ {
	// 			switch buf[i] {
	// 			case '{':
	// 				openCurlyCount++
	// 				break
	// 			case '}':
	// 				openCurlyCount--
	// 				break
	// 			case '\n':
	// 				lineCount++
	// 				break
	// 			}
	// 			if openCurlyCount == 0 {
	// 				result = append(result, string(buf[i]))
	// 				return strings.Join(result, ""), lineCount, leadingText
	// 			}
	// 		}
	// 	} else {
	// 		switch buf[i] {
	// 		case '(':
	// 			openParenCount++
	// 			result = append(result, string(buf[i]))
	// 			if leadingText == "" {
	// 				leadingText = strings.TrimSpace(buf[0:i])
	// 			}
	// 			break
	// 		case '[':
	// 			openBraceCount++
	// 			result = append(result, string(buf[i]))
	// 			if leadingText == "" {
	// 				leadingText = strings.TrimSpace(buf[0:i])
	// 			}
	// 			break
	// 		case '{':
	// 			openCurlyCount++
	// 			result = append(result, string(buf[i]))
	// 			if leadingText == "" {
	// 				leadingText = strings.TrimSpace(buf[0:i])
	// 			}
	// 			break
	// 		case ';':
	// 			result = append(result, string(buf[i]))
	// 			if leadingText == "" {
	// 				leadingText = strings.TrimSpace(buf[0:i])
	// 			}
	// 			return strings.Join(result, ""), lineCount, leadingText
	// 		case '=':
	// 			if i < len(buf)-1 && buf[i+1] == '>' {
	// 				result = append(result, "=>")
	// 			} else {
	// 				result = append(result, string(buf[i]))
	// 			}
	// 			if leadingText == "" {
	// 				leadingText = strings.TrimSpace(buf[0:i])
	// 			}
	// 			break
	// 		case '\n':
	// 			lineCount++
	// 			break
	// 		}
	// 	}
	// }

	return strings.Join(result, ""), lineCount, leadingText, nil
}

func (c *Class) markMethod(lineNum int, methodName string, entityType EntityType) (*Entity, error) {
	if !strings.HasSuffix(methodName, "(") {
		return nil, fmt.Errorf("markMethod: %q must end with the open parenthesis '('", methodName)
	}

	entity := &Entity{
		name:       methodName,
		entityType: entityType,
	}

	// Identify all lines within the main (or factory) constructor.
	lineOffset := strings.Index(c.classBody, c.lines[lineNum].line)
	inLineOffset := strings.Index(c.lines[lineNum].line, methodName)
	relOpenParenOffset := lineOffset + inLineOffset + len(methodName) - 1
	if c.classBody[relOpenParenOffset] != '(' {
		return nil, fmt.Errorf("expected open parenthesis at relative offset %v but got %v", relOpenParenOffset, c.classBody[relOpenParenOffset:])
	}

	absOpenParenOffset := c.openCurlyOffset + relOpenParenOffset
	absCloseParenOffset, err := c.e.findMatchingBracket(absOpenParenOffset)
	if err != nil {
		return nil, err
	}

	relCloseParenOffset := absCloseParenOffset - c.openCurlyOffset
	if c.classBody[relCloseParenOffset] != ')' {
		return nil, fmt.Errorf("expected close parenthesis at relative offset %v but got %v", relCloseParenOffset, c.classBody[relCloseParenOffset:])
	}

	curlyDeltaOffset := strings.Index(c.classBody[relCloseParenOffset:], "{")
	semicolonOffset := strings.Index(c.classBody[relCloseParenOffset:], ";")
	var nextOffset int
	if curlyDeltaOffset < 0 || (curlyDeltaOffset >= 0 && semicolonOffset >= 0 && semicolonOffset < curlyDeltaOffset) { // no body.
		nextOffset = relCloseParenOffset + semicolonOffset
	} else {
		absOpenCurlyOffset := absCloseParenOffset + curlyDeltaOffset
		absCloseCurlyOffset, err := c.e.findMatchingBracket(absOpenCurlyOffset)
		if err != nil {
			return nil, err
		}

		nextOffset = absCloseCurlyOffset - c.openCurlyOffset
	}
	constructorBuf := c.classBody[lineOffset : nextOffset+1]
	numLines := len(strings.Split(constructorBuf, "\n"))
	for i := 0; i < numLines; i++ {
		if c.lines[lineNum+i].entityType >= MainConstructor && c.lines[lineNum+i].entityType != entityType {
			if err := c.repairIncorrectlyLabeledLine(lineNum + i); err != nil {
				return nil, err
			}
		}
		c.lines[lineNum+i].entityType = entityType
		entity.lines = append(entity.lines, c.lines[lineNum+i])
	}

	// Preserve the comment lines leading up to the method.
	for lineNum--; lineNum > 0; lineNum-- {
		if isComment(c.lines[lineNum]) || strings.HasPrefix(c.lines[lineNum].stripped, "@") {
			c.lines[lineNum].entityType = entityType
			entity.lines = append([]*Line{c.lines[lineNum]}, entity.lines...)
			continue
		}
		break
	}

	return entity, nil
}
