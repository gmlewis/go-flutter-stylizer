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
	var lineStartOffset int
	for i, line := range lines {
		e.lines = append(e.lines, NewLine(line, lineStartOffset, i))
		lineStartOffset += len(line) + 1
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

	absOffset := openOffset + 1
	lineIndex, relStrippedOffset := e.findLineIndexAtOffset(absOffset)
	if lineIndex >= len(e.lines) {
		return 0, fmt.Errorf("findLineIndexAtOffset(%v) returned lineIndex(=%v) past last line %v", absOffset, lineIndex, len(e.lines))
	}

	e.logf("lineIndex=%v, relStrippedOffset=%v, stripped=%q(%v), line=%q(%v)", lineIndex, relStrippedOffset, e.lines[lineIndex].stripped, len(e.lines[lineIndex].stripped), e.lines[lineIndex].line, len(e.lines[lineIndex].line))

	reader := strings.NewReader("")
	if relStrippedOffset < 0 {
		return 0, fmt.Errorf("findLineIndexAtOffset(%v) returned bad relStrippedOffset(=%v) for stripped line %q(%v)", absOffset, relStrippedOffset, e.lines[lineIndex].stripped, len(e.lines[lineIndex].stripped))
	}
	if relStrippedOffset < len(e.lines[lineIndex].stripped) {
		reader = strings.NewReader(e.lines[lineIndex].stripped[relStrippedOffset:])
	}

	cursor := &Cursor{
		e:         e,
		absOffset: absOffset,
		lineIndex: lineIndex,

		relStrippedOffset: relStrippedOffset,

		reader: reader,
	}
	if _, err := cursor.advanceUntil(searchFor); err != nil {
		return 0, fmt.Errorf("advanceUntil(%q): %v", searchFor, err)
	}

	return cursor.absOffset - 1 - len(cursor.runeBuf), nil
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

// logf logs the line if verbose is true.
func (e *Editor) logf(fmtStr string, args ...interface{}) {
	if e.Verbose {
		log.Printf(fmtStr, args...)
	}
}
