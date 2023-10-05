#!/usr/bin/env bash

set -euo pipefail

docker kill telebot || true
docker rm telebot || true

VERSION=$(git describe --tags 2>/dev/null || true)
TAG=""
if [[ ! -z "$VERSION" ]]; then
    TAG=":${VERSION}"
fi

docker image rm hxstarrys/telebot${TAG}

exit 0
