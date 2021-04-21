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
	"strings"
)

// Class represents a Dart class.
type Class struct {
	e *Editor

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
// openCurlyOffset is the position of the "{" for that class.
// closeCurlyOffset is the position of the "}" for that class.
// groupAndSortGetterMethods determines how getter methods are processed.
func NewClass(editor *Editor, className string, openCurlyOffset int,
	closeCurlyOffset int, groupAndSortGetterMethods bool) *Class {
	lessThanOffset := strings.Index(className, "<")
	if lessThanOffset >= 0 { // Strip off <T>.
		className = className[0:lessThanOffset]
	}

	return &Class{
		e:                editor,
		className:        className,
		openCurlyOffset:  openCurlyOffset,
		closeCurlyOffset: closeCurlyOffset,

		groupAndSortGetterMethods: groupAndSortGetterMethods,
	}
}

func (c *Class) FindFeatures() {
	for i, line := range c.e.lines {
		// Change a blank line following a comment to a SingleLineComment in
		// order to keep it with the following entity.
		if line.entityType == Unknown && len(line.stripped) == 0 {
			line.entityType = BlankLine
		}
		if i > 1 && c.e.lines[i-1].entityType == BlankLine && isComment(c.e.lines[i-2]) {
			c.e.lines[i-1].entityType = SingleLineComment
		}
	}

	c.identifyMultiLineComments()
	c.identifyMainConstructor()
	c.identifyNamedConstructors()
	c.identifyOverrideMethodsAndVars()
	c.identifyOthers()

	if c.e.Verbose {
		for i := 0; i < len(c.e.lines); i++ {
			line := c.e.lines[i]
			c.e.logf("line #%v type=%v: %v", i, line.entityType, line.line)
		}
	}
}

func (c *Class) genStripped(startLine int) string {
	var strippedLines []string
	for i := startLine; i < len(c.e.lines); i++ {
		strippedLines = append(strippedLines, c.e.lines[i].stripped)
	}
	return strings.Join(strippedLines, "\n")
}

func (c *Class) identifyMultiLineComments() {
	var inComment bool
	for i := 1; i < len(c.e.lines); i++ {
		line := c.e.lines[i]
		if line.entityType != Unknown {
			continue
		}
		if inComment {
			c.e.lines[i].entityType = MultiLineComment
			// Note: a multiline comment followed by code on the same
			// line is not supported.
			endComment := strings.Index(line.stripped, "*/")
			if endComment >= 0 {
				inComment = false
				if strings.LastIndex(line.stripped, "/*") > endComment+1 {
					inComment = true
				}
			}
			continue
		}
		startComment := strings.Index(line.stripped, "/*")
		if startComment >= 0 {
			inComment = true
			c.e.lines[i].entityType = MultiLineComment
			if strings.LastIndex(line.stripped, "*/") > startComment+1 {
				inComment = false
			}
		}
	}
}

func (c *Class) identifyMainConstructor() {
	//     const className = c.className + '('
	//     for (let i = 1; i < c.e.lines.length; i++) {
	//       const line = c.e.lines[i]
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
	//         if (c.e.lines[i].entityType > MainConstructor) {
	//           c.repairIncorrectlyLabeledLine(i)
	//         }
	//         c.e.lines[i].entityType = MainConstructor
	//         c.theConstructor = await c.markMethod(i, className, MainConstructor)
	//         break
	//       }
	//     }
}

func (c *Class) identifyNamedConstructors() {
	//     const className = c.className + '.'
	//     for (let i = 1; i < c.e.lines.length; i++) {
	//       const line = c.e.lines[i]
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
	//         if (c.e.lines[i].entityType >= MainConstructor && c.e.lines[i].entityType !== NamedConstructor) {
	//           c.repairIncorrectlyLabeledLine(i)
	//         }
	//         c.e.lines[i].entityType = NamedConstructor
	//         const entity = await c.markMethod(i, namedConstructor, NamedConstructor)
	//         c.namedConstructors.push(entity)
	//       }
	//     }
}

func (c *Class) identifyOverrideMethodsAndVars() {
	//     for (let i = 1; i < c.e.lines.length; i++) {
	//       const line = c.e.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }

	//       if (line.stripped.startsWith('@override') && i < c.e.lines.length - 1) {
	//         const offset = c.e.lines[i + 1].stripped.indexOf('(')
	//         if (offset >= 0) {
	//           // Include open paren in name.
	//           const ss = c.e.lines[i + 1].stripped.substring(0, offset + 1)
	//           // Search for beginning of method name.
	//           const nameOffset = ss.lastIndexOf(' ') + 1
	//           const name = ss.substring(nameOffset)
	//           const entityType = (name === 'build(') ? BuildMethod : OverrideMethod
	//           if (c.e.lines[i].entityType >= MainConstructor && c.e.lines[i].entityType !== entityType) {
	//             c.repairIncorrectlyLabeledLine(i)
	//           }
	//           c.e.lines[i].entityType = entityType
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
	//           if (c.e.lines[i + 1].stripped.indexOf('{') >= 0) {
	//             const lineOffset = c.fullBuf.indexOf(c.e.lines[i + 1].line)
	//             const inLineOffset = c.e.lines[i + 1].line.indexOf('{')
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
	//               if (c.e.lines[lineNum + j].entityType >= MainConstructor && c.e.lines[lineNum + j].entityType !== entity.entityType) {
	//                 c.repairIncorrectlyLabeledLine(lineNum + j)
	//               }
	//               c.e.lines[lineNum + j].entityType = entity.entityType
	//               entity.lines.push(c.e.lines[lineNum + j])
	//             }
	//           } else {
	//             // Does not have a body - if it has no fat arrow, it is a variable.
	//             if (c.e.lines[i + 1].stripped.indexOf('=>') < 0) {
	//               entity.entityType = OverrideVariable
	//             }
	//             // Find next ';', marking entityType forward.
	//             for (let j = i + 1; j < c.e.lines.length; j++) {
	//               if (c.e.lines[j].entityType >= MainConstructor && c.e.lines[j].entityType !== entity.entityType) {
	//                 c.repairIncorrectlyLabeledLine(j)
	//               }
	//               c.e.lines[j].entityType = entity.entityType
	//               entity.lines.push(c.e.lines[j])
	//               const semicolonOffset = c.e.lines[j].stripped.indexOf(';')
	//               if (semicolonOffset >= 0) {
	//                 break
	//               }
	//             }
	//           }
	//           // Preserve the comment lines leading up to the method.
	//           for (lineNum--; lineNum > 0; lineNum--) {
	//             if (isComment(c.e.lines[lineNum]) || c.e.lines[lineNum].stripped.startsWith('@')) {
	//               c.e.lines[lineNum].entityType = entity.entityType
	//               entity.lines.unshift(c.e.lines[lineNum])
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
	//     for (let i = 1; i < c.e.lines.length; i++) {
	//       const line = c.e.lines[i]
	//       if (line.entityType !== Unknown) {
	//         continue
	//       }

	//       const entity = await c.scanMethod(i)
	//       if (entity.entityType === Unknown) {
	//         continue
	//       }

	//       // Preserve the comment lines leading up to the entity.
	//       for (let lineNum = i - 1; lineNum > 0; lineNum--) {
	//         if (isComment(c.e.lines[lineNum])) {
	//           c.e.lines[lineNum].entityType = entity.entityType
	//           entity.lines.unshift(c.e.lines[lineNum])
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
	//       if (c.e.lines[lineNum + i].entityType >= MainConstructor && c.e.lines[lineNum + i].entityType !== entity.entityType) {
	//         c.repairIncorrectlyLabeledLine(lineNum + i)
	//       }
	//       c.e.lines[lineNum + i].entityType = entity.entityType
	//       entity.lines.push(c.e.lines[lineNum + i])
	//     }

	return entity
}

func (c *Class) repairIncorrectlyLabeledLine(lineNum int) {
	//     const incorrectLabel = c.e.lines[lineNum].entityType
	//     switch (incorrectLabel) {
	//       case NamedConstructor:
	//         for (let i = 0; i < c.namedConstructors.length; i++) {
	//           const el = c.namedConstructors[i]
	//           for (let j = 0; j < el.lines.length; j++) {
	//             const line = el.lines[j]
	//             if (line !== c.e.lines[lineNum]) { continue }
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
//     const lineOffset = c.fullBuf.indexOf(c.e.lines[lineNum].line)
//     const inLineOffset = c.e.lines[lineNum].line.indexOf(methodName)
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
//       if (c.e.lines[lineNum + i].entityType >= MainConstructor && c.e.lines[lineNum + i].entityType !== entityType) {
//         c.repairIncorrectlyLabeledLine(lineNum + i)
//       }
//       c.e.lines[lineNum + i].entityType = entityType
//       entity.lines.push(c.e.lines[lineNum + i])
//     }

//     // Preserve the comment lines leading up to the method.
//     for (lineNum--; lineNum > 0; lineNum--) {
//       if (isComment(c.e.lines[lineNum]) || c.e.lines[lineNum].stripped.startsWith('@')) {
//         c.e.lines[lineNum].entityType = entityType
//         entity.lines.unshift(c.e.lines[lineNum])
//         continue
//       }
//       break
//     }
//     return entity
//   }
