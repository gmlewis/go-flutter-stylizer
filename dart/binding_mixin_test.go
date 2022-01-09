/*
Copyright © 2022 Glenn M. Lewis

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

//go:embed testfiles/binding_mixin.dart.txt
var binding_mixin_dart_txt string

//go:embed testfiles/binding_mixin_want.txt
var binding_mixin_want_txt string

func TestBindingMixinExample(t *testing.T) {
	source := binding_mixin_dart_txt
	wantSource := binding_mixin_want_txt

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,                 // line #15: {
		OtherMethod,             // line #16:   /// The current [SemanticsBinding], if one has been created.
		OtherMethod,             // line #17:   static SemanticsBinding? get instance => _instance;
		StaticPrivateVariable,   // line #18:   static SemanticsBinding? _instance;
		BlankLine,               // line #19:
		OverrideMethod,          // line #20:   @override
		OverrideMethod,          // line #21:   void initInstances() {
		OverrideMethod,          // line #22:     super.initInstances();
		OverrideMethod,          // line #23:     _instance = this;
		OverrideMethod,          // line #24:     _accessibilityFeatures = window.accessibilityFeatures;
		OverrideMethod,          // line #25:   }
		BlankLine,               // line #26:
		OtherMethod,             // line #27:   /// Called when the platform accessibility features change.
		OtherMethod,             // line #28:   ///
		OtherMethod,             // line #29:   /// See [dart:ui.PlatformDispatcher.onAccessibilityFeaturesChanged].
		OtherMethod,             // line #30:   @protected
		OtherMethod,             // line #31:   void handleAccessibilityFeaturesChanged() {
		OtherMethod,             // line #32:     _accessibilityFeatures = window.accessibilityFeatures;
		OtherMethod,             // line #33:   }
		BlankLine,               // line #34:
		OtherMethod,             // line #35:   /// Creates an empty semantics update builder.
		OtherMethod,             // line #36:   ///
		OtherMethod,             // line #37:   /// The caller is responsible for filling out the semantics node updates.
		OtherMethod,             // line #38:   ///
		OtherMethod,             // line #39:   /// This method is used by the [SemanticsOwner] to create builder for all its
		OtherMethod,             // line #40:   /// semantics updates.
		OtherMethod,             // line #41:   ui.SemanticsUpdateBuilder createSemanticsUpdateBuilder() {
		OtherMethod,             // line #42:     return ui.SemanticsUpdateBuilder();
		OtherMethod,             // line #43:   }
		BlankLine,               // line #44:
		OtherMethod,             // line #45:   /// The currently active set of [AccessibilityFeatures].
		OtherMethod,             // line #46:   ///
		OtherMethod,             // line #47:   /// This is initialized the first time [runApp] is called and updated whenever
		OtherMethod,             // line #48:   /// a flag is changed.
		OtherMethod,             // line #49:   ///
		OtherMethod,             // line #50:   /// To listen to changes to accessibility features, create a
		OtherMethod,             // line #51:   /// [WidgetsBindingObserver] and listen to
		OtherMethod,             // line #52:   /// [WidgetsBindingObserver.didChangeAccessibilityFeatures].
		OtherMethod,             // line #53:   ui.AccessibilityFeatures get accessibilityFeatures => _accessibilityFeatures;
		PrivateInstanceVariable, // line #54:   late ui.AccessibilityFeatures _accessibilityFeatures;
		BlankLine,               // line #55:
		OtherMethod,             // line #56:   /// The platform is requesting that animations be disabled or simplified.
		OtherMethod,             // line #57:   ///
		OtherMethod,             // line #58:   /// This setting can be overridden for testing or debugging by setting
		OtherMethod,             // line #59:   /// [debugSemanticsDisableAnimations].
		OtherMethod,             // line #60:   bool get disableAnimations {
		OtherMethod,             // line #61:     bool value = _accessibilityFeatures.disableAnimations;
		OtherMethod,             // line #62:     assert(() {
		OtherMethod,             // line #63:       if (debugSemanticsDisableAnimations != null)
		OtherMethod,             // line #64:         value = debugSemanticsDisableAnimations!;
		OtherMethod,             // line #65:       return true;
		OtherMethod,             // line #66:     }());
		OtherMethod,             // line #67:     return value;
		OtherMethod,             // line #68:   }
		BlankLine,               // line #69:
	}

	runFullStylizer(t, opts, source, wantSource, want)
}
