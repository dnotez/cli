#!/bin/bash

echo "Building CLI ..."

export GOPATH=`pwd`
go get -d -v ./...

go clean
#go build -o bin/pl pl
go build -o bin/dz src/main.go

ls -lsh bin