package dart

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed basic_classes.dart.txt
var basicClasses string

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
		"build()": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "build()") + 5,
			want:       strings.Index(basicClasses, "build()") + 6,
		},
		"build() {}": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "build() {}") + 8,
			want:       strings.Index(basicClasses, "build() {}") + 9,
		},
		"Class1();": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "Class1();") + 6,
			want:       strings.Index(basicClasses, "Class1();") + 7,
		},
		"Class1.fromNum();": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "Class1.fromNum();") + 14,
			want:       strings.Index(basicClasses, "Class1.fromNum();") + 15,
		},
		"var myfunc = (int n) => n;": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "var myfunc = (int n) => n;") + 13,
			want:       strings.Index(basicClasses, "var myfunc = (int n) => n;") + 19,
		},
		"toString()": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "toString()") + 8,
			want:       strings.Index(basicClasses, "toString()") + 9,
		},
		"toString() {": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "toString() {") + 11,
			want:       strings.Index(basicClasses, "toString() {") + 85,
		},
		"print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 5,
			want:       strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 50,
		},
		"${sqrt(2)}": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "${sqrt(2)}") + 1,
			want:       strings.Index(basicClasses, "${sqrt(2)}") + 9,
		},
		"sqrt(2)": {
			class:      bc,
			openOffset: strings.Index(basicClasses, "sqrt(2)") + 4,
			want:       strings.Index(basicClasses, "sqrt(2)") + 6,
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
