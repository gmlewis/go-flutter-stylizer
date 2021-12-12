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
	"io"
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
	separatePrivateMethods    bool

	theConstructor          *Entity
	namedConstructors       []*Entity
	staticVariables         []*Entity
	instanceVariables       []*Entity
	overrideVariables       []*Entity
	staticPrivateVariables  []*Entity
	privateVariables        []*Entity
	overrideMethods         []*Entity
	otherAllOrPublicMethods []*Entity
	otherPrivateMethods     []*Entity
	buildMethod             *Entity
	getterMethods           []*Entity
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
	groupAndSortGetterMethods, separatePrivateMethods bool) *Class {
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
	// c.e.logf("p[%v]=%q(%v), closeCurlyOffset=%v, len(classBody)=%v, last classLine=%#v", len(p)-1, p[len(p)-1], len(p[len(p)-1]), closeCurlyOffset, len(classBody), classLines[len(classLines)-1])

	// Replace last line for truncated classBody:
	ll := classLines[len(classLines)-1]
	classLines[len(classLines)-1] = &Line{
		line:              p[len(p)-1],
		stripped:          strings.TrimSpace(p[len(p)-1]),
		strippedOffset:    ll.strippedOffset,
		startOffset:       ll.startOffset,
		endOffset:         ll.endOffset,
		entityType:        BlankLine,
		originalIndex:     ll.originalIndex,
		isCommentOrString: ll.isCommentOrString,
	}

	return &Class{
		e:                editor,
		classBody:        classBody,
		lines:            classLines,
		className:        className,
		openCurlyOffset:  openCurlyOffset,
		closeCurlyOffset: closeCurlyOffset,

		groupAndSortGetterMethods: groupAndSortGetterMethods,
		separatePrivateMethods:    separatePrivateMethods,
	}
}

func (c *Class) FindFeatures() error {
	for i, line := range c.lines {
		// Change a blank line following a full-line comment to a SingleLineComment in
		// order to keep it with the following entity.
		if line.entityType == Unknown && line.stripped == "" {
			c.e.logf("FindFeatures: marking line %v as type BlankLine", i+1)
			line.entityType = BlankLine
		}
		if i > 1 && c.lines[i-1].entityType == BlankLine && isComment(c.lines[i-2]) && c.lines[i-2].stripped == "" {
			c.e.logf("FindFeatures: marking line %v as type SingleLineComment", i)
			c.lines[i-1].entityType = SingleLineComment
		}
	}

	// if c.e.Verbose {
	// 	c.e.logf("\n\nBEFORE IDENTIFICATION:")
	// 	for i := 0; i < len(c.lines); i++ {
	// 		line := c.lines[i]
	// 		c.e.logf("line #%v type=%v: %v", i+1, line.entityType, line.line)
	// 	}
	// }

	// c.identifyMultiLineComments()  // This step is not needed, as the editor marked these already.
	if err := c.identifyDecoratorsAsComments(); err != nil {
		return fmt.Errorf("identifyDecoratorsAsComments: %v", err)
	}
	if err := c.identifyMainConstructor(); err != nil {
		return fmt.Errorf("identifyMainConstructor: %v", err)
	}
	if err := c.identifyNamedConstructors(); err != nil {
		return fmt.Errorf("identifyNamedConstructors: %v", err)
	}
	if err := c.identifyOverrideMethodsAndVars(); err != nil {
		return fmt.Errorf("identifyOverrideMethodsAndVars: %v", err)
	}
	if err := c.identifyOthers(); err != nil {
		return fmt.Errorf("identifyOthers: %v", err)
	}

	if c.e.Verbose {
		for i, line := range c.lines {
			c.e.logf("line #%v type=%v: %v", i+1, line.entityType, line.line)
		}
	}

	// sanity check... better to catch errors early.
	for i, line := range c.lines {
		if i > 0 && line.entityType == Unknown {
			return fmt.Errorf("programming error: unhandled line #%v: %v", i+1, line.line)
		}
	}

	return nil
}

