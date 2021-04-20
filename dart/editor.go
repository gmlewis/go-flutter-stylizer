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

import "log"

// findMatchingBracket finds the matching closing "}" (for "{") or ")" (for "(")
// given the offset of the opening rune.
func (c *Class) findMatchingBracket(openOffset int) int {
	open := c.fullBuf[openOffset : openOffset+1]
	if open != "{" && open != "(" {
		log.Fatalf("ERROR: findMatchingBracket(%v) called on non-bracket %q. Please file a bug report on the GitHub issue tracker.", openOffset, open)
	}
	searchFor := "}"
	if open == "(" {
		searchFor = ")"
	}

	return c.findClosing(openOffset, searchFor)
}

// findClosing finds searchFor and returns its absolute offset.
func (c *Class) findClosing(openOffset int, searchFor string) int {
	lineIndex, relOffset := c.findLineIndexAtOffset(openOffset)
	log.Printf("lineIndex=%v, relOffset=%v", lineIndex, relOffset)
	return 0
}

// findLineIndexAtOffset finds the line index and relative offset for the
// absolute offset within the text buffer.
func (c *Class) findLineIndexAtOffset(openOffset int) (int, int) {
	for i, line := range c.lines {
		if len(line.line) > openOffset {
			return i, openOffset
		}
		openOffset -= len(line.line) + 1
	}
	return len(c.lines), openOffset
}
