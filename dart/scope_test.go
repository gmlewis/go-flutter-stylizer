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

//go:embed testfiles/scope.dart.txt
var scope_dart_txt string

//go:embed testfiles/scope_want.txt
var scope_want_txt string

func TestScope_GetClasses(t *testing.T) {
	e, err := NewEditor(scope_dart_txt, Options{})
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	classes, err := e.GetClasses()
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 10; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}
}

func TestScope_Class1(t *testing.T) {
	source := scope_dart_txt[560:769]
	// wantSource := scope_want_txt[560:769]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,         // line #1: {
		MainConstructor, // line #2:   ClassScope(Scope parent, ClassElement element) : super(parent) {
		MainConstructor, // line #3:     element.accessors.forEach(_addPropertyAccessor);
		MainConstructor, // line #4:     element.methods.forEach(_addGetter);
		MainConstructor, // line #5:   }
		BlankLine,       // line #6:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class2(t *testing.T) {
	source := scope_dart_txt[769:1027]
	// wantSource := scope_want_txt[560:769]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,         // line #2: {
		MainConstructor, // line #3:   ConstructorInitializerScope(Scope parent, ConstructorElement element)
		MainConstructor, // line #4:       : super(parent) {
		MainConstructor, // line #5:     element.parameters.forEach(_addGetter);
		MainConstructor, // line #6:   }
		BlankLine,       // line #7:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class3(t *testing.T) {
	source := scope_dart_txt[1027:2248]
	wantSource := scope_want_txt[1027:2248]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,                 // line #2: {
		PrivateInstanceVariable, // line #3:   final Scope _parent;
		PrivateInstanceVariable, // line #4:   final Map<String, Element> _getters = {};
		PrivateInstanceVariable, // line #5:   final Map<String, Element> _setters = {};
		BlankLine,               // line #6:
		MainConstructor,         // line #7:   EnclosedScope(Scope parent) : _parent = parent;
		BlankLine,               // line #8:
		OtherMethod,             // line #9:   Scope get parent => _parent;
		BlankLine,               // line #10:
		OverrideMethod,          // line #11:   @Deprecated('Use lookup2() that is closer to the language specification')
		OverrideMethod,          // line #12:   @override
		OverrideMethod,          // line #13:   Element lookup({@required String id, @required bool setter}) {
		OverrideMethod,          // line #14:     var result = lookup2(id);
		OverrideMethod,          // line #15:     return setter ? result.setter : result.getter;
		OverrideMethod,          // line #16:   }
		BlankLine,               // line #17:
		OverrideMethod,          // line #18:   @override
		OverrideMethod,          // line #19:   ScopeLookupResult lookup2(String id) {
		OverrideMethod,          // line #20:     var getter = _getters[id];
		OverrideMethod,          // line #21:     var setter = _setters[id];
		OverrideMethod,          // line #22:     if (getter != null || setter != null) {
		OverrideMethod,          // line #23:       return ScopeLookupResult(getter, setter);
		OverrideMethod,          // line #24:     }
		OverrideMethod,          // line #25:
		OverrideMethod,          // line #26:     return _parent.lookup2(id);
		OverrideMethod,          // line #27:   }
		BlankLine,               // line #28:
		OtherMethod,             // line #29:   void _addGetter(Element element) {
		OtherMethod,             // line #30:     _addTo(_getters, element);
		OtherMethod,             // line #31:   }
		BlankLine,               // line #32:
		OtherMethod,             // line #33:   void _addPropertyAccessor(PropertyAccessorElement element) {
		OtherMethod,             // line #34:     if (element.isGetter) {
		OtherMethod,             // line #35:       _addGetter(element);
		OtherMethod,             // line #36:     } else {
		OtherMethod,             // line #37:       _addSetter(element);
		OtherMethod,             // line #38:     }
		OtherMethod,             // line #39:   }
		BlankLine,               // line #40:
		OtherMethod,             // line #41:   void _addSetter(Element element) {
		OtherMethod,             // line #42:     _addTo(_setters, element);
		OtherMethod,             // line #43:   }
		BlankLine,               // line #44:
		OtherMethod,             // line #45:   void _addTo(Map<String, Element> map, Element element) {
		OtherMethod,             // line #46:     var id = element.displayName;
		OtherMethod,             // line #47:     map[id] ??= element;
		OtherMethod,             // line #48:   }
		BlankLine,               // line #49:
	}

	// runParsePhase(t, opts, source, [][]EntityType{want})
	runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class4(t *testing.T) {
	source := scope_dart_txt[2248:2521]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,         // line #2: {
		MainConstructor, // line #3:   ExtensionScope(
		MainConstructor, // line #4:     Scope parent,
		MainConstructor, // line #5:     ExtensionElement element,
		MainConstructor, // line #6:   ) : super(parent) {
		MainConstructor, // line #7:     element.accessors.forEach(_addPropertyAccessor);
		MainConstructor, // line #8:     element.methods.forEach(_addGetter);
		MainConstructor, // line #9:   }
		BlankLine,       // line #10:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class5(t *testing.T) {
	source := scope_dart_txt[2521:2818]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,         // line #1: {
		MainConstructor, // line #2:   FormalParameterScope(
		MainConstructor, // line #3:     Scope parent,
		MainConstructor, // line #4:     List<ParameterElement> elements,
		MainConstructor, // line #5:   ) : super(parent) {
		MainConstructor, // line #6:     for (var parameter in elements) {
		MainConstructor, // line #7:       if (parameter is! FieldFormalParameterElement) {
		MainConstructor, // line #8:         _addGetter(parameter);
		MainConstructor, // line #9:       }
		MainConstructor, // line #10:     }
		MainConstructor, // line #11:   }
		BlankLine,       // line #12:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class6(t *testing.T) {
	source := scope_dart_txt[2818:5338]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,                 // line #1: {
		PrivateInstanceVariable, // line #2:   final LibraryElement _element;
		InstanceVariable,        // line #3:   final List<ExtensionElement> extensions = [];
		BlankLine,               // line #4:
		MainConstructor,         // line #5:   LibraryScope(LibraryElement element)
		MainConstructor,         // line #6:       : _element = element,
		MainConstructor,         // line #7:         super(_LibraryImportScope(element)) {
		MainConstructor,         // line #8:     extensions.addAll((_parent as _LibraryImportScope).extensions);
		MainConstructor,         // line #9:
		MainConstructor,         // line #10:     _element.prefixes.forEach(_addGetter);
		MainConstructor,         // line #11:     _element.units.forEach(_addUnitElements);
		MainConstructor,         // line #12:   }
		BlankLine,               // line #13:
		OtherMethod,             // line #14:   bool shouldIgnoreUndefined({
		OtherMethod,             // line #15:     @required String prefix,
		OtherMethod,             // line #16:     @required String name,
		OtherMethod,             // line #17:   }) {
		OtherMethod,             // line #18:     Iterable<NamespaceCombinator> getShowCombinators(
		OtherMethod,             // line #19:         ImportElement importElement) {
		OtherMethod,             // line #20:       return importElement.combinators.whereType<ShowElementCombinator>();
		OtherMethod,             // line #21:     }
		OtherMethod,             // line #22:
		OtherMethod,             // line #23:     if (prefix != null) {
		OtherMethod,             // line #24:       for (var importElement in _element.imports) {
		OtherMethod,             // line #25:         if (importElement.prefix?.name == prefix &&
		OtherMethod,             // line #26:             importElement.importedLibrary?.isSynthetic != false) {
		OtherMethod,             // line #27:           var showCombinators = getShowCombinators(importElement);
		OtherMethod,             // line #28:           if (showCombinators.isEmpty) {
		OtherMethod,             // line #29:             return true;
		OtherMethod,             // line #30:           }
		OtherMethod,             // line #31:           for (ShowElementCombinator combinator in showCombinators) {
		OtherMethod,             // line #32:             if (combinator.shownNames.contains(name)) {
		OtherMethod,             // line #33:               return true;
		OtherMethod,             // line #34:             }
		OtherMethod,             // line #35:           }
		OtherMethod,             // line #36:         }
		OtherMethod,             // line #37:       }
		OtherMethod,             // line #38:     } else {
		OtherMethod,             // line #39:       // TODO(scheglov) merge for(s).
		OtherMethod,             // line #40:       for (var importElement in _element.imports) {
		OtherMethod,             // line #41:         if (importElement.prefix == null &&
		OtherMethod,             // line #42:             importElement.importedLibrary?.isSynthetic != false) {
		OtherMethod,             // line #43:           for (ShowElementCombinator combinator
		OtherMethod,             // line #44:               in getShowCombinators(importElement)) {
		OtherMethod,             // line #45:             if (combinator.shownNames.contains(name)) {
		OtherMethod,             // line #46:               return true;
		OtherMethod,             // line #47:             }
		OtherMethod,             // line #48:           }
		OtherMethod,             // line #49:         }
		OtherMethod,             // line #50:       }
		OtherMethod,             // line #51:
		OtherMethod,             // line #52:       if (name.startsWith(r'_$')) {
		OtherMethod,             // line #53:         for (var partElement in _element.parts) {
		OtherMethod,             // line #54:           if (partElement.isSynthetic &&
		OtherMethod,             // line #55:               isGeneratedSource(partElement.source)) {
		OtherMethod,             // line #56:             return true;
		OtherMethod,             // line #57:           }
		OtherMethod,             // line #58:         }
		OtherMethod,             // line #59:       }
		OtherMethod,             // line #60:     }
		OtherMethod,             // line #61:
		OtherMethod,             // line #62:     return false;
		OtherMethod,             // line #63:   }
		BlankLine,               // line #64:
		OtherMethod,             // line #65:   void _addExtension(ExtensionElement element) {
		OtherMethod,             // line #66:     _addGetter(element);
		OtherMethod,             // line #67:     if (!extensions.contains(element)) {
		OtherMethod,             // line #68:       extensions.add(element);
		OtherMethod,             // line #69:     }
		OtherMethod,             // line #70:   }
		BlankLine,               // line #71:
		OtherMethod,             // line #72:   void _addUnitElements(CompilationUnitElement compilationUnit) {
		OtherMethod,             // line #73:     compilationUnit.accessors.forEach(_addPropertyAccessor);
		OtherMethod,             // line #74:     compilationUnit.enums.forEach(_addGetter);
		OtherMethod,             // line #75:     compilationUnit.extensions.forEach(_addExtension);
		OtherMethod,             // line #76:     compilationUnit.functions.forEach(_addGetter);
		OtherMethod,             // line #77:     compilationUnit.functionTypeAliases.forEach(_addGetter);
		OtherMethod,             // line #78:     compilationUnit.mixins.forEach(_addGetter);
		OtherMethod,             // line #79:     compilationUnit.types.forEach(_addGetter);
		OtherMethod,             // line #80:   }
		BlankLine,               // line #81:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class7(t *testing.T) {
	source := scope_dart_txt[5338:5486]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,         // line #1: {
		MainConstructor, // line #2:   LocalScope(Scope parent) : super(parent);
		BlankLine,       // line #3:
		OtherMethod,     // line #4:   void add(Element element) {
		OtherMethod,     // line #5:     _addGetter(element);
		OtherMethod,     // line #6:   }
		BlankLine,       // line #7:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class8(t *testing.T) {
	source := scope_dart_txt[5486:8326]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,                 // line #1: {
		PrivateInstanceVariable, // line #2:   final LibraryElement _library;
		PrivateInstanceVariable, // line #3:   final Map<String, Element> _getters = {};
		PrivateInstanceVariable, // line #4:   final Map<String, Element> _setters = {};
		PrivateInstanceVariable, // line #5:   final Set<ExtensionElement> _extensions = {};
		PrivateInstanceVariable, // line #6:   LibraryElement _deferredLibrary;
		BlankLine,               // line #7:
		MainConstructor,         // line #8:   PrefixScope(this._library, PrefixElement prefix) {
		MainConstructor,         // line #9:     for (var import in _library.imports) {
		MainConstructor,         // line #10:       if (import.prefix == prefix) {
		MainConstructor,         // line #11:         var elements = impl.NamespaceBuilder().getImportedElements(import);
		MainConstructor,         // line #12:         elements.forEach(_add);
		MainConstructor,         // line #13:         if (import.isDeferred) {
		MainConstructor,         // line #14:           _deferredLibrary ??= import.importedLibrary;
		MainConstructor,         // line #15:         }
		MainConstructor,         // line #16:       }
		MainConstructor,         // line #17:     }
		MainConstructor,         // line #18:   }
		BlankLine,               // line #19:
		OverrideMethod,          // line #20:   @Deprecated('Use lookup2() that is closer to the language specification')
		OverrideMethod,          // line #21:   @override
		OverrideMethod,          // line #22:   Element lookup({@required String id, @required bool setter}) {
		OverrideMethod,          // line #23:     var result = lookup2(id);
		OverrideMethod,          // line #24:     return setter ? result.setter : result.getter;
		OverrideMethod,          // line #25:   }
		BlankLine,               // line #26:
		OverrideMethod,          // line #27:   @override
		OverrideMethod,          // line #28:   ScopeLookupResult lookup2(String id) {
		OverrideMethod,          // line #29:     if (_deferredLibrary != null && id == FunctionElement.LOAD_LIBRARY_NAME) {
		OverrideMethod,          // line #30:       return ScopeLookupResult(_deferredLibrary.loadLibraryFunction, null);
		OverrideMethod,          // line #31:     }
		OverrideMethod,          // line #32:
		OverrideMethod,          // line #33:     var getter = _getters[id];
		OverrideMethod,          // line #34:     var setter = _setters[id];
		OverrideMethod,          // line #35:     return ScopeLookupResult(getter, setter);
		OverrideMethod,          // line #36:   }
		BlankLine,               // line #37:
		OtherMethod,             // line #38:   void _add(Element element) {
		OtherMethod,             // line #39:     if (element is PropertyAccessorElement && element.isSetter) {
		OtherMethod,             // line #40:       _addTo(map: _setters, element: element);
		OtherMethod,             // line #41:     } else {
		OtherMethod,             // line #42:       _addTo(map: _getters, element: element);
		OtherMethod,             // line #43:       if (element is ExtensionElement) {
		OtherMethod,             // line #44:         _extensions.add(element);
		OtherMethod,             // line #45:       }
		OtherMethod,             // line #46:     }
		OtherMethod,             // line #47:   }
		BlankLine,               // line #48:
		OtherMethod,             // line #49:   void _addTo({
		OtherMethod,             // line #50:     @required Map<String, Element> map,
		OtherMethod,             // line #51:     @required Element element,
		OtherMethod,             // line #52:   }) {
		OtherMethod,             // line #53:     var id = element.displayName;
		OtherMethod,             // line #54:
		OtherMethod,             // line #55:     var existing = map[id];
		OtherMethod,             // line #56:     if (existing != null && existing != element) {
		OtherMethod,             // line #57:       map[id] = _merge(existing, element);
		OtherMethod,             // line #58:       return;
		OtherMethod,             // line #59:     }
		OtherMethod,             // line #60:
		OtherMethod,             // line #61:     map[id] = element;
		OtherMethod,             // line #62:   }
		BlankLine,               // line #63:
		OtherMethod,             // line #64:   Element _merge(Element existing, Element other) {
		OtherMethod,             // line #65:     if (_isSdkElement(existing)) {
		OtherMethod,             // line #66:       if (!_isSdkElement(other)) {
		OtherMethod,             // line #67:         return other;
		OtherMethod,             // line #68:       }
		OtherMethod,             // line #69:     } else {
		OtherMethod,             // line #70:       if (_isSdkElement(other)) {
		OtherMethod,             // line #71:         return existing;
		OtherMethod,             // line #72:       }
		OtherMethod,             // line #73:     }
		OtherMethod,             // line #74:
		OtherMethod,             // line #75:     var conflictingElements = <Element>{};
		OtherMethod,             // line #76:     _addElement(conflictingElements, existing);
		OtherMethod,             // line #77:     _addElement(conflictingElements, other);
		OtherMethod,             // line #78:
		OtherMethod,             // line #79:     return MultiplyDefinedElementImpl(
		OtherMethod,             // line #80:       _library.context,
		OtherMethod,             // line #81:       _library.session,
		OtherMethod,             // line #82:       conflictingElements.first.name,
		OtherMethod,             // line #83:       conflictingElements.toList(),
		OtherMethod,             // line #84:     );
		OtherMethod,             // line #85:   }
		BlankLine,               // line #86:
		OtherMethod,             // line #87:   static void _addElement(
		OtherMethod,             // line #88:     Set<Element> conflictingElements,
		OtherMethod,             // line #89:     Element element,
		OtherMethod,             // line #90:   ) {
		OtherMethod,             // line #91:     if (element is MultiplyDefinedElementImpl) {
		OtherMethod,             // line #92:       conflictingElements.addAll(element.conflictingElements);
		OtherMethod,             // line #93:     } else {
		OtherMethod,             // line #94:       conflictingElements.add(element);
		OtherMethod,             // line #95:     }
		OtherMethod,             // line #96:   }
		BlankLine,               // line #97:
		OtherMethod,             // line #98:   static bool _isSdkElement(Element element) {
		OtherMethod,             // line #99:     if (element is DynamicElementImpl || element is NeverElementImpl) {
		OtherMethod,             // line #100:       return true;
		OtherMethod,             // line #101:     }
		OtherMethod,             // line #102:     if (element is MultiplyDefinedElement) {
		OtherMethod,             // line #103:       return false;
		OtherMethod,             // line #104:     }
		OtherMethod,             // line #105:     return element.library.isInSdk;
		OtherMethod,             // line #106:   }
		BlankLine,               // line #107:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class9(t *testing.T) {
	source := scope_dart_txt[8326:8519]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,         // line #1: {
		MainConstructor, // line #2:   TypeParameterScope(
		MainConstructor, // line #3:     Scope parent,
		MainConstructor, // line #4:     List<TypeParameterElement> elements,
		MainConstructor, // line #5:   ) : super(parent) {
		MainConstructor, // line #6:     elements.forEach(_addGetter);
		MainConstructor, // line #7:   }
		BlankLine,       // line #8:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}

func TestScope_Class10(t *testing.T) {
	source := scope_dart_txt[8519:9323]
	// wantSource := scope_want_txt[2248:]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,                 // line #1: {
		PrivateInstanceVariable, // line #2:   final LibraryElement _library;
		PrivateInstanceVariable, // line #3:   final PrefixScope _nullPrefixScope;
		PrivateInstanceVariable, // line #4:   List<ExtensionElement> _extensions;
		BlankLine,               // line #5:
		MainConstructor,         // line #6:   _LibraryImportScope(LibraryElement library)
		MainConstructor,         // line #7:       : _library = library,
		MainConstructor,         // line #8:         _nullPrefixScope = PrefixScope(library, null);
		BlankLine,               // line #9:
		OtherMethod,             // line #10:   List<ExtensionElement> get extensions {
		OtherMethod,             // line #11:     return _extensions ??= {
		OtherMethod,             // line #12:       ..._nullPrefixScope._extensions,
		OtherMethod,             // line #13:       for (var prefix in _library.prefixes)
		OtherMethod,             // line #14:         ...(prefix.scope as PrefixScope)._extensions,
		OtherMethod,             // line #15:     }.toList();
		OtherMethod,             // line #16:   }
		BlankLine,               // line #17:
		OverrideMethod,          // line #18:   @Deprecated('Use lookup2() that is closer to the language specification')
		OverrideMethod,          // line #19:   @override
		OverrideMethod,          // line #20:   Element lookup({@required String id, @required bool setter}) {
		OverrideMethod,          // line #21:     throw UnimplementedError();
		OverrideMethod,          // line #22:   }
		BlankLine,               // line #23:
		OverrideMethod,          // line #24:   @override
		OverrideMethod,          // line #25:   ScopeLookupResult lookup2(String id) {
		OverrideMethod,          // line #26:     return _nullPrefixScope.lookup2(id);
		OverrideMethod,          // line #27:   }
		BlankLine,               // line #28:
	}

	runParsePhase(t, opts, source, [][]EntityType{want})
	// runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}
