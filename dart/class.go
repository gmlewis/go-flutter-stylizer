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

import "strings"

// TextEditor represents an editor that can manipulate text.
type TextEditor interface {
}

// Class represents a Dart class.
type Class struct {
	editor                    TextEditor
	className                 string
	classOffset               int
	openCurlyOffset           int
	closeCurlyOffset          int
	groupAndSortGetterMethods bool
	fullBuf                   string
	lines                     []*Line // Line 0 is always the open curly brace.

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
func NewClass(editor TextEditor, className string, classOffset int,
	openCurlyOffset int, closeCurlyOffset int, groupAndSortGetterMethods bool) *Class {
	lessThanOffset := strings.Index(className, "<")
	if lessThanOffset >= 0 { // Strip off <T>.
		className = className[0:lessThanOffset]
	}

	return &Class{
		editor:                    editor,
		className:                 className,
		classOffset:               classOffset,
		openCurlyOffset:           openCurlyOffset,
		closeCurlyOffset:          closeCurlyOffset,
		groupAndSortGetterMethods: groupAndSortGetterMethods,
	}
}

func (c *Class) FindFeatures(buf string) {
	c.fullBuf = buf
	lines := strings.Split(c.fullBuf, "\n")
	var lineOffset int
	for _, line := range lines {
		c.lines = append(c.lines, NewLine(line, lineOffset))
		lineOffset += len(line)
		// Change a blank line following a comment to a SingleLineComment in
		// order to keep it with the following entity.
		numLines := len(c.lines)
		if numLines > 1 &&
			c.lines[numLines-1].entityType == BlankLine &&
			isComment(c.lines[numLines-2]) {
			c.lines[numLines-1].entityType = SingleLineComment
		}
	}

	c.identifyMultiLineComments()
	c.identifyMainConstructor()
	c.identifyNamedConstructors()
	c.identifyOverrideMethodsAndVars()
	c.identifyOthers()

	// c.lines.forEach((line, index) => console.log(`line #${index} type=${EntityType[line.entityType]}: ${line.line}`));
}

func (c *Class) genStripped(startLine int) string {
	var strippedLines []string
	for i := startLine; i < len(c.lines); i++ {
		strippedLines = append(strippedLines, c.lines[i].stripped)
	}
	return strings.Join(strippedLines, "\n")
}

func (c *Class) identifyMultiLineComments() {
	//     let inComment = false
	//     for (let i = 1; i < c.lines.length; i++) {
	//       const line = c.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }
	//       if (inComment) {
	//         c.lines[i].entityType = MultiLineComment
	//         // Note: a multiline comment followed by code on the same
	//         // line is not supported.
	//         const endComment = line.stripped.indexOf("*/")
	//         if (endComment >= 0) {
	//           inComment = false
	//           if (line.stripped.lastIndexOf("/*") > endComment + 1) {
	//             inComment = true
	//           }
	//         }
	//         continue
	//       }
	//       const startComment = line.stripped.indexOf("/*")
	//       if (startComment >= 0) {
	//         inComment = true
	//         c.lines[i].entityType = MultiLineComment
	//         if (line.stripped.lastIndexOf("*/") > startComment + 1) {
	//           inComment = false
	//         }
	//       }
	//     }
}

func (c *Class) identifyMainConstructor() {
	//     const className = c.className + '('
	//     for (let i = 1; i < c.lines.length; i++) {
	//       const line = c.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }
	//       const offset = line.stripped.indexOf(className)
	//       if (offset >= 0) {
	//         if (offset > 0) {
	//           const char = line.stripped.substring(offset - 1, offset)
	//           if (char !== ' ' && char !== '\t') {
	//             continue
	//           }
	//         }
	//         if (c.lines[i].entityType > MainConstructor) {
	//           c.repairIncorrectlyLabeledLine(i)
	//         }
	//         c.lines[i].entityType = MainConstructor
	//         c.theConstructor = await c.markMethod(i, className, MainConstructor)
	//         break
	//       }
	//     }
}

func (c *Class) identifyNamedConstructors() {
	//     const className = c.className + '.'
	//     for (let i = 1; i < c.lines.length; i++) {
	//       const line = c.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }
	//       const offset = line.stripped.indexOf(className)
	//       if (offset >= 0) {
	//         if (offset > 0) {
	//           const char = line.stripped.substring(offset - 1, offset)
	//           if (line.stripped[0] === '?' || line.stripped[0] === ':' || (char !== ' ' && char !== '\t')) {
	//             continue
	//           }
	//         }
	//         const openParenOffset = offset + line.stripped.substring(offset).indexOf('(')
	//         const namedConstructor = line.stripped.substring(offset, openParenOffset + 1)  // Include open parenthesis.
	//         if (c.lines[i].entityType >= MainConstructor && c.lines[i].entityType !== NamedConstructor) {
	//           c.repairIncorrectlyLabeledLine(i)
	//         }
	//         c.lines[i].entityType = NamedConstructor
	//         const entity = await c.markMethod(i, namedConstructor, NamedConstructor)
	//         c.namedConstructors.push(entity)
	//       }
	//     }
}

func (c *Class) identifyOverrideMethodsAndVars() {
	//     for (let i = 1; i < c.lines.length; i++) {
	//       const line = c.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }

	//       if (line.stripped.startsWith('@override') && i < c.lines.length - 1) {
	//         const offset = c.lines[i + 1].stripped.indexOf('(')
	//         if (offset >= 0) {
	//           // Include open paren in name.
	//           const ss = c.lines[i + 1].stripped.substring(0, offset + 1)
	//           // Search for beginning of method name.
	//           const nameOffset = ss.lastIndexOf(' ') + 1
	//           const name = ss.substring(nameOffset)
	//           const entityType = (name === 'build(') ? BuildMethod : OverrideMethod
	//           if (c.lines[i].entityType >= MainConstructor && c.lines[i].entityType !== entityType) {
	//             c.repairIncorrectlyLabeledLine(i)
	//           }
	//           c.lines[i].entityType = entityType
	//           const entity = await c.markMethod(i + 1, name, entityType)
	//           if (name === 'build(') {
	//             c.buildMethod = entity
	//           } else {
	//             c.overrideMethods.push(entity)
	//           }
	//         } else {
	//           const entity = new DartEntity()
	//           entity.entityType = OverrideMethod
	//           let lineNum = i + 1
	//           // No open paren - could be a getter. See if it has a body.
	//           if (c.lines[i + 1].stripped.indexOf('{') >= 0) {
	//             const lineOffset = c.fullBuf.indexOf(c.lines[i + 1].line)
	//             const inLineOffset = c.lines[i + 1].line.indexOf('{')
	//             const relOpenCurlyOffset = lineOffset + inLineOffset
	//             assert.strictEqual(c.fullBuf[relOpenCurlyOffset], '{', 'Expected open curly bracket at relative offset')
	//             const absOpenCurlyOffset = c.openCurlyOffset + relOpenCurlyOffset
	//             const absCloseCurlyOffset = await findMatchingBracket(c.editor, absOpenCurlyOffset)
	//             const relCloseCurlyOffset = absCloseCurlyOffset - c.openCurlyOffset
	//             assert.strictEqual(c.fullBuf[relCloseCurlyOffset], '}', 'Expected close curly bracket at relative offset')
	//             const nextOffset = absCloseCurlyOffset - c.openCurlyOffset
	//             const bodyBuf = c.fullBuf.substring(lineOffset, nextOffset + 1)
	//             const numLines = bodyBuf.split('\n').length
	//             for (let j = 0; j < numLines; j++) {
	//               if (c.lines[lineNum + j].entityType >= MainConstructor && c.lines[lineNum + j].entityType !== entity.entityType) {
	//                 c.repairIncorrectlyLabeledLine(lineNum + j)
	//               }
	//               c.lines[lineNum + j].entityType = entity.entityType
	//               entity.lines.push(c.lines[lineNum + j])
	//             }
	//           } else {
	//             // Does not have a body - if it has no fat arrow, it is a variable.
	//             if (c.lines[i + 1].stripped.indexOf('=>') < 0) {
	//               entity.entityType = OverrideVariable
	//             }
	//             // Find next ';', marking entityType forward.
	//             for (let j = i + 1; j < c.lines.length; j++) {
	//               if (c.lines[j].entityType >= MainConstructor && c.lines[j].entityType !== entity.entityType) {
	//                 c.repairIncorrectlyLabeledLine(j)
	//               }
	//               c.lines[j].entityType = entity.entityType
	//               entity.lines.push(c.lines[j])
	//               const semicolonOffset = c.lines[j].stripped.indexOf(';')
	//               if (semicolonOffset >= 0) {
	//                 break
	//               }
	//             }
	//           }
	//           // Preserve the comment lines leading up to the method.
	//           for (lineNum--; lineNum > 0; lineNum--) {
	//             if (isComment(c.lines[lineNum]) || c.lines[lineNum].stripped.startsWith('@')) {
	//               c.lines[lineNum].entityType = entity.entityType
	//               entity.lines.unshift(c.lines[lineNum])
	//               continue
	//             }
	//             break
	//           }
	//           if (entity.entityType === OverrideVariable) {
	//             c.overrideVariables.push(entity)
	//           } else {
	//             c.overrideMethods.push(entity)
	//           }
	//         }
	//       }
	//     }
}

func (c *Class) identifyOthers() {
	//     for (let i = 1; i < c.lines.length; i++) {
	//       const line = c.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }

	//       const entity = await c.scanMethod(i)
	//       if (entity.entityType === Unknown) {
	//         continue
	//       }

	//       // Preserve the comment lines leading up to the entity.
	//       for (let lineNum = i - 1; lineNum > 0; lineNum--) {
	//         if (isComment(c.lines[lineNum])) {
	//           c.lines[lineNum].entityType = entity.entityType
	//           entity.lines.unshift(c.lines[lineNum])
	//           continue
	//         }
	//         break
	//       }

	//       switch (entity.entityType) {
	//         case OtherMethod:
	//           c.otherMethods.push(entity)
	//           break
	//         case GetterMethod:
	//           c.getterMethods.push(entity)
	//           break
	//         case StaticVariable:
	//           c.staticVariables.push(entity)
	//           break
	//         case StaticPrivateVariable:
	//           c.staticPrivateVariables.push(entity)
	//           break
	//         case InstanceVariable:
	//           c.instanceVariables.push(entity)
	//           break
	//         case OverrideVariable:
	//           c.overrideVariables.push(entity)
	//           break
	//         case PrivateInstanceVariable:
	//           c.privateVariables.push(entity)
	//           break
	//         default:
	//           console.log('UNEXPECTED EntityType=', entity.entityType)
	//           break
	//       }
	//     }
}

func (c *Class) scanMethod(lineNum int) *Entity {
	entity := &Entity{}

	//     const buf = c.genStripped(lineNum)
	//     const result = c.findSequence(buf)
	//     const sequence = result[0]
	//     const lineCount = result[1]
	//     const leadingText = result[2]

	//     const nameParts = leadingText.split(' ')
	//     let staticKeyword = false
	//     let privateVar = false
	//     if (nameParts.length > 0) {
	//       entity.name = nameParts[nameParts.length - 1]
	//       if (entity.name.startsWith('_')) {
	//         privateVar = true
	//       }
	//       if (nameParts[0] === 'static') {
	//         staticKeyword = true
	//       }
	//     }
	//     entity.entityType = InstanceVariable
	//     switch (true) {
	//       case privateVar && staticKeyword:
	//         entity.entityType = StaticPrivateVariable
	//         break
	//       case staticKeyword:
	//         entity.entityType = StaticVariable
	//         break
	//       case privateVar:
	//         entity.entityType = PrivateInstanceVariable
	//         break
	//     }

	//     switch (sequence) {
	//       case '(){}':
	//         entity.entityType = OtherMethod
	//         break

	//       case '();':  // instance variable or abstract method.
	//         if (!leadingText.endsWith(' Function')) {
	//           entity.entityType = OtherMethod
	//         }
	//         break

	//       case '=(){}':
	//         entity.entityType = OtherMethod
	//         break

	//       default:
	//         if (sequence.indexOf('=>') >= 0) {
	//           entity.entityType = OtherMethod
	//         }
	//         break
	//     }

	//     // Force getters to be methods.
	//     if (leadingText.indexOf(' get ') >= 0) {
	//       if (c.groupAndSortGetterMethods) {
	//         entity.entityType = GetterMethod
	//       } else {
	//         entity.entityType = OtherMethod
	//       }
	//     }

	//     for (let i = 0; i <= lineCount; i++) {
	//       if (c.lines[lineNum + i].entityType >= MainConstructor && c.lines[lineNum + i].entityType !== entity.entityType) {
	//         c.repairIncorrectlyLabeledLine(lineNum + i)
	//       }
	//       c.lines[lineNum + i].entityType = entity.entityType
	//       entity.lines.push(c.lines[lineNum + i])
	//     }

	return entity
}

func (c *Class) repairIncorrectlyLabeledLine(lineNum int) {
	//     const incorrectLabel = c.lines[lineNum].entityType
	//     switch (incorrectLabel) {
	//       case NamedConstructor:
	//         for (let i = 0; i < c.namedConstructors.length; i++) {
	//           const el = c.namedConstructors[i]
	//           for (let j = 0; j < el.lines.length; j++) {
	//             const line = el.lines[j]
	//             if (line !== c.lines[lineNum]) { continue }
	//             c.namedConstructors[i].lines.splice(j, 1)
	//             if (c.namedConstructors[i].lines.length === 0) {
	//               c.namedConstructors.splice(i)
	//             }
	//             return
	//           }
	//         }
	//         break
	//       default:
	//         console.log(`repairIncorrectlyLabeledLine: Unhandled case ${incorrectLabel}. Please report on GitHub Issue Tracker with example test case.`)
	//         break
	//     }
}

func (c *Class) findSequence(buf string) (string, int, string) {
	var result []string

	var leadingText string
	var lineCount int
	//     let openParenCount = 0
	//     let openBraceCount = 0
	//     let openCurlyCount = 0
	//     for (let i = 0; i < buf.length; i++) {
	//       if (openParenCount > 0) {
	//         for (; i < buf.length; i++) {
	//           switch (buf[i]) {
	//             case '(':
	//               openParenCount++
	//               break
	//             case ')':
	//               openParenCount--
	//               break
	//             case '\n':
	//               lineCount++
	//               break
	//           }
	//           if (openParenCount === 0) {
	//             result.push(buf[i])
	//             break
	//           }
	//         }
	//       } else if (openBraceCount > 0) {
	//         for (; i < buf.length; i++) {
	//           switch (buf[i]) {
	//             case '[':
	//               openBraceCount++
	//               break
	//             case ']':
	//               openBraceCount--
	//               break
	//             case '\n':
	//               lineCount++
	//               break
	//           }
	//           if (openBraceCount === 0) {
	//             result.push(buf[i])
	//             return [result.join(''), lineCount, leadingText]
	//           }
	//         }
	//       } else if (openCurlyCount > 0) {
	//         for (; i < buf.length; i++) {
	//           switch (buf[i]) {
	//             case '{':
	//               openCurlyCount++
	//               break
	//             case '}':
	//               openCurlyCount--
	//               break
	//             case '\n':
	//               lineCount++
	//               break
	//           }
	//           if (openCurlyCount === 0) {
	//             result.push(buf[i])
	//             return [result.join(''), lineCount, leadingText]
	//           }
	//         }
	//       } else {
	//         switch (buf[i]) {
	//           case '(':
	//             openParenCount++
	//             result.push(buf[i])
	//             if (leadingText === '') {
	//               leadingText = buf.substring(0, i).trim()
	//             }
	//             break
	//           case '[':
	//             openBraceCount++
	//             result.push(buf[i])
	//             if (leadingText === '') {
	//               leadingText = buf.substring(0, i).trim()
	//             }
	//             break
	//           case '{':
	//             openCurlyCount++
	//             result.push(buf[i])
	//             if (leadingText === '') {
	//               leadingText = buf.substring(0, i).trim()
	//             }
	//             break
	//           case ';':
	//             result.push(buf[i])
	//             if (leadingText === '') {
	//               leadingText = buf.substring(0, i).trim()
	//             }
	//             return [result.join(''), lineCount, leadingText]
	//           case '=':
	//             if (i < buf.length - 1 && buf[i + 1] === '>') {
	//               result.push('=>')
	//             } else {
	//               result.push(buf[i])
	//             }
	//             if (leadingText === '') {
	//               leadingText = buf.substring(0, i).trim()
	//             }
	//             break
	//           case '\n':
	//             lineCount++
	//             break
	//         }
	//       }
	//     }
	return strings.Join(result, ""), lineCount, leadingText
}

//   func (c *Class) markMethod(lineNum int, methodName string, entityType: EntityType): Promise<DartEntity> {
//     assert.strictEqual(true, methodName.endsWith('('), 'markMethod: ' + methodName + ' must end with the open parenthesis.')
//     const entity = new DartEntity
//     entity.entityType = entityType
//     entity.lines = []
//     entity.name = methodName

//     // Identify all lines within the main (or factory) constructor.
//     const lineOffset = c.fullBuf.indexOf(c.lines[lineNum].line)
//     const inLineOffset = c.lines[lineNum].line.indexOf(methodName)
//     const relOpenParenOffset = lineOffset + inLineOffset + methodName.length - 1
//     assert.strictEqual(c.fullBuf[relOpenParenOffset], '(', 'Expected open parenthesis at relative offset')

//     const absOpenParenOffset = c.openCurlyOffset + relOpenParenOffset
//     const absCloseParenOffset = await findMatchingBracket(c.editor, absOpenParenOffset)
//     const relCloseParenOffset = absCloseParenOffset - c.openCurlyOffset
//     assert.strictEqual(c.fullBuf[relCloseParenOffset], ')', 'Expected close parenthesis at relative offset')

//     const curlyDeltaOffset = c.fullBuf.substring(relCloseParenOffset).indexOf('{')
//     const semicolonOffset = c.fullBuf.substring(relCloseParenOffset).indexOf(';')
//     let nextOffset = 0
//     if (curlyDeltaOffset < 0 || (curlyDeltaOffset >= 0 && semicolonOffset >= 0 && semicolonOffset < curlyDeltaOffset)) { // no body.
//       nextOffset = relCloseParenOffset + semicolonOffset
//     } else {
//       const absOpenCurlyOffset = absCloseParenOffset + curlyDeltaOffset
//       const absCloseCurlyOffset = await findMatchingBracket(c.editor, absOpenCurlyOffset)
//       nextOffset = absCloseCurlyOffset - c.openCurlyOffset
//     }
//     const constructorBuf = c.fullBuf.substring(lineOffset, nextOffset + 1)
//     const numLines = constructorBuf.split('\n').length
//     for (let i = 0; i < numLines; i++) {
//       if (c.lines[lineNum + i].entityType >= MainConstructor && c.lines[lineNum + i].entityType !== entityType) {
//         c.repairIncorrectlyLabeledLine(lineNum + i)
//       }
//       c.lines[lineNum + i].entityType = entityType
//       entity.lines.push(c.lines[lineNum + i])
//     }

//     // Preserve the comment lines leading up to the method.
//     for (lineNum--; lineNum > 0; lineNum--) {
//       if (isComment(c.lines[lineNum]) || c.lines[lineNum].stripped.startsWith('@')) {
//         c.lines[lineNum].entityType = entityType
//         entity.lines.unshift(c.lines[lineNum])
//         continue
//       }
//       break
//     }
//     return entity
//   }