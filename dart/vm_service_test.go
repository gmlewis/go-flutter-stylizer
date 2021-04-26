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

//go:embed testfiles/vm_service.dart.txt
var vm_service_dart_txt string

//go:embed testfiles/vm_service_want.txt
var vm_service_want_txt string

func TestVMService_GetClasses(t *testing.T) {
	e, err := NewEditor(vm_service_dart_txt, false)
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	c := &Client{}
	classes, err := c.GetClasses(e, false)
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 96; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}
}

func TestVMService_Source(t *testing.T) {
	source := vm_service_dart_txt
	wantSource := vm_service_want_txt

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	runFullStylizer(t, opts, source, wantSource, nil)
	runFullStylizer(t, opts, wantSource, wantSource, nil)
}

func TestVMService_Class96(t *testing.T) {
	source := vm_service_dart_txt[218716:]
	wantSource := vm_service_want_txt[218643:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,          // line #1: {
		OtherMethod,      // line #2:   static VM parse(Map<String, dynamic> json) =>
		OtherMethod,      // line #3:       json == null ? null : VM._fromJson(json);
		BlankLine,        // line #4:
		InstanceVariable, // line #5:   /// A name identifying this vm. Not guaranteed to be unique.
		InstanceVariable, // line #6:   String name;
		BlankLine,        // line #7:
		InstanceVariable, // line #8:   /// Word length on target architecture (e.g. 32, 64).
		InstanceVariable, // line #9:   int architectureBits;
		BlankLine,        // line #10:
		InstanceVariable, // line #11:   /// The CPU we are actually running on.
		InstanceVariable, // line #12:   String hostCPU;
		BlankLine,        // line #13:
		InstanceVariable, // line #14:   /// The operating system we are running on.
		InstanceVariable, // line #15:   String operatingSystem;
		BlankLine,        // line #16:
		InstanceVariable, // line #17:   /// The CPU we are generating code for.
		InstanceVariable, // line #18:   String targetCPU;
		BlankLine,        // line #19:
		InstanceVariable, // line #20:   /// The Dart VM version string.
		InstanceVariable, // line #21:   String version;
		BlankLine,        // line #22:
		InstanceVariable, // line #23:   /// The process id for the VM.
		InstanceVariable, // line #24:   int pid;
		BlankLine,        // line #25:
		InstanceVariable, // line #26:   /// The time that the VM started in milliseconds since the epoch.
		InstanceVariable, // line #27:   ///
		InstanceVariable, // line #28:   /// Suitable to pass to DateTime.fromMillisecondsSinceEpoch.
		InstanceVariable, // line #29:   int startTime;
		BlankLine,        // line #30:
		InstanceVariable, // line #31:   /// A list of isolates running in the VM.
		InstanceVariable, // line #32:   List<IsolateRef> isolates;
		BlankLine,        // line #33:
		InstanceVariable, // line #34:   /// A list of isolate groups running in the VM.
		InstanceVariable, // line #35:   List<IsolateGroupRef> isolateGroups;
		BlankLine,        // line #36:
		InstanceVariable, // line #37:   /// A list of system isolates running in the VM.
		InstanceVariable, // line #38:   List<IsolateRef> systemIsolates;
		BlankLine,        // line #39:
		InstanceVariable, // line #40:   /// A list of isolate groups which contain system isolates running in the VM.
		InstanceVariable, // line #41:   List<IsolateGroupRef> systemIsolateGroups;
		BlankLine,        // line #42:
		MainConstructor,  // line #43:   VM({
		MainConstructor,  // line #44:     @required this.name,
		MainConstructor,  // line #45:     @required this.architectureBits,
		MainConstructor,  // line #46:     @required this.hostCPU,
		MainConstructor,  // line #47:     @required this.operatingSystem,
		MainConstructor,  // line #48:     @required this.targetCPU,
		MainConstructor,  // line #49:     @required this.version,
		MainConstructor,  // line #50:     @required this.pid,
		MainConstructor,  // line #51:     @required this.startTime,
		MainConstructor,  // line #52:     @required this.isolates,
		MainConstructor,  // line #53:     @required this.isolateGroups,
		MainConstructor,  // line #54:     @required this.systemIsolates,
		MainConstructor,  // line #55:     @required this.systemIsolateGroups,
		MainConstructor,  // line #56:   });
		BlankLine,        // line #57:
		NamedConstructor, // line #58:   VM._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
		NamedConstructor, // line #59:     name = json['name'];
		NamedConstructor, // line #60:     architectureBits = json['architectureBits'];
		NamedConstructor, // line #61:     hostCPU = json['hostCPU'];
		NamedConstructor, // line #62:     operatingSystem = json['operatingSystem'];
		NamedConstructor, // line #63:     targetCPU = json['targetCPU'];
		NamedConstructor, // line #64:     version = json['version'];
		NamedConstructor, // line #65:     pid = json['pid'];
		NamedConstructor, // line #66:     startTime = json['startTime'];
		NamedConstructor, // line #67:     isolates = List<IsolateRef>.from(
		NamedConstructor, // line #68:         createServiceObject(json['isolates'], const ['IsolateRef']) ?? []);
		NamedConstructor, // line #69:     isolateGroups = List<IsolateGroupRef>.from(
		NamedConstructor, // line #70:         createServiceObject(json['isolateGroups'], const ['IsolateGroupRef']) ??
		NamedConstructor, // line #71:             []);
		NamedConstructor, // line #72:     systemIsolates = List<IsolateRef>.from(
		NamedConstructor, // line #73:         createServiceObject(json['systemIsolates'], const ['IsolateRef']) ??
		NamedConstructor, // line #74:             []);
		NamedConstructor, // line #75:     systemIsolateGroups = List<IsolateGroupRef>.from(createServiceObject(
		NamedConstructor, // line #76:             json['systemIsolateGroups'], const ['IsolateGroupRef']) ??
		NamedConstructor, // line #77:         []);
		NamedConstructor, // line #78:   }
		BlankLine,        // line #79:
		OverrideMethod,   // line #80:   @override
		OverrideMethod,   // line #81:   Map<String, dynamic> toJson() {
		OverrideMethod,   // line #82:     var json = <String, dynamic>{};
		OverrideMethod,   // line #83:     json['type'] = 'VM';
		OverrideMethod,   // line #84:     json.addAll({
		OverrideMethod,   // line #85:       'name': name,
		OverrideMethod,   // line #86:       'architectureBits': architectureBits,
		OverrideMethod,   // line #87:       'hostCPU': hostCPU,
		OverrideMethod,   // line #88:       'operatingSystem': operatingSystem,
		OverrideMethod,   // line #89:       'targetCPU': targetCPU,
		OverrideMethod,   // line #90:       'version': version,
		OverrideMethod,   // line #91:       'pid': pid,
		OverrideMethod,   // line #92:       'startTime': startTime,
		OverrideMethod,   // line #93:       'isolates': isolates.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #94:       'isolateGroups': isolateGroups.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #95:       'systemIsolates': systemIsolates.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #96:       'systemIsolateGroups':
		OverrideMethod,   // line #97:           systemIsolateGroups.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #98:     });
		OverrideMethod,   // line #99:     return json;
		OverrideMethod,   // line #100:   }
		BlankLine,        // line #101:
		OtherMethod,      // line #102:   String toString() => '[VM]';
		BlankLine,        // line #103:
	}

	runFullStylizer(t, opts, source, wantSource, want)

	wantAfter := []EntityType{
		Unknown,          // line #1: {
		MainConstructor,  // line #43:   VM({
		MainConstructor,  // line #44:     @required this.name,
		MainConstructor,  // line #45:     @required this.architectureBits,
		MainConstructor,  // line #46:     @required this.hostCPU,
		MainConstructor,  // line #47:     @required this.operatingSystem,
		MainConstructor,  // line #48:     @required this.targetCPU,
		MainConstructor,  // line #49:     @required this.version,
		MainConstructor,  // line #50:     @required this.pid,
		MainConstructor,  // line #51:     @required this.startTime,
		MainConstructor,  // line #52:     @required this.isolates,
		MainConstructor,  // line #53:     @required this.isolateGroups,
		MainConstructor,  // line #54:     @required this.systemIsolates,
		MainConstructor,  // line #55:     @required this.systemIsolateGroups,
		MainConstructor,  // line #56:   });
		BlankLine,        // line #57:
		NamedConstructor, // line #58:   VM._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
		NamedConstructor, // line #59:     name = json['name'];
		NamedConstructor, // line #60:     architectureBits = json['architectureBits'];
		NamedConstructor, // line #61:     hostCPU = json['hostCPU'];
		NamedConstructor, // line #62:     operatingSystem = json['operatingSystem'];
		NamedConstructor, // line #63:     targetCPU = json['targetCPU'];
		NamedConstructor, // line #64:     version = json['version'];
		NamedConstructor, // line #65:     pid = json['pid'];
		NamedConstructor, // line #66:     startTime = json['startTime'];
		NamedConstructor, // line #67:     isolates = List<IsolateRef>.from(
		NamedConstructor, // line #68:         createServiceObject(json['isolates'], const ['IsolateRef']) ?? []);
		NamedConstructor, // line #69:     isolateGroups = List<IsolateGroupRef>.from(
		NamedConstructor, // line #70:         createServiceObject(json['isolateGroups'], const ['IsolateGroupRef']) ??
		NamedConstructor, // line #71:             []);
		NamedConstructor, // line #72:     systemIsolates = List<IsolateRef>.from(
		NamedConstructor, // line #73:         createServiceObject(json['systemIsolates'], const ['IsolateRef']) ??
		NamedConstructor, // line #74:             []);
		NamedConstructor, // line #75:     systemIsolateGroups = List<IsolateGroupRef>.from(createServiceObject(
		NamedConstructor, // line #76:             json['systemIsolateGroups'], const ['IsolateGroupRef']) ??
		NamedConstructor, // line #77:         []);
		NamedConstructor, // line #78:   }
		BlankLine,        // line #79:
		InstanceVariable, // line #8:   /// Word length on target architecture (e.g. 32, 64).
		InstanceVariable, // line #9:   int architectureBits;
		BlankLine,        // line #10:
		InstanceVariable, // line #11:   /// The CPU we are actually running on.
		InstanceVariable, // line #12:   String hostCPU;
		BlankLine,        // line #13:
		InstanceVariable, // line #34:   /// A list of isolate groups running in the VM.
		InstanceVariable, // line #35:   List<IsolateGroupRef> isolateGroups;
		BlankLine,        // line #36:
		InstanceVariable, // line #31:   /// A list of isolates running in the VM.
		InstanceVariable, // line #32:   List<IsolateRef> isolates;
		BlankLine,        // line #33:
		InstanceVariable, // line #5:   /// A name identifying this vm. Not guaranteed to be unique.
		InstanceVariable, // line #6:   String name;
		BlankLine,        // line #7:
		InstanceVariable, // line #14:   /// The operating system we are running on.
		InstanceVariable, // line #15:   String operatingSystem;
		BlankLine,        // line #16:
		InstanceVariable, // line #23:   /// The process id for the VM.
		InstanceVariable, // line #24:   int pid;
		BlankLine,        // line #25:
		InstanceVariable, // line #26:   /// The time that the VM started in milliseconds since the epoch.
		InstanceVariable, // line #27:   ///
		InstanceVariable, // line #28:   /// Suitable to pass to DateTime.fromMillisecondsSinceEpoch.
		InstanceVariable, // line #29:   int startTime;
		BlankLine,        // line #30:
		InstanceVariable, // line #37:   /// A list of system isolates running in the VM.
		InstanceVariable, // line #38:   List<IsolateRef> systemIsolates;
		BlankLine,        // line #39:
		InstanceVariable, // line #40:   /// A list of isolate groups which contain system isolates running in the VM.
		InstanceVariable, // line #41:   List<IsolateGroupRef> systemIsolateGroups;
		BlankLine,        // line #42:
		InstanceVariable, // line #17:   /// The CPU we are generating code for.
		InstanceVariable, // line #18:   String targetCPU;
		BlankLine,        // line #19:
		InstanceVariable, // line #20:   /// The Dart VM version string.
		InstanceVariable, // line #21:   String version;
		BlankLine,        // line #22:
		OverrideMethod,   // line #80:   @override
		OverrideMethod,   // line #81:   Map<String, dynamic> toJson() {
		OverrideMethod,   // line #82:     var json = <String, dynamic>{};
		OverrideMethod,   // line #83:     json['type'] = 'VM';
		OverrideMethod,   // line #84:     json.addAll({
		OverrideMethod,   // line #85:       'name': name,
		OverrideMethod,   // line #86:       'architectureBits': architectureBits,
		OverrideMethod,   // line #87:       'hostCPU': hostCPU,
		OverrideMethod,   // line #88:       'operatingSystem': operatingSystem,
		OverrideMethod,   // line #89:       'targetCPU': targetCPU,
		OverrideMethod,   // line #90:       'version': version,
		OverrideMethod,   // line #91:       'pid': pid,
		OverrideMethod,   // line #92:       'startTime': startTime,
		OverrideMethod,   // line #93:       'isolates': isolates.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #94:       'isolateGroups': isolateGroups.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #95:       'systemIsolates': systemIsolates.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #96:       'systemIsolateGroups':
		OverrideMethod,   // line #97:           systemIsolateGroups.map((f) => f.toJson()).toList(),
		OverrideMethod,   // line #98:     });
		OverrideMethod,   // line #99:     return json;
		OverrideMethod,   // line #100:   }
		BlankLine,        // line #101:
		OtherMethod,      // line #2:   static VM parse(Map<String, dynamic> json) =>
		OtherMethod,      // line #3:       json == null ? null : VM._fromJson(json);
		BlankLine,        // line #4:
		OtherMethod,      // line #102:   String toString() => '[VM]';
		BlankLine,        // line #103:
	}

	runParsePhase(t, opts, wantSource, wantAfter)
	runFullStylizer(t, opts, wantSource, wantSource, nil)
}
