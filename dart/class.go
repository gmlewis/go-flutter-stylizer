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
// editor is the editor used to parse the class.
// className is the name of the class.
// openCurlyOffset is the position of the "{" for that class.
// closeCurlyOffset is the position of the "}" for that class.
// groupAndSortGetterMethods determines how getter methods are processed.
func NewClass(editor *Editor, className string,
	openCurlyOffset, closeCurlyOffset int,
	groupAndSortGetterMethods bool) *Class {
	lessThanOffset := strings.Index(className, "<")
	if lessThanOffset >= 0 { // Strip off <T>.
		className = className[0:lessThanOffset]
	}

	classBody := editor.fullBuf[openCurlyOffset:closeCurlyOffset]
	p := strings.Split(classBody, "\n")
	numClassLines := len(p)

	lineIndex, _ := editor.findLineIndexAtOffset(openCurlyOffset)
	classLines := editor.lines[lineIndex : lineIndex+numClassLines]

	p0Stripped := strings.TrimSpace(p[0])
	classLines[0] = &Line{
		line:              p[0], // Keep only the open curly brace and what follows.
		stripped:          p0Stripped,
		strippedOffset:    classLines[0].strippedOffset + len(p0Stripped) - len(classLines[0].stripped),
		originalIndex:     classLines[0].originalIndex,
		startOffset:       openCurlyOffset,
		endOffset:         classLines[0].endOffset,
		entityType:        classLines[0].entityType,
		isCommentOrString: classLines[0].isCommentOrString,
	}
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

func (c *Client) GetClasses(editor *Editor, groupAndSortGetterMethods bool) ([]*Class, error) {
	var classes []*Class

	for _, line := range editor.lines {
		if (strings.HasPrefix(line.line, "void main(") || strings.HasPrefix(line.line, "main(")) && strings.HasSuffix(line.line, ") {") {
			// Find the end of main, marking sections as comments or strings to avoid them below.
			openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, line.startOffset)
			editor.findMatchingBracket(openCurlyOffset)
			continue
		}

		mm := matchClassRE.FindStringSubmatch(line.line)
		if len(mm) != 2 {
			continue
		}

		className := mm[1]
		classOffset := line.startOffset
		openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, classOffset)

		if editor.fullBuf[openCurlyOffset] == ';' {
			continue
		}

		if openCurlyOffset <= classOffset {
			return nil, fmt.Errorf(`expected "{" after "class" at offset %v`, classOffset)
		}

		if line.entityType != Unknown || line.isCommentOrString {
			editor.logf("\n\nSkipping new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v due to entityType/comment/string", className, classOffset, openCurlyOffset, line)
			continue
		}

		// This is a hack, but helps with files like `localizations_utils.dart` in
		// the Flutter distribution.
		if strings.ContainsAny(line.stripped, "$'") {
			editor.logf("\n\nSkipping new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v due to appearance of $ or '", className, classOffset, openCurlyOffset, line)
			continue
		}

		editor.logf("\n\nFound new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v", className, classOffset, openCurlyOffset, line)
		closeCurlyOffset, err := editor.findMatchingBracket(openCurlyOffset)
		if err != nil {
			return nil, err
		}

		if closeCurlyOffset <= openCurlyOffset {
			return nil, fmt.Errorf(`expected "}" after "{" at offset %v`, openCurlyOffset)
		}
		editor.logf("\n\nFound end of class %q at closeCurlyOffset=%v", className, closeCurlyOffset)

		dartClass := NewClass(editor, className, openCurlyOffset, closeCurlyOffset, groupAndSortGetterMethods)
		if err := dartClass.FindFeatures(); err != nil {
			return nil, err
		}

		classes = append(classes, dartClass)
	}

	return classes, nil
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
	if err := c.identifyDeprecatedAsComments(); err != nil {
		return err
	}
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
			c.e.logf("line #%v type=%v: %v", i+1, line.entityType, line.line)
		}
	}

	return nil
}

