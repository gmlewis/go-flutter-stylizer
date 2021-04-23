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
	"reflect"
	"testing"
)

//go:embed testfiles/flow_analysis.dart.txt
var flowAnalysis string

func TestGetClasses(t *testing.T) {
	e := NewEditor(flowAnalysis)

	c := &Client{}
	classes, err := c.getClasses(e, false)
	if err != nil {
		t.Fatalf("getClasses: %v", err)
	}

	if got, want := len(classes), 22; got != want {
		t.Errorf("got %v classes", len(classes))
	}

	for i, class := range flowAnalysisWantClasses {
		if !reflect.DeepEqual(classes[i], class) {
			t.Errorf("class #%v (%q) differs; got:\n%#v\nwant:%#v", i+1, classes[i].className, classes[i], class)
		}
	}
}

var flowAnalysisWantClasses = []*Class{
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
}
