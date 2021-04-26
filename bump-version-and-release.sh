#!/bin/bash -ex
go mod tidy
go generate ./...
go test ./...
go vet ./...
bump_version patch cmd/flutter-stylizer/main.go
./install.sh
git push
GOOS=linux GOARCH=amd64 go build -o ${HOME}/Downloads/flutter-stylizer-for-linux ./cmd/flutter-stylizer
GOOS=windows GOARCH=386 go build -o ${HOME}/Downloads/flutter-stylizer-for-windows.exe ./cmd/flutter-stylizer
GOOS=darwin GOARCH=amd64 go build -o ${HOME}/Downloads/flutter-stylizer-for-mac ./cmd/flutter-stylizer
