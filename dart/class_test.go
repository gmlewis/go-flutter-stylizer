package dart

import (
	"testing"
)

func TestFindFeatures_linux_mac(t *testing.T) {
	bc, bcOCO, bcCCO := setupEditor(t, "class Class1 {", basicClasses)

	_, err := bc.findMatchingBracket(bcOCO)
	if err != nil {
		t.Fatalf("findMatchingBracket(%v) = %v, want nil", bcOCO, err)
	}
	uc := NewClass(bc, "Class1", bcOCO, bcCCO, false)

	want := []EntityType{
		Unknown,           // line #1: // This is a test file. After running this through `Flutter Stylizer`,
		Unknown,           // line #2: // the output should match the content of file `basic_classes_out.dart`.
		BlankLine,         // line #3:
		Unknown,           // line #4: import 'dart:math';
		BlankLine,         // line #5:
		Unknown,           // line #6: // Class1 is the first class.
		Unknown,           // line #7: class Class1 {
		SingleLineComment, // line #8:   // _pvi is a private instance variable.
		Unknown,           // line #9:   List<String> _pvi = ['one', 'two'];
		Unknown,           // line #10:   @override
		Unknown,           // line #11:   build() {}  // build method
		BlankLine,         // line #12:
		SingleLineComment, // line #13:   // This is a random single-line comment somewhere in the class.
		SingleLineComment, // line #14:
		SingleLineComment, // line #15:   // _spv is a static private variable.
		Unknown,           // line #16:   static final String _spv = 'spv';
		BlankLine,         // line #17:
		MultiLineComment,  // line #18:   /* This is a
		MultiLineComment,  // line #19:    * random multi-
		MultiLineComment,  // line #20:    * line comment
		MultiLineComment,  // line #21:    * somewhere in the middle
		MultiLineComment,  // line #22:    * of the class */
		SingleLineComment, // line #23:
		SingleLineComment, // line #24:   // _spvni is a static private variable with no initializer.
		Unknown,           // line #25:   static double _spvni;
		Unknown,           // line #26:   int _pvini;
		Unknown,           // line #27:   static int sv;
		Unknown,           // line #28:   int v;
		Unknown,           // line #29:   final double fv = 42.0;
		Unknown,           // line #30:   Class1();
		Unknown,           // line #31:   Class1.fromNum();
		Unknown,           // line #32:   var myfunc = (int n) => n;
		Unknown,           // line #33:   get vv => v; // getter
		Unknown,           // line #34:   @override
		Unknown,           // line #35:   toString() {
		Unknown,           // line #36:     print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');
		Unknown,           // line #37:     return '';
		Unknown,           // line #38:   }
		BlankLine,         // line #39:
		SingleLineComment, // line #40:   // "Here is 'where we add ${ text to "trip 'up' ''' the ${dart parser}.
		MultiLineComment,  // line #41:   /*
		MultiLineComment,  // line #42:     '''
		MultiLineComment,  // line #43:     """
		MultiLineComment,  // line #44:     //
		MultiLineComment,  // line #45:   */
		Unknown,           // line #46:   const a = """
		Unknown,           // line #47:    '${b}
		Unknown,           // line #48:    '''
		Unknown,           // line #49:   """
		Unknown,           // line #50:   const b = '''
		Unknown,           // line #51:     {  (  ))) """ {{{} ))))
		Unknown,           // line #52:   '''
		Unknown,           // line #53:   const c = { '{{{((... """ ${'((('}' }
		Unknown,           // line #54: }
		BlankLine,         // line #55:
	}

	t.Run("*nix file", func(t *testing.T) {
		uc.FindFeatures()

		if len(bc.lines) != len(want) {
			t.Fatalf("lines mismatch: got = %v, want %v", len(bc.lines), len(want))
		}

		for i, line := range bc.lines {
			// fmt.Printf("%v, // line #%v: %v\n", line.entityType, i+1, line.line)
			if line.entityType != want[i] {
				t.Errorf("line #%v: entityType = %v, want %v", i+1, line.entityType, want[i])
			}
		}
	})
}

