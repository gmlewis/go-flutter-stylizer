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
	"fmt"
	"io"
	"strings"
)

// BraceLevel represents the current brace level the cursor is within.
type BraceLevel int

const (
	BraceUnknown BraceLevel = iota

	BraceNormal       // Most recent unmatched '{' was not string literal related.
	BraceSingle       // Most recent unmatched '{' was `'...${`.
	BraceDouble       // Most recent unmatched '{' was `"...${`.
	BraceTripleSingle // Most recent unmatched '{' was `'''...${`.
	BraceTripleDouble // Most recent unmatched '{' was `"""...${`.
)

// Cursor represents an editor cursor and is used to advance through
// the Dart source code.
type Cursor struct {
	e *Editor

	absOffset int
	lineIndex int

	relStrippedOffset int

	reader *strings.Reader

	braceLevels        []BraceLevel
	parenLevels        int
	inSingleQuote      bool
	inDoubleQuote      bool
	inTripleSingle     bool
	inTripleDouble     bool
	inMultiLineComment int // multiline comments can nest
	stringIsRaw        bool

	// UnreadRune can not be called twice in succession, so we need to keep track
	// of our lookahead stack ourselves.
	runeBuf []rune

	// classLineIndices keeps track of line indices where a class or abstract class starts.
	classLineIndices []int
}

func (c *Cursor) String() string {
	var line, stripped string
	if c.lineIndex < len(c.e.lines) {
		line = c.e.lines[c.lineIndex].line
		stripped = c.e.lines[c.lineIndex].stripped
	}
	return fmt.Sprintf(`{absOffset=%v, lineIndex=%v, relStrippedOffset=%v, stripped=%q(%v), line=%q(%v), raw=%v, '=%v, "=%v, '''=%v, """=%v, /*=%v, (=%v, braceLevels=%#v}`,
		c.absOffset, c.lineIndex, c.relStrippedOffset, stripped, len(stripped), line, len(line), c.stringIsRaw,
		c.inSingleQuote, c.inDoubleQuote, c.inTripleSingle, c.inTripleDouble, c.inMultiLineComment, c.parenLevels, c.braceLevels)
}

