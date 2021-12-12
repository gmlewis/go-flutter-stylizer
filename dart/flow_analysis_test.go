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
	"testing"
)

//go:embed testfiles/flow_analysis.dart.txt
var flowAnalysis string

func TestFlowAnalysis_GetClasses(t *testing.T) {
	e, err := NewEditor(flowAnalysis, false)
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	classes, err := e.GetClasses(false, false)
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 22; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}
}
