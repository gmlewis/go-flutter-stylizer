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
	"log"
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
	d *DartEditor

	absOffset int
	lineIndex int
	relOffset int
	reader    *strings.Reader

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
	if c.lineIndex < len(c.d.lines) {
		line = c.d.lines[c.lineIndex].line
		stripped = c.d.lines[c.lineIndex].stripped
	}
	return fmt.Sprintf(`{absOffet=%v, lineIndex=%v, relOffset=%v, stripped=%q(%v), line=%q(%v), '=%v, "=%v, '''=%v, """=%v, /*=%v, (=%v, braceLevels=%#v}`,
		c.absOffset, c.lineIndex, c.relOffset, stripped, len(stripped), line, len(line),
		c.inSingleQuote, c.inDoubleQuote, c.inTripleSingle, c.inTripleDouble, c.inMultiLineComment, c.parenLevels, c.braceLevels)
}

// advanceUntil searches for the provided string, stepping through
// the Dart source code (and obeying its grammatical nuances) until found.
func (c *Cursor) advanceUntil(searchFor string) {
	c.advanceToNextFeature() // Advance past the opening bracket to not count it.
	for {
		nf := c.advanceToNextFeature()
		log.Printf("nf=%q searchFor=%q, abs=%v, ind=%v, rel=%v", nf, searchFor, c.absOffset, c.lineIndex, c.relOffset)
		if nf == searchFor &&
			!c.inSingleQuote && !c.inDoubleQuote && !c.inTripleSingle && !c.inTripleDouble &&
			c.parenLevels == 0 && len(c.braceLevels) == 0 {
			return
		}
		switch nf {
		case "/*":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			c.inMultiLineComment = true
			log.Printf("inMultiLineComment=true: cursor=%v", c)
		case "*/":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble {
				continue
			}
			if !c.inMultiLineComment {
				log.Fatalf("ERROR: Found */ before /*: cursor=%v", c)
			}
			c.inMultiLineComment = false
			log.Printf("inMultiLineComment=false: cursor=%v", c)
		case "'''":
			if c.inMultiLineComment {
				continue
			}
			if c.inSingleQuote {
				log.Fatalf("ERROR: Found ''' after ': cursor=%v", c)
			}
			if c.inDoubleQuote || c.inTripleDouble {
				continue
			}
			c.inTripleSingle = !c.inTripleSingle
			log.Printf("inTripleSingle: cursor=%v", c)
		case `"""`:
			if c.inMultiLineComment {
				continue
			}
			if c.inDoubleQuote {
				log.Fatalf(`ERROR: Found """ after ": cursor=%v`, c)
			}
			if c.inSingleQuote || c.inTripleSingle {
				continue
			}
			c.inTripleDouble = !c.inTripleDouble
			log.Printf("inTripleDouble: cursor=%v", c)
		case "${":
			switch {
			case c.inMultiLineComment:
			case c.inSingleQuote:
				c.inSingleQuote = false
				c.braceLevels = append(c.braceLevels, BraceSingle)
				log.Printf("${: inSingleQuote: cursor=%v", c)
			case c.inDoubleQuote:
				c.inDoubleQuote = false
				c.braceLevels = append(c.braceLevels, BraceDouble)
				log.Printf("${: inDoubleQuote: cursor=%v", c)
			case c.inTripleSingle:
				c.inTripleSingle = false
				c.braceLevels = append(c.braceLevels, BraceTripleSingle)
				log.Printf("${: inTripleSingle: cursor=%v", c)
			case c.inTripleDouble:
				c.inTripleDouble = false
				c.braceLevels = append(c.braceLevels, BraceTripleDouble)
				log.Printf("${: inTripleDouble: cursor=%v", c)
			default:
				log.Fatalf("ERROR: Found ${ outside of a string: cursor=%v", c)
			}
		case "'":
			if c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.inSingleQuote = !c.inSingleQuote
			log.Printf("inSingleQuote: cursor=%v", c)
		case `"`:
			if c.inSingleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.inDoubleQuote = !c.inDoubleQuote
			log.Printf("inDoubleQuote: cursor=%v", c)
		case "(":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.parenLevels++
			log.Printf("parenLevels++: cursor=%v", c)
		case ")":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleDouble || c.inTripleSingle || c.inMultiLineComment {
				continue
			}
			c.parenLevels--
			log.Printf("parenLevels--: cursor=%v", c)
		case "{":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			c.braceLevels = append(c.braceLevels, BraceNormal)
			log.Printf("{: cursor=%v", c)
		case "}":
			if c.inSingleQuote || c.inDoubleQuote || c.inTripleSingle || c.inTripleDouble || c.inMultiLineComment {
				continue
			}
			if len(c.braceLevels) == 0 {
				log.Fatalf("ERROR: Found } before {: cursor=%v", c)
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
				log.Fatalf("ERROR: Unknown braceLevel %v: cursor=%v", braceLevel, c)
			}
			log.Printf("}: cursor=%v", c)
		}
	}
}