// findNext finds the next occurrence of one of the searchFor terms
// within the classLevelText (possibly spanning multiple lines).
//
// It returns a "features" string which is all class-level text up to
// and including the searched-for string.
//
// It also returns the absolute offset index (into the editor.fullBuf)
// of the last character returned.
//
// If no search terms can be found, the error io.EOF is returned.
func (c *Class) findNext(lineNum int, searchFor ...string) (features string, classLineNum int, absOffsetIndex int, err error) {
	classLineNum = lineNum

	for ; classLineNum < len(c.lines); classLineNum++ {
		line := c.lines[classLineNum]
		features += line.classLevelText
		classLevelIndex := -1
		var s string
		for si, ss := range searchFor {
			i := strings.Index(line.classLevelText, ss)
			if i < 0 {
				continue
			}
			if si != 0 && classLevelIndex >= 0 && i >= classLevelIndex {
				continue
			}

			classLevelIndex = i
			s = ss
		}

		if classLevelIndex >= 0 {
			if i := strings.Index(features, s); i > 0 {
				features = features[0 : i+len(s)]
			}

			if classLevelIndex >= len(line.classLevelTextOffsets) {
				return "", 0, 0, fmt.Errorf("programming error: classLevelIndex = %v but should be less than %v, line=%#v", classLevelIndex, len(line.classLevelTextOffsets), line)
			}

			absOffsetIndex := line.classLevelTextOffsets[classLevelIndex]

			return features, classLineNum, absOffsetIndex, nil
		}

		features += " " // instead of newline.
	}

	return "", 0, 0, io.EOF
}

func (c *Class) identifyDecoratorsAsComments() error {
	const override = "@override"

	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown || line.isCommentOrString || line.classLevelText == "" {
			continue
		}

		lineText := strings.TrimSpace(line.classLevelText)
		if !strings.HasPrefix(lineText, "@") || strings.HasPrefix(lineText, override) {
			continue
		}

		// Making the simplifying assumption here that if a decorator has arguments,
		// they should start on the same line as the decorator.
		lineIndex := i
		if strings.Contains(lineText, "(") {
			var features string
			var err error
			features, lineIndex, _, err = c.findNext(i, " ", "=", ";", "{", "(")
			if err != nil {
				lineIndex = i // Not an error, just reached EOF.
			}

			p := strings.Split(features, " ")
			if strings.HasSuffix(features, "(") && (len(p) == 1 || p[1] == "(") { // decorator includes args after '('
				var err error
				_, lineIndex, _, err = c.findNext(i, ")")
				if err != nil {
					return fmt.Errorf("unable to find end of decorator ')' with args from line #%v", c.lines[0].originalIndex+i+1)
				}
			}
		}

		for i <= lineIndex {
			c.e.logf("identifyDecoratorsAsComments: marking decorator line #%v as type SingleLineComment", i+1)
			c.lines[i].entityType = SingleLineComment
			i++
		}
		i--
	}

	return nil
}

