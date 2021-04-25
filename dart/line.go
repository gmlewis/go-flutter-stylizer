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

// Line represents a line of Dart code.
type Line struct {
	line           string // original, unmodified line
	stripped       string // removes comments and surrounding whitespace
	strippedOffset int    // offset to start of stripped line, compared to 'line'
	classLevelText string // preserved text at braceLevel==1 - Note that this is untrimmed.

	classLevelOffsets []int // absolute offsets for each character within classLevelText.

	originalIndex int

	startOffset int
	endOffset   int
	entityType  EntityType

	isCommentOrString bool // used when searching for new classes
}

// NewLine returns a new Line.
func NewLine(line string, startOffset, originalIndex int) *Line {
	stripped := strings.TrimSpace(line)
	entityType := Unknown
	if stripped == "" {
		entityType = BlankLine
	}
	strippedOffset := strings.Index(line, stripped)

	return &Line{
		line:           line,
		stripped:       stripped,
		strippedOffset: strippedOffset,

		originalIndex: originalIndex,

		startOffset: startOffset,
		endOffset:   startOffset + len(line),
		entityType:  entityType,
	}
}

func isComment(line *Line) bool {
	return line.entityType == SingleLineComment || line.entityType == MultiLineComment
}
