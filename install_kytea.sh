#!/bin/bash -x
set -o errexit
set -o nounset
#git clone https://github.com/tetsuok/kytea libkytea
cd libkytea
. ci_config.sh
bash ./setup_gtest.sh
bash ./build_with_gyp.sh Release make
if [[ "$TARGET_OS" = "Darwin" ]]; then
  cp out_mac/Release/libkytea.a ..
else
  cp out_unix/Release/libkytea.a ..
fi
cp -fr src/include/kytea .

ln -sf `pwd`/data/model.bin ../sample
