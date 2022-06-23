#!/bin/bash -ex

./test-all.sh

go run cmd/bump-version/main.go patch cli/cli.go

export VERSION=$(grep 'const VERSION' cli/cli.go | sed -e 's/^.* //' -e "s/'//g" -e 's/"//g')
git commit -sam "Version ${VERSION}"

# Re-tag the updated version:
git tag -f "${VERSION}"

./install.sh
git push && git push --tags

GOOS=linux GOARCH=amd64 go build -o ${HOME}/Downloads/flutter-stylizer-for-linux ./cmd/flutter-stylizer
GOOS=windows GOARCH=386 go build -o ${HOME}/Downloads/flutter-stylizer-for-windows.exe ./cmd/flutter-stylizer
GOOS=darwin GOARCH=amd64 go build -o ${HOME}/Downloads/flutter-stylizer-for-mac ./cmd/flutter-stylizer
