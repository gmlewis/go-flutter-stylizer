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
	"fmt"
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

	runFullStylizer(t, opts, source, wantSource, want)
}

func TestBlockKindExample_StaysTheSame(t *testing.T) {
	wantSource := block_kind_want_txt

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	// want := []EntityType{
	// 	Unknown,          // line #5: {
	// 	NamedConstructor, // line #6:   const BlockKind._(this.name, this.useNameForMissingBlock);
	// 	BlankLine,        // line #7:
	// 	StaticVariable,   // line #8:   static const BlockKind catchClause =
	// 	StaticVariable,   // line #9:       const BlockKind._('catch clause', /* useNameForMissingBlock = */ true);
	// 	BlankLine,        // line #10:
	// 	StaticVariable,   // line #11:   static const BlockKind classDeclaration = const BlockKind._(
	// 	StaticVariable,   // line #12:       'class declaration', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #13:
	// 	StaticVariable,   // line #14:   static const BlockKind enumDeclaration = const BlockKind._(
	// 	StaticVariable,   // line #15:       'enum declaration', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #16:
	// 	StaticVariable,   // line #17:   static const BlockKind extensionDeclaration = const BlockKind._(
	// 	StaticVariable,   // line #18:       'extension declaration', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #19:
	// 	StaticVariable,   // line #20:   static const BlockKind finallyClause =
	// 	StaticVariable,   // line #21:       const BlockKind._('finally clause', /* useNameForMissingBlock = */ true);
	// 	BlankLine,        // line #22:
	// 	StaticVariable,   // line #23:   static const BlockKind functionBody =
	// 	StaticVariable,   // line #24:       const BlockKind._('function body', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #25:
	// 	StaticVariable,   // line #26:   static const BlockKind invalid =
	// 	StaticVariable,   // line #27:       const BlockKind._('invalid', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #28:
	// 	StaticVariable,   // line #29:   static const BlockKind mixinDeclaration = const BlockKind._(
	// 	StaticVariable,   // line #30:       'mixin declaration', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #31:
	// 	StaticVariable,   // line #32:   static const BlockKind statement =
	// 	StaticVariable,   // line #33:       const BlockKind._('statement', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #34:
	// 	StaticVariable,   // line #35:   static const BlockKind switchStatement = const BlockKind._(
	// 	StaticVariable,   // line #36:       'switch statement', /* useNameForMissingBlock = */ false);
	// 	BlankLine,        // line #37:
	// 	StaticVariable,   // line #38:   static const BlockKind tryStatement =
	// 	StaticVariable,   // line #39:       const BlockKind._('try statement', /* useNameForMissingBlock = */ true);
	// 	BlankLine,        // line #40:
	// 	InstanceVariable, // line #41:   final String name;
	// 	InstanceVariable, // line #42:   final bool useNameForMissingBlock;
	// 	BlankLine,        // line #43:
	// 	OtherMethod,      // line #44:   /// Returns the name to use for this block if it is missing in
	// 	OtherMethod,      // line #45:   /// [templateExpectedClassOrMixinBody].
	// 	OtherMethod,      // line #46:   ///
	// 	OtherMethod,      // line #47:   /// If `null` the generic [templateExpectedButGot] is used instead.
	// 	OtherMethod,      // line #48:   String get missingBlockName => useNameForMissingBlock ? name : null;
	// 	BlankLine,        // line #49:
	// 	OtherMethod,      // line #50:   String toString() => 'BlockKind($name)';
	// 	BlankLine,        // line #51:
	// }

	classes := runFullStylizer(t, opts, wantSource, wantSource, nil)

	if len(classes) != 1 {
		t.Fatalf("got %v classes; want 1", len(classes))
	}

	// got := classes[0]

	// compareEntity(t, "theConstructor", got.theConstructor, &Entity{})
	// compareEntities(t, "namedConstructors", got.namedConstructors, []*Entity{})
	// compareEntities(t, "staticVariables", got.staticVariables, []*Entity{})
	// compareEntities(t, "instanceVariables", got.instanceVariables, []*Entity{})
	// compareEntities(t, "overrideVariables", got.overrideVariables, []*Entity{})
	// compareEntities(t, "staticPrivateVariables", got.staticPrivateVariables, []*Entity{})
	// compareEntities(t, "privateVariables", got.privateVariables, []*Entity{})
	// compareEntities(t, "overrideMethods", got.overrideMethods, []*Entity{})
	// compareEntities(t, "otherMethods", got.otherMethods, []*Entity{})
	// compareEntity(t, "buildMethod", got.buildMethod, &Entity{})
	// compareEntities(t, "getterMethods", got.getterMethods, []*Entity{})
}

func compareEntities(t *testing.T, name string, got, want []*Entity) {
	t.Helper()

	if got == nil {
		return
	}

	if len(got) != len(want) {
		t.Errorf("entities %q: got %v, want %v", name, len(got), len(want))

		for i := 0; i < len(got); i++ {
			var e *Entity
			if i < len(want) {
				e = want[i]
			}
			compareEntity(t, fmt.Sprintf("%v[%v]", name, i+1), got[i], e)
		}
	}
}

func compareEntity(t *testing.T, name string, got, want *Entity) {
	t.Helper()

	if got == nil {
		return
	}

	if want == nil {
		want = &Entity{}
	}

	if len(got.lines) != len(want.lines) {
		t.Errorf("entity %q: got %v lines, want %v", name, len(got.lines), len(want.lines))

		for i := 0; i < len(got.lines); i++ {
			line := got.lines[i]
			if i < len(want.lines) && line.line != want.lines[i].line {
				t.Errorf("entity %q: line #%v: got:\n%v\nwant:\n%v", name, i+1, line.line, want.lines[i].line)
			} else {
				t.Errorf("entity %q: line #%v: got:\n%v", name, i+1, line.line)
			}
		}
	}
}
