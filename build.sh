#!/bin/bash

echo "Building CLI ..."

OUT=/tmp/dnotez-cli
mkdir -p ${OUT}
rm -fr ${OUT}/*
export GOPATH=${OUT}
go get -d -v ./...

go clean
#go build -o bin/pl pl
go build -o ${OUT}/bin/dz main.go

ls -lsh bin