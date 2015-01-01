#!/bin/bash

echo "Building CLI ..."

export GOPATH=`pwd`

go get

go build -o bin/pl pl
go build -o bin/dz main