// parse parses the Dart source, identifies line entity types in the source,
// keeps track of matching pairs, and returns a list of class line indices.
func (c *Cursor) parse(matchingPairs MatchingPairsMap) (err error) {
	var lastFeature string
	var nf string // nextFeature
	var matchingPairStack []*MatchingPair

	for {
		lastFeature = nf
		nf, err = c.advanceToNextFeature()
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("advanceToNextFeature: %v", err)
			}
			return nil
		}

		c.e.logf("nf=%q matchingPairStack=%v, abs=%v, ind=%v, rel=%v", nf, len(matchingPairStack), c.absOffset, c.lineIndex, c.relStrippedOffset)

		switch nf {
		case "//":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment > 0 {
				continue
			}
			c.relStrippedOffset -= 2
			c.absOffset -= 2
			beforeLen := c.relStrippedOffset
			c.e.lines[c.lineIndex].stripped = strings.TrimSpace(c.e.lines[c.lineIndex].stripped[0:c.relStrippedOffset])
			afterLen := len(c.e.lines[c.lineIndex].stripped)
			c.absOffset -= beforeLen - afterLen
			// Reset the reader because we chopped off the stripped line and have no more to parse on this line.
			c.reader = strings.NewReader("")
			if afterLen == 0 {
				c.e.logf("parse: marking line #%v as type SingleLineComment", c.lineIndex+1)
				c.e.lines[c.lineIndex].entityType = SingleLineComment
			}
			// c.e.logf("STRIPPED MODIFIED! singleLineComment=true: stripped=%q(%v), beforeLen=%v, afterLen=%v, cursor=%v", c.e.lines[c.lineIndex].stripped, len(c.e.lines[c.lineIndex].stripped), beforeLen, afterLen, c)
		case "/*":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble {
				continue
			}
			c.inMultiLineComment++
			c.e.logf("advanceUntil: marking line #%v as type MultiLineComment", c.lineIndex+1)
			c.e.lines[c.lineIndex].entityType = MultiLineComment
			c.e.logf("inMultiLineComment=%v: cursor=%v", c.inMultiLineComment, c)
			matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
		case "*/":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble {
				continue
			}
			if c.inMultiLineComment == 0 {
				return fmt.Errorf("ERROR: Found */ before /*: cursor=%v", c)
			}
			c.e.logf("advanceUntil: marking line %v as type MultiLineComment", c.lineIndex+1)
			c.e.lines[c.lineIndex].entityType = MultiLineComment
			c.inMultiLineComment--
			c.e.logf("inMultiLineComment=%v: cursor=%v", c.inMultiLineComment, c)
			matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
		case "'''":
			if c.inMultiLineComment > 0 {
				continue
			}
			if c.inSingleQuote {
				return fmt.Errorf("ERROR: Found ''' after ': cursor=%v", c)
			}
			if c.inDoubleQuote || c.inTripleDouble {
				continue
			}
			c.inTripleSingle = !c.inTripleSingle
			c.stringIsRaw = c.inTripleSingle && lastFeature == "r"
			c.e.logf("inTripleSingle: cursor=%v", c)
			if c.inTripleSingle {
				if c.stringIsRaw {
					matchingPairStack = c.NewMatchingPair("r'''", matchingPairs, matchingPairStack)
				} else {
					matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
				}
			} else {
				matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
			}
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1, c.absOffset, c.absOffset+1)
			}
		case `"""`:
			if c.inMultiLineComment > 0 {
				continue
			}
			if c.inDoubleQuote {
				return fmt.Errorf(`ERROR: Found """ after ": cursor=%v`, c)
			}
			if c.inSingleQuote || c.inTripleSingle {
				continue
			}
			c.inTripleDouble = !c.inTripleDouble
			c.stringIsRaw = c.inTripleDouble && lastFeature == "r"
			c.e.logf("inTripleDouble: cursor=%v", c)
			if c.inTripleDouble {
				if c.stringIsRaw {
					matchingPairStack = c.NewMatchingPair(`r"""`, matchingPairs, matchingPairStack)
				} else {
					matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
				}
			} else {
				matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
			}
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1, c.absOffset, c.absOffset+1)
			}
		case "${":
			switch {
			case c.inMultiLineComment > 0:
			case c.inSingleQuote:
				if c.stringIsRaw {
					continue
				}
				c.inSingleQuote = false
				c.braceLevels = append(c.braceLevels, BraceSingle)
				c.e.logf("${: inSingleQuote: cursor=%v", c)
				matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
			case c.inDoubleQuote:
				if c.stringIsRaw {
					continue
				}
				c.inDoubleQuote = false
				c.braceLevels = append(c.braceLevels, BraceDouble)
				c.e.logf("${: inDoubleQuote: cursor=%v", c)
				matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
			case c.inTripleSingle:
				if c.stringIsRaw {
					continue
				}
				c.inTripleSingle = false
				c.braceLevels = append(c.braceLevels, BraceTripleSingle)
				c.e.logf("${: inTripleSingle: cursor=%v", c)
				matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
			case c.inTripleDouble:
				if c.stringIsRaw {
					continue
				}
				c.inTripleDouble = false
				c.braceLevels = append(c.braceLevels, BraceTripleDouble)
				c.e.logf("${: inTripleDouble: cursor=%v", c)
				matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
			default:
				return fmt.Errorf("ERROR: Found ${ outside of a string: cursor=%v", c)
			}
		case "'":
			if c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment > 0 {
				continue
			}
			c.inSingleQuote = !c.inSingleQuote
			c.stringIsRaw = c.inSingleQuote && lastFeature == "r"
			c.e.logf("inSingleQuote: lastFeature=%q, cursor=%v", lastFeature, c)
			if c.inSingleQuote {
				if c.stringIsRaw {
					matchingPairStack = c.NewMatchingPair("r'", matchingPairs, matchingPairStack)
				} else {
					matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
				}
			} else {
				matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
			}
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
		case `"`:
			if c.inSingleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment > 0 {
				continue
			}
			c.inDoubleQuote = !c.inDoubleQuote
			c.stringIsRaw = c.inDoubleQuote && lastFeature == "r"
			c.e.logf("inDoubleQuote: cursor=%v", c)
			if c.inDoubleQuote {
				if c.stringIsRaw {
					matchingPairStack = c.NewMatchingPair(`r"`, matchingPairs, matchingPairStack)
				} else {
					matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
				}
			} else {
				matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
			}
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
		case "(":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment > 0 {
				continue
			}
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
			c.parenLevels++
			c.e.logf("parenLevels++: cursor=%v", c)
			matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
		case ")":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment > 0 {
				continue
			}
			c.parenLevels--
			c.e.logf("parenLevels--: cursor=%v", c)
			matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
		case "{":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment > 0 {
				continue
			}
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
			c.braceLevels = append(c.braceLevels, BraceNormal)
			c.e.logf("{: cursor=%v", c)
			matchingPairStack = c.NewMatchingPair(nf, matchingPairs, matchingPairStack)
		case "}":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment > 0 {
				continue
			}
			if len(c.braceLevels) == 0 {
				return fmt.Errorf("ERROR: Found } before {: cursor=%v", c)
			}
			braceLevel := c.braceLevels[len(c.braceLevels)-1]
			c.braceLevels = c.braceLevels[:len(c.braceLevels)-1]
			switch braceLevel {
			case BraceNormal:
			case BraceSingle:
				c.inSingleQuote = true
			case BraceDouble:
				c.inDoubleQuote = true
			case BraceTripleSingle:
				c.inTripleSingle = true
			case BraceTripleDouble:
				c.inTripleDouble = true
			default:
				return fmt.Errorf("ERROR: Unknown braceLevel %v: cursor=%v", braceLevel, c)
			}
			c.e.logf("}: cursor=%v", c)
			matchingPairStack = c.CloseMatchingPair(nf, matchingPairStack)
			if c.parenLevels == 0 && len(c.braceLevels) == 1 {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
		default:
			if c.atTopOfBraceLevel(1) {
				c.e.lines[c.lineIndex].classLevelText += nf
				c.e.lines[c.lineIndex].classLevelTextOffsets = append(c.e.lines[c.lineIndex].classLevelTextOffsets, c.absOffset-1)
			}
		}
	}
}

