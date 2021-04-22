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
	inMultiLineComment bool

	// UnreadRune can not be called twice in succession, so we need to keep track
	// of our lookahead stack ourselves.
	runeBuf []rune
}

func (c *Cursor) String() string {
	var line, stripped string
	if c.lineIndex < len(c.e.lines) {
		line = c.e.lines[c.lineIndex].line
		stripped = c.e.lines[c.lineIndex].stripped
	}
	return fmt.Sprintf(`{absOffet=%v, lineIndex=%v, relStrippedOffset=%v, stripped=%q(%v), line=%q(%v), '=%v, "=%v, '''=%v, """=%v, /*=%v, (=%v, braceLevels=%#v}`,
		c.absOffset, c.lineIndex, c.relStrippedOffset, stripped, len(stripped), line, len(line),
		c.inSingleQuote, c.inDoubleQuote, c.inTripleSingle, c.inTripleDouble, c.inMultiLineComment, c.parenLevels, c.braceLevels)
}

// advanceUntil searches for one of the provided strings, stepping through
// the Dart source code (and obeying its grammatical nuances) until found.
//
// It also adds a list of "interesting" top-level features encountered as it
// processes the data, effectively filtering out the contents of
// strings, bodies, and comments.
func (c *Cursor) advanceUntil(searchFor ...string) ([]string, error) {
	var features []string
	for {
		nf, err := c.advanceToNextFeature()
		if err != nil {
			return nil, err
		}

		c.e.logf("nf=%q searchFor=%#v, abs=%v, ind=%v, rel=%v", nf, searchFor, c.absOffset, c.lineIndex, c.relStrippedOffset)
		var foundIt bool
		for _, sf := range searchFor {
			if nf == sf {
				foundIt = true
				break
			}
		}

		if foundIt &&
			!c.inSingleQuote && !c.inDoubleQuote && !c.inTripleSingle && !c.inTripleDouble &&
			c.parenLevels == 0 &&
			len(c.braceLevels) == 0 {
			return append(features, nf), nil
		}

		switch nf {
		case "//":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			c.relStrippedOffset -= 2
			c.absOffset -= 2
			beforeLen := c.relStrippedOffset
			c.e.lines[c.lineIndex].stripped = strings.TrimSpace(c.e.lines[c.lineIndex].stripped[0:c.relStrippedOffset])
			afterLen := len(c.e.lines[c.lineIndex].stripped)
			c.absOffset -= beforeLen - afterLen
			// Reset the reader because we chopped off the stripped line.
			c.reader = strings.NewReader("")
			if afterLen == 0 {
				c.e.lines[c.lineIndex].entityType = SingleLineComment
			}
			c.e.logf("STRIPPED MODIFIED! singleLineComment=true: stripped=%q(%v), beforeLen=%v, afterLen=%v, cursor=%v", c.e.lines[c.lineIndex].stripped, len(c.e.lines[c.lineIndex].stripped), beforeLen, afterLen, c)
			continue
		case "/*":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			c.inMultiLineComment = true
			c.e.lines[c.lineIndex].entityType = MultiLineComment
			c.e.logf("inMultiLineComment=true: cursor=%v", c)
			continue
		case "*/":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble {
				continue
			}
			if !c.inMultiLineComment {
				return nil, fmt.Errorf("ERROR: Found */ before /*: cursor=%v", c)
			}
			c.e.lines[c.lineIndex].entityType = MultiLineComment
			c.inMultiLineComment = false
			c.e.logf("inMultiLineComment=false: cursor=%v", c)
			continue
		case "'''":
			if c.inMultiLineComment {
				continue
			}
			if c.inSingleQuote {
				return nil, fmt.Errorf("ERROR: Found ''' after ': cursor=%v", c)
			}
			if c.inDoubleQuote || c.inTripleDouble {
				continue
			}
			c.inTripleSingle = !c.inTripleSingle
			c.e.logf("inTripleSingle: cursor=%v", c)
			continue
		case `"""`:
			if c.inMultiLineComment {
				continue
			}
			if c.inDoubleQuote {
				return nil, fmt.Errorf(`ERROR: Found """ after ": cursor=%v`, c)
			}
			if c.inSingleQuote || c.inTripleSingle {
				continue
			}
			c.inTripleDouble = !c.inTripleDouble
			c.e.logf("inTripleDouble: cursor=%v", c)
			continue
		case "${":
			switch {
			case c.inMultiLineComment:
			case c.inSingleQuote:
				c.inSingleQuote = false
				c.braceLevels = append(c.braceLevels, BraceSingle)
				c.e.logf("${: inSingleQuote: cursor=%v", c)
			case c.inDoubleQuote:
				c.inDoubleQuote = false
				c.braceLevels = append(c.braceLevels, BraceDouble)
				c.e.logf("${: inDoubleQuote: cursor=%v", c)
			case c.inTripleSingle:
				c.inTripleSingle = false
				c.braceLevels = append(c.braceLevels, BraceTripleSingle)
				c.e.logf("${: inTripleSingle: cursor=%v", c)
			case c.inTripleDouble:
				c.inTripleDouble = false
				c.braceLevels = append(c.braceLevels, BraceTripleDouble)
				c.e.logf("${: inTripleDouble: cursor=%v", c)
			default:
				return nil, fmt.Errorf("ERROR: Found ${ outside of a string: cursor=%v", c)
			}
			continue
		case "'":
			if c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.inSingleQuote = !c.inSingleQuote
			c.e.logf("inSingleQuote: cursor=%v", c)
			continue
		case `"`:
			if c.inSingleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.inDoubleQuote = !c.inDoubleQuote
			c.e.logf("inDoubleQuote: cursor=%v", c)
			continue
		case "(":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			if c.parenLevels == 0 {
				features = append(features, nf)
			}
			c.parenLevels++
			c.e.logf("parenLevels++: cursor=%v", c)
		case ")":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.parenLevels--
			c.e.logf("parenLevels--: cursor=%v", c)
		case "{":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			if len(c.braceLevels) == 0 {
				features = append(features, nf)
			}
			c.braceLevels = append(c.braceLevels, BraceNormal)
			c.e.logf("{: cursor=%v", c)
		case "}":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			if len(c.braceLevels) == 0 {
				return nil, fmt.Errorf("ERROR: Found } before {: cursor=%v", c)
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
				return nil, fmt.Errorf("ERROR: Unknown braceLevel %v: cursor=%v", braceLevel, c)
			}
			c.e.logf("}: cursor=%v", c)
		}

		if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment || c.parenLevels > 0 || len(c.braceLevels) > 0 {
			continue
		}
		features = append(features, nf)

		if foundIt && len(searchFor) > 1 {
			return features, nil
		}
	}
}

