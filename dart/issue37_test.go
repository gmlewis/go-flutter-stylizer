/*
Copyright © 2022 Glenn M. Lewis

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
	_ "embed"
	"testing"
)

//go:embed testfiles/issue37.dart.txt
var issue37_dart_txt string

//go:embed testfiles/issue37_want.txt
var issue37_want_txt string

//go:embed testfiles/issue37_want2.txt
var issue37_want2_txt string

func TestIssue37_case1(t *testing.T) {
	source := issue37_dart_txt
	wantSource := issue37_want_txt

	opts := &Options{
		MemberOrdering: []string{
			"public-constructor",
			"named-constructors",
			"public-static-variables",
			"private-static-variables",
			"public-instance-variables",
			"public-override-variables",
			"private-instance-variables",
			"public-override-methods",
			"public-other-methods",
			"private-other-methods",
			"build-method",
		},
	}

	want := [][]EntityType{
		{
			Unknown,          // line #1: {
			InstanceVariable, // line #2:   double left, top, width, height;
			BlankLine,        // line #3:
			MainConstructor,  // line #4:   Rectangle(this.left, this.top, this.width, this.height);
			BlankLine,        // line #5:
			OtherMethod,      // line #6:   Rectangle operator +(Rectangle v) => Rectangle(left, top, width + v.width, height + v.height);
			BlankLine,        // line #7:
			OtherMethod,      // line #8:   double get right => left + width;
			OtherMethod,      // line #9:   set right(double value) => left = value - width;
			OtherMethod,      // line #10:   double get bottom => top + height;
			OtherMethod,      // line #11:   set bottom(double value) => top = value - height;
			BlankLine,        // line #12:
		},
	}

	runFullStylizer(t, opts, source, wantSource, want)
}

func TestIssue37_case2(t *testing.T) {
	source := issue37_dart_txt
	wantSource := issue37_want2_txt

	opts := &Options{
		GroupAndSortVariableTypes: true,
		MemberOrdering: []string{
			"public-constructor",
			"named-constructors",
			"public-static-variables",
			"public-static-properties",
			"public-instance-variables",
			"public-instance-properties",
			"public-override-variables",
			"public-override-properties",
			"private-static-variables",
			"private-static-properties",
			"private-instance-variables",
			"private-instance-properties",
			"public-override-methods",
			"public-other-methods",
			"private-other-methods",
			"operators",
			"build-method",
		},
		SortOtherMethods: true,
	}

	want := [][]EntityType{
		{
			Unknown,          // line #1: {
			InstanceVariable, // line #2:   double left, top, width, height;
			BlankLine,        // line #3:
			MainConstructor,  // line #4:   Rectangle(this.left, this.top, this.width, this.height);
			BlankLine,        // line #5:
			OtherMethod,      // line #6:   Rectangle operator +(Rectangle v) => Rectangle(left, top, width + v.width, height + v.height);
			BlankLine,        // line #7:
			OtherMethod,      // line #8:   double get right => left + width;
			OtherMethod,      // line #9:   set right(double value) => left = value - width;
			OtherMethod,      // line #10:   double get bottom => top + height;
			OtherMethod,      // line #11:   set bottom(double value) => top = value - height;
			BlankLine,        // line #12:
		},
	}

	runFullStylizer(t, opts, source, wantSource, want)
}
