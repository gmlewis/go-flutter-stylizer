/*
Copyright © 2021 Glenn M. Lewis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliee.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dart

import (
	"fmt"
	"log"
	"regexp"
	"runtime/debug"
	"strings"
)

var (
	matchClassOrMixinRE       = regexp.MustCompile(`^(?:abstract\s+)?(class|mixin)\s+(\S+).*$`)
	matchClassOrMixinOrEnumRE = regexp.MustCompile(`^(?:abstract\s+)?(class|mixin|enum)\s+(\S+).*$`)
)

// Editor represents a text editor that understands Dart syntax.
type Editor struct {
	fullBuf   string
	lines     []*Line
	eofOffset int

	opts Options

	matchingPairs MatchingPairsMap
	// classLineIndices contains line indices where a class or abstract class starts.
	classLineIndices []int

	classMatcher *regexp.Regexp
}

// NewEditor returns a new Editor.
func NewEditor(buf string, opts Options) (*Editor, error) {
	e := &Editor{
		fullBuf:       buf,
		matchingPairs: MatchingPairsMap{},
		classMatcher:  matchClassOrMixinRE,
		opts:          opts.validate(),
	}

	if opts.ProcessEnumsLikeClasses {
		e.classMatcher = matchClassOrMixinOrEnumRE
	}

	e.fullBuf = buf
	lines := strings.Split(e.fullBuf, "\n")
	var lineStartOffset int
	for i, line := range lines {
		e.lines = append(e.lines, NewLine(line, lineStartOffset, i))
		lineStartOffset += len(line) + 1
		// // Change a blank line following a comment to a SingleLineComment in
		// // order to keep it with the following entity.
		// numLines := len(e.lines)
		// if numLines > 1 && e.lines[numLines-1].entityType == BlankLine && isComment(e.lines[numLines-2]) {
		// 	e.lines[numLines-1].entityType = SingleLineComment
		// }
	}

	cursor := &Cursor{
		e:      e,
		reader: strings.NewReader(e.lines[0].stripped),
	}
	if err := cursor.parse(e.matchingPairs); err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}

	e.eofOffset = cursor.absOffset
	e.classLineIndices = cursor.classLineIndices

	return e, nil
}

func (e *Editor) GetClasses() ([]*Class, error) {
	var classes []*Class

	for _, lineIndex := range e.classLineIndices {
		line := e.lines[lineIndex]
		mm := e.classMatcher.FindStringSubmatch(line.line)
		if len(mm) != 3 {
			return nil, fmt.Errorf("programming error: expected class on line #%v, got %q", lineIndex+1, line.line)
		}

		classType := mm[1]
		className := mm[2]
		classOffset := line.startOffset
		openCurlyOffset := e.findStartOfClass(classOffset)
		if e.fullBuf[openCurlyOffset] == ';' { // this is valid and can be ignored: class D = Object with Function;
			continue
		}

		e.logf("\n\nFound new %v %q at classOffset=%v, openCurlyOffset=%v, line=%#v", classType, className, classOffset, openCurlyOffset, line)
		pair, ok := e.matchingPairs[openCurlyOffset]
		if !ok {
			return nil, fmt.Errorf("programming error: no matching pair found at openCurlyOffset %v", openCurlyOffset)
		}

		closeCurlyOffset := pair.closeAbsOffset
		e.logf("\n\nFound end of %v %q at closeCurlyOffset=%v", classType, className, closeCurlyOffset)

		dartClass := NewClass(e, classType, className, openCurlyOffset, closeCurlyOffset, e.opts)
		if err := dartClass.FindFeatures(); err != nil {
			return nil, err
		}

		classes = append(classes, dartClass)
	}

	return classes, nil
}

// findStartOfClass returns the absolute offset of the next top-level '{' or ';'
// starting at offset startOffset.
//
// Note that this class definition is valid: "class D = Object with Function;"
func (e *Editor) findStartOfClass(startOffset int) int {
	for startOffset < e.eofOffset {
		if e.fullBuf[startOffset] == '{' || e.fullBuf[startOffset] == ';' {
			return startOffset
		}
		if pair, ok := e.matchingPairs[startOffset]; ok {
			startOffset = pair.closeAbsOffset + 1
			continue
		}
		startOffset++
	}

	// offset := strings.Index(e.fullBuf[startOffset:], "{")
	// if semiOffset := strings.Index(e.fullBuf[startOffset:], ";"); offset < 0 || (semiOffset >= 0 && semiOffset < offset) {
	// 	offset = semiOffset // this is valid:
	// }
	// return startOffset + offset

	debug.PrintStack()
	log.Fatalf("programming error: findStartOfClass(%v) should not reach here", startOffset)
	return 0
}

// findLineIndexAtOffset finds the line index and relative offset for the
// absolute offset within the text buffer. Note that the returned
// relStrippedOffset may be greater than the length of the stripped string
// or even negative, before the stripped string.
func (e *Editor) findLineIndexAtOffset(absOffset int) (int, int) {
	for i, line := range e.lines {
		if absOffset <= line.endOffset {
			relStrippedOffset := absOffset - line.startOffset - line.strippedOffset
			return i, relStrippedOffset
		}
	}
	return len(e.lines), 0
}

// findClassAbsoluteStart finds the very first absolute offset either at the
// beginning of the buffer or after the last blank line before a class is defined.
func (e *Editor) findClassAbsoluteStart(dc *Class) int {
	result := dc.openCurlyOffset
	for ; result > 0; result-- {
		if result > 1 && e.fullBuf[result-2:result] == "\n\n" {
			break
		}
		if result > 3 && e.fullBuf[result-4:result] == "\r\n\r\n" {
			break
		}
	}
	return result
}

// logf logs the line if verbose is true.
func (e *Editor) logf(fmtStr string, args ...interface{}) {
	if e.opts.Verbose {
		log.Printf(fmtStr, args...)
	}
}
