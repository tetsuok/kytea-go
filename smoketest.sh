#!/bin/bash -x
PWD=`pwd`
export CGO_LDFLAGS="-L$PWD"
cd sample
go build
./sample
cd ..
