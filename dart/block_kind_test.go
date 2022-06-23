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
	_ "embed"
	"testing"
)

//go:embed testfiles/block_kind.dart.txt
var block_kind_dart_txt string

//go:embed testfiles/block_kind_want.txt
var block_kind_want_txt string

func TestBlockKindExample(t *testing.T) {
	source := block_kind_dart_txt
	wantSource := block_kind_want_txt

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,          // line #5: {
		InstanceVariable, // line #6:   final String name;
		BlankLine,        // line #7:
		InstanceVariable, // line #8:   final bool useNameForMissingBlock;
		BlankLine,        // line #9:
		NamedConstructor, // line #10:   const BlockKind._(this.name, this.useNameForMissingBlock);
		BlankLine,        // line #11:
		OtherMethod,      // line #12:   /// Returns the name to use for this block if it is missing in
		OtherMethod,      // line #13:   /// [templateExpectedClassOrMixinBody].
		OtherMethod,      // line #14:   ///
		OtherMethod,      // line #15:   /// If `null` the generic [templateExpectedButGot] is used instead.
		OtherMethod,      // line #16:   String get missingBlockName => useNameForMissingBlock ? name : null;
		BlankLine,        // line #17:
		OtherMethod,      // line #18:   String toString() => 'BlockKind($name)';
		BlankLine,        // line #19:
		StaticVariable,   // line #20:   static const BlockKind catchClause =
		StaticVariable,   // line #21:       const BlockKind._('catch clause', /* useNameForMissingBlock = */ true);
		StaticVariable,   // line #22:   static const BlockKind classDeclaration = const BlockKind._(
		StaticVariable,   // line #23:       'class declaration', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #24:   static const BlockKind enumDeclaration = const BlockKind._(
		StaticVariable,   // line #25:       'enum declaration', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #26:   static const BlockKind extensionDeclaration = const BlockKind._(
		StaticVariable,   // line #27:       'extension declaration', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #28:   static const BlockKind finallyClause =
		StaticVariable,   // line #29:       const BlockKind._('finally clause', /* useNameForMissingBlock = */ true);
		StaticVariable,   // line #30:   static const BlockKind functionBody =
		StaticVariable,   // line #31:       const BlockKind._('function body', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #32:   static const BlockKind invalid =
		StaticVariable,   // line #33:       const BlockKind._('invalid', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #34:   static const BlockKind mixinDeclaration = const BlockKind._(
		StaticVariable,   // line #35:       'mixin declaration', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #36:   static const BlockKind statement =
		StaticVariable,   // line #37:       const BlockKind._('statement', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #38:   static const BlockKind switchStatement = const BlockKind._(
		StaticVariable,   // line #39:       'switch statement', /* useNameForMissingBlock = */ false);
		StaticVariable,   // line #40:   static const BlockKind tryStatement =
		StaticVariable,   // line #41:       const BlockKind._('try statement', /* useNameForMissingBlock = */ true);
		BlankLine,        // line #42:
	}

	runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestBlockKindExample_StaysTheSame(t *testing.T) {
	wantSource := block_kind_want_txt

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	classes := runFullStylizer(t, opts, wantSource, wantSource, nil)

	if len(classes) != 1 {
		t.Fatalf("got %v classes; want 1", len(classes))
	}
}
