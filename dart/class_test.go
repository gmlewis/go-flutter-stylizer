package dart

import (
	"testing"
)

func TestFindFeatures_linux_mac(t *testing.T) {
	bc, bcLineOffset, bcOCO, bcCCO := setupEditor(t, "class Class1 {", basicClasses)

	_, err := bc.findMatchingBracket(bcOCO)
	if err != nil {
		t.Fatalf("findMatchingBracket(%v) = %v, want nil", bcOCO, err)
	}
	uc := NewClass(bc, "Class1", bcLineOffset, bcOCO, bcCCO, false)

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
	wc := NewClass(wz, "Class1", wzLineOffset, wzOCO, wzCCO, false)

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