func (c *Class) identifyMainConstructor() error {
	className := c.className + "("
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown || line.isCommentOrString || line.classLevelText == "" {
			continue
		}

		features, lineIndex, lastCharAbsOffset, err := c.findNext(i, "=", "{", ";", "(")
		if err != nil {
			// An error here is OK... it just means we reached EOF.
			return nil
		}

		advanceToNextLineIndex := func() {
			if !strings.HasSuffix(features, "}") && !strings.HasSuffix(features, ";") {
				features, lineIndex, lastCharAbsOffset, err = c.findNext(i, "}", ";")
				if err != nil {
					// An error here is OK... it just means we reached EOF.
					return
				}
			}
			for i+1 < len(c.lines) && i+1 <= lineIndex {
				i++
			}
		}

		if !strings.HasSuffix(features, "(") {
			advanceToNextLineIndex()
			continue
		}

		if c.e.fullBuf[lastCharAbsOffset] != '(' {
			return fmt.Errorf("programming error: expected fullBuf[%v]='(' but got '%c'", lastCharAbsOffset, c.e.fullBuf[lastCharAbsOffset])
		}

		classNameIndex := strings.Index(features, className)
		if classNameIndex < 0 {
			advanceToNextLineIndex()
			continue
		}

		if classNameIndex > 0 {
			char := features[classNameIndex-1 : classNameIndex]
			if char != " " && char != "\t" {
				advanceToNextLineIndex()
				continue
			}
		}

		if line.entityType > MainConstructor {
			if err := c.repairIncorrectlyLabeledLine(i); err != nil {
				return err
			}
		}
		c.e.logf("identifyMainConstructor: marking line %v as type MainConstructor", i+1)
		line.entityType = MainConstructor

		c.e.logf("identifyMainConstructor: calling markMethod(line #%v, className=%q, MainConstructor)", i+1, className)
		c.theConstructor, err = c.markMethod(i, className, MainConstructor, lastCharAbsOffset)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func (c *Class) identifyNamedConstructors() error {
	className := c.className + "."
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown || line.isCommentOrString || line.classLevelText == "" {
			continue
		}

		lineNum := i

		features, lineIndex, lastCharAbsOffset, err := c.findNext(lineNum, "=", "{", ";", "(")
		if err != nil {
			// An error here is OK... it just means we reached EOF.
			return nil
		}

		advanceToNextLineIndex := func() {
			if !strings.HasSuffix(features, "}") && !strings.HasSuffix(features, ";") {
				features, lineIndex, lastCharAbsOffset, err = c.findNext(i, "}", ";")
				if err != nil {
					// An error here is OK... it just means we reached EOF.
					return
				}
			}
			for i+1 < len(c.lines) && i+1 <= lineIndex {
				i++
			}
		}

		if !strings.HasSuffix(features, "(") {
			advanceToNextLineIndex()
			continue
		}

		if c.e.fullBuf[lastCharAbsOffset] != '(' {
			return fmt.Errorf("programming error: expected fullBuf[%v]='(' but got '%c'", lastCharAbsOffset, c.e.fullBuf[lastCharAbsOffset])
		}

		classNameIndex := strings.Index(features, className)
		if classNameIndex < 0 {
			advanceToNextLineIndex()
			continue
		}

		leadingText := features[0:classNameIndex]
		namedConstructor := features[classNameIndex:]
		c.e.logf("identifyNamedConstructors: leadingText=%q, classNameIndex=%v, namedConstructor=%q, lastCharAbsOffset=%v, features=%q", leadingText, classNameIndex, namedConstructor, lastCharAbsOffset, features)

		if strings.ContainsAny(leadingText, "?:") {
			advanceToNextLineIndex()
			continue
		}

		if c.lines[i].entityType >= MainConstructor && c.lines[i].entityType != NamedConstructor {
			if err := c.repairIncorrectlyLabeledLine(i); err != nil {
				return err
			}
		}
		c.e.logf("identifyNamedConstructors: marking line %v as type NamedConstructor", i+1)
		c.lines[i].entityType = NamedConstructor

		c.e.logf("identifyNamedConstructor: calling markMethod(line #%v, namedConstructor=%q, NamedConstructor)", i+1, namedConstructor)
		entity, err := c.markMethod(i, namedConstructor, NamedConstructor, lastCharAbsOffset)
		if err != nil {
			return err
		}
		c.namedConstructors = append(c.namedConstructors, entity)
	}

	return nil
}

