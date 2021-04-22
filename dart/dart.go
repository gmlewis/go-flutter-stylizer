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
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

// Options represents the configuration options for the Dart processor.
type Options struct {
	Diff    bool
	List    bool
	Quiet   bool
	Verbose bool
	Write   bool

	GroupAndSortGetterMethods bool

	MemberOrdering   []string
	SortOtherMethods bool
}

// Client represents a Dart processor.
type Client struct {
	opts Options
}

var defaultMemberOrdering = []string{
	"public-constructor",
	"named-constructors",
	"public-static-variables",
	"public-instance-variables",
	"public-override-variables",
	"private-static-variables",
	"private-instance-variables",
	"public-override-methods",
	"public-other-methods",
	"build-method",
}

// New returns a new Dart processor.
func New(opts Options) *Client {
	if !validateMemberOrdering(opts.MemberOrdering) {
		opts.MemberOrdering = defaultMemberOrdering
	}

	return &Client{opts: opts}
}

// StylizeFile sylizes a single Dart file using the provided options.
// If any differences were found, true is returned.
func (c *Client) StylizeFile(filename string) (bool, error) {
	var buf string
	{
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return false, err
		}
		buf = string(b)
	}

	e := NewEditor(buf)
	e.Verbose = c.opts.Verbose
	classes, err := c.getClasses(e, c.opts.GroupAndSortGetterMethods)
	if err != nil {
		return false, err
	}
	if !c.opts.Quiet && len(classes) > 0 {
		log.Printf("Found %v classes in file %v", len(classes), filename)
	}

	edits := c.generateEdits(classes)
	if !c.opts.Quiet {
		log.Printf("%v classes need rewriting.", len(edits))
	}

	if c.opts.List {
		if len(edits) > 0 {
			fmt.Printf("%v differs.\n", filename)
		}
		return len(edits) != 0, nil
	}

	newBuf := buf
	if len(edits) > 0 {
		newBuf = c.rewriteClasses(buf, edits)
	}

	switch {
	case c.opts.Diff:
		if len(edits) > 0 {
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(buf),
				B:        difflib.SplitLines(newBuf),
				FromFile: "Original",
				ToFile:   "Stylized",
				Context:  3,
				Eol:      "\n",
			}
			result, err := difflib.GetUnifiedDiffString(diff)
			if err != nil {
				return true, err
			}
			fmt.Printf("%v\n", strings.Replace(result, "\t", " ", -1))
		}
	case c.opts.Write:
		if len(edits) > 0 {
			if err := ioutil.WriteFile(filename, []byte(newBuf), 0644); err != nil {
				return true, err
			}
			if !c.opts.Quiet {
				log.Printf("Successfully wrote %v bytes to file %v", len(newBuf), filename)
			}
		}
	default: // dump file to stdout
		fmt.Print(newBuf)
	}

	return len(edits) != 0, nil
}

func validateMemberOrdering(memberOrdering []string) bool {
	if len(memberOrdering) != len(defaultMemberOrdering) {
		log.Printf("flutterStylizer.memberOrdering must have %v values. Ignoring and using defaults.", len(defaultMemberOrdering))
		return false
	}

	lookup := map[string]bool{}
	for _, s := range defaultMemberOrdering {
		lookup[s] = true
	}

	seen := map[string]bool{}
	for _, el := range memberOrdering {
		if !lookup[el] {
			log.Printf("Unknown member %q in flutterStylizer.memberOrdering. Ignoring and using defaults.", el)
			return false
		}
		if seen[el] {
			log.Printf("Duplicate member %q in flutterStylizer.memberOrdering. Ignoring and using defaults.", el)
			return false
		}
		seen[el] = true
	}

	return true
}

var (
	matchClassRE = regexp.MustCompile(`^(?:abstract\s+)?class\s+(\S+).*$`)
)

func findOpenCurlyOffset(buf string, startOffset int) int {
	offset := strings.Index(buf[startOffset:], "{")
	if semiOffset := strings.Index(buf[startOffset:], ";"); offset < 0 || (semiOffset >= 0 && semiOffset < offset) {
		offset = semiOffset // this is valid: class D = Object with Function;
	}
	return startOffset + offset
}

func (c *Client) getClasses(editor *Editor, groupAndSortGetterMethods bool) ([]*Class, error) {
	var classes []*Class

	for i, line := range editor.lines {
		if (strings.HasPrefix(line.line, "void main(") || strings.HasPrefix(line.line, "main(")) && strings.HasSuffix(line.line, ") {") {
			// Find the end of main, marking sections as comments or strings to avoid them below.
			openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, line.startOffset)
			editor.findMatchingBracket(openCurlyOffset)
			continue
		}

		mm := matchClassRE.FindStringSubmatch(line.line)
		if len(mm) != 2 {
			continue
		}

		className := mm[1]
		classOffset := line.startOffset
		openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, classOffset)

		if editor.fullBuf[openCurlyOffset] == ';' {
			continue
		}

		if openCurlyOffset <= classOffset {
			return nil, fmt.Errorf(`expected "{" after "class" at offset %v`, classOffset)
		}

		if line.entityType != Unknown || line.isCommentOrString {
			editor.logf("\n\nSkipping new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v due to entityType/comment/string", className, classOffset, openCurlyOffset, line)
			continue
		}

		// This is a hack, but helps with files like `localizations_utils.dart` in
		// the Flutter distribution.
		if strings.ContainsAny(line.stripped, "$'") {
			editor.logf("\n\nSkipping new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v due to appearance of $ or '", className, classOffset, openCurlyOffset, line)
			continue
		}

		editor.logf("\n\nFound new class %q at classOffset=%v, openCurlyOffset=%v, line=%#v", className, classOffset, openCurlyOffset, line)
		closeCurlyOffset, err := editor.findMatchingBracket(openCurlyOffset)
		if err != nil {
			return nil, err
		}

		if closeCurlyOffset <= openCurlyOffset {
			return nil, fmt.Errorf(`expected "}" after "{" at offset %v`, openCurlyOffset)
		}
		editor.logf("\n\nFound end of class %q at closeCurlyOffset=%v", className, closeCurlyOffset)

		dartClass := NewClass(editor, className, i, openCurlyOffset, closeCurlyOffset, groupAndSortGetterMethods)
		dartClass.FindFeatures()
		classes = append(classes, dartClass)
	}

	return classes, nil
}

// logf logs the line if verbose is true.
func (c *Client) logf(fmtStr string, args ...interface{}) {
	if c.opts.Verbose {
		log.Printf(fmtStr, args...)
	}
}
