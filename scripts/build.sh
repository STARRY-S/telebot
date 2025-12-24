#!/usr/bin/env bash

set -euo pipefail

cd $(dirname $0)/../
source scripts/version.sh

mkdir -p build
cd build

set -x

CGO_ENABLED=0 \
    go build -ldflags "${BUILD_FLAG:-}" ..

set +x

echo "--------------------------"
ls -alh telebot
echo "--------------------------"

echo "build: Done"
