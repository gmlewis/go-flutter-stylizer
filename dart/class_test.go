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
	"testing"
)

func TestClassesAreFound(t *testing.T) {
	const source = `// test.dart
class myClass extends Widget {

}`

	e := NewEditor(source)
	c := &Client{}
	got, err := c.GetClasses(e, false)
	if err != nil {
		t.Fatal(err)
	}

	if want := 1; len(got) != want {
		t.Errorf("GetClasses = %v, want %v", got, want)
	}

	if want := 3; len(got[0].lines) != want {
		t.Errorf("GetClasses lines= %v, want %v", got, want)
	}
}

func TestNamedConstructorsAreKeptIntact(t *testing.T) {
	const source = `class AnimationController extends Animation<double>
with AnimationEagerListenerMixin, AnimationLocalListenersMixin, AnimationLocalStatusListenersMixin {
	AnimationController.unbounded({
	double value = 0.0,
	this.duration,
	this.debugLabel,
	@required TickerProvider vsync,
	this.animationBehavior = AnimationBehavior.preserve,
	}) : assert(value != null),
			assert(vsync != null),
			lowerBound = double.negativeInfinity,
			upperBound = double.infinity,
			_direction = _AnimationDirection.forward {
	_ticker = vsync.createTicker(_tick);
	_internalSetValue(value);
	}
}`

	e := NewEditor(source)
	c := &Client{}
	got, err := c.GetClasses(e, false)
	if err != nil {
		t.Fatal(err)
	}

	if want := 1; len(got) != want {
		t.Errorf("GetClasses = %v, want %v", got, want)
	}

	want := []EntityType{
		Unknown,          // line #1: {
		NamedConstructor, // line #2: 	AnimationController.unbounded({
		NamedConstructor, // line #3: 	double value = 0.0,
		NamedConstructor, // line #4: 	this.duration,
		NamedConstructor, // line #5: 	this.debugLabel,
		NamedConstructor, // line #6: 	@required TickerProvider vsync,
		NamedConstructor, // line #7: 	this.animationBehavior = AnimationBehavior.preserve,
		NamedConstructor, // line #8: 	}) : assert(value != null),
		NamedConstructor, // line #9: 			assert(vsync != null),
		NamedConstructor, // line #10: 			lowerBound = double.negativeInfinity,
		NamedConstructor, // line #11: 			upperBound = double.infinity,
		NamedConstructor, // line #12: 			_direction = _AnimationDirection.forward {
		NamedConstructor, // line #13: 	_ticker = vsync.createTicker(_tick);
		NamedConstructor, // line #14: 	_internalSetValue(value);
		NamedConstructor, // line #15: 	}
		BlankLine,        // line #16:
	}

	if len(got[0].lines) != len(want) {
		t.Errorf("getClasses lines= %v, want %v", got, want)
	}

	for i := 0; i < len(got[0].lines); i++ {
		line := got[0].lines[i]
		if line.entityType != want[i] {
			t.Errorf("line #%v: got entityType %v, want %v: %v", i+1, line.entityType, want[i], line.line)
		}
	}
}

func TestFindFeatures_linux_mac(t *testing.T) {
	bc, bcLineOffset, bcOCO, bcCCO := setupEditor(t, "class Class1 {", basicClasses)

	_, err := bc.findMatchingBracket(bcOCO)
	if err != nil {
		t.Fatalf("findMatchingBracket(%v) = %v, want nil", bcOCO, err)
	}
	uc := NewClass(bc, "Class1", bcOCO, bcCCO, false)

	want := []EntityType{
		Unknown,                 // line #7: class Class1 {
		PrivateInstanceVariable, // line #8:   // _pvi is a private instance variable.
		PrivateInstanceVariable, // line #9:   List<String> _pvi = ['one', 'two'];
		BuildMethod,             // line #10:   @override
		BuildMethod,             // line #11:   build() {} // build method
		BlankLine,               // line #12:
		StaticPrivateVariable,   // line #13:   // This is a random single-line comment somewhere in the class.
		StaticPrivateVariable,   // line #14:
		StaticPrivateVariable,   // line #15:   // _spv is a static private variable.
		StaticPrivateVariable,   // line #16:   static final String _spv = 'spv';
		BlankLine,               // line #17:
		StaticPrivateVariable,   // line #18:   /* This is a
		StaticPrivateVariable,   // line #19:    * random multi-
		StaticPrivateVariable,   // line #20:    * line comment
		StaticPrivateVariable,   // line #21:    * somewhere in the middle
		StaticPrivateVariable,   // line #22:    * of the class */
		StaticPrivateVariable,   // line #23:
		StaticPrivateVariable,   // line #24:   // _spvni is a static private variable with no initializer.
		StaticPrivateVariable,   // line #25:   static double _spvni = 0;
		PrivateInstanceVariable, // line #26:   int _pvini = 1;
		StaticVariable,          // line #27:   static int sv = 0;
		InstanceVariable,        // line #28:   int v = 2;
		InstanceVariable,        // line #29:   final double fv = 42.0;
		MainConstructor,         // line #30:   Class1();
		NamedConstructor,        // line #31:   Class1.fromNum();
		OtherMethod,             // line #32:   var myfunc = (int n) => n;
		OtherMethod,             // line #33:   get vv => v; // getter
		OverrideMethod,          // line #34:   @override
		OverrideMethod,          // line #35:   toString() {
		OverrideMethod,          // line #36:     print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');
		OverrideMethod,          // line #37:     return '';
		OverrideMethod,          // line #38:   }
		BlankLine,               // line #39:
		StaticVariable,          // line #40:   // "Here is 'where we add ${ text to "trip 'up' ''' the ${dart parser}.
		StaticVariable,          // line #41:   /*
		StaticVariable,          // line #42:     '''
		StaticVariable,          // line #43:     """
		StaticVariable,          // line #44:     //
		StaticVariable,          // line #45:   */
		StaticVariable,          // line #46:   static const a = """;
		StaticVariable,          // line #47:    '${b};
		StaticVariable,          // line #48:    ''' ;
		StaticVariable,          // line #49:   """;
		StaticVariable,          // line #50:   static const b = ''';
		StaticVariable,          // line #51:     {  (  ))) """ {{{} ))));
		StaticVariable,          // line #52:   ''';
		StaticVariable,          // line #53:   static const c = { '{{{((... """ ${'((('};'};
		BlankLine,               // line #54: }
	}

	t.Run("*nix file", func(t *testing.T) {
		if err := uc.FindFeatures(); err != nil {
			t.Fatalf("FindFeatures: %v", err)
		}

		if len(uc.lines) != len(want) {
			t.Fatalf("lines mismatch: got = %v, want %v", len(uc.lines), len(want))
		}

		for i, line := range uc.lines {
			// fmt.Printf("%v, // line #%v: %v\n", line.entityType, i+1, line.line)
			if line.entityType != want[i] {
				t.Errorf("line #%v: entityType = %v, want %v", bcLineOffset+i+1, line.entityType, want[i])
			}
		}
	})
}

