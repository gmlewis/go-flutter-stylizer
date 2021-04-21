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
	"log"
	"strings"
)

// DartEditor represents a text editor that understands Dart syntax.
type DartEditor struct {
	fullBuf string
	lines   []*Line
}

// NewEditor returns a new DartEditor.
func NewEditor(buf string) *DartEditor {
	d := &DartEditor{
		fullBuf: buf,
	}

	d.fullBuf = buf
	lines := strings.Split(d.fullBuf, "\n")
	var lineOffset int
	for _, line := range lines {
		d.lines = append(d.lines, NewLine(line, lineOffset))
		lineOffset += len(line)
		// Change a blank line following a comment to a SingleLineComment in
		// order to keep it with the following entity.
		numLines := len(d.lines)
		if numLines > 1 &&
			d.lines[numLines-1].entityType == BlankLine &&
			isComment(d.lines[numLines-2]) {
			d.lines[numLines-1].entityType = SingleLineComment
		}
	}

	for i := 0; i < len(d.lines); i++ {
		line := d.lines[i]
		log.Printf("line #%v type=%v: %v", i, line.entityType, line.line)
	}

	return d
}

// findMatchingBracket finds the matching closing "}" (for "{") or ")" (for "(")
// given the offset of the opening rune.
func (d *DartEditor) findMatchingBracket(openOffset int) int {
	open := d.fullBuf[openOffset : openOffset+1]
	if open != "{" && open != "(" {
		log.Fatalf("ERROR: findMatchingBracket(%v) called on non-bracket %q. Please file a bug report on the GitHub issue tracker.", openOffset, open)
	}
	searchFor := "}"
	if open == "(" {
		searchFor = ")"
	}

	return d.findClosing(openOffset, searchFor)
}

// findClosing finds searchFor and returns its absolute offset.
func (d *DartEditor) findClosing(openOffset int, searchFor string) int {
	lineIndex, relOffset := d.findLineIndexAtOffset(openOffset)
	log.Printf("lineIndex=%v, relOffset=%v, stripped=%q(%v), line=%q(%v)", lineIndex, relOffset, d.lines[lineIndex].stripped, len(d.lines[lineIndex].stripped), d.lines[lineIndex].line, len(d.lines[lineIndex].line))
	cursor := &Cursor{
		d:         d,
		absOffset: openOffset,
		lineIndex: lineIndex,
		relOffset: relOffset,
		reader:    strings.NewReader(d.lines[lineIndex].stripped[relOffset:]),
	}
	cursor.advanceUntil(searchFor)
	return cursor.absOffset
}

// findLineIndexAtOffset finds the line index and relative offset for the
// absolute offset within the text buffer.
func (d *DartEditor) findLineIndexAtOffset(openOffset int) (int, int) {
	for i, line := range d.lines {
		if len(line.line) > openOffset {
			return i, openOffset
		}
		openOffset -= len(line.line) + 1
	}
	return len(d.lines), openOffset
}
