#!/bin/bash

set -euo pipefail

cd $(dirname $0)/../
source scripts/version.sh

TAG="${VERSION:-latest}"

set -x
docker build \
    --build-arg GITCOMMIT=${GITCOMMIT} \
    --build-arg VERSION=${VERSION} \
    -f package/Dockerfile \
    -t hxstarrys/telebot:${TAG} .
