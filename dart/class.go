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
