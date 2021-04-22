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
	if len(opts.MemberOrdering) == 0 {
		opts.MemberOrdering = defaultMemberOrdering
	}

	return &Client{opts: opts}
}

// StylizeFile sylizes a single Dart file using the provided options.
func (c *Client) StylizeFile(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	e := NewEditor(string(buf))
	e.Verbose = c.opts.Verbose
	classes, err := c.getClasses(e, c.opts.GroupAndSortGetterMethods)
	if err != nil {
		return err
	}
	if !c.opts.Quiet {
		log.Printf("Found %v classes in file %v", len(classes), filename)
	}

	return nil
}

var (
	matchClassRE = regexp.MustCompile(`^(?:abstract\s+)?class\s+(\S+).*$`)
)

func findOpenCurlyOffset(buf string, startOffset int) int {
	offset := strings.Index(buf[startOffset:], "{")
	return startOffset + offset
}

func (c *Client) getClasses(editor *Editor, groupAndSortGetterMethods bool) ([]*Class, error) {
	var classes []*Class

	for i, line := range editor.lines {
		mm := matchClassRE.FindStringSubmatch(line.stripped)
		if len(mm) != 2 {
			continue
		}
		className := mm[1]
		classOffset := line.startOffset + line.strippedOffset
		openCurlyOffset := findOpenCurlyOffset(editor.fullBuf, classOffset)
		if openCurlyOffset <= classOffset {
			return nil, fmt.Errorf(`expected "{" after "class" at offset %v`, classOffset)
		}

		closeCurlyOffset, err := editor.findMatchingBracket(openCurlyOffset)
		if err != nil {
			return nil, err
		}

		if closeCurlyOffset <= openCurlyOffset {
			return nil, fmt.Errorf(`expected "}" after "{" at offset %v`, openCurlyOffset)
		}

		dartClass := NewClass(editor, className, i, openCurlyOffset, closeCurlyOffset, groupAndSortGetterMethods)
		dartClass.FindFeatures()
		classes = append(classes, dartClass)
	}

	return classes, nil
}
