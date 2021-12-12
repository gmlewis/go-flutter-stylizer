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

func TestSeparatePrivateMethods_false(t *testing.T) {
	e, err := NewEditor(separatePrivateMethodsSource, false)
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	classes, err := e.GetClasses(false, false)
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 1; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}

	c := New(Options{})
	edits := c.generateEdits(classes)
	newBuf := c.rewriteClasses(separatePrivateMethodsSource, edits)

	if newBuf != separatePrivateMethodsSource {
		t.Errorf("rewriteClasses =\n%v\nwant:\n%v", newBuf, separatePrivateMethodsSource)
	}
}

func TestSeparatePrivateMethods_true(t *testing.T) {
	e, err := NewEditor(separatePrivateMethodsSource, false)
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	classes, err := e.GetClasses(false, true)
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 1; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}

	opts := Options{
		SeparatePrivateMethods: true,
		MemberOrdering: []string{
			"public-constructor",
			"named-constructors",
			"public-static-variables",
			"public-instance-variables",
			"public-override-variables",
			"private-static-variables",
			"private-instance-variables",
			"public-override-methods",
			"private-other-methods",
			"public-other-methods",
			"build-method",
		},
	}
	c := New(opts)
	edits := c.generateEdits(classes)
	newBuf := c.rewriteClasses(separatePrivateMethodsSource, edits)

	if newBuf != separatePrivateMethodsWant {
		t.Errorf("rewriteClasses =\n%v\nwant:\n%v", newBuf, separatePrivateMethodsWant)
	}
}

var separatePrivateMethodsSource = `
class MyClass {
  myMethod(){};

  _myMethod(){};
}
`

var separatePrivateMethodsWant = `
class MyClass {
  _myMethod(){};

  myMethod(){};
}
`
