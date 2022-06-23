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

//go:embed testfiles/issue20.dart.txt
var issue20_dart_txt string

//go:embed testfiles/issue20_want.txt
var issue20_want_txt string

func TestIssue20(t *testing.T) {
	source := issue20_dart_txt
	wantSource := issue20_want_txt

	opts := &Options{
		MemberOrdering: []string{
			"public-static-variables",
			"public-instance-variables",
			"public-override-variables",
			"private-static-variables",
			"private-instance-variables",
			"public-constructor",
			"named-constructors",
			"public-override-methods",
			"public-other-methods",
			"build-method",
		},
	}

	want := []EntityType{
		Unknown,                 // {
		InstanceVariable,        //   final String typeName;
		PrivateInstanceVariable, //   final int _attributes;
		InstanceVariable,        //   final int baseTypeToken;
		InstanceVariable,        //   final TypeIdentifier? typeSpec;
		BlankLine,               //
		PrivateInstanceVariable, //   final _interfaces = <TypeDef>[];
		PrivateInstanceVariable, //   final _fields = <Field>[];
		PrivateInstanceVariable, //   final _methods = <Method>[];
		PrivateInstanceVariable, //   final _properties = <Property>[];
		PrivateInstanceVariable, //   final _events = <Event>[];
		BlankLine,               //
		OtherMethod,             //   TypeVisibility get typeVisibility =>
		OtherMethod,             //       TypeVisibility.values[_attributes & CorTypeAttr.tdVisibilityMask];
		BlankLine,               //
		OtherMethod,             //   TypeLayout get typeLayout {
		OtherMethod,             //     switch (_attributes & CorTypeAttr.tdLayoutMask) {
		OtherMethod,             //       case CorTypeAttr.tdAutoLayout:
		OtherMethod,             //         return TypeLayout.auto;
		OtherMethod,             //       case CorTypeAttr.tdSequentialLayout:
		OtherMethod,             //         return TypeLayout.sequential;
		OtherMethod,             //       case CorTypeAttr.tdExplicitLayout:
		OtherMethod,             //         return TypeLayout.explicit;
		OtherMethod,             //       default:
		OtherMethod,             //         throw WinmdException('Attribute missing type layout information');
		OtherMethod,             //     }
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Is the type a class?
		OtherMethod,             //   bool get isClass =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdClassSemanticsMask == CorTypeAttr.tdClass;
		BlankLine,               //
		OtherMethod,             //   /// Is the type an interface?
		OtherMethod,             //   bool get isInterface =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdClassSemanticsMask == CorTypeAttr.tdInterface;
		BlankLine,               //
		OtherMethod,             //   bool get isAbstract =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdAbstract == CorTypeAttr.tdAbstract;
		BlankLine,               //
		OtherMethod,             //   bool get isSealed =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdSealed == CorTypeAttr.tdSealed;
		BlankLine,               //
		OtherMethod,             //   bool get isSpecialName =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdSpecialName == CorTypeAttr.tdSpecialName;
		BlankLine,               //
		OtherMethod,             //   bool get isImported =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdImport == CorTypeAttr.tdImport;
		BlankLine,               //
		OtherMethod,             //   bool get isSerializable =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdSerializable == CorTypeAttr.tdSerializable;
		BlankLine,               //
		OtherMethod,             //   bool get isWindowsRuntime =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdWindowsRuntime ==
		OtherMethod,             //       CorTypeAttr.tdWindowsRuntime;
		BlankLine,               //
		OtherMethod,             //   bool get isRTSpecialName =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdRTSpecialName == CorTypeAttr.tdRTSpecialName;
		BlankLine,               //
		OtherMethod,             //   StringFormat get stringFormat {
		OtherMethod,             //     switch (_attributes & CorTypeAttr.tdStringFormatMask) {
		OtherMethod,             //       case CorTypeAttr.tdAnsiClass:
		OtherMethod,             //         return StringFormat.ansi;
		OtherMethod,             //       case CorTypeAttr.tdUnicodeClass:
		OtherMethod,             //         return StringFormat.unicode;
		OtherMethod,             //       case CorTypeAttr.tdAutoClass:
		OtherMethod,             //         return StringFormat.auto;
		OtherMethod,             //       case CorTypeAttr.tdCustomFormatClass:
		OtherMethod,             //         return StringFormat.custom;
		OtherMethod,             //       default:
		OtherMethod,             //         throw WinmdException('Attribute missing string format information');
		OtherMethod,             //     }
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   bool get isBeforeFieldInit =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdBeforeFieldInit ==
		OtherMethod,             //       CorTypeAttr.tdBeforeFieldInit;
		BlankLine,               //
		OtherMethod,             //   bool get isForwarder =>
		OtherMethod,             //       _attributes & CorTypeAttr.tdForwarder == CorTypeAttr.tdForwarder;
		BlankLine,               //
		OtherMethod,             //   /// Is the type a delegate?
		OtherMethod,             //   bool get isDelegate => parent?.typeName == 'System.MulticastDelegate';
		BlankLine,               //
		OtherMethod,             //   /// Retrieve class layout information.
		OtherMethod,             //   ///
		OtherMethod,             //   /// This includes the packing alignment, the minimum class size, and the field
		OtherMethod,             //   /// layout (e.g. for sparsely or overlapping structs).
		OtherMethod,             //   ClassLayout get classLayout => ClassLayout(reader, token);
		BlankLine,               //
		OtherMethod,             //   /// Is the type a non-Windows Runtime type, such as System.Object or
		OtherMethod,             //   /// IInspectable?
		OtherMethod,             //   ///
		OtherMethod,             //   /// More information at:
		OtherMethod,             //   /// https://docs.microsoft.com/en-us/uwp/winrt-cref/winmd-files#type-system-encoding
		OtherMethod,             //   bool get isSystemType => systemTokens.containsValue(typeName);
		BlankLine,               //
		MainConstructor,         //   /// Create a typedef.
		MainConstructor,         //   ///
		MainConstructor,         //   /// Typically, typedefs should be obtained from a [WinmdScope] object rather
		MainConstructor,         //   /// than being created directly.
		MainConstructor,         //   TypeDef(IMetaDataImport2 reader,
		MainConstructor,         //       [int token = 0,
		MainConstructor,         //       this.typeName = '',
		MainConstructor,         //       this._attributes = 0,
		MainConstructor,         //       this.baseTypeToken = 0,
		MainConstructor,         //       this.typeSpec])
		MainConstructor,         //       : super(reader, token);
		BlankLine,               //
		NamedConstructor,        //   /// Creates a typedef object from its given token.
		NamedConstructor,        //   factory TypeDef.fromToken(IMetaDataImport2 reader, int token) {
		NamedConstructor,        //     switch (token & 0xFF000000) {
		NamedConstructor,        //       case CorTokenType.mdtTypeRef:
		NamedConstructor,        //         return TypeDef.fromTypeRefToken(reader, token);
		NamedConstructor,        //       case CorTokenType.mdtTypeDef:
		NamedConstructor,        //         return TypeDef.fromTypeDefToken(reader, token);
		NamedConstructor,        //       case CorTokenType.mdtTypeSpec:
		NamedConstructor,        //         return TypeDef.fromTypeSpecToken(reader, token);
		NamedConstructor,        //       default:
		NamedConstructor,        //         throw WinmdException('Unrecognized token ${token.toHexString(32)}');
		NamedConstructor,        //     }
		NamedConstructor,        //   }
		BlankLine,               //
		NamedConstructor,        //   /// Instantiate a typedef from a TypeDef token.
		NamedConstructor,        //   factory TypeDef.fromTypeDefToken(IMetaDataImport2 reader, int typeDefToken) {
		NamedConstructor,        //     final szTypeDef = stralloc(MAX_STRING_SIZE);
		NamedConstructor,        //     final pchTypeDef = calloc<ULONG>();
		NamedConstructor,        //     final pdwTypeDefFlags = calloc<DWORD>();
		NamedConstructor,        //     final ptkExtends = calloc<mdToken>();
		NamedConstructor,        //
		NamedConstructor,        //     try {
		NamedConstructor,        //       final hr = reader.GetTypeDefProps(typeDefToken, szTypeDef,
		NamedConstructor,        //           MAX_STRING_SIZE, pchTypeDef, pdwTypeDefFlags, ptkExtends);
		NamedConstructor,        //
		NamedConstructor,        //       if (SUCCEEDED(hr)) {
		NamedConstructor,        //         return TypeDef(reader, typeDefToken, szTypeDef.toDartString(),
		NamedConstructor,        //             pdwTypeDefFlags.value, ptkExtends.value);
		NamedConstructor,        //       } else {
		NamedConstructor,        //         throw WindowsException(hr);
		NamedConstructor,        //       }
		NamedConstructor,        //     } finally {
		NamedConstructor,        //       free(pchTypeDef);
		NamedConstructor,        //       free(pdwTypeDefFlags);
		NamedConstructor,        //       free(ptkExtends);
		NamedConstructor,        //       free(szTypeDef);
		NamedConstructor,        //     }
		NamedConstructor,        //   }
		BlankLine,               //
		NamedConstructor,        //   /// Instantiate a typedef from a TypeRef token.
		NamedConstructor,        //   ///
		NamedConstructor,        //   /// Unless the TypeRef token is `IInspectable`, the COM parent interface for
		NamedConstructor,        //   /// Windows Runtime classes, the TypeRef is used to obtain the host scope
		NamedConstructor,        //   /// metadata file, from which the TypeDef can be found and returned.
		NamedConstructor,        //   factory TypeDef.fromTypeRefToken(IMetaDataImport2 reader, int typeRefToken) {
		NamedConstructor,        //     final ptkResolutionScope = calloc<mdToken>();
		NamedConstructor,        //     final szName = stralloc(MAX_STRING_SIZE);
		NamedConstructor,        //     final pchName = calloc<ULONG>();
		NamedConstructor,        //
		NamedConstructor,        //     try {
		NamedConstructor,        //       final hr = reader.GetTypeRefProps(
		NamedConstructor,        //           typeRefToken, ptkResolutionScope, szName, MAX_STRING_SIZE, pchName);
		NamedConstructor,        //
		NamedConstructor,        //       if (SUCCEEDED(hr)) {
		NamedConstructor,        //         final typeName = szName.toDartString();
		NamedConstructor,        //         try {
		NamedConstructor,        //           final newScope = MetadataStore.getScopeForType(typeName);
		NamedConstructor,        //           return newScope.findTypeDef(typeName)!;
		NamedConstructor,        //         } catch (exception) {
		NamedConstructor,        //           // a token like IInspectable is out of reach of GetTypeRefProps, since it is
		NamedConstructor,        //           // a plain COM object. These objects are returned as system types.
		NamedConstructor,        //           if (systemTokens.containsKey(typeRefToken)) {
		NamedConstructor,        //             return TypeDef(reader, 0, systemTokens[typeRefToken]!);
		NamedConstructor,        //           }
		NamedConstructor,        //           if (systemTokens.containsValue(typeName)) {
		NamedConstructor,        //             return TypeDef(reader, 0, typeName);
		NamedConstructor,        //           }
		NamedConstructor,        //           // Perhaps we can find it in the current scope after all (for example,
		NamedConstructor,        //           // it's a nested class)
		NamedConstructor,        //           try {
		NamedConstructor,        //             final typedef = TypeDef.fromTypeDefToken(reader, typeRefToken);
		NamedConstructor,        //             return typedef;
		NamedConstructor,        //           } catch (exception) {
		NamedConstructor,        //             throw WinmdException(
		NamedConstructor,        //                 'Unable to find scope for $typeName [${typeRefToken.toHexString(32)}]...');
		NamedConstructor,        //           }
		NamedConstructor,        //         }
		NamedConstructor,        //       } else {
		NamedConstructor,        //         throw WindowsException(hr);
		NamedConstructor,        //       }
		NamedConstructor,        //     } finally {
		NamedConstructor,        //       free(ptkResolutionScope);
		NamedConstructor,        //       free(szName);
		NamedConstructor,        //       free(pchName);
		NamedConstructor,        //     }
		NamedConstructor,        //   }
		BlankLine,               //
		NamedConstructor,        //   /// Instantiate a typedef from a TypeSpec token.
		NamedConstructor,        //   factory TypeDef.fromTypeSpecToken(
		NamedConstructor,        //       IMetaDataImport2 reader, int typeSpecToken) {
		NamedConstructor,        //     final ppvSig = calloc<PCCOR_SIGNATURE>();
		NamedConstructor,        //     final pcbSig = calloc<ULONG>();
		NamedConstructor,        //
		NamedConstructor,        //     try {
		NamedConstructor,        //       final hr =
		NamedConstructor,        //           reader.GetTypeSpecFromToken(typeSpecToken, ppvSig.cast(), pcbSig);
		NamedConstructor,        //       final signature = ppvSig.value.asTypedList(pcbSig.value);
		NamedConstructor,        //       final typeTuple = parseTypeFromSignature(signature, reader);
		NamedConstructor,        //
		NamedConstructor,        //       if (SUCCEEDED(hr)) {
		NamedConstructor,        //         return TypeDef(
		NamedConstructor,        //             reader, typeSpecToken, '', 0, 0, typeTuple.typeIdentifier);
		NamedConstructor,        //       } else {
		NamedConstructor,        //         throw WindowsException(hr);
		NamedConstructor,        //       }
		NamedConstructor,        //     } finally {
		NamedConstructor,        //       free(ppvSig);
		NamedConstructor,        //       free(pcbSig);
		NamedConstructor,        //     }
		NamedConstructor,        //   }
		BlankLine,               //
		OtherMethod,             //   /// Converts an individual interface into a type.
		OtherMethod,             //   TypeDef processInterfaceToken(int token) {
		OtherMethod,             //     final ptkClass = calloc<mdTypeDef>();
		OtherMethod,             //     final ptkIface = calloc<mdToken>();
		OtherMethod,             //
		OtherMethod,             //     try {
		OtherMethod,             //       final hr = reader.GetInterfaceImplProps(token, ptkClass, ptkIface);
		OtherMethod,             //       if (SUCCEEDED(hr)) {
		OtherMethod,             //         final interfaceToken = ptkIface.value;
		OtherMethod,             //         return TypeDef.fromToken(reader, interfaceToken);
		OtherMethod,             //       } else {
		OtherMethod,             //         throw WindowsException(hr);
		OtherMethod,             //       }
		OtherMethod,             //     } finally {
		OtherMethod,             //       free(ptkClass);
		OtherMethod,             //       free(ptkIface);
		OtherMethod,             //     }
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Enumerate all interfaces that this type implements.
		OtherMethod,             //   List<TypeDef> get interfaces {
		OtherMethod,             //     if (_interfaces.isEmpty) {
		OtherMethod,             //       final phEnum = calloc<HCORENUM>();
		OtherMethod,             //       final rImpls = calloc<mdInterfaceImpl>();
		OtherMethod,             //       final pcImpls = calloc<ULONG>();
		OtherMethod,             //
		OtherMethod,             //       try {
		OtherMethod,             //         var hr = reader.EnumInterfaceImpls(phEnum, token, rImpls, 1, pcImpls);
		OtherMethod,             //         while (hr == S_OK) {
		OtherMethod,             //           final interfaceToken = rImpls.value;
		OtherMethod,             //
		OtherMethod,             //           _interfaces.add(processInterfaceToken(interfaceToken));
		OtherMethod,             //           hr = reader.EnumInterfaceImpls(phEnum, token, rImpls, 1, pcImpls);
		OtherMethod,             //         }
		OtherMethod,             //       } finally {
		OtherMethod,             //         reader.CloseEnum(phEnum.value);
		OtherMethod,             //         free(phEnum);
		OtherMethod,             //         free(rImpls);
		OtherMethod,             //         free(pcImpls);
		OtherMethod,             //       }
		OtherMethod,             //     }
		OtherMethod,             //     return _interfaces;
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Enumerate all fields contained within this type.
		OtherMethod,             //   List<Field> get fields {
		OtherMethod,             //     if (_fields.isEmpty) {
		OtherMethod,             //       final phEnum = calloc<HCORENUM>();
		OtherMethod,             //       final rgFields = calloc<mdFieldDef>();
		OtherMethod,             //       final pcTokens = calloc<ULONG>();
		OtherMethod,             //
		OtherMethod,             //       try {
		OtherMethod,             //         var hr = reader.EnumFields(phEnum, token, rgFields, 1, pcTokens);
		OtherMethod,             //         while (hr == S_OK) {
		OtherMethod,             //           final fieldToken = rgFields.value;
		OtherMethod,             //
		OtherMethod,             //           _fields.add(Field.fromToken(reader, fieldToken));
		OtherMethod,             //           hr = reader.EnumFields(phEnum, token, rgFields, 1, pcTokens);
		OtherMethod,             //         }
		OtherMethod,             //       } finally {
		OtherMethod,             //         reader.CloseEnum(phEnum.value);
		OtherMethod,             //         free(phEnum);
		OtherMethod,             //         free(rgFields);
		OtherMethod,             //         free(pcTokens);
		OtherMethod,             //       }
		OtherMethod,             //     }
		OtherMethod,             //     return _fields;
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Enumerate all methods contained within this type.
		OtherMethod,             //   List<Method> get methods {
		OtherMethod,             //     if (_methods.isEmpty) {
		OtherMethod,             //       final phEnum = calloc<HCORENUM>();
		OtherMethod,             //       final rgMethods = calloc<mdMethodDef>();
		OtherMethod,             //       final pcTokens = calloc<ULONG>();
		OtherMethod,             //
		OtherMethod,             //       try {
		OtherMethod,             //         var hr = reader.EnumMethods(phEnum, token, rgMethods, 1, pcTokens);
		OtherMethod,             //         while (hr == S_OK) {
		OtherMethod,             //           final methodToken = rgMethods.value;
		OtherMethod,             //
		OtherMethod,             //           _methods.add(Method.fromToken(reader, methodToken));
		OtherMethod,             //           hr = reader.EnumMethods(phEnum, token, rgMethods, 1, pcTokens);
		OtherMethod,             //         }
		OtherMethod,             //       } finally {
		OtherMethod,             //         reader.CloseEnum(phEnum.value);
		OtherMethod,             //         free(phEnum);
		OtherMethod,             //         free(rgMethods);
		OtherMethod,             //         free(pcTokens);
		OtherMethod,             //       }
		OtherMethod,             //     }
		OtherMethod,             //     return _methods;
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Enumerate all properties contained within this type.
		OtherMethod,             //   List<Property> get properties {
		OtherMethod,             //     if (_properties.isEmpty) {
		OtherMethod,             //       final phEnum = calloc<HCORENUM>();
		OtherMethod,             //       final rgProperties = calloc<mdProperty>();
		OtherMethod,             //       final pcProperties = calloc<ULONG>();
		OtherMethod,             //
		OtherMethod,             //       try {
		OtherMethod,             //         var hr =
		OtherMethod,             //             reader.EnumProperties(phEnum, token, rgProperties, 1, pcProperties);
		OtherMethod,             //         while (hr == S_OK) {
		OtherMethod,             //           final propertyToken = rgProperties.value;
		OtherMethod,             //
		OtherMethod,             //           _properties.add(Property.fromToken(reader, propertyToken));
		OtherMethod,             //           hr = reader.EnumMethods(phEnum, token, rgProperties, 1, pcProperties);
		OtherMethod,             //         }
		OtherMethod,             //       } finally {
		OtherMethod,             //         reader.CloseEnum(phEnum.value);
		OtherMethod,             //         free(phEnum);
		OtherMethod,             //         free(rgProperties);
		OtherMethod,             //         free(pcProperties);
		OtherMethod,             //       }
		OtherMethod,             //     }
		OtherMethod,             //     return _properties;
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Enumerate all events contained within this type.
		OtherMethod,             //   List<Event> get events {
		OtherMethod,             //     if (_events.isEmpty) {
		OtherMethod,             //       final phEnum = calloc<HCORENUM>();
		OtherMethod,             //       final rgEvents = calloc<mdEvent>();
		OtherMethod,             //       final pcEvents = calloc<ULONG>();
		OtherMethod,             //
		OtherMethod,             //       try {
		OtherMethod,             //         var hr = reader.EnumEvents(phEnum, token, rgEvents, 1, pcEvents);
		OtherMethod,             //         while (hr == S_OK) {
		OtherMethod,             //           final eventToken = rgEvents.value;
		OtherMethod,             //
		OtherMethod,             //           _events.add(Event.fromToken(reader, eventToken));
		OtherMethod,             //           hr = reader.EnumEvents(phEnum, token, rgEvents, 1, pcEvents);
		OtherMethod,             //         }
		OtherMethod,             //       } finally {
		OtherMethod,             //         reader.CloseEnum(phEnum.value);
		OtherMethod,             //         free(phEnum);
		OtherMethod,             //         free(rgEvents);
		OtherMethod,             //         free(pcEvents);
		OtherMethod,             //       }
		OtherMethod,             //     }
		OtherMethod,             //     return _events;
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Get a field matching the name, if one exists.
		OtherMethod,             //   ///
		OtherMethod,             //   /// Returns null if the field is not found.
		OtherMethod,             //   Field? findField(String fieldName) {
		OtherMethod,             //     try {
		OtherMethod,             //       return fields.firstWhere((field) => field.name == fieldName);
		OtherMethod,             //     } on StateError {
		OtherMethod,             //       return null;
		OtherMethod,             //     }
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Get a method matching the name, if one exists.
		OtherMethod,             //   ///
		OtherMethod,             //   /// Returns null if the method is not found.
		OtherMethod,             //   Method? findMethod(String methodName) {
		OtherMethod,             //     final szName = methodName.toNativeUtf16();
		OtherMethod,             //     final pmb = calloc<mdMethodDef>();
		OtherMethod,             //
		OtherMethod,             //     try {
		OtherMethod,             //       final hr = reader.FindMethod(token, szName, nullptr, 0, pmb);
		OtherMethod,             //       if (SUCCEEDED(hr)) {
		OtherMethod,             //         return Method.fromToken(reader, pmb.value);
		OtherMethod,             //       } else if (hr == CLDB_E_RECORD_NOTFOUND) {
		OtherMethod,             //         return null;
		OtherMethod,             //       } else {
		OtherMethod,             //         throw COMException(hr);
		OtherMethod,             //       }
		OtherMethod,             //     } finally {
		OtherMethod,             //       free(szName);
		OtherMethod,             //       free(pmb);
		OtherMethod,             //     }
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Gets the type referencing this type's superclass.
		OtherMethod,             //   TypeDef? get parent =>
		OtherMethod,             //       token == 0 ? null : TypeDef.fromToken(reader, baseTypeToken);
		BlankLine,               //
		OtherMethod,             //   String? getCustomGUIDAttribute(String guidAttributeName) {
		OtherMethod,             //     final ptrAttributeName = guidAttributeName.toNativeUtf16();
		OtherMethod,             //     final ppData = calloc<Pointer<BYTE>>();
		OtherMethod,             //     final pcbData = calloc<ULONG>();
		OtherMethod,             //
		OtherMethod,             //     try {
		OtherMethod,             //       final hr = reader.GetCustomAttributeByName(
		OtherMethod,             //           token, ptrAttributeName, ppData.cast(), pcbData);
		OtherMethod,             //       if (SUCCEEDED(hr)) {
		OtherMethod,             //         final blob = ppData.value;
		OtherMethod,             //         if (pcbData.value > 0) {
		OtherMethod,             //           final returnValue = blob.elementAt(2).cast<GUID>();
		OtherMethod,             //           return returnValue.ref.toString();
		OtherMethod,             //         }
		OtherMethod,             //       }
		OtherMethod,             //     } finally {
		OtherMethod,             //       free(ptrAttributeName);
		OtherMethod,             //       free(ppData);
		OtherMethod,             //       free(pcbData);
		OtherMethod,             //     }
		OtherMethod,             //   }
		BlankLine,               //
		OtherMethod,             //   /// Get the GUID for this type.
		OtherMethod,             //   ///
		OtherMethod,             //   /// Returns null if a GUID couldn't be found.
		OtherMethod,             //   String? get guid {
		OtherMethod,             //     var guid =
		OtherMethod,             //         getCustomGUIDAttribute('Windows.Foundation.Metadata.GuidAttribute');
		OtherMethod,             //     guid ??= getCustomGUIDAttribute('Windows.Win32.Interop.GuidAttribute');
		OtherMethod,             //
		OtherMethod,             //     return guid;
		OtherMethod,             //   }
		BlankLine,               //
		OverrideMethod,          //   @override
		OverrideMethod,          //   String toString() => 'TypeDef: $typeName';
		BlankLine,               //
	}

	runFullStylizer(t, opts, source, wantSource, [][]EntityType{want})
}
