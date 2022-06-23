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

//go:embed testfiles/issue26_case1.dart.txt
var issue26_case1_dart_txt string

//go:embed testfiles/issue26_case1.want.txt
var issue26_case1_want_txt string

func TestIssue26_Case1(t *testing.T) {
	source := issue26_case1_dart_txt
	wantSource := issue26_case1_want_txt

	opts := &Options{
		GroupAndSortGetterMethods: true,
		SortOtherMethods:          true,
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
		Unknown,                 // line #5: {
		MainConstructor,         // line #6:   SnackbarService();
		BlankLine,               // line #7:
		PrivateInstanceVariable, // line #8:   final Map<dynamic, SnackBar Function(SnackBarConfigBase)> _snackbars = {};
		OtherMethod,             // line #9:   SnackBar Function(SnackBarConfigBase) getSnackbar(dynamic key) {
		OtherMethod,             // line #10:     return _snackbars[key]!;
		OtherMethod,             // line #11:   }
		BlankLine,               // line #12:
		InstanceVariable,        // line #13:   final GlobalKey<ScaffoldMessengerState> snackbarKey = GlobalKey<ScaffoldMessengerState>();
		BlankLine,               // line #14:
		OtherMethod,             // line #15:   bool containsKey(dynamic key) {
		OtherMethod,             // line #16:     return _snackbars.containsKey(key);
		OtherMethod,             // line #17:   }
		BlankLine,               // line #18:
		OtherMethod,             // line #19:   void registerSnackbar(
		OtherMethod,             // line #20:       {required dynamic key, required SnackBar Function(SnackBarConfigBase) snackbarBuilder}) {
		OtherMethod,             // line #21:     _snackbars[key] = snackbarBuilder;
		OtherMethod,             // line #22:   }
		BlankLine,               // line #23:
	}

	runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

//go:embed testfiles/issue26_case2.dart.txt
var issue26_case2_dart_txt string

//go:embed testfiles/issue26_case2.want.txt
var issue26_case2_want_txt string

func TestIssue26_Case2(t *testing.T) {
	source := issue26_case2_dart_txt
	wantSource := issue26_case2_want_txt

	opts := &Options{
		GroupAndSortGetterMethods: true,
		SortOtherMethods:          true,
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
		PrivateInstanceVariable, // line #2:   final Map<dynamic, Widget Function(DialogRequest)> _dialogs = {};
		PrivateInstanceVariable, // line #3:   late Function(DialogRequest) _dialogHandler;
		BlankLine,               // line #4:
		PrivateInstanceVariable, // line #5:   late Completer<DialogReponse>? _dialogCompleter;
		PrivateInstanceVariable, // line #6:   late final NavigatorService _navigatorService;
		BlankLine,               // line #7:
	}

	runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}
