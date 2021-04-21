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
	"strings"
)

// Editor represents a text editor that understands Dart syntax.
type Editor struct {
	fullBuf string
	lines   []*Line

	Verbose bool
}

// NewEditor returns a new Editor.
func NewEditor(buf string) *Editor {
	e := &Editor{
		fullBuf: buf,
	}

	e.fullBuf = buf
	lines := strings.Split(e.fullBuf, "\n")
	var lineOffset int
	for i, line := range lines {
		e.lines = append(e.lines, NewLine(line, lineOffset, i))
		lineOffset += len(line)
		// Change a blank line following a comment to a SingleLineComment in
		// order to keep it with the following entity.
		numLines := len(e.lines)
		if numLines > 1 && e.lines[numLines-1].entityType == BlankLine && isComment(e.lines[numLines-2]) {
			e.lines[numLines-1].entityType = SingleLineComment
		}
	}

	return e
}

// findMatchingBracket finds the matching closing "}" (for "{") or ")" (for "(")
// given the offset of the opening rune.
func (e *Editor) findMatchingBracket(openOffset int) (int, error) {
	open := e.fullBuf[openOffset : openOffset+1]
	if open != "{" && open != "(" {
		return 0, fmt.Errorf("findMatchingBracket(%v) called on non-bracket %q. Please file a bug report on the GitHub issue tracker", openOffset, open)
	}
	searchFor := "}"
	if open == "(" {
		searchFor = ")"
	}

	e.logf("findClosing(openOffset=%v, searchFor=%q), open=%q", openOffset, searchFor, open)

	lineIndex, relOffset := e.findLineIndexAtOffset(openOffset)
	e.logf("lineIndex=%v, relOffset=%v, stripped=%q(%v), line=%q(%v)", lineIndex, relOffset, e.lines[lineIndex].stripped, len(e.lines[lineIndex].stripped), e.lines[lineIndex].line, len(e.lines[lineIndex].line))
	if e.lines[lineIndex].stripped[relOffset:relOffset+1] != open {
		return 0, fmt.Errorf("stripped[%v]=%q != open[%v]=%q", relOffset, e.lines[lineIndex].stripped[relOffset:relOffset+1], openOffset, open)
	}

	cursor := &Cursor{
		e:         e,
		absOffset: openOffset + 1,
		lineIndex: lineIndex,
		relOffset: relOffset + 1,
		reader:    strings.NewReader(e.lines[lineIndex].stripped[relOffset+1:]),
	}
	if _, err := cursor.advanceUntil(searchFor); err != nil {
		return 0, fmt.Errorf("advanceUntil(%q): %v", searchFor, err)
	}

	return cursor.absOffset - 1 - len(cursor.runeBuf), nil
}

// findLineIndexAtOffset finds the line index and relative offset for the
// absolute offset within the text buffer.
func (e *Editor) findLineIndexAtOffset(openOffset int) (int, int) {
	for i, line := range e.lines {
		if len(line.line) > openOffset {
			return i, openOffset - line.strippedOffset
		}
		openOffset -= len(line.line) + 1
	}
	return len(e.lines), openOffset
}

// logf logs the line if verbose is true.
func (e *Editor) logf(fmtStr string, args ...interface{}) {
	if e.Verbose {
		log.Printf(fmtStr, args...)
	}
}
