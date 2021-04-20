package dart

import (
	"strings"
	"testing"
)

func TestFindMatchingBracket(t *testing.T) {
	bc, bcOCO, bcCCO := setupClass(t, "class Class1 {", basicClasses)

	tests := map[string]struct {
		class      *Class
		openOffset int
		want       int
	}{
		"top-level": {
			class:      bc,
			openOffset: bcOCO,
			want:       bcCCO,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.class.findMatchingBracket(tt.openOffset)
			if got != tt.want {
				t.Errorf("findMatchingBracket(%v) = %v, want %v", tt.openOffset, got, tt.want)
			}
		})
	}
}

func setupClass(t *testing.T, searchFor, buf string) (*Class, int, int) {
	t.Helper()
	classOffset := strings.Index(buf, searchFor)
	openCurlyOffset := classOffset + len(searchFor) - 1
	closeCurlyOffset := len(buf) - 2

	if got, want := buf[classOffset:openCurlyOffset+1], searchFor; got != want {
		t.Errorf("offset error: buf[%v:%v] = %q, want %q", classOffset, openCurlyOffset+1, got, want)
	}
	if got, want := buf[closeCurlyOffset:closeCurlyOffset+1], "}"; got != want {
		t.Errorf("offset error: buf[%v:%v] = %q, want %q", closeCurlyOffset, closeCurlyOffset+1, got, want)
	}

	c := NewClass("Class1", classOffset, openCurlyOffset, closeCurlyOffset, false)
	c.FindFeatures(buf)
	return c, openCurlyOffset, closeCurlyOffset
}

var basicClasses = `// This is a test file. After running this through "Flutter Stylizer",
// the output should match the content of file "basic_classes_out.dart".

import 'dart:math';

// Class1 is the first class.
class Class1 {
  // _pvi is a private instance variable.
  List<String> _pvi = ['one', 'two'];
  @override
  build() {}  // build method

  // This is a random single-line comment somewhere in the class.

  // _spv is a static private variable.
  static final String _spv = 'spv';

  /* This is a
   * random multi-
   * line comment
   * somewhere in the middle
   * of the class */

  // _spvni is a static private variable with no initializer.
  static double _spvni;
  int _pvini;
  static int sv;
  int v;
  final double fv = 42.0;
  Class1();
  Class1.fromNum();
  var myfunc = (int n) => n;
  get vv => v; // getter
  @override
  toString() {
    print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');
    return '';
  }
}
`
