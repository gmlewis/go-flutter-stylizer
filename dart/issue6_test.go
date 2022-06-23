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

//go:embed testfiles/issue6.dart.txt
var issue6_dart_txt string

//go:embed testfiles/issue6_want.txt
var issue6_want_txt string

func TestIssue6_case1(t *testing.T) {
	source := issue6_dart_txt

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
		Unknown,                 // line #1: {
		StaticVariable,          // line #2:   static const kReleaseMode = true;
		PrivateInstanceVariable, // line #3:   final Map<String, dynamic> _initialValues = kReleaseMode
		PrivateInstanceVariable, // line #4:       ? {}
		PrivateInstanceVariable, // line #5:       : {
		PrivateInstanceVariable, // line #6:           'hubType': 'test',
		PrivateInstanceVariable, // line #7:           'ht.localIP': '192.168.1.1',
		PrivateInstanceVariable, // line #8:           'ht.makerAppID': '2233',
		PrivateInstanceVariable, // line #9:           'ht.accessToken': '7de3-yyyyyy-xxxxxxxxxxx',
		PrivateInstanceVariable, // line #10:         };
		BlankLine,               // line #11:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
}

func TestIssue6_case2(t *testing.T) {
	source := issue6_dart_txt
	wantSource := issue6_want_txt

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
		Unknown,                 // line #1: {
		StaticVariable,          // line #2:   static const kReleaseMode = true;
		PrivateInstanceVariable, // line #3:   final Map<String, dynamic> _initialValues = kReleaseMode
		PrivateInstanceVariable, // line #4:       ? {}
		PrivateInstanceVariable, // line #5:       : {
		PrivateInstanceVariable, // line #6:           'hubType': 'test',
		PrivateInstanceVariable, // line #7:           'ht.localIP': '192.168.1.1',
		PrivateInstanceVariable, // line #8:           'ht.makerAppID': '2233',
		PrivateInstanceVariable, // line #9:           'ht.accessToken': '7de3-yyyyyy-xxxxxxxxxxx',
		PrivateInstanceVariable, // line #10:         };
		BlankLine,               // line #11:
	}

	runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}
