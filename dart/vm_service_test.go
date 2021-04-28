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
	"strings"
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

	classes, err := e.GetClasses(false)
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
	source := vm_service_dart_txt[strings.Index(vm_service_dart_txt, "class VM "):]
	wantSource := vm_service_want_txt[strings.Index(vm_service_want_txt, "class VM "):]

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

func getClass(src, from, to string) string {
	start := strings.Index(src, from)
	end := strings.Index(src, to)
	return src[start:end]
}

func TestVMService_RPCErrorClass(t *testing.T) {
	source := getClass(vm_service_dart_txt, "class RPCError ", "class SentinelException ")
	wantSource := getClass(vm_service_want_txt, "class RPCError ", "class SentinelException ")

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,          // line #1: {
		StaticVariable,   // line #2:   /// Application specific error codes.
		StaticVariable,   // line #3:   static const int kServerError = -32000;
		BlankLine,        // line #4:
		StaticVariable,   // line #5:   /// The JSON sent is not a valid Request object.
		StaticVariable,   // line #6:   static const int kInvalidRequest = -32600;
		BlankLine,        // line #7:
		StaticVariable,   // line #8:   /// The method does not exist or is not available.
		StaticVariable,   // line #9:   static const int kMethodNotFound = -32601;
		BlankLine,        // line #10:
		StaticVariable,   // line #11:   /// Invalid method parameter(s), such as a mismatched type.
		StaticVariable,   // line #12:   static const int kInvalidParams = -32602;
		BlankLine,        // line #13:
		StaticVariable,   // line #14:   /// Internal JSON-RPC error.
		StaticVariable,   // line #15:   static const int kInternalError = -32603;
		BlankLine,        // line #16:
		OtherMethod,      // line #17:   static RPCError parse(String callingMethod, dynamic json) {
		OtherMethod,      // line #18:     return RPCError(callingMethod, json['code'], json['message'], json['data']);
		OtherMethod,      // line #19:   }
		BlankLine,        // line #20:
		InstanceVariable, // line #21:   final String callingMethod;
		InstanceVariable, // line #22:   final int code;
		InstanceVariable, // line #23:   final String message;
		InstanceVariable, // line #24:   final Map data;
		BlankLine,        // line #25:
		MainConstructor,  // line #26:   RPCError(this.callingMethod, this.code, this.message, [this.data]);
		BlankLine,        // line #27:
		NamedConstructor, // line #28:   RPCError.withDetails(this.callingMethod, this.code, this.message,
		NamedConstructor, // line #29:       {Object details})
		NamedConstructor, // line #30:       : data = details == null ? null : <String, dynamic>{} {
		NamedConstructor, // line #31:     if (details != null) {
		NamedConstructor, // line #32:       data['details'] = details;
		NamedConstructor, // line #33:     }
		NamedConstructor, // line #34:   }
		BlankLine,        // line #35:
		OtherMethod,      // line #36:   String get details => data == null ? null : data['details'];
		BlankLine,        // line #37:
		OtherMethod,      // line #38:   /// Return a map representation of this error suitable for converstion to
		OtherMethod,      // line #39:   /// json.
		OtherMethod,      // line #40:   Map<String, dynamic> toMap() {
		OtherMethod,      // line #41:     Map<String, dynamic> map = {
		OtherMethod,      // line #42:       'code': code,
		OtherMethod,      // line #43:       'message': message,
		OtherMethod,      // line #44:     };
		OtherMethod,      // line #45:     if (data != null) {
		OtherMethod,      // line #46:       map['data'] = data;
		OtherMethod,      // line #47:     }
		OtherMethod,      // line #48:     return map;
		OtherMethod,      // line #49:   }
		BlankLine,        // line #50:
		OtherMethod,      // line #51:   String toString() {
		OtherMethod,      // line #52:     if (details == null) {
		OtherMethod,      // line #53:       return '$callingMethod: ($code) $message';
		OtherMethod,      // line #54:     } else {
		OtherMethod,      // line #55:       return '$callingMethod: ($code) $message\n$details';
		OtherMethod,      // line #56:     }
		OtherMethod,      // line #57:   }
		BlankLine,        // line #58:
	}

	runFullStylizer(t, opts, source, wantSource, want)

	// wantAfter := []EntityType{
	// 	Unknown,          // line #1: {
	// 	StaticVariable,   // line #2:   /// Application specific error codes.
	// 	StaticVariable,   // line #3:   static const int kServerError = -32000;
	// 	BlankLine,        // line #4:
	// 	StaticVariable,   // line #5:   /// The JSON sent is not a valid Request object.
	// 	StaticVariable,   // line #6:   static const int kInvalidRequest = -32600;
	// 	BlankLine,        // line #7:
	// 	StaticVariable,   // line #8:   /// The method does not exist or is not available.
	// 	StaticVariable,   // line #9:   static const int kMethodNotFound = -32601;
	// 	BlankLine,        // line #10:
	// 	StaticVariable,   // line #11:   /// Invalid method parameter(s), such as a mismatched type.
	// 	StaticVariable,   // line #12:   static const int kInvalidParams = -32602;
	// 	BlankLine,        // line #13:
	// 	StaticVariable,   // line #14:   /// Internal JSON-RPC error.
	// 	StaticVariable,   // line #15:   static const int kInternalError = -32603;
	// 	BlankLine,        // line #16:
	// 	OtherMethod,      // line #17:   static RPCError parse(String callingMethod, dynamic json) {
	// 	OtherMethod,      // line #18:     return RPCError(callingMethod, json['code'], json['message'], json['data']);
	// 	OtherMethod,      // line #19:   }
	// 	BlankLine,        // line #20:
	// 	InstanceVariable, // line #21:   final String callingMethod;
	// 	InstanceVariable, // line #22:   final int code;
	// 	InstanceVariable, // line #23:   final String message;
	// 	InstanceVariable, // line #24:   final Map data;
	// 	BlankLine,        // line #25:
	// 	MainConstructor,  // line #26:   RPCError(this.callingMethod, this.code, this.message, [this.data]);
	// 	BlankLine,        // line #27:
	// 	NamedConstructor, // line #28:   RPCError.withDetails(this.callingMethod, this.code, this.message,
	// 	NamedConstructor, // line #29:       {Object details})
	// 	NamedConstructor, // line #30:       : data = details == null ? null : <String, dynamic>{} {
	// 	NamedConstructor,          // line #31:     if (details != null) {
	// 	NamedConstructor,          // line #32:       data['details'] = details;
	// 	NamedConstructor,          // line #33:     }
	// 	InstanceVariable, // line #34:   }
	// 	BlankLine,        // line #35:
	// 	OtherMethod,      // line #36:   String get details => data == null ? null : data['details'];
	// 	BlankLine,        // line #37:
	// 	OtherMethod,      // line #38:   /// Return a map representation of this error suitable for converstion to
	// 	OtherMethod,      // line #39:   /// json.
	// 	OtherMethod,      // line #40:   Map<String, dynamic> toMap() {
	// 	OtherMethod,      // line #41:     Map<String, dynamic> map = {
	// 	OtherMethod,      // line #42:       'code': code,
	// 	OtherMethod,      // line #43:       'message': message,
	// 	OtherMethod,      // line #44:     };
	// 	OtherMethod,      // line #45:     if (data != null) {
	// 	OtherMethod,      // line #46:       map['data'] = data;
	// 	OtherMethod,      // line #47:     }
	// 	OtherMethod,      // line #48:     return map;
	// 	OtherMethod,      // line #49:   }
	// 	BlankLine,        // line #50:
	// 	OtherMethod,      // line #51:   String toString() {
	// 	OtherMethod,      // line #52:     if (details == null) {
	// 	OtherMethod,      // line #53:       return '$callingMethod: ($code) $message';
	// 	OtherMethod,      // line #54:     } else {
	// 	OtherMethod,      // line #55:       return '$callingMethod: ($code) $message\n$details';
	// 	OtherMethod,      // line #56:     }
	// 	OtherMethod,      // line #57:   }
	// 	BlankLine,        // line #58:
	// }

	// runParsePhase(t, opts, wantSource, wantAfter)
	// runFullStylizer(t, opts, wantSource, wantSource, nil)
}
