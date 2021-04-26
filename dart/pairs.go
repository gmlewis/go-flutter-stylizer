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
	"log"
	"runtime/debug"
)

// MatchingPairsMap represents a map of matching pairs keyed by
// the openAbsOffset value.
type MatchingPairsMap map[int]*MatchingPair

// MatchingPair represents a matching pair of features in the Dart
// source code, such as r'''=>''', (=>), {=>}, etc.
type MatchingPair struct {
	open  string
	close string

	openAbsOffset  int
	closeAbsOffset int

	openLineIndex  int
	closeLineIndex int

	openRelStrippedOffset  int
	closeRelStrippedOffset int

	pairStackDepth          int
	parentPairOpenAbsOffset int
}

// NewMatchingPair adds the start of a new matching pair onto the stack.
func (c *Cursor) NewMatchingPair(open string, matchingPairs MatchingPairsMap, matchingPairStack []*MatchingPair) []*MatchingPair {
	absOffset := c.absOffset - len(open)
	pair := &MatchingPair{
		open:                  open,
		openAbsOffset:         absOffset,
		openLineIndex:         c.lineIndex,
		openRelStrippedOffset: c.relStrippedOffset,
		pairStackDepth:        len(matchingPairStack),
	}

	if got := c.e.fullBuf[pair.openAbsOffset : pair.openAbsOffset+len(open)]; got != open {
		debug.PrintStack()
		log.Fatalf("programming error: openAbsOffset = %q, want %q", got, open)
	}

	if i := len(matchingPairStack); i > 0 {
		pair.parentPairOpenAbsOffset = matchingPairStack[i-1].openAbsOffset
	}

	matchingPairs[absOffset] = pair

	return append(matchingPairStack, pair)
}

// CloseMatchingPair closes the last open matching pair on the stack.
func (c *Cursor) CloseMatchingPair(close string, matchingPairStack []*MatchingPair) []*MatchingPair {
	pair := matchingPairStack[len(matchingPairStack)-1]

	pair.close = close
	pair.closeAbsOffset = c.absOffset - len(close)
	pair.closeLineIndex = c.lineIndex
	pair.closeRelStrippedOffset = c.relStrippedOffset

	if got := c.e.fullBuf[pair.closeAbsOffset : pair.closeAbsOffset+len(close)]; got != close {
		debug.PrintStack()
		log.Fatalf("programming error: closeAbsOffset = %q, want %q", got, close)
	}

	return matchingPairStack[0 : len(matchingPairStack)-1]
}
