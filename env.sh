#!/bin/bash
PWD=`pwd`
export CGO_LDFLAGS="-L$PWD"
export CGO_CFLAGS="-I$PWD"
