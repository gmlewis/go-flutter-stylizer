package dart

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed basic_classes.dart.txt
var basicClasses string

func setupEditor(t *testing.T, searchFor, buf string) (*DartEditor, int, int) {
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

	d := NewEditor(buf)
	return d, openCurlyOffset, closeCurlyOffset
}

func TestFindMatchingBracket(t *testing.T) {
	bc, bcOCO, bcCCO := setupEditor(t, "class Class1 {", basicClasses)

	tests := map[string]struct {
		editor     *DartEditor
		openOffset int
		want       int
	}{
		"top-level": {
			editor:     bc,
			openOffset: bcOCO,
			want:       bcCCO,
		},
		"build()": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "build()") + 5,
			want:       strings.Index(basicClasses, "build()") + 6,
		},
		"build() {}": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "build() {}") + 8,
			want:       strings.Index(basicClasses, "build() {}") + 9,
		},
		"Class1();": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "Class1();") + 6,
			want:       strings.Index(basicClasses, "Class1();") + 7,
		},
		"Class1.fromNum();": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "Class1.fromNum();") + 14,
			want:       strings.Index(basicClasses, "Class1.fromNum();") + 15,
		},
		"var myfunc = (int n) => n;": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "var myfunc = (int n) => n;") + 13,
			want:       strings.Index(basicClasses, "var myfunc = (int n) => n;") + 19,
		},
		"toString()": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "toString()") + 8,
			want:       strings.Index(basicClasses, "toString()") + 9,
		},
		"toString() {": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "toString() {") + 11,
			want:       strings.Index(basicClasses, "toString() {") + 85,
		},
		"print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 5,
			want:       strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 50,
		},
		"${sqrt(2)}": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "${sqrt(2)}") + 1,
			want:       strings.Index(basicClasses, "${sqrt(2)}") + 9,
		},
		"sqrt(2)": {
			editor:     bc,
			openOffset: strings.Index(basicClasses, "sqrt(2)") + 4,
			want:       strings.Index(basicClasses, "sqrt(2)") + 6,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.editor.findMatchingBracket(tt.openOffset)
			if got != tt.want {
				t.Errorf("findMatchingBracket(%v) = %v, want %v", tt.openOffset, got, tt.want)
			}
		})
	}
}

func TestFindLineIndexAtOffset(t *testing.T) {
	bc, bcOCO, _ := setupEditor(t, "class Class1 {", basicClasses)

	tests := map[string]struct {
		editor        *DartEditor
		openOffset    int
		wantLineIndex int
		wantRelOffset int
	}{
		"top-level": {
			editor:        bc,
			openOffset:    bcOCO,
			wantLineIndex: 6,
			wantRelOffset: 13,
		},
		"build()": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "build()") + 5,
			wantLineIndex: 10,
			wantRelOffset: 7,
		},
		"build() {}": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "build() {}") + 8,
			wantLineIndex: 10,
			wantRelOffset: 10,
		},
		"Class1();": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "Class1();") + 6,
			wantLineIndex: 29,
			wantRelOffset: 8,
		},
		"Class1.fromNum();": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "Class1.fromNum();") + 14,
			wantLineIndex: 30,
			wantRelOffset: 16,
		},
		"var myfunc = (int n) => n;": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "var myfunc = (int n) => n;") + 13,
			wantLineIndex: 31,
			wantRelOffset: 15,
		},
		"toString()": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "toString()") + 8,
			wantLineIndex: 34,
			wantRelOffset: 10,
		},
		"toString() {": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "toString() {") + 11,
			wantLineIndex: 34,
			wantRelOffset: 13,
		},
		"print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 5,
			wantLineIndex: 35,
			wantRelOffset: 9,
		},
		"${sqrt(2)}": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "${sqrt(2)}") + 1,
			wantLineIndex: 35,
			wantRelOffset: 44,
		},
		"sqrt(2)": {
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "sqrt(2)") + 4,
			wantLineIndex: 35,
			wantRelOffset: 49,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			gotLineIndex, gotRelOffset := tt.editor.findLineIndexAtOffset(tt.openOffset)
			if gotLineIndex != tt.wantLineIndex {
				t.Errorf("findMatchingBracket(%v) lineIndex = %v, want %v", tt.openOffset, gotLineIndex, tt.wantLineIndex)
			}
			if gotRelOffset != tt.wantRelOffset {
				t.Errorf("findMatchingBracket(%v) relOffset = %v, want %v", tt.openOffset, gotRelOffset, tt.wantRelOffset)
			}
		})
	}
}
