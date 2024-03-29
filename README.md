# flutter-stylizer - Mothballed/Archived Project

[![CodeQL](https://github.com/gmlewis/go-flutter-stylizer/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/gmlewis/go-flutter-stylizer/actions/workflows/codeql-analysis.yml)

`flutter-stylizer` is a stand-alone statically-linked Go executable that
organizes your Flutter classes in an opinionated and consistent manner.

It is based on the [Flutter Stylizer VSCode plugin](
https://github.com/gmlewis/flutter-stylizer) but is
written in Go instead of TypeScript.

This makes it particularly well-suited to be used in GitHub Actions,
Bitbucket, or other CI/CD pipelines, to maintain Flutter code repositories.

## Status

After 5 years, 119K installs, 4.9/5 stars, and zero sponsors, this project
is now mothballed and archived. Feel free to fork the project to create your
own version of the plugin and add your own customizations to your liking.

This has been run on the Flutter source code base (consisting of 8003 *.dart
files) located here (on Linux):
https://flutter.dev/docs/get-started/install/linux
and produces no further changes after the initial run.

```
$ cd ~/development
$ time tar xf ~/Downloads/flutter_linux_2.2.3-stable.tar.xz

real	0m24.000s
user	0m23.249s
sys	0m5.379s

$ time flutter-stylizer -q -w ./...

real	0m12.916s
user	0m15.289s
sys	0m1.202s

$ time flutter-stylizer -q -d ./...

real	0m11.828s
user	0m14.362s
sys	0m0.745s
```

## Installation

`flutter-stylizer` runs on Mac, Windows, and Linux.

### Option 1 (for non-Go devs or CI/CD pipelines)

* Download the appropriate precompiled statically-linked binary
  executable for your platform of the latest release, located here:
  https://github.com/gmlewis/go-flutter-stylizer/releases
* On Mac and Linux, make the file executable:
  * `chmod a+x flutter-stylizer`

### Option 2 (for Go devs or to build from source)

* One-time setup for using Go on your system:
  * Install Go from https://golang.org
  * Add two lines to your `${HOME}/.bashrc`:
    * `export GOPATH=${HOME}/go`
    * `PATH=${PATH}:${HOME}/go/bin`
  * Create a directory for the Go executables: `mkdir -p ${GOPATH}/bin`

* To install `flutter-stylizer`, run from a bash-like shell or terminal:

```
$ go get -u github.com/gmlewis/go-flutter-stylizer/cmd/flutter-stylizer
```

## Usage:

```
$ flutter-stylizer --help
...

Usage:
  flutter-stylizer [flags] [path ... | ./...]

Flags:
      --config string   config file (default is $HOME/.flutter-stylizer.yaml)
      --debug           dump insane levels of details to debug what is going on
  -d, --diff            display diffs (cannot be used with -l or -w); exit code 1 on diffs
  -h, --help            help for flutter-stylizer
  -l, --list            list files whose formatting differs from flutter-stylizer's (cannot be used with -d or -w); exit code 1 on diffs
  -q, --quiet           don't print unless there are errors
  -v, --verbose         write progress details to stderr
  -w, --write           write result to (source) file instead of stdout (cannot be used with -d or -l); exit code 0 on diffs
```

## Examples:

### Stylize all Dart files in directory tree

```
$ flutter-stylizer -w ./...
```

Note that this command will only write files when the stylizer finds
changes to make.

### Show stylizer differences for all Dart files in directory tree

```
$ flutter-stylizer -d ./...
```

Note that this command does not modify any files.
Also if any differences were found, it will exit with code 1
which is useful for breaking pipelines.

### List all Dart files with stylizer differences in directory tree

```
$ flutter-stylizer -l ./...
```

Note that this command does not modify any files.
Also if any differences were found, it will exit with code 1
which is useful for breaking pipelines.

### Show stylized output for a source file

```
$ flutter-stylizer lib/path/to/file.dart
```

Note that this command makes no changes to the file, but prints the
stylized output to stdout.

## Features

`flutter-stylizer` organizes the class(es) within a `*.dart` file
in the following manner (with a blank line separating these parts):

* The main (possibly factory) constructor is listed first, if it exists.
  - (`public-constructor` in configuration)
* Any named constructors are listed next, in sorted order.
  - (`named-constructors` in configuration)
* Any static (class) variables are listed next, in sorted order.
  - (`public-static-variables` in configuration)
* Any instance variables are listed next, in sorted order.
  - (`public-instance-variables` in configuration)
* Any `@override` variables are listed next, in sorted order.
  - (`public-override-variables` in configuration)
* Any private static (class) variables are listed next, in sorted order.
  - (`private-static-variables` in configuration)
* Any private instance variables are listed next, in sorted order.
  - (`private-instance-variables` in configuration)
* Any `@override` methods are listed next, in sorted order.
  - (`public-override-methods` in configuration)
* Any other methods are listed next in their original (unchanged) order.
  (As of version `v0.0.19`, two new flags affect this section; see below.)
  - (`public-other-methods` in configuration)
* If `private-other-methods` is (optionally) specified, these will be sorted
  separately from `public-other-methods`.
* The `build` method is listed last.
  - (`build-method` in configuration)

I have found that developer productivity increases when all code in
large projects follows a consistent and opinionated style.

Additionally, bringing new developers into a team with a large code base
is easier when the code is consistently written and therefore easier
to navigate and understand.

Without tooling to enforce a consistent style, developing code is less fun.
Having an automated tool to do this ugly work for you, however, makes
coding a lot more enjoyable, as you don't have to worry about the rules,
but can just run the plugin on file save, and the rules are automatically
enforced.

## Configuration

To override the default order of the stylizer, add a section to your
`${HOME}/.flutter-stylizer.yaml` file (or any `.yaml` file and use the
`--config` flag to point to it), like this:

```
groupAndSortGetterMethods: false
groupAndSortVariableTypes: false
memberOrdering:
  public-constructor
  named-constructors
  public-static-variables
  public-instance-variables
  public-override-variables
  private-static-variables
  private-instance-variables
  public-override-methods
  public-other-methods
  private-other-methods
  build-method
processEnumsLikeClasses: false
sortClassesWithinFile: false
sortOtherMethods: false
```

And then rearrange member names as desired.

Note that as of `v0.0.19`, two new flags were added to modify the
behavior of the "public-other-methods" as requested in #18:

- `groupAndSortGetterMethods` (default: `false`)
  - Whether to group getters separately (before 'public-other-methods')
    and sort them within their group.

- `sortOtherMethods` (default: `false`)
  - Whether to sort the 'public-other-methods' within their group.

As of `v0.1.5`, a new `private-other-methods` field was added. (See #3.)
If not specified, private methods will continue to be grouped within
the `public-other-methods` section.

As of `v0.1.8`, a new option flag was added:

- `groupAndSortVariableTypes` (default: `false`)
  - Whether to group public variables separately by type and sort
    them within their groups. Types are: "final", "optional" (`?`), and "normal".

As of `v0.1.12`, a new option flag was added:

- `sortClassesWithinFile` (default: `false`)
  - Whether to sort multiple classes within each file.

As of `v0.1.15`, a new option flag was added:
- `processEnumsLikeClasses` (default: `false`)
  - Whether to process enums identically to how classes are processed.

These features are experimental and should be used with caution.
Please file any bugs you find on the [GitHub issue tracker].

### Exclude files using file globbing

To exclude files from processing, add a section with one or more
[glob-style](https://pub.dev/packages/glob#syntax) patterns to your
`${HOME}/.flutter-stylizer.yaml` file (or any `.yaml` file and use the
`--config` flag to point to it), like this:

```
exclude:
  - "**/*.g.dart"
  - "**/*.gr.dart"
  - "**/*.freezed.dart"
```

## Limitations

This program does not have a full-featured Dart syntax tree parser.
As a result, it may come across Dart code that it doesn't handle properly.
See the [Known Issues](#known-issues) section below for more details.

## Reporting Problems

It is my goal to be able to use this program on large group projects, so
every attempt has been made to make this robust. If, however, problems
are found, please raise issues on the [GitHub issue tracker] for this repo
along with a (short) example demonstrating the "before" and "after" results
of running this plugin on the example code.

Even better, please submit a PR with your new "before"/"after" example coded-up
as a unit test along with the code to fix the problem, and I'll try to
incorporate the fix into the repo.

***Please remember to state which version of the program you are using, what platform/OS,
   and include your configuration settings!***

[GitHub issue tracker]: https://github.com/gmlewis/go-flutter-stylizer/issues

## Known Issues

* `flutter-stylizer` is line-oriented. It is meant to be run on code that
  is nicely separated by lines.  The `dartfmt` tool typically makes
  sane-looking code, and this is the type of code that is being targeted
  by this program.
* Code that follows the end of a multiline comment on the same
  line is not supported. Unusual code like this will most likely not ever be
  supported even though the Dart compiler can handle it.

## Release Notes

### v0.1.16

- Fix [issue #31](https://github.com/gmlewis/flutter-stylizer/issues/31) for private vars.

### v0.1.15

- Add new option:
  - `processEnumsLikeClasses` (default: `false`)

### v0.1.14

- Fix sort order of class names when `sortClassesWithinFile: true`.

### v0.1.13

- Fix [issue #8](https://github.com/gmlewis/go-flutter-stylizer/issues/8).

### v0.1.12

- Add new option [issue #8](https://github.com/gmlewis/go-flutter-stylizer/issues/8):
  - `sortClassesWithinFile` (default: `false`)

### v0.1.11

- Fix [issue #6](https://github.com/gmlewis/go-flutter-stylizer/issues/6).

### v0.1.8

- Add new option [issue #31](https://github.com/gmlewis/flutter-stylizer/issues/31):
  - `groupAndSortVariableTypes` (default: `false`)

### v0.1.7

- Fix [issue #26](https://github.com/gmlewis/flutter-stylizer/issues/26) caused by `Function()`.

### v0.1.6

- Process all `mixin` blocks in addition to all `class` blocks.

### v0.1.5

- `private-other-methods` can be added to the member ordering.

### v0.1.4

- Add "exclude" file globbing in the config file.

### v0.1.1

- Initial release. This version was run on the Flutter code base and
  should be ready for use.

----------------------------------------------------------------------

**Enjoy!**

----------------------------------------------------------------------

# License

Copyright 2020 Glenn M. Lewis. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