// advanceToNextFeature moves the cursor to the next rune and returns
// it as a string, keeping track of where it is within the grammar.
// If it encounters a triple single- or triple double-quote or a "${" while
// within a string, it returns it as a whole string.
func (c *Cursor) advanceToNextFeature() string {
	var r rune
	var size int
	var err error

	if len(c.runeBuf) > 0 {
		r = c.runeBuf[len(c.runeBuf)-1]
		size = len(string(r))
		c.runeBuf = c.runeBuf[0 : len(c.runeBuf)-1]
	} else {
		r, size, err = c.reader.ReadRune()
		if err != nil {
			c.advanceToNextLine()
			return c.advanceToNextFeature()
		}
	}

	c.absOffset += size
	c.relOffset += size

	if size > 1 { // a utf-8 rune of no interest; return it.
		return string(r)
	}

	switch r {
	case '\'':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "'"
		}
		if nsize != 1 || nr != '\'' {
			c.runeBuf = append(c.runeBuf, nr)
			return "'"
		}
		nr, nsize, err = c.reader.ReadRune()
		if err != nil {
			c.runeBuf = append(c.runeBuf, nr)
			return "'"
		}
		if nsize != 1 || nr != '\'' {
			c.runeBuf = append(c.runeBuf, r)
			c.runeBuf = append(c.runeBuf, nr)
			return "'"
		}
		c.absOffset += 2
		c.relOffset += 2
		return "'''"
	case '"':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return `"`
		}
		if nsize != 1 || nr != '"' {
			c.runeBuf = append(c.runeBuf, nr)
			return `"`
		}
		nr, nsize, err = c.reader.ReadRune()
		if err != nil {
			c.runeBuf = append(c.runeBuf, nr)
			return `"`
		}
		if nsize != 1 || nr != '"' {
			c.runeBuf = append(c.runeBuf, r)
			c.runeBuf = append(c.runeBuf, nr)
			return `"`
		}
		c.absOffset += 2
		c.relOffset += 2
		return `"""`
	case '$':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "$"
		}
		if nsize != 1 || nr != '{' {
			c.runeBuf = append(c.runeBuf, nr)
			return "$"
		}
		c.absOffset++
		c.relOffset++
		return "${"
	case '/':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "/"
		}
		if nsize != 1 || nr != '*' {
			c.runeBuf = append(c.runeBuf, nr)
			return "/"
		}
		c.absOffset++
		c.relOffset++
		return "/*"
	case '*':
		nr, nsize, err := c.reader.ReadRune()
		if err != nil {
			return "*"
		}
		if nsize != 1 || nr != '/' {
			c.runeBuf = append(c.runeBuf, nr)
			return "*"
		}
		c.absOffset++
		c.relOffset++
		return "*/"
	}

	return string(r)
}

// advanceToNextLine advances the cursor to the next line.
func (c *Cursor) advanceToNextLine() {
	lineLengthDiff := len(c.d.lines[c.lineIndex].line) - len(c.d.lines[c.lineIndex].stripped)
	c.absOffset += lineLengthDiff
	c.lineIndex++
	c.relOffset = 0
	if c.lineIndex >= len(c.d.lines) {
		log.Fatalf("ERROR: advanceToNextLine went past EOF: cursor=%v", c)
	}
	c.reader = strings.NewReader(c.d.lines[c.lineIndex].stripped)
}
