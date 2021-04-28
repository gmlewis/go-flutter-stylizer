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

//go:embed testfiles/pubspec.dart.txt
var pubspec_dart_txt string

//go:embed testfiles/pubspec_want.txt
var pubspec_want_txt string

func TestPubspec_GetClasses(t *testing.T) {
	e, err := NewEditor(pubspec_dart_txt, false)
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	classes, err := e.GetClasses(false)
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 1; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}
}

func TestPubspec_Class1(t *testing.T) {
	source := pubspec_dart_txt
	wantSource := pubspec_want_txt

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,          // line #14: {
		InstanceVariable, // line #15:   // TODO: executables
		InstanceVariable, // line #16:
		InstanceVariable, // line #17:   final String name;
		BlankLine,        // line #18:
		InstanceVariable, // line #19:   @JsonKey(fromJson: _versionFromString)
		InstanceVariable, // line #20:   final Version version;
		BlankLine,        // line #21:
		InstanceVariable, // line #22:   final String description;
		BlankLine,        // line #23:
		InstanceVariable, // line #24:   /// This should be a URL pointing to the website for the package.
		InstanceVariable, // line #25:   final String homepage;
		BlankLine,        // line #26:
		InstanceVariable, // line #27:   /// Specifies where to publish this package.
		InstanceVariable, // line #28:   ///
		InstanceVariable, // line #29:   /// Accepted values: `null`, `'none'` or an `http` or `https` URL.
		InstanceVariable, // line #30:   ///
		InstanceVariable, // line #31:   /// [More information](https://dart.dev/tools/pub/pubspec#publish_to).
		InstanceVariable, // line #32:   final String publishTo;
		BlankLine,        // line #33:
		InstanceVariable, // line #34:   /// Optional field to specify the source code repository of the package.
		InstanceVariable, // line #35:   /// Useful when a package has both a home page and a repository.
		InstanceVariable, // line #36:   final Uri repository;
		BlankLine,        // line #37:
		InstanceVariable, // line #38:   /// Optional field to a web page where developers can report new issues or
		InstanceVariable, // line #39:   /// view existing ones.
		InstanceVariable, // line #40:   final Uri issueTracker;
		BlankLine,        // line #41:
		OtherMethod,      // line #42:   /// If there is exactly 1 value in [authors], returns it.
		OtherMethod,      // line #43:   ///
		OtherMethod,      // line #44:   /// If there are 0 or more than 1, returns `null`.
		OtherMethod,      // line #45:   @Deprecated(
		OtherMethod,      // line #46:       'Here for completeness, but not recommended. Use `authors` instead.')
		OtherMethod,      // line #47:   String get author {
		OtherMethod,      // line #48:     if (authors.length == 1) {
		OtherMethod,      // line #49:       return authors.single;
		OtherMethod,      // line #50:     }
		OtherMethod,      // line #51:     return null;
		OtherMethod,      // line #52:   }
		BlankLine,        // line #53:
		InstanceVariable, // line #54:   final List<String> authors;
		InstanceVariable, // line #55:   final String documentation;
		BlankLine,        // line #56:
		InstanceVariable, // line #57:   @JsonKey(fromJson: _environmentMap)
		InstanceVariable, // line #58:   final Map<String, VersionConstraint> environment;
		BlankLine,        // line #59:
		InstanceVariable, // line #60:   @JsonKey(fromJson: parseDeps, nullable: false)
		InstanceVariable, // line #61:   final Map<String, Dependency> dependencies;
		BlankLine,        // line #62:
		InstanceVariable, // line #63:   @JsonKey(fromJson: parseDeps, nullable: false)
		InstanceVariable, // line #64:   final Map<String, Dependency> devDependencies;
		BlankLine,        // line #65:
		InstanceVariable, // line #66:   @JsonKey(fromJson: parseDeps, nullable: false)
		InstanceVariable, // line #67:   final Map<String, Dependency> dependencyOverrides;
		BlankLine,        // line #68:
		InstanceVariable, // line #69:   /// Optional configuration specific to [Flutter](https://flutter.io/)
		InstanceVariable, // line #70:   /// packages.
		InstanceVariable, // line #71:   ///
		InstanceVariable, // line #72:   /// May include
		InstanceVariable, // line #73:   /// [assets](https://flutter.io/docs/development/ui/assets-and-images)
		InstanceVariable, // line #74:   /// and other settings.
		InstanceVariable, // line #75:   final Map<String, dynamic> flutter;
		BlankLine,        // line #76:
		MainConstructor,  // line #77:   /// If [author] and [authors] are both provided, their values are combined
		MainConstructor,  // line #78:   /// with duplicates eliminated.
		MainConstructor,  // line #79:   Pubspec(
		MainConstructor,  // line #80:     this.name, {
		MainConstructor,  // line #81:     this.version,
		MainConstructor,  // line #82:     this.publishTo,
		MainConstructor,  // line #83:     String author,
		MainConstructor,  // line #84:     List<String> authors,
		MainConstructor,  // line #85:     Map<String, VersionConstraint> environment,
		MainConstructor,  // line #86:     this.homepage,
		MainConstructor,  // line #87:     this.repository,
		MainConstructor,  // line #88:     this.issueTracker,
		MainConstructor,  // line #89:     this.documentation,
		MainConstructor,  // line #90:     this.description,
		MainConstructor,  // line #91:     Map<String, Dependency> dependencies,
		MainConstructor,  // line #92:     Map<String, Dependency> devDependencies,
		MainConstructor,  // line #93:     Map<String, Dependency> dependencyOverrides,
		MainConstructor,  // line #94:     this.flutter,
		MainConstructor,  // line #95:   })  : authors = _normalizeAuthors(author, authors),
		MainConstructor,  // line #96:         environment = environment ?? const {},
		MainConstructor,  // line #97:         dependencies = dependencies ?? const {},
		MainConstructor,  // line #98:         devDependencies = devDependencies ?? const {},
		MainConstructor,  // line #99:         dependencyOverrides = dependencyOverrides ?? const {} {
		MainConstructor,  // line #100:     if (name == null || name.isEmpty) {
		MainConstructor,  // line #101:       throw ArgumentError.value(name, 'name', '"name" cannot be empty.');
		MainConstructor,  // line #102:     }
		MainConstructor,  // line #103:
		MainConstructor,  // line #104:     if (publishTo != null && publishTo != 'none') {
		MainConstructor,  // line #105:       try {
		MainConstructor,  // line #106:         final targetUri = Uri.parse(publishTo);
		MainConstructor,  // line #107:         if (!(targetUri.isScheme('http') || targetUri.isScheme('https'))) {
		MainConstructor,  // line #108:           throw const FormatException('Must be an http or https URL.');
		MainConstructor,  // line #109:         }
		MainConstructor,  // line #110:       } on FormatException catch (e) {
		MainConstructor,  // line #111:         throw ArgumentError.value(publishTo, 'publishTo', e.message);
		MainConstructor,  // line #112:       }
		MainConstructor,  // line #113:     }
		MainConstructor,  // line #114:   }
		BlankLine,        // line #115:
		NamedConstructor, // line #116:   factory Pubspec.fromJson(Map json, {bool lenient = false}) {
		NamedConstructor, // line #117:     lenient ??= false;
		NamedConstructor, // line #118:
		NamedConstructor, // line #119:     if (lenient) {
		NamedConstructor, // line #120:       while (json.isNotEmpty) {
		NamedConstructor, // line #121:         // Attempting to remove top-level properties that cause parsing errors.
		NamedConstructor, // line #122:         try {
		NamedConstructor, // line #123:           return _$PubspecFromJson(json);
		NamedConstructor, // line #124:         } on CheckedFromJsonException catch (e) {
		NamedConstructor, // line #125:           if (e.map == json && json.containsKey(e.key)) {
		NamedConstructor, // line #126:             json = Map.from(json)..remove(e.key);
		NamedConstructor, // line #127:             continue;
		NamedConstructor, // line #128:           }
		NamedConstructor, // line #129:           rethrow;
		NamedConstructor, // line #130:         }
		NamedConstructor, // line #131:       }
		NamedConstructor, // line #132:     }
		NamedConstructor, // line #133:
		NamedConstructor, // line #134:     return _$PubspecFromJson(json);
		NamedConstructor, // line #135:   }
		BlankLine,        // line #136:
		NamedConstructor, // line #137:   /// Parses source [yaml] into [Pubspec].
		NamedConstructor, // line #138:   ///
		NamedConstructor, // line #139:   /// When [lenient] is set, top-level property-parsing or type cast errors are
		NamedConstructor, // line #140:   /// ignored and `null` values are returned.
		NamedConstructor, // line #141:   factory Pubspec.parse(String yaml, {sourceUrl, bool lenient = false}) {
		NamedConstructor, // line #142:     lenient ??= false;
		NamedConstructor, // line #143:
		NamedConstructor, // line #144:     return checkedYamlDecode(
		NamedConstructor, // line #145:         yaml, (map) => Pubspec.fromJson(map, lenient: lenient),
		NamedConstructor, // line #146:         sourceUrl: sourceUrl);
		NamedConstructor, // line #147:   }
		BlankLine,        // line #148:
		OtherMethod,      // line #149:   static List<String> _normalizeAuthors(String author, List<String> authors) {
		OtherMethod,      // line #150:     final value = <String>{};
		OtherMethod,      // line #151:     if (author != null) {
		OtherMethod,      // line #152:       value.add(author);
		OtherMethod,      // line #153:     }
		OtherMethod,      // line #154:     if (authors != null) {
		OtherMethod,      // line #155:       value.addAll(authors);
		OtherMethod,      // line #156:     }
		OtherMethod,      // line #157:     return value.toList();
		OtherMethod,      // line #158:   }
		BlankLine,        // line #159:
	}

	runFullStylizer(t, opts, source, wantSource, want)
	runFullStylizer(t, opts, wantSource, wantSource, nil)
}