func (c *Class) identifyOverrideMethodsAndVars() error {
	for i := 1; i < len(c.lines); i++ {
		if c.lines[i].entityType != Unknown || c.lines[i].isCommentOrString || c.lines[i].classLevelText == "" {
			continue
		}

		if !strings.HasPrefix(strings.TrimSpace(c.lines[i].classLevelText), "@override") || i >= len(c.lines)-1 {
			continue
		}

		lineNum := i

		features, lineIndex, lastCharAbsOffset, err := c.findNext(lineNum, "=>", "=", "{", ";", "(")
		if err != nil {
			return fmt.Errorf("expected valid @override method on line #%v: %v", lineNum+1, err)
		}

		if strings.HasPrefix(features, "operator") || strings.Contains(features, " operator") {
			// redo the search, but don't include "=" since "operator" is
			// a reserved keyword and must be an OverrideMethod.
			features, lineIndex, lastCharAbsOffset, err = c.findNext(lineNum, "{", "(", ";")
			if err != nil || features == "" {
				return fmt.Errorf("expected valid @override operator method on line #%v: %v", lineNum+1, err)
			}
		}

		f := func(i int) string {
			v := strings.TrimSpace(features[:i])
			nameOffset := strings.LastIndex(v, " ")
			if nameOffset >= 0 {
				return v[nameOffset+1:]
			}
			return v
		}

		if strings.HasSuffix(features, "(") {
			name := f(len(features) - 1)
			entityType := OverrideMethod
			if name == "build" {
				entityType = BuildMethod
			}

			c.e.logf("identifyOverrideMethodsAndVars: calling markMethod(line #%v, name=%q, %v), line=%q", i+1, name, entityType, c.lines[i].line)
			entity, err := c.markMethod(lineNum, name+"(", entityType, lastCharAbsOffset)
			if err != nil {
				return err
			}
			if name == "build" {
				c.buildMethod = entity
			} else {
				c.overrideMethods = append(c.overrideMethods, entity)
			}
			continue
		}

		entity := &Entity{
			entityType: OverrideMethod,
		}

		// No open paren - could be a getter. See if it has a body.
		if strings.HasSuffix(features, "{") {
			if c.e.fullBuf[lastCharAbsOffset] != '{' {
				return fmt.Errorf("programming error: expected '{' at offset %v but got %c", lastCharAbsOffset, c.e.fullBuf[lastCharAbsOffset])
			}

			entity.name = f(len(features) - 1)
		} else {
			// Does not have a body - if it has no fat arrow, it is a variable.
			if strings.HasSuffix(features, "=>") {
				entity.name = f(len(features) - 2)
			} else {
				entity.entityType = OverrideVariable
				entity.name = f(len(features) - 1)
			}

			if !strings.HasSuffix(features, ";") {
				features, lineIndex, lastCharAbsOffset, err = c.findNext(lineNum, ";")
				if err != nil || features == "" {
					return fmt.Errorf("expected trailing ';' for @override operator method on line #%v: %v", lineNum+1, err)
				}
			}
		}

		entity, err = c.markBody(entity, lineNum, entity.entityType, lineIndex, lastCharAbsOffset)
		if err != nil {
			return err
		}

		if entity.entityType == OverrideVariable {
			c.overrideVariables = append(c.overrideVariables, entity)
		} else {
			c.overrideMethods = append(c.overrideMethods, entity)
		}
	}

	return nil
}

