#!/bin/bash -x
PWD=`pwd`
export CGO_LDFLAGS="-L$PWD"
export CGO_CFLAGS="-I$PWD"
cd sample
go build
./sample
cd ..
