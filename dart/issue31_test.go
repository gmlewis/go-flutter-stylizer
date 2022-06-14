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

//go:embed testfiles/issue31.dart.txt
var issue31_dart_txt string

//go:embed testfiles/issue31_want.txt
var issue31_want_txt string

func TestIssue31_case1(t *testing.T) {
	source := issue31_dart_txt

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

	want := []EntityType{
		Unknown,          // line #1: {
		MainConstructor,  // line #2:   SnackbarService();
		BlankLine,        // line #3:
		InstanceVariable, // line #4:   final varA1 = 'varA1';
		InstanceVariable, // line #5:   String? varA2;
		InstanceVariable, // line #6:   String varA3 = 'varA3';
		InstanceVariable, // line #7:   final varB1 = 'varB1';
		InstanceVariable, // line #8:   String? varB2;
		InstanceVariable, // line #9:   String varB3 = 'varB3';
		InstanceVariable, // line #10:   final varC1 = 'varC1';
		InstanceVariable, // line #11:   String? varC2;
		InstanceVariable, // line #12:   String varC3 = 'varC3';
		BlankLine,        // line #13:
	}

	runParsePhase(t, opts, source, want)
}

func TestIssue31_case2(t *testing.T) {
	source := issue31_dart_txt
	wantSource := issue31_want_txt

	opts := &Options{
		GroupAndSortVariableTypes: true,
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

	want := []EntityType{
		Unknown,          // line #1: {
		MainConstructor,  // line #2:   SnackbarService();
		BlankLine,        // line #3:
		InstanceVariable, // line #4:   final varA1 = 'varA1';
		InstanceVariable, // line #5:   String? varA2;
		InstanceVariable, // line #6:   String varA3 = 'varA3';
		InstanceVariable, // line #7:   final varB1 = 'varB1';
		InstanceVariable, // line #8:   String? varB2;
		InstanceVariable, // line #9:   String varB3 = 'varB3';
		InstanceVariable, // line #10:   final varC1 = 'varC1';
		InstanceVariable, // line #11:   String? varC2;
		InstanceVariable, // line #12:   String varC3 = 'varC3';
		BlankLine,        // line #13:
	}

	runFullStylizer(t, opts, source, wantSource, want)
}
