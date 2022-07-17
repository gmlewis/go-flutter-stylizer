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
	"path/filepath"
	"strings"

	"github.com/pmezard/go-difflib/difflib"
)

// Options represents the configuration options for the Dart processor.
type Options struct {
	Debug   bool
	Diff    bool
	List    bool
	Quiet   bool
	Verbose bool
	Write   bool

	GroupAndSortGetterMethods bool
	GroupAndSortVariableTypes bool

	MemberOrdering          []string
	ProcessEnumsLikeClasses bool
	SortClassesWithinFile   bool
	SortOtherMethods        bool

	hasPublicStaticProperties    bool
	hasPublicInstanceProperties  bool
	hasPublicOverrideProperties  bool
	hasPrivateStaticProperties   bool
	hasPrivateInstanceProperties bool
	hasPrivateOtherMethods       bool
	hasOperators                 bool
}

// Client represents a Dart processor.
type Client struct {
	editor *Editor
	opts   Options
}

// New returns a new Dart processor.
func New(e *Editor, opts Options) *Client {
	return &Client{editor: e, opts: opts.validate()}
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

	e, err := NewEditor(buf, c.opts)
	if err != nil {
		return false, fmt.Errorf("NewEditor: %w", err)
	}
	c.editor = e

	classes, err := e.GetClasses()
	if err != nil {
		return false, err
	}
	if !c.opts.Quiet && c.opts.Verbose && len(classes) > 0 {
		log.Printf("Found %v classes in file %v", len(classes), filename)
	}

	edits := c.generateEdits(classes)
	if !c.opts.Quiet && c.opts.Verbose {
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
			fromFile := filepath.Join("original", filename)
			toFile := filepath.Join("stylized", filename)
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(buf),
				B:        difflib.SplitLines(newBuf),
				FromFile: fromFile,
				ToFile:   toFile,
				Context:  3,
				Eol:      "\n",
			}
			result, err := difflib.GetUnifiedDiffString(diff)
			if err != nil {
				return true, err
			}
			fmt.Printf("\ndiff %v %v\n%v\n", fromFile, toFile, strings.Replace(result, "\t", " ", -1))
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

// logf logs the line if debug is true.
func (c *Client) logf(fmtStr string, args ...interface{}) {
	if c.opts.Debug {
		log.Printf(fmtStr, args...)
	}
}