func (c *Class) identifyDeprecatedAsComments() error {
	const depStr = "@deprecated("

	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown {
			continue
		}

		lower := strings.ToLower(line.stripped)
		if !strings.HasPrefix(lower, depStr[0:len(depStr)-1]) {
			continue
		}

		c.lines[i].entityType = SingleLineComment
		if strings.HasPrefix(lower, depStr) {
			if _, err := c.markMethod(i, line.stripped[0:len(depStr)], SingleLineComment); err != nil {
				return err
			}
		}
	}

	return nil
}

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

			c.e.logf("identifyMainConstructor: calling markMethod(line #%v, className=%q, MainConstructor)", i+1, className)
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
			if namedConstructor == "" {
				continue
			}

			if c.lines[i].entityType >= MainConstructor && c.lines[i].entityType != NamedConstructor {
				if err := c.repairIncorrectlyLabeledLine(i); err != nil {
					return err
				}
			}
			c.lines[i].entityType = NamedConstructor

			c.e.logf("identifyNamedConstructor: calling markMethod(line #%v, namedConstructor=%q, NamedConstructor)", i+1, namedConstructor)
			entity, err := c.markMethod(i, namedConstructor, NamedConstructor)
			if err != nil {
				return err
			}
			c.namedConstructors = append(c.namedConstructors, entity)
		}
	}

	return nil
}

func (c *Class) findNext(lineNum int, searchFor ...string) ([]string, *Cursor, error) {
	startLine := c.lines[lineNum]
	cursor := &Cursor{
		e:         c.e,
		absOffset: startLine.startOffset + startLine.strippedOffset,
		lineIndex: startLine.originalIndex,

		relStrippedOffset: 0,

		reader: strings.NewReader(c.e.lines[startLine.originalIndex].stripped),
	}
	c.e.logf("findNext(%v, %#v): %#v, BEFORE: cursor=%v", lineNum+1, searchFor, startLine, cursor)
	features, err := cursor.advanceUntil(searchFor...)
	if err != nil || len(features) == 0 {
		return nil, nil, fmt.Errorf(`advanceUntil(%#v): %v`, searchFor, err)
	}
	c.e.logf("findNext(%v, %#v): AFTER: cursor=%v, features=%v", lineNum+1, searchFor, cursor, strings.Join(features, ""))

	return features, cursor, nil
}

