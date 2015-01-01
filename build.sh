#!/bin/bash

echo "Building CLI ..."

export GOPATH=`pwd`
cd src
go get -v
cd -

#go build -o bin/pl pl
go build -o bin/dz src/main.go