// advanceToNextFeature moves the cursor to the next rune and returns
// it as a string, keeping track of where it is within the grammar.
// If it encounters a triple single- or triple double-quote or a "${" while
// within a string, it returns it as a whole string.
func (c *Cursor) advanceToNextFeature() (string, error) {
	getRune := func() (rune, int, error) {
		if len(c.runeBuf) > 0 {
			c.e.logf("Grabbing rune from runeBuf... before=%#v", c.runeBuf)
			r := c.runeBuf[0]
			size := len(string(r))
			c.runeBuf = c.runeBuf[1:]
			c.absOffset += size
			c.relStrippedOffset += size
			c.e.logf("rune=%c, size=%v, after=%#v", r, size, c.runeBuf)
			return r, size, nil
		}

		r, size, err := c.reader.ReadRune()
		if err != nil {
			return 0, 0, err
		}
		c.absOffset += size
		c.relStrippedOffset += size
		return r, size, nil
	}

	r, size, err := getRune()
	if err != nil {
		if err := c.advanceToNextLine(); err != nil {
			if !c.atTopOfBraceLevel(0) {
				return "", fmt.Errorf("parse error: reached EOF, cursor=%v", c)
			}
			return "", err
		}
		r, size, err = ' ', 1, nil // replace newline with single space
	}

	if size > 1 { // a utf-8 rune of no interest; return it.
		return string(r), nil
	}

	pushRune := func(r rune) {
		c.runeBuf = append(c.runeBuf, r)
		size := len(string(r))
		c.absOffset -= size
		c.relStrippedOffset -= size
		c.e.logf("Pushing rune=%c, size=%v to runeBuf: after=%#v", r, len(string(r)), c.runeBuf)
	}

	switch r {
	case '\\':
		if c.stringIsRaw || c.inMultiLineComment > 0 {
			return string(r), nil
		}
		nr, _, err := getRune()
		if err != nil {
			return `\`, nil
		}
		return fmt.Sprintf("\\%c", nr), nil
	case '\'':
		if (c.stringIsRaw || c.inDoubleQuote || c.inTripleDouble) && !c.inTripleSingle {
			return string(r), nil
		}
		nr, nsize, err := getRune()
		if err != nil {
			return "'", nil
		}
		if nsize != 1 || nr != '\'' {
			pushRune(nr)
			return "'", nil
		}
		// r == nr == '\' at this point.
		nr, nsize, err = getRune()
		if err != nil {
			pushRune(r)
			return "'", nil
		}
		if nsize != 1 || nr != '\'' {
			pushRune(r)
			pushRune(nr)
			return "'", nil
		}
		return "'''", nil
	case '"':
		if (c.stringIsRaw || c.inSingleQuote || c.inTripleSingle) && !c.inTripleDouble {
			return string(r), nil
		}
		nr, nsize, err := getRune()
		if err != nil {
			return `"`, nil
		}
		if nsize != 1 || nr != '"' {
			pushRune(nr)
			return `"`, nil
		}
		// r == nr == '"' at this point.
		nr, nsize, err = getRune()
		if err != nil {
			pushRune(r)
			return `"`, nil
		}
		if nsize != 1 || nr != '"' {
			pushRune(r)
			pushRune(nr)
			return `"`, nil
		}
		return `"""`, nil
	case '$':
		if c.stringIsRaw {
			return string(r), nil
		}
		nr, nsize, err := getRune()
		if err != nil {
			return "$", nil
		}
		if nsize != 1 || nr != '{' {
			pushRune(nr)
			return "$", nil
		}
		return "${", nil
	case '/':
		if c.stringIsRaw {
			return string(r), nil
		}
		nr, nsize, err := getRune()
		if err != nil {
			return "/", nil
		}
		if nsize != 1 || (nr != '*' && nr != '/') {
			pushRune(nr)
			return "/", nil
		}
		if nr == '/' {
			return "//", nil
		}
		return "/*", nil
	case '*':
		if c.stringIsRaw {
			return string(r), nil
		}
		nr, nsize, err := getRune()
		if err != nil {
			return "*", nil
		}
		if nsize != 1 || nr != '/' {
			pushRune(nr)
			return "*", nil
		}
		return "*/", nil
	}

	return string(r), nil
}

