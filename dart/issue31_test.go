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

	want := [][]EntityType{
		{
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
		},
		{
			Unknown,                 // line #15: {
			MainConstructor,         // line #16:   User({
			MainConstructor,         // line #17:     required this.firstNameFinal,
			MainConstructor,         // line #18:     required this.idFinal,
			MainConstructor,         // line #19:     required this.lastNameFinal,
			MainConstructor,         // line #20:     required this.middleNameFinal,
			MainConstructor,         // line #21:     required this.phoneNumberFinal,
			MainConstructor,         // line #22:     required this.usernameFinal,
			MainConstructor,         // line #23:     required this.firstNameRegular,
			MainConstructor,         // line #24:     required this.idRegular,
			MainConstructor,         // line #25:     required this.lastNameRegular,
			MainConstructor,         // line #26:     required this.usernameRegular,
			MainConstructor,         // line #27:   });
			BlankLine,               // line #28:
			InstanceVariable,        // line #29:   final String firstNameFinal;
			InstanceVariable,        // line #30:   final int idFinal;
			InstanceVariable,        // line #31:   final String lastNameFinal;
			InstanceVariable,        // line #32:   final String? middleNameFinal;
			InstanceVariable,        // line #33:   final String? phoneNumberFinal;
			InstanceVariable,        // line #34:   final String usernameFinal;
			BlankLine,               // line #35:
			InstanceVariable,        // line #36:   String firstNameRegular;
			InstanceVariable,        // line #37:   int idRegular;
			InstanceVariable,        // line #38:   String lastNameRegular;
			InstanceVariable,        // line #39:   String usernameRegular;
			BlankLine,               // line #40:
			InstanceVariable,        // line #41:   int? ageOptional;
			InstanceVariable,        // line #42:   String? birthdateOptional;
			InstanceVariable,        // line #43:   String? emailOptional;
			InstanceVariable,        // line #44:   String? middleNameRegular;
			InstanceVariable,        // line #45:   String? phoneNumberRegular;
			BlankLine,               // line #46:
			PrivateInstanceVariable, // line #47:   int? _agePrivate;
			PrivateInstanceVariable, // line #48:   String? _birthdatePrivate;
			PrivateInstanceVariable, // line #49:   String? _emailPrivate;
			PrivateInstanceVariable, // line #50:   final String _firstNamePrivate = 'Secret';
			PrivateInstanceVariable, // line #51:   final int _idPrivate = 0;
			PrivateInstanceVariable, // line #52:   final String _lastNamePrivate = 'Secret';
			PrivateInstanceVariable, // line #53:   String? _middleNamePrivate;
			PrivateInstanceVariable, // line #54:   final String _phoneNumberPrivate = 'Secret';
			PrivateInstanceVariable, // line #55:   final String _usernamePrivate = 'Secret';
			BlankLine,               // line #56:
		},
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
			"public-instance-variables",
			"public-override-variables",
			"private-static-variables",
			"private-instance-variables",
			"public-other-methods",
			"public-override-methods",
			"private-other-methods",
			"build-method",
		},
		SortOtherMethods: true,
	}

	want := [][]EntityType{
		{
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
		},
		{
			Unknown,                 // line #15: {
			MainConstructor,         // line #16:   User({
			MainConstructor,         // line #17:     required this.firstNameFinal,
			MainConstructor,         // line #18:     required this.idFinal,
			MainConstructor,         // line #19:     required this.lastNameFinal,
			MainConstructor,         // line #20:     required this.middleNameFinal,
			MainConstructor,         // line #21:     required this.phoneNumberFinal,
			MainConstructor,         // line #22:     required this.usernameFinal,
			MainConstructor,         // line #23:     required this.firstNameRegular,
			MainConstructor,         // line #24:     required this.idRegular,
			MainConstructor,         // line #25:     required this.lastNameRegular,
			MainConstructor,         // line #26:     required this.usernameRegular,
			MainConstructor,         // line #27:   });
			BlankLine,               // line #28:
			InstanceVariable,        // line #29:   final String firstNameFinal;
			InstanceVariable,        // line #30:   final int idFinal;
			InstanceVariable,        // line #31:   final String lastNameFinal;
			InstanceVariable,        // line #32:   final String? middleNameFinal;
			InstanceVariable,        // line #33:   final String? phoneNumberFinal;
			InstanceVariable,        // line #34:   final String usernameFinal;
			BlankLine,               // line #35:
			InstanceVariable,        // line #36:   String firstNameRegular;
			InstanceVariable,        // line #37:   int idRegular;
			InstanceVariable,        // line #38:   String lastNameRegular;
			InstanceVariable,        // line #39:   String usernameRegular;
			BlankLine,               // line #40:
			InstanceVariable,        // line #41:   int? ageOptional;
			InstanceVariable,        // line #42:   String? birthdateOptional;
			InstanceVariable,        // line #43:   String? emailOptional;
			InstanceVariable,        // line #44:   String? middleNameRegular;
			InstanceVariable,        // line #45:   String? phoneNumberRegular;
			BlankLine,               // line #46:
			PrivateInstanceVariable, // line #47:   int? _agePrivate;
			PrivateInstanceVariable, // line #48:   String? _birthdatePrivate;
			PrivateInstanceVariable, // line #49:   String? _emailPrivate;
			PrivateInstanceVariable, // line #50:   final String _firstNamePrivate = 'Secret';
			PrivateInstanceVariable, // line #51:   final int _idPrivate = 0;
			PrivateInstanceVariable, // line #52:   final String _lastNamePrivate = 'Secret';
			PrivateInstanceVariable, // line #53:   String? _middleNamePrivate;
			PrivateInstanceVariable, // line #54:   final String _phoneNumberPrivate = 'Secret';
			PrivateInstanceVariable, // line #55:   final String _usernamePrivate = 'Secret';
			BlankLine,               // line #56:
		},
	}

	runFullStylizer(t, opts, source, wantSource, want)
}
