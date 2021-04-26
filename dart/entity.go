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

// Entity represents a single, independent feature of a Dart.Class.
type Entity struct {
	entityType EntityType
	lines      []*Line
	name       string // Used for sorting, but could be "".
}

// EntityType represents a type of Dart Line.
type EntityType int

const (
	Unknown EntityType = iota
	BlankLine
	SingleLineComment
	MultiLineComment
	MainConstructor
	NamedConstructor
	StaticVariable
	InstanceVariable
	OverrideVariable
	StaticPrivateVariable
	PrivateInstanceVariable
	OverrideMethod
	OtherMethod
	BuildMethod
	GetterMethod
)

func (e EntityType) String() string {
	switch e {
	default:
		return "Unknown"
	case BlankLine:
		return "BlankLine"
	case SingleLineComment:
		return "SingleLineComment"
	case MultiLineComment:
		return "MultiLineComment"
	case MainConstructor:
		return "MainConstructor"
	case NamedConstructor:
		return "NamedConstructor"
	case StaticVariable:
		return "StaticVariable"
	case InstanceVariable:
		return "InstanceVariable"
	case OverrideVariable:
		return "OverrideVariable"
	case StaticPrivateVariable:
		return "StaticPrivateVariable"
	case PrivateInstanceVariable:
		return "PrivateInstanceVariable"
	case OverrideMethod:
		return "OverrideMethod"
	case OtherMethod:
		return "OtherMethod"
	case BuildMethod:
		return "BuildMethod"
	case GetterMethod:
		return "GetterMethod"
	}
}
