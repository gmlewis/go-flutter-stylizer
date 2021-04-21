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

	tests := []struct {
		name       string
		editor     *DartEditor
		verbose    bool
		openOffset int
		want       int
	}{
		{
			name:       "top-level",
			editor:     bc,
			openOffset: bcOCO,
			want:       bcCCO,
		},
		{
			name:       "build()",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "build()") + 5,
			want:       strings.Index(basicClasses, "build()") + 6,
		},
		{
			name:       "build() {}",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "build() {}") + 8,
			want:       strings.Index(basicClasses, "build() {}") + 9,
		},
		{
			name:       "Class1();",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "Class1();") + 6,
			want:       strings.Index(basicClasses, "Class1();") + 7,
		},
		{
			name:       "Class1.fromNum();",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "Class1.fromNum();") + 14,
			want:       strings.Index(basicClasses, "Class1.fromNum();") + 15,
		},
		{
			name:       "var myfunc = (int n) => n;",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "var myfunc = (int n) => n;") + 13,
			want:       strings.Index(basicClasses, "var myfunc = (int n) => n;") + 19,
		},
		{
			name:       "toString()",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "toString()") + 8,
			want:       strings.Index(basicClasses, "toString()") + 9,
		},
		{
			name:       "toString() {",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "toString() {") + 11,
			want:       strings.Index(basicClasses, "toString() {") + 87,
		},
		{
			name:       "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 5,
			want:       strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 50,
		},
		{
			name:       "${sqrt(2)}",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "${sqrt(2)}") + 1,
			want:       strings.Index(basicClasses, "${sqrt(2)}") + 9,
		},
		{
			name:       "sqrt(2)",
			editor:     bc,
			openOffset: strings.Index(basicClasses, "sqrt(2)") + 4,
			want:       strings.Index(basicClasses, "sqrt(2)") + 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.editor.Verbose = tt.verbose
			got, err := tt.editor.findMatchingBracket(tt.openOffset)
			if err != nil {
				t.Fatalf("findMatchingBracket(%v) = %v, want nil", tt.openOffset, err)
			}

			if got != tt.want {
				t.Errorf("findMatchingBracket(%v) = %v, want %v", tt.openOffset, got, tt.want)
			}
		})
	}
}

func TestFindLineIndexAtOffset(t *testing.T) {
	bc, bcOCO, _ := setupEditor(t, "class Class1 {", basicClasses)

	tests := []struct {
		name          string
		editor        *DartEditor
		openOffset    int
		wantLineIndex int
		wantRelOffset int
	}{
		{
			name:          "top-level",
			editor:        bc,
			openOffset:    bcOCO,
			wantLineIndex: 6,
			wantRelOffset: 13,
		},
		{
			name:          "build()",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "build()") + 5,
			wantLineIndex: 10,
			wantRelOffset: 5,
		},
		{
			name:          "build() {}",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "build() {}") + 8,
			wantLineIndex: 10,
			wantRelOffset: 8,
		},
		{
			name:          "Class1();",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "Class1();") + 6,
			wantLineIndex: 29,
			wantRelOffset: 6,
		},
		{
			name:          "Class1.fromNum();",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "Class1.fromNum();") + 14,
			wantLineIndex: 30,
			wantRelOffset: 14,
		},
		{
			name:          "var myfunc = (int n) => n;",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "var myfunc = (int n) => n;") + 13,
			wantLineIndex: 31,
			wantRelOffset: 13,
		},
		{
			name:          "toString()",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "toString()") + 8,
			wantLineIndex: 34,
			wantRelOffset: 8,
		},
		{
			name:          "toString() {",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "toString() {") + 11,
			wantLineIndex: 34,
			wantRelOffset: 11,
		},
		{
			name:          "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "print('$_pvi, $_spv, $_spvni, $_pvini, ${sqrt(2)}');") + 5,
			wantLineIndex: 35,
			wantRelOffset: 5,
		},
		{
			name:          "${sqrt(2)}",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "${sqrt(2)}") + 1,
			wantLineIndex: 35,
			wantRelOffset: 40,
		},
		{
			name:          "sqrt(2)",
			editor:        bc,
			openOffset:    strings.Index(basicClasses, "sqrt(2)") + 4,
			wantLineIndex: 35,
			wantRelOffset: 45,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLineIndex, gotRelOffset := tt.editor.findLineIndexAtOffset(tt.openOffset)
			if gotLineIndex != tt.wantLineIndex {
				t.Errorf("findLineIndexAtOffset(%v) lineIndex = %v, want %v", tt.openOffset, gotLineIndex, tt.wantLineIndex)
			}
			if gotRelOffset != tt.wantRelOffset {
				t.Errorf("findLineIndexAtOffset(%v) relOffset = %v, want %v", tt.openOffset, gotRelOffset, tt.wantRelOffset)
			}
		})
	}
}
