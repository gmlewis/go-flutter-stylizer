#!/bin/bash -ex
go mod tidy
go test ./...
go vet ./...
pushd cmd/flutter-stylizer
go install
popd
