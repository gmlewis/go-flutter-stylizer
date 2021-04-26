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
	"log"
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

	for _, lineIndex := range editor.classLineIndices {
		line := editor.lines[lineIndex]
		mm := matchClassRE.FindStringSubmatch(line.line)
		if len(mm) != 2 {
			return nil, fmt.Errorf("programming error: expected class on line #%v, got %q", lineIndex+1, line.line)
		}

		className := mm[1]
		classOffset := line.startOffset
		openCurlyOffset := editor.findStartOfClass(classOffset)
		if editor.fullBuf[openCurlyOffset] == ';' { // this is valid and can be ignored: class D = Object with Function;
			continue
		}

		editor.logf("\n\nFound new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v", className, classOffset, openCurlyOffset, line)
		pair, ok := editor.matchingPairs[openCurlyOffset]
		if !ok {
			return nil, fmt.Errorf("programming error: no matching pair found at openCurlyOffset %v", openCurlyOffset)
		}

		closeCurlyOffset := pair.closeAbsOffset
		editor.logf("\n\nFound end of class %q at closeCurlyOffset=%v", className, closeCurlyOffset)

		dartClass := NewClass(editor, className, openCurlyOffset, closeCurlyOffset, groupAndSortGetterMethods)
		if err := dartClass.FindFeatures(); err != nil {
			return nil, err
		}

		classes = append(classes, dartClass)
	}

	// for _, line := range editor.lines {
	// 	if (strings.HasPrefix(line.line, "void main(") || strings.HasPrefix(line.line, "main(")) && strings.HasSuffix(line.line, ") {") {
	// 		// Find the end of main, marking sections as comments or strings to avoid them below.
	// 		openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, line.startOffset)
	// 		editor.findMatchingBracket(openCurlyOffset)
	// 		continue
	// 	}

	// 	mm := matchClassRE.FindStringSubmatch(line.line)
	// 	if len(mm) != 2 {
	// 		continue
	// 	}

	// 	className := mm[1]
	// 	classOffset := line.startOffset
	// 	openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, classOffset)

	// 	if editor.fullBuf[openCurlyOffset] == ';' {
	// 		continue
	// 	}

	// 	if openCurlyOffset <= classOffset {
	// 		return nil, fmt.Errorf(`expected "{" after "class" at offset %v`, classOffset)
	// 	}

	// 	if line.entityType != Unknown || line.isCommentOrString {
	// 		editor.logf("\n\nSkipping new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v due to entityType/comment/string", className, classOffset, openCurlyOffset, line)
	// 		continue
	// 	}

	// 	// This is a hack, but helps with files like `localizations_utils.dart` in
	// 	// the Flutter distribution.
	// 	if strings.ContainsAny(line.stripped, "$'") {
	// 		editor.logf("\n\nSkipping new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v due to appearance of $ or '", className, classOffset, openCurlyOffset, line)
	// 		continue
	// 	}

	// 	editor.logf("\n\nFound new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v", className, classOffset, openCurlyOffset, line)
	// 	closeCurlyOffset, err := editor.findMatchingBracket(openCurlyOffset)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	if closeCurlyOffset <= openCurlyOffset {
	// 		return nil, fmt.Errorf(`expected "}" after "{" at offset %v`, openCurlyOffset)
	// 	}
	// 	editor.logf("\n\nFound end of class %q at closeCurlyOffset=%v", className, closeCurlyOffset)

	// 	dartClass := NewClass(editor, className, openCurlyOffset, closeCurlyOffset, groupAndSortGetterMethods)
	// 	if err := dartClass.FindFeatures(); err != nil {
	// 		return nil, err
	// 	}

	// 	classes = append(classes, dartClass)
	// }

	return classes, nil
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
	if err := c.identifyDeprecatedAsComments(); err != nil {
		return fmt.Errorf("identifyDeprecatedAsComments: %v", err)
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
		if line.entityType != Unknown || line.isCommentOrString || line.classLevelText == "" {
			continue
		}

		lower := strings.ToLower(line.classLevelText)
		if !strings.HasPrefix(lower, depStr[0:len(depStr)-1]) {
			continue
		}

		c.e.logf("identifyDeprecatedAsComments: marking @deprecated line %v as type SingleLineComment", i+1)
		c.lines[i].entityType = SingleLineComment
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

		offset := strings.Index(line.classLevelText, className)
		if offset >= 0 {
			if offset > 0 {
				char := line.classLevelText[offset-1 : offset]
				if char != " " && char != "\t" {
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

			absOpenParenOffset := line.classLevelTextOffsets[offset+len(className)-1]
			c.e.logf("identifyMainConstructor: calling markMethod(line #%v, className=%q, MainConstructor)", i+1, className)
			var err error
			c.theConstructor, err = c.markMethod(i, className, MainConstructor, absOpenParenOffset)
			if err != nil {
				return err
			}
			break
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
			if si == 0 || classLevelIndex < 0 || i < classLevelIndex {
				classLevelIndex = i
				s = ss
			}
		}

		if classLevelIndex >= 0 {
			if i := strings.Index(features, s); i > 0 {
				features = features[0 : i+len(s)]
			}

			if classLevelIndex >= len(line.classLevelTextOffsets) {
				log.Fatalf("programming error: classLevelIndex = %v but should be less than %v, line=%#v", classLevelIndex, len(line.classLevelTextOffsets), line)
			}

			absOffsetIndex := line.classLevelTextOffsets[classLevelIndex]

			return features, classLineNum, absOffsetIndex, nil
		}

		features += " " // instead of newline.
	}

	return "", 0, 0, io.EOF
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
			for i+1 < len(c.lines) && i <= lineIndex {
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

		lineNum := i + 1

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
				return v[nameOffset:]
			}
			return v
		}

		if strings.HasSuffix(features, "(") {
			name := f(len(features))
			entityType := OverrideMethod
			if name == "build(" {
				entityType = BuildMethod
			}

			c.e.logf("identifyOverrideMethodsAndVars: calling markMethod(line #%v, name=%q, %v), line=%q", i+1, name, entityType, c.lines[i].line)
			entity, err := c.markMethod(i, name, entityType, lastCharAbsOffset)
			if err != nil {
				return err
			}
			if name == "build(" {
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

			c.e.logf("identifyOverrideMethodsAndVars: calling markBody(line #%v, name=OverrideMethod, %v), line=%q", i+1, entity.name, c.lines[i].line)
			entity, err = c.markBody(entity, i, OverrideMethod, lineIndex, lastCharAbsOffset)
			if err != nil {
				return err
			}
		} else {
			// Does not have a body - if it has no fat arrow, it is a variable.
			if !strings.HasSuffix(features, "=>") {
				entity.name = f(len(features) - 2)
				entity.entityType = OverrideVariable
			} else {
				entity.name = f(len(features) - 1)
			}

			if !strings.HasSuffix(features, ";") {
				features, lineIndex, lastCharAbsOffset, err = c.findNext(lineNum, ";")
				if err != nil || features == "" {
					return fmt.Errorf("expected trailing ';' for @override operator method on line #%v: %v", lineNum+1, err)
				}
			}

			entity, err = c.markBody(entity, i, entity.entityType, lineIndex, lastCharAbsOffset)
			if err != nil {
				return err
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
	// case MainConstructor:
	// 	el := c.theConstructor
	// 	for j := 0; j < len(el.lines); j++ {
	// 		line := el.lines[j]
	// 		if line != c.lines[lineNum] {
	// 			continue
	// 		}
	// 		c.theConstructor.lines = append(c.theConstructor.lines[:j], c.theConstructor.lines[j+1:]...)
	// 		if len(c.theConstructor.lines) == 0 {
	// 			c.theConstructor = nil
	// 		}
	// 		return nil
	// 	}
	// case NamedConstructor:
	// 	for i := 0; i < len(c.namedConstructors); i++ {
	// 		el := c.namedConstructors[i]
	// 		for j := 0; j < len(el.lines); j++ {
	// 			line := el.lines[j]
	// 			if line != c.lines[lineNum] {
	// 				continue
	// 			}
	// 			c.namedConstructors[i].lines = append(c.namedConstructors[i].lines[:j], c.namedConstructors[i].lines[j+1:]...)
	// 			if len(c.namedConstructors[i].lines) == 0 {
	// 				c.namedConstructors = append(c.namedConstructors[:i], c.namedConstructors[i+1:]...)
	// 			}
	// 			return nil
	// 		}
	// 	}
	// case OverrideMethod:
	// 	for i := 0; i < len(c.overrideMethods); i++ {
	// 		el := c.overrideMethods[i]
	// 		for j := 0; j < len(el.lines); j++ {
	// 			line := el.lines[j]
	// 			if line != c.lines[lineNum] {
	// 				continue
	// 			}
	// 			c.overrideMethods[i].lines = append(c.overrideMethods[i].lines[:j], c.overrideMethods[i].lines[j+1:]...)
	// 			if len(c.overrideMethods[i].lines) == 0 {
	// 				c.overrideMethods = append(c.overrideMethods[:i], c.overrideMethods[i+1:]...)
	// 			}
	// 			return nil
	// 		}
	// 	}
	// case OverrideVariable:
	// for i := 0; i < len(c.overrideVariables); i++ {
	// 	el := c.overrideVariables[i]
	// 	for j := 0; j < len(el.lines); j++ {
	// 		line := el.lines[j]
	// 		if line != c.lines[lineNum] {
	// 			continue
	// 		}
	// 		c.overrideVariables[i].lines = append(c.overrideVariables[i].lines[:j], c.overrideVariables[i].lines[j+1:]...)
	// 		if len(c.overrideVariables[i].lines) == 0 {
	// 			c.overrideVariables = append(c.overrideVariables[:i], c.overrideVariables[i+1:]...)
	// 		}
	// 		return nil
	// 	}
	// }
	default:
		return fmt.Errorf("repairIncorrectlyLabeledLine: line #%v, unhandled case %v. Please report on GitHub Issue Tracker with example test case.", lineNum+1, incorrectLabel)
	}

	// return nil
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

	features, classLineIndex, lastCharAbsOffset, err := c.findNext(pair.closeLineIndex-c.lines[0].originalIndex, "=>", "{", ";")
	if err != nil {
		return nil, fmt.Errorf("expected method body starting at classLineIndex=%v: %v", pair.closeLineIndex-c.lines[0].originalIndex, err)
	}

	if strings.HasSuffix(features, "=>") {
		features, classLineIndex, lastCharAbsOffset, err = c.findNext(pair.closeLineIndex-c.lines[0].originalIndex, "{", ";")
	}

	return c.markBody(entity, classLineNum, entityType, classLineIndex, lastCharAbsOffset)
}

// markBody marks an entire body with the same entityType.
// classLineNum is the starting class line index of the body.
// classLineIndex is the ending class line index of the body (if ";" was used).
// lastCharAbsOffset must either point to the body's opening "{" or to its ending ";".
func (c *Class) markBody(entity *Entity, classLineNum int, entityType EntityType, classLineIndex, lastCharAbsOffset int) (*Entity, error) {
	if c.e.fullBuf[lastCharAbsOffset] == '{' {
		pair, ok := c.e.matchingPairs[lastCharAbsOffset]
		if !ok {
			return nil, fmt.Errorf("expected matching '}' pair at lastCharAbsOffset=%v", lastCharAbsOffset)
		}
		classLineIndex = pair.closeLineIndex - c.lines[0].originalIndex
	}

	for i := classLineNum; i <= classLineIndex; i++ {
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
	for classLineNum--; classLineNum > 0; classLineNum-- {
		if isComment(c.lines[classLineNum]) || strings.HasPrefix(c.lines[classLineNum].stripped, "@") {
			c.e.logf("markMethod: marking comment line %v as type %v", classLineNum+1, entityType)
			c.lines[classLineNum].entityType = entityType
			entity.lines = append([]*Line{c.lines[classLineNum]}, entity.lines...)
			continue
		}
		break
	}

	return entity, nil
}