func (c *Class) identifyOthers() error {
	for i := 1; i < len(c.lines); i++ {
		line := c.lines[i]
		if line.entityType != Unknown || line.isCommentOrString || line.classLevelText == "" {
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
				c.e.logf("identifyOthers: marking line %v as type %v", lineNum+1, entity.entityType)
				c.lines[lineNum].entityType = entity.entityType
				entity.lines = append([]*Line{c.lines[lineNum]}, entity.lines...)
				continue
			}
			break
		}

		switch entity.entityType {
		case OtherMethod:
			if c.separatePrivateMethods && entity.isPrivateMethod() {
				c.otherPrivateMethods = append(c.otherPrivateMethods, entity)
			} else {
				c.otherAllOrPublicMethods = append(c.otherAllOrPublicMethods, entity)
			}
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
	case staticKeyword:
		entity.entityType = StaticVariable
	case privateVar:
		entity.entityType = PrivateInstanceVariable
	}

	switch sequence {
	case "(){}":
		entity.entityType = OtherMethod

	case "();": // instance variable or abstract method.
		if !strings.HasSuffix(leadingText, " Function") {
			entity.entityType = OtherMethod
		}

	case "=(){}":
		entity.entityType = OtherMethod

	default:
		if strings.Index(sequence, "=>") >= 0 {
			entity.entityType = OtherMethod
		}
	}

	// Force getters to be methods.
	if strings.Index(leadingText, " get ") >= 0 {
		if c.groupAndSortGetterMethods {
			entity.entityType = GetterMethod
		} else {
			entity.entityType = OtherMethod
		}
	}

	for i := 0; i < lineCount; i++ {
		if lineNum+i >= len(c.lines) {
			break
		}

		if c.lines[lineNum+i].entityType >= MainConstructor && c.lines[lineNum+i].entityType != entity.entityType {
			c.e.logf("scanMethod: Changing line #%v from type %v to type %v", lineNum+i+1, c.lines[lineNum+i].entityType, entity.entityType)
			if err := c.repairIncorrectlyLabeledLine(lineNum + i); err != nil {
				return nil, err
			}
		}

		c.e.logf("scanMethod: marking line %v as type %v", lineNum+i+1, entity.entityType)
		c.lines[lineNum+i].entityType = entity.entityType
		entity.lines = append(entity.lines, c.lines[lineNum+i])
	}

	return entity, nil
}

func (c *Class) repairIncorrectlyLabeledLine(lineNum int) error {
	incorrectLabel := c.lines[lineNum].entityType
	switch incorrectLabel {
	default:
		return fmt.Errorf("repairIncorrectlyLabeledLine: class %q, class line #%v, file line #%v, unhandled case %v. Please report on GitHub Issue Tracker with example test case.", c.className, lineNum+1, c.lines[0].originalIndex+lineNum+1, incorrectLabel)
	}
}

func (c *Class) findSequence(lineNum int) (string, int, string, error) {
	var result string

	features, lineIndex, _, err := c.findNext(lineNum, ";", "}")
	if err != nil || features == "" {
		return "", 0, "", fmt.Errorf(`findNext: %v`, err)
	}

	buildLeadingText := true
	var buildStr string
	for i, f := range features {
		if strings.ContainsAny(string(f), "()[]{}=;") {
			buildLeadingText = false
			if f == '=' && i < len(features)-1 && features[i+1] == '>' {
				result += "=>"
				continue
			}
			result += string(f)
		}
		if buildLeadingText {
			buildStr += string(f)
		}
	}
	leadingText := strings.TrimSpace(buildStr)
	lineCount := lineIndex - lineNum + 1

	return result, lineCount, leadingText, nil
}