func TestFindFeatures_windoze(t *testing.T) {
	wz, wzOCO, wzCCO := setupEditor(t, "class Class1 {", bcWindoze)

	_, err := wz.findMatchingBracket(wzOCO)
	if err != nil {
		t.Fatalf("findMatchingBracket(%v) = %v, want nil", wzOCO, err)
	}
	uc := NewClass(wz, "Class1", wzOCO, wzCCO, false)

	want := []EntityType{
		Unknown,           // line #1: // This is a test file. After running this through `Flutter Stylizer`,
		Unknown,           // line #2: // the output should match the content of file `basic_classes_out.dart`.
		BlankLine,         // line #3:
		Unknown,           // line #4: import 'dart:math';
		BlankLine,         // line #5:
		Unknown,           // line #6: // Class1 is the first class.
		Unknown,           // line #7: class Class1 {
		SingleLineComment, // line #8:   // _pvi is a private instance variable.
		Unknown,           // line #9:   List<String> _pvi = ['one', 'two'];
		Unknown,           // line #10:   @override
		Unknown,           // line #11:   build() {}  // build method
		BlankLine,         // line #12:
		SingleLineComment, // line #13:   // This is a random single-line comment somewhere in the class.
		SingleLineComment, // line #14:
		SingleLineComment, // line #15:   // _spv is a static private variable.
		Unknown,           // line #16:   static final String _spv = 'spv';
		BlankLine,         // line #17:
		MultiLineComment,  // line #18:   /* This is a
		MultiLineComment,  // line #19:    * random multi-
		MultiLineComment,  // line #20:    * line comment
		MultiLineComment,  // line #21:    * somewhere in the middle
		MultiLineComment,  // line #22:    * of the class */
		SingleLineComment, // line #23:
		SingleLineComment, // line #24:   // _spvni is a static private variable with no initializer.
		Unknown,           // line #25:   static double _spvni;
		Unknown,           // line #26:   int _pvini;
		Unknown,           // line #27:   static int sv;
		Unknown,           // line #28:   int v;
		Unknown,           // line #29:   final double fv = 42.0;
		Unknown,           // line #30:   Class1();
		Unknown,           // line #31:   Class1.fromNum();
		Unknown,           // line #32:   var myfunc = (int n) => n;
		Unknown,           // line #33:   get vv => v; // getter
		Unknown,           // line #34:   @override
		Unknown,           // line #35:   toString() {
		Unknown,           // line #36:     print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');
		Unknown,           // line #37:     return '';
		Unknown,           // line #38:   }
		BlankLine,         // line #39:
		SingleLineComment, // line #40:   // "Here is 'where we add ${ text to "trip 'up' ''' the ${dart parser}.
		MultiLineComment,  // line #41:   /*
		MultiLineComment,  // line #42:     '''
		MultiLineComment,  // line #43:     """
		MultiLineComment,  // line #44:     //
		MultiLineComment,  // line #45:   */
		Unknown,           // line #46:   const a = """
		Unknown,           // line #47:    '${b}
		Unknown,           // line #48:    '''
		Unknown,           // line #49:   """
		Unknown,           // line #50:   const b = '''
		Unknown,           // line #51:     {  (  ))) """ {{{} ))))
		Unknown,           // line #52:   '''
		Unknown,           // line #53:   const c = { '{{{((... """ ${'((('}' }
		Unknown,           // line #54: }
		BlankLine,         // line #55:
	}

	t.Run("windoze file", func(t *testing.T) {
		uc.FindFeatures()

		if len(wz.lines) != len(want) {
			t.Fatalf("lines mismatch: got = %v, want %v", len(wz.lines), len(want))
		}

		for i, line := range wz.lines {
			// fmt.Printf("%v, // line #%v: %v\n", line.entityType, i+1, line.line)
			if line.entityType != want[i] {
				t.Errorf("line #%v: entityType = %v, want %v", i+1, line.entityType, want[i])
			}
		}
	})
}
