#!/bin/bash -ex
go test ./...
go vet ./...
pushd cmd/flutter-stylizer
go install
popd
