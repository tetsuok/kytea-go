#!/bin/bash -x
set -o errexit
set -o nounset

PWD=`pwd`

git clone https://github.com/tetsuok/kytea libkytea
cd libkytea
. ci_config.sh
bash ./setup_gtest.sh
bash ./gyp/install_gyp.sh
bash ./build_with_gyp.sh Release make
if [[ "$TARGET_OS" = "Darwin" ]]; then
  cp out_mac/Release/libkytea.a ..
else
  cp out_unix/Release/obj.target/gyp/libkytea.a ..
fi
cp -fr src/include/kytea .

ln -sf $PWD/data/model.bin ../sample
cd ..
