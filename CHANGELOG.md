# Change Log

All notable changes to `flutter-stylizer` will be documented in
this file.

Please make sure you have a backup (preferably in git) of your code before running
`flutter-stylizer` in case it doesn't handle your code properly.


## [v0.1.16] - 2022-06-28

- Fix [issue #31](https://github.com/gmlewis/flutter-stylizer/issues/31) for private vars.

## [v0.1.15] - 2022-06-27

- Add new option:
  - `processEnumsLikeClasses` (default: `false`)

## [v0.1.14] - 2022-06-23

- Fix sort order of class names when `sortClassesWithinFile: true`.

## [v0.1.13] - 2022-06-23

- Fix [issue #8](https://github.com/gmlewis/go-flutter-stylizer/issues/8).

## [v0.1.12] - 2022-06-22

- Add new option [issue #8](https://github.com/gmlewis/go-flutter-stylizer/issues/8):
  - `sortClassesWithinFile` (default: `false`)

## [v0.1.11] - 2022-06-22

- Fix [issue #6](https://github.com/gmlewis/go-flutter-stylizer/issues/6).

## [v0.1.8] - 2022-06-14

- Add new option [issue #31](https://github.com/gmlewis/flutter-stylizer/issues/31):
  - `groupAndSortVariableTypes` (default: `false`)

## [v0.1.7] - 2022-01-12

- Fix [issue #26](https://github.com/gmlewis/flutter-stylizer/issues/26) caused by `Function()`.

## [v0.1.6] - 2022-01-08

- Process all `mixin` blocks in addition to all `class` blocks.

## [v0.1.5] - 2021-12-12

- `private-other-methods` can be added to the member ordering.

## [v0.1.4] - 2021-11-03

- Add "exclude" file globbing in the config file.

## [v0.1.1] - 2021-04-26

- Initial release. This version was successfully run on the Flutter code base
  and should be ready for use.