func (c *Class) identifyOverrideMethodsAndVars() error {
	for i := 1; i < len(c.lines); i++ {
		if c.lines[i].entityType != Unknown {
			continue
		}

		if !strings.HasPrefix(c.lines[i].stripped, "@override") || i >= len(c.lines)-1 {
			continue
		}

		lineNum := i + 1

		features, cursor, err := c.findNext(lineNum, "=", "{", "(", ";")
		if err != nil || len(features) == 0 {
			return fmt.Errorf(`findNext: %v`, err)
		}

		if strings.Contains(strings.Join(features, ""), " operator ") {
			// redo the search, but don't include "=" since "operator" is
			// a reserved keyword and must be an OverrideMethod.
			features, cursor, err = c.findNext(lineNum, "{", "(", ";")
			if err != nil || len(features) == 0 {
				return fmt.Errorf(`findNext: %v`, err)
			}
		}

		if features[len(features)-1] == "(" {
			// Include open paren in name.
			ss := strings.Join(features, "")
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

			c.e.logf("identifyOverrideMethodsAndVars: calling markMethod(line #%v, name=%q, %v), stripped=%q", lineNum+1, name, entityType, c.lines[lineNum].stripped)
			entity, err := c.markMethod(lineNum, name, entityType)
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

			// No open paren - could be a getter. See if it has a body.
			if features[len(features)-1] == "{" {
				// lineOffset := strings.Index(c.classBody, c.lines[lineNum].line)
				// inLineOffset := strings.Index(c.lines[lineNum].line, "{")
				// relOpenCurlyOffset := lineOffset + inLineOffset

				// if c.classBody[relOpenCurlyOffset] != '{' {
				// 	return fmt.Errorf("expected open curly bracket at relative offset %v but got %q", relOpenCurlyOffset, c.classBody[relOpenCurlyOffset:])
				// }

				absOpenCurlyOffset := cursor.absOffset - 1
				if c.e.fullBuf[absOpenCurlyOffset] != '{' {
					return fmt.Errorf("identifyOverrideMethodsAndVars: expected '{' at offset %v but got %c", absOpenCurlyOffset, c.e.fullBuf[absOpenCurlyOffset])
				}

				c.e.logf("identifyOverrideMethodsAndVars: calling findMatchingBracket(%v)", absOpenCurlyOffset)
				absCloseCurlyOffset, err := c.e.findMatchingBracket(absOpenCurlyOffset)
				if err != nil {
					return err
				}

				relCloseCurlyOffset := absCloseCurlyOffset - c.openCurlyOffset

				if c.classBody[relCloseCurlyOffset] != '}' {
					return fmt.Errorf("expected '}' at relative offset %v but got %c: %q", relCloseCurlyOffset, c.classBody[relCloseCurlyOffset], c.classBody[relCloseCurlyOffset:])
				}

				lineOffset := c.lines[lineNum].startOffset - c.openCurlyOffset
				bodyBuf := c.classBody[lineOffset : relCloseCurlyOffset+1]
				numLines := len(strings.Split(bodyBuf, "\n"))
				for j := 0; j < numLines; j++ {
					if c.lines[lineNum+j].entityType >= MainConstructor && c.lines[lineNum+j].entityType != entity.entityType {
						if err := c.repairIncorrectlyLabeledLine(lineNum + j); err != nil {
							return err
						}
					}
					c.e.logf("identifyOverrideMethodsAndVars: marking body line %v as type %v", lineNum+j+1, entity.entityType)
					c.lines[lineNum+j].entityType = entity.entityType
					entity.lines = append(entity.lines, c.lines[lineNum+j])
				}
			} else {
				// Does not have a body - if it has no fat arrow, it is a variable.
				if features[len(features)-1] != "=" ||
					(cursor.absOffset < len(c.e.fullBuf) && c.e.fullBuf[cursor.absOffset] != '>') {
					entity.entityType = OverrideVariable
				}

				features, cursor, err := c.findNext(lineNum, ";")
				if err != nil || len(features) == 0 {
					return fmt.Errorf(`findNext: %v`, err)
				}

				numLines := cursor.lineIndex - c.lines[lineNum].originalIndex
				for j := lineNum; j <= lineNum+numLines; j++ {
					if c.lines[j].entityType >= MainConstructor && c.lines[j].entityType != entity.entityType {
						if err := c.repairIncorrectlyLabeledLine(j); err != nil {
							return err
						}
					}
					c.e.logf("identifyOverrideMethodsAndVars: marking line %v as type %v", j+1, entity.entityType)
					c.lines[j].entityType = entity.entityType
					entity.lines = append(entity.lines, c.lines[j])
				}
			}

			// Preserve the comment lines leading up to the method.
			for lineNum--; lineNum > 0; lineNum-- {
				if isComment(c.lines[lineNum]) || strings.HasPrefix(c.lines[lineNum].stripped, "@") {
					c.e.logf("identifyOverrideMethodsAndVars: marking comment line %v as type %v", lineNum+1, entity.entityType)
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
	case MainConstructor:
		el := c.theConstructor
		for j := 0; j < len(el.lines); j++ {
			line := el.lines[j]
			if line != c.lines[lineNum] {
				continue
			}
			c.theConstructor.lines = append(c.theConstructor.lines[:j], c.theConstructor.lines[j+1:]...)
			if len(c.theConstructor.lines) == 0 {
				c.theConstructor = nil
			}
			return nil
		}
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
	case OverrideMethod:
		for i := 0; i < len(c.overrideMethods); i++ {
			el := c.overrideMethods[i]
			for j := 0; j < len(el.lines); j++ {
				line := el.lines[j]
				if line != c.lines[lineNum] {
					continue
				}
				c.overrideMethods[i].lines = append(c.overrideMethods[i].lines[:j], c.overrideMethods[i].lines[j+1:]...)
				if len(c.overrideMethods[i].lines) == 0 {
					c.overrideMethods = append(c.overrideMethods[:i], c.overrideMethods[i+1:]...)
				}
				return nil
			}
		}
	case OverrideVariable:
		for i := 0; i < len(c.overrideVariables); i++ {
			el := c.overrideVariables[i]
			for j := 0; j < len(el.lines); j++ {
				line := el.lines[j]
				if line != c.lines[lineNum] {
					continue
				}
				c.overrideVariables[i].lines = append(c.overrideVariables[i].lines[:j], c.overrideVariables[i].lines[j+1:]...)
				if len(c.overrideVariables[i].lines) == 0 {
					c.overrideVariables = append(c.overrideVariables[:i], c.overrideVariables[i+1:]...)
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

	features, cursor, err := c.findNext(lineNum, ";", "}")
	if err != nil || len(features) == 0 {
		return "", 0, "", fmt.Errorf(`findNext: %v`, err)
	}

	buildLeadingText := true
	var buildStr []string
	for i, f := range features {
		if strings.ContainsAny(f, "()[]{}=;") {
			buildLeadingText = false
			if f == "=" && i < len(features)-1 && features[i+1] == ">" {
				f = "=>"
			}
			result = append(result, f)
		}
		if buildLeadingText {
			buildStr = append(buildStr, f)
		}
	}
	leadingText := strings.TrimSpace(strings.Join(buildStr, ""))
	lineCount := cursor.lineIndex - c.lines[lineNum].originalIndex

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

	line := c.lines[lineNum]
	for {
		index := strings.Index(line.line, methodName)
		if index >= 0 {
			break
		}
		lineNum++
		if lineNum >= len(c.lines) {
			return nil, fmt.Errorf("expected to find methodName=%q, but ran out of lines", methodName)
		}
		line = c.lines[lineNum]
	}

	absOpenParenOffset := line.startOffset + strings.Index(line.line, methodName) + len(methodName) - 1
	if c.e.fullBuf[absOpenParenOffset] != '(' {
		return nil, fmt.Errorf("expected absOpenParenOffset=%v to be '(' but got '%c': line=%q", absOpenParenOffset, c.e.fullBuf[absOpenParenOffset], line.line)
	}
	lineOffset := line.startOffset - c.openCurlyOffset

	c.e.logf("markMethod(line #%v, methodName=%q, entityType=%v): absOpenParenOffset=%v, lineOffset=%v, calling findMatchingBracket(absOpenParenOffset=%v), line=%q", lineNum+1, methodName, entityType, absOpenParenOffset, lineOffset, absOpenParenOffset, line.line)
	absCloseParenOffset, err := c.e.findMatchingBracket(absOpenParenOffset)
	if err != nil {
		return nil, err
	}

	relCloseParenOffset := absCloseParenOffset - c.openCurlyOffset
	if c.classBody[relCloseParenOffset] != ')' {
		return nil, fmt.Errorf("expected close parenthesis at relative offset %v but got %v", relCloseParenOffset, c.classBody[relCloseParenOffset:])
	}

	fatArrowOffset := strings.Index(c.classBody[relCloseParenOffset:], "=>")
	curlyDeltaOffset := strings.Index(c.classBody[relCloseParenOffset:], "{")
	semicolonOffset := strings.Index(c.classBody[relCloseParenOffset:], ";")
	c.e.logf("fatArrowOffset=%v, curlyDeltaOffset=%v, semicolonOffset=%v", fatArrowOffset, curlyDeltaOffset, semicolonOffset)

	var nextOffset int
	if curlyDeltaOffset < 0 ||
		(curlyDeltaOffset >= 0 && fatArrowOffset >= 0 && semicolonOffset >= 0 && fatArrowOffset < curlyDeltaOffset && fatArrowOffset < semicolonOffset) ||
		(curlyDeltaOffset >= 0 && semicolonOffset >= 0 && semicolonOffset < curlyDeltaOffset) { // no curly-braced body.
		nextOffset = relCloseParenOffset + semicolonOffset
	} else {
		absOpenCurlyOffset := absCloseParenOffset + curlyDeltaOffset
		c.e.logf("markMethod: calling findMatchingBracket(absOpenCurlyOffset=%v)", absOpenCurlyOffset)
		absCloseCurlyOffset, err := c.e.findMatchingBracket(absOpenCurlyOffset)
		if err != nil {
			return nil, err
		}

		nextOffset = absCloseCurlyOffset - c.openCurlyOffset
	}

	if nextOffset+1 > len(c.classBody) {
		nextOffset = len(c.classBody) - 1
	}

	constructorBuf := c.classBody[lineOffset : nextOffset+1]
	c.e.logf("markMethod: constructorBuf[%v:%v]=%q", lineOffset, nextOffset+1, constructorBuf)
	numLines := len(strings.Split(constructorBuf, "\n"))
	for i := 0; i < numLines; i++ {
		if lineNum+i >= len(c.lines) {
			break
		}

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
