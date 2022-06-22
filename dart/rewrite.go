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
	"sort"
	"strings"
)

// Edit represents an edit of an editor buffer.
type Edit struct {
	startPos int
	endPos   int
	text     string
}

func (c *Client) generateEdits(classes []*Class) []*Edit {
	var edits []*Edit

	for i := len(classes) - 1; i >= 0; i-- {
		dc := classes[i]
		startPos := dc.openCurlyOffset
		endPos := dc.closeCurlyOffset

		lines, changesMade := c.reorderClass(dc)
		if !changesMade {
			continue
		}

		edit := &Edit{
			startPos: startPos,
			endPos:   endPos,
			text:     strings.Join(lines, "\n"),
		}

		edits = append(edits, edit)
	}

	return edits
}

func (c *Client) reorderClass(dc *Class) ([]string, bool) {
	var lines []string

	lines = append(lines, dc.lines[0].line) // Curly brace.

	addEntity := func(entity *Entity, separateEntities bool) {
		if entity == nil {
			return
		}

		for _, line := range entity.lines {
			lines = append(lines, line.line)
		}

		if separateEntities || len(entity.lines) > 1 {
			if len(lines) > 0 && lines[len(lines)-1] != "" {
				lines = append(lines, "")
				c.logf("reorderClass.addEntity(%v): adding blank line #%v", entity.entityType, len(lines))
			}
		}
	}

	addEntities := func(entities []*Entity, separateEntities bool) {
		if len(entities) == 0 {
			return
		}
		for _, e := range entities {
			addEntity(e, separateEntities)
		}
		if !separateEntities && len(lines) > 0 && lines[len(lines)-1] != "" {
			lines = append(lines, "")
			c.logf("reorderClass.addEntities(%v): separateEntities=%v, adding blank line #%v", entities[0].entityType, separateEntities, len(lines))
		}
	}

	addEntitiesByVarTypes := func(entities []*Entity) {
		var finalVars, normalVars, optionalVars []*Entity
		for _, e := range entities {
			stripped := e.lines[0].stripped
			if strings.Contains(stripped, "final ") {
				finalVars = append(finalVars, e)
			} else if strings.Contains(stripped, "?") {
				optionalVars = append(optionalVars, e)
			} else {
				normalVars = append(normalVars, e)
			}
		}
		addEntities(finalVars, false)
		addEntities(normalVars, false)
		addEntities(optionalVars, false)
	}

	sortFunc := func(slice []*Entity) func(a, b int) bool {
		return func(a, b int) bool { return slice[a].name < slice[b].name }
	}

	for order, el := range c.opts.MemberOrdering {
		if !c.opts.Quiet {
			c.logf("Ordering step #%v: placing all '%v'...", order+1, el)
		}

		// Strip trailing blank lines.
		for len(lines) > 2 && lines[len(lines)-1] == "" && lines[len(lines)-2] == "" {
			c.logf("reorderClass(el='%v'): removing blank line #%v", el, len(lines))
			lines = lines[0 : len(lines)-1]
		}

		switch el {
		case "public-constructor":
			addEntity(dc.theConstructor, true)
		case "named-constructors":
			sort.SliceStable(dc.namedConstructors, sortFunc(dc.namedConstructors))
			addEntities(dc.namedConstructors, true)
		case "public-static-variables":
			sort.SliceStable(dc.staticVariables, sortFunc(dc.staticVariables))
			addEntities(dc.staticVariables, false)
		case "public-instance-variables":
			sort.SliceStable(dc.instanceVariables, sortFunc(dc.instanceVariables))
			if c.opts.GroupAndSortVariableTypes {
				addEntitiesByVarTypes(dc.instanceVariables)
			} else {
				addEntities(dc.instanceVariables, false)
			}
		case "public-override-variables":
			sort.SliceStable(dc.overrideVariables, sortFunc(dc.overrideVariables))
			addEntities(dc.overrideVariables, false)
		case "private-static-variables":
			sort.SliceStable(dc.staticPrivateVariables, sortFunc(dc.staticPrivateVariables))
			addEntities(dc.staticPrivateVariables, false)
		case "private-instance-variables":
			sort.SliceStable(dc.privateVariables, sortFunc(dc.privateVariables))
			addEntities(dc.privateVariables, false)
		case "public-override-methods":
			sort.SliceStable(dc.overrideMethods, sortFunc(dc.overrideMethods))
			addEntities(dc.overrideMethods, true)
		case "private-other-methods":
			if c.opts.SeparatePrivateMethods {
				if c.opts.SortOtherMethods {
					sort.SliceStable(dc.otherPrivateMethods, sortFunc(dc.otherPrivateMethods))
				}
				addEntities(dc.otherPrivateMethods, true)
			}
		case "public-other-methods":
			if c.opts.GroupAndSortGetterMethods {
				sort.SliceStable(dc.getterMethods, sortFunc(dc.getterMethods))
				addEntities(dc.getterMethods, false)
			}

			if c.opts.SortOtherMethods {
				sort.SliceStable(dc.otherAllOrPublicMethods, sortFunc(dc.otherAllOrPublicMethods))
			}
			addEntities(dc.otherAllOrPublicMethods, true)

			// Preserve random single-line and multi-line comments.
			for i := 1; i < len(dc.lines); i++ {
				var foundComment bool
				for ; i < len(dc.lines) && isComment(dc.lines[i]); i++ {
					lines = append(lines, dc.lines[i].line)
					foundComment = true
				}
				if foundComment {
					lines = append(lines, "")
				}
			}
		case "build-method":
			addEntity(dc.buildMethod, true)
		}
	}

	if !c.opts.Quiet && len(lines) > 0 {
		c.logf("Ordering done. Placed %v lines.", len(lines))
	}

	var changesMade bool
	if len(dc.lines) != len(lines) {
		changesMade = true
	} else {
		for i := 0; i < len(dc.lines); i++ {
			if dc.lines[i].line != lines[i] {
				changesMade = true
				break
			}
		}
	}

	return lines, changesMade
}

func (c *Client) rewriteClasses(buf string, edits []*Edit) string {
	newBuf := buf
	for _, edit := range edits {
		newBuf = strings.Join([]string{newBuf[0:edit.startPos], edit.text, newBuf[edit.endPos:]}, "")
	}
	return newBuf
}
