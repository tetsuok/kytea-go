#!/bin/bash -x
set -o errexit
set -o nounset

PWD=`pwd`
PREFIX=/usr/local

git clone https://github.com/tetsuok/kytea libkytea
cd libkytea
. ci_config.sh
bash ./setup_gtest.sh
bash ./gyp/install_gyp.sh
bash ./build_with_gyp.sh Release make

# Copying a shared library
SHARED_EXT=so
if [[ "$TARGET_OS" = "Darwin" ]]; then
  SHARED_EXT=dylib
fi
SHARED_LIB=libkytea.$SHARED_EXT

SHARED_PATH=out_unix/Release/lib.target
if [[ "$TARGET_OS" = "Darwin" ]]; then
  echo "OS X is not supported yet"
  exit 1
fi
sudo cp $SHARED_PATH/$SHARED_LIB $PREFIX/lib

# Copying headers
if ! [[ -e $PREFIX/include ]]; then
  sudo mkdir -p $PREFIX/include
fi
sudo cp -fr src/include/kytea $PREFIX/include

ln -sf $PWD/data/model.bin ../sample
cd ..
