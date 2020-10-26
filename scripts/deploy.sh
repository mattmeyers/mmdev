#!/usr/bin/env bash

pushd "$(dirname "$0")" > /dev/null 2>&1

cd ..
hugo --minify && rsync -r public matt@172.104.30.170:/home/matt/www/mmdev

popd > /dev/null 2>&1