func (c *Cursor) atTopOfBraceLevel(braceLevel int) bool {
	if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment > 0 || c.parenLevels > 0 {
		return false
	}
	return len(c.braceLevels) == braceLevel
}

// advanceToNextLine advances the cursor to the next line.
// It returns io.EOF when it reaches the end of the file.
//
// It also detects the start of class lines.
func (c *Cursor) advanceToNextLine() error {
	if c.lineIndex == 0 && matchClassRE.FindStringSubmatch(c.e.lines[c.lineIndex].line) != nil {
		c.classLineIndices = append(c.classLineIndices, c.lineIndex)
	}

	// c.e.lines[c.lineIndex].classLevelText = strings.TrimSpace(c.e.lines[c.lineIndex].classLevelText)

	c.lineIndex++
	if c.lineIndex >= len(c.e.lines) {
		return io.EOF
	}

	if c.atTopOfBraceLevel(0) && matchClassRE.FindStringSubmatch(c.e.lines[c.lineIndex].line) != nil {
		c.classLineIndices = append(c.classLineIndices, c.lineIndex)
	}

	c.absOffset = c.e.lines[c.lineIndex].startOffset + c.e.lines[c.lineIndex].strippedOffset
	c.relStrippedOffset = 0

	if c.e.lines[c.lineIndex].stripped != "" && c.e.fullBuf[c.absOffset] != c.e.lines[c.lineIndex].stripped[0] {
		return fmt.Errorf("programming error: fullBuf[%v]=%c, want %c", c.absOffset, c.e.fullBuf[c.absOffset], c.e.lines[c.lineIndex].stripped[0])
	}

	c.reader = strings.NewReader(c.e.lines[c.lineIndex].stripped)
	if c.inMultiLineComment > 0 {
		c.e.logf("advanceToNextLine: marking line #%v as MultiLineComment", c.lineIndex+1)
		c.e.lines[c.lineIndex].entityType = MultiLineComment
	}
	if c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment > 0 {
		c.e.logf("advanceToNextLine: marking line #%v as isCommentOrString", c.lineIndex+1)
		c.e.lines[c.lineIndex].isCommentOrString = true
	}
	return nil
}