// markMethod marks an entire method with the same entityType.
// methodName must end with "(" and absOpenParenOffset must point to that open paren.
func (c *Class) markMethod(classLineNum int, methodName string, entityType EntityType, absOpenParenOffset int) (*Entity, error) {
	if !strings.HasSuffix(methodName, "(") {
		return nil, fmt.Errorf("programming error: markMethod: %q must end with the open parenthesis '('", methodName)
	}

	entity := &Entity{
		name:       methodName,
		entityType: entityType,
	}

	pair, ok := c.e.matchingPairs[absOpenParenOffset]
	if !ok {
		return nil, fmt.Errorf("programming error: expected '()' pair at absOpenParenOffset=%v, line=%#v", absOpenParenOffset, c.lines[classLineNum])
	}
	classCloseLineIndex := c.classCloseLineIndex(pair)

	features, classLineIndex, lastCharAbsOffset, err := c.findNext(classCloseLineIndex, "=>", "{", ";")
	if err != nil {
		return nil, fmt.Errorf("expected method body starting at classCloseLineIndex=%v: %v", classCloseLineIndex, err)
	}

	if strings.HasSuffix(features, "{") {
		c.e.logf("markMethod %q: moving past initializers: classLineIndex #%v, features=%v", methodName, classLineIndex+c.lines[0].originalIndex+1, features)
		for classLineIndex < len(c.lines)-1 && !strings.HasSuffix(c.lines[classLineIndex].classLevelText, " {") && !strings.HasSuffix(c.lines[classLineIndex].classLevelText, "}") {
			classLineIndex++
		}
		v := len(c.lines[classLineIndex].classLevelTextOffsets)
		if v != len(c.lines[classLineIndex].classLevelText) {
			return nil, fmt.Errorf("programming error: line #%v: classLevelText=%v != classLevelTextOffsets=%v", classLineIndex+1, len(c.e.lines[classLineIndex].classLevelText), len(c.e.lines[classLineIndex].classLevelTextOffsets))
		}
		if v > 0 {
			lastCharAbsOffset = c.lines[classLineIndex].classLevelTextOffsets[v-1]
		}

		c.e.logf("markMethod %q: after move past initializers: lastCharAbsOffset=%v, classLineIndex #%v, classLevelText=%v", methodName, lastCharAbsOffset, classLineIndex+c.lines[0].originalIndex+1, c.lines[classLineIndex].classLevelText)
	}

	if strings.HasSuffix(features, "=>") {
		features, classLineIndex, lastCharAbsOffset, err = c.findNext(classCloseLineIndex, "{", ";")
		if err != nil {
			return nil, fmt.Errorf("expected fat arrow method body starting at classCloseLineIndex=%v: %v", classCloseLineIndex, err)
		}
	}

	return c.markBody(entity, classLineNum, entityType, classLineIndex, lastCharAbsOffset)
}

func (c *Class) classCloseLineIndex(pair *MatchingPair) int {
	return pair.closeLineIndex - c.lines[0].originalIndex
}

// markBody marks an entire body with the same entityType.
// startClassLineNum is the starting class line index of the body.
// endClassLineNum is the ending class line index of the body (if ";" was used).
// lastCharAbsOffset must either point to the body's opening "{" or to its ending ";".
func (c *Class) markBody(entity *Entity, startClassLineNum int, entityType EntityType, endClassLineNum, lastCharAbsOffset int) (*Entity, error) {
	if c.e.fullBuf[lastCharAbsOffset] == '{' {
		pair, ok := c.e.matchingPairs[lastCharAbsOffset]
		if !ok {
			return nil, fmt.Errorf("expected matching '}' pair at lastCharAbsOffset=%v", lastCharAbsOffset)
		}
		if pair.open != "{" || pair.close != "}" {
			return nil, fmt.Errorf("programming error: expected '{' but got pair=%#v", pair)
		}
		endClassLineNum = c.classCloseLineIndex(pair)
	}

	c.e.logf("markBody marking lines #%v-%v as %v ...", startClassLineNum+1, endClassLineNum+1, entityType)
	for i := startClassLineNum; i <= endClassLineNum; i++ {
		if i >= len(c.lines) {
			break
		}

		if c.lines[i].entityType >= MainConstructor && c.lines[i].entityType != entityType {
			if err := c.repairIncorrectlyLabeledLine(i); err != nil {
				return nil, err
			}
		}

		c.e.logf("markMethod: marking line #%v as type %v", i+1, entityType)
		c.lines[i].entityType = entityType
		entity.lines = append(entity.lines, c.lines[i])
	}

	// Preserve the comment lines leading up to the method.
	for startClassLineNum--; startClassLineNum > 0; startClassLineNum-- {
		if isComment(c.lines[startClassLineNum]) || strings.HasPrefix(c.lines[startClassLineNum].stripped, "@") {
			c.e.logf("markMethod: marking comment line %v as type %v", startClassLineNum+1, entityType)
			c.lines[startClassLineNum].entityType = entityType
			entity.lines = append([]*Line{c.lines[startClassLineNum]}, entity.lines...)
			continue
		}
		break
	}

	return entity, nil
}