func TestFindFeatures_windoze(t *testing.T) {
	wz, wzLineOffset, wzOCO, wzCCO := setupEditor(t, "class Class1 {", bcWindoze)

	_, err := wz.findMatchingBracket(wzOCO)
	if err != nil {
		t.Fatalf("findMatchingBracket(%v) = %v, want nil", wzOCO, err)
	}
	wc := NewClass(wz, "Class1", wzOCO, wzCCO, false)

	want := []EntityType{
		Unknown,                 // line #7: class Class1 {
		PrivateInstanceVariable, // line #8:   // _pvi is a private instance variable.
		PrivateInstanceVariable, // line #9:   List<String> _pvi = ['one', 'two'];
		BuildMethod,             // line #10:   @override
		BuildMethod,             // line #11:   build() {} // build method
		BlankLine,               // line #12:
		StaticPrivateVariable,   // line #13:   // This is a random single-line comment somewhere in the class.
		StaticPrivateVariable,   // line #14:
		StaticPrivateVariable,   // line #15:   // _spv is a static private variable.
		StaticPrivateVariable,   // line #16:   static final String _spv = 'spv';
		BlankLine,               // line #17:
		StaticPrivateVariable,   // line #18:   /* This is a
		StaticPrivateVariable,   // line #19:    * random multi-
		StaticPrivateVariable,   // line #20:    * line comment
		StaticPrivateVariable,   // line #21:    * somewhere in the middle
		StaticPrivateVariable,   // line #22:    * of the class */
		StaticPrivateVariable,   // line #23:
		StaticPrivateVariable,   // line #24:   // _spvni is a static private variable with no initializer.
		StaticPrivateVariable,   // line #25:   static double _spvni = 0;
		PrivateInstanceVariable, // line #26:   int _pvini = 1;
		StaticVariable,          // line #27:   static int sv = 0;
		InstanceVariable,        // line #28:   int v = 2;
		InstanceVariable,        // line #29:   final double fv = 42.0;
		MainConstructor,         // line #30:   Class1();
		NamedConstructor,        // line #31:   Class1.fromNum();
		OtherMethod,             // line #32:   var myfunc = (int n) => n;
		OtherMethod,             // line #33:   get vv => v; // getter
		OverrideMethod,          // line #34:   @override
		OverrideMethod,          // line #35:   toString() {
		OverrideMethod,          // line #36:     print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');
		OverrideMethod,          // line #37:     return '';
		OverrideMethod,          // line #38:   }
		BlankLine,               // line #39:
		StaticVariable,          // line #40:   // "Here is 'where we add ${ text to "trip 'up' ''' the ${dart parser}.
		StaticVariable,          // line #41:   /*
		StaticVariable,          // line #42:     '''
		StaticVariable,          // line #43:     """
		StaticVariable,          // line #44:     //
		StaticVariable,          // line #45:   */
		StaticVariable,          // line #46:   static const a = """;
		StaticVariable,          // line #47:    '${b};
		StaticVariable,          // line #48:    ''' ;
		StaticVariable,          // line #49:   """;
		StaticVariable,          // line #50:   static const b = ''';
		StaticVariable,          // line #51:     {  (  ))) """ {{{} ))));
		StaticVariable,          // line #52:   ''';
		StaticVariable,          // line #53:   static const c = { '{{{((... """ ${'((('};'};
		BlankLine,               // line #54: }
	}

	t.Run("windoze file", func(t *testing.T) {
		if err := wc.FindFeatures(); err != nil {
			t.Fatalf("FindFeatures: %v", err)
		}

		if len(wc.lines) != len(want) {
			t.Fatalf("lines mismatch: got = %v, want %v", len(wc.lines), len(want))
		}

		for i, line := range wc.lines {
			// fmt.Printf("%v, // line #%v: %v\n", line.entityType, i+1, line.line)
			if line.entityType != want[i] {
				t.Errorf("line #%v: entityType = %v, want %v", wzLineOffset+i+1, line.entityType, want[i])
			}
		}
	})
}
