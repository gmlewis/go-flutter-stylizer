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

// Package dart implements methods for manipulating Dart code.
package dart

import (
	"regexp"
	"strings"
)

var (
	commentRE = regexp.MustCompile(`^(.*?)\s*\/\/.*$`)
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

// Line represents a line of Dart code.
type Line struct {
	line        string
	stripped    string
	startOffset int
	endOffset   int
	entityType  EntityType
}

// NewLine returns a new Line.
func NewLine(line string, startOffset int) *Line {
	m := commentRE.FindStringSubmatch(line)
	stripped := strings.TrimSpace(line)
	if len(m) == 2 {
		stripped = strings.TrimSpace(m[1])
	}
	entityType := BlankLine
	if len(stripped) == 0 && strings.Index(line, "//") >= 0 {
		entityType = SingleLineComment
	}

	return &Line{
		line:        line,
		stripped:    stripped,
		startOffset: startOffset,
		endOffset:   startOffset + len(line) - 1,
		entityType:  entityType,
	}
}

func isComment(line *Line) bool {
	return line.entityType == SingleLineComment || line.entityType == MultiLineComment
}
