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

// EntityType represents a type of Dart Line.
type EntityType int

const (
	Unknown EntityType = iota
	BlankLine
	SingleLineComment
	MultiLineComment
	MainConstructor
	NamedConstructor
	StaticVariable
	InstanceVariable
	OverrideVariable
	StaticPrivateVariable
	PrivateInstanceVariable
	OverrideMethod
	OtherMethod
	BuildMethod
	GetterMethod
)

func (e EntityType) String() string {
	switch e {
	default:
		return "Unknown"
	case BlankLine:
		return "BlankLine"
	case SingleLineComment:
		return "SingleLineComment"
	case MultiLineComment:
		return "MultiLineComment"
	case MainConstructor:
		return "MainConstructor"
	case NamedConstructor:
		return "NamedConstructor"
	case StaticVariable:
		return "StaticVariable"
	case InstanceVariable:
		return "InstanceVariable"
	case OverrideVariable:
		return "OverrideVariable"
	case StaticPrivateVariable:
		return "StaticPrivateVariable"
	case PrivateInstanceVariable:
		return "PrivateInstanceVariable"
	case OverrideMethod:
		return "OverrideMethod"
	case OtherMethod:
		return "OtherMethod"
	case BuildMethod:
		return "BuildMethod"
	case GetterMethod:
		return "GetterMethod"
	}
}

// Line represents a line of Dart code.
type Line struct {
	line           string
	stripped       string
	strippedOffset int

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
	if len(stripped) == 0 {
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