// advanceToNextFeature moves the cursor to the next rune and returns
// it as a string, keeping track of where it is within the grammar.
// If it encounters a triple single- or triple double-quote or a "${" while
// within a string, it returns it as a whole string.
func (c *Cursor) advanceToNextFeature() (string, error) {
	var r rune
	var size int
	var err error

	if len(c.runeBuf) > 0 {
		r = c.runeBuf[0]
		size = len(string(r))
		c.runeBuf = c.runeBuf[1:]
	} else {
		r, size, err = c.reader.ReadRune()
		if err != nil {
			err := c.advanceToNextLine()
			if err != nil {
				return "", err
			}
			return c.advanceToNextFeature()
		}
	}

	c.absOffset += size
	c.relStrippedOffset += size

	if size > 1 { // a utf-8 rune of no interest; return it.
		return string(r), nil
	}

	switch r {
	case '\'':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "'", nil
		}
		if nsize != 1 || nr != '\'' {
			c.runeBuf = append(c.runeBuf, nr)
			return "'", nil
		}
		nr, nsize, err = c.reader.ReadRune()
		if err != nil {
			c.runeBuf = append(c.runeBuf, nr)
			return "'", nil
		}
		if nsize != 1 || nr != '\'' {
			c.runeBuf = append(c.runeBuf, r)
			c.runeBuf = append(c.runeBuf, nr)
			return "'", nil
		}
		c.absOffset += 2
		c.relStrippedOffset += 2
		return "'''", nil
	case '"':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return `"`, nil
		}
		if nsize != 1 || nr != '"' {
			c.runeBuf = append(c.runeBuf, nr)
			return `"`, nil
		}
		nr, nsize, err = c.reader.ReadRune()
		if err != nil {
			c.runeBuf = append(c.runeBuf, nr)
			return `"`, nil
		}
		if nsize != 1 || nr != '"' {
			c.runeBuf = append(c.runeBuf, r)
			c.runeBuf = append(c.runeBuf, nr)
			return `"`, nil
		}
		c.absOffset += 2
		c.relStrippedOffset += 2
		return `"""`, nil
	case '$':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "$", nil
		}
		if nsize != 1 || nr != '{' {
			c.runeBuf = append(c.runeBuf, nr)
			return "$", nil
		}
		c.absOffset++
		c.relStrippedOffset++
		return "${", nil
	case '/':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "/", nil
		}
		if nsize != 1 || (nr != '*' && nr != '/') {
			c.runeBuf = append(c.runeBuf, nr)
			return "/", nil
		}
		c.absOffset++
		c.relStrippedOffset++
		if nr == '/' {
			return "//", nil
		}
		return "/*", nil
	case '*':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "*", nil
		}
		if nsize != 1 || nr != '/' {
			c.runeBuf = append(c.runeBuf, nr)
			return "*", nil
		}
		c.absOffset++
		c.relStrippedOffset++
		return "*/", nil
	}

	return string(r), nil
}

// advanceToNextLine advances the cursor to the next line.
func (c *Cursor) advanceToNextLine() error {
	lineLengthDiff := len(c.e.lines[c.lineIndex].line) - len(c.e.lines[c.lineIndex].stripped)
	lineLengthDiff++ // newline
	c.absOffset += lineLengthDiff - c.e.lines[c.lineIndex].strippedOffset
	c.e.logf("advanceToNextLine: lineLengthDiff=%v, BEFORE absOffset=%v", lineLengthDiff, c.absOffset)
	c.lineIndex++

	if c.lineIndex >= len(c.e.lines) {
		return fmt.Errorf("advanceToNextLine went past EOF: cursor=%v", c)
	}

	c.absOffset += c.e.lines[c.lineIndex].strippedOffset
	c.e.logf("AFTER absOffset=%v", c.absOffset)
	c.relStrippedOffset = 0

	c.reader = strings.NewReader(c.e.lines[c.lineIndex].stripped)
	if c.inMultiLineComment {
		c.e.lines[c.lineIndex].entityType = MultiLineComment
	}
	return nil
}
