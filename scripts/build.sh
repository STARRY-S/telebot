#!/usr/bin/env bash

set -euo pipefail

cd $(dirname $0)/../

if [[ ! -e "config.yaml" ]]; then
    cp config.yaml.example config.yaml
fi

GITCOMMIT=$(git rev-parse HEAD || true)
VERSION=$(git describe --tags 2>/dev/null || true)
BUILD_FLAG=""
if [[ "${GITCOMMIT}" != "UNKNOW" ]]; then
    BUILD_FLAG="${BUILD_FLAG} -X 'github.com/STARRY-S/telebot/pkg/utils.gitCommit=${GITCOMMIT}'"
fi
BUILD_FLAG="${BUILD_FLAG} -X 'github.com/STARRY-S/telebot/pkg/utils.version=${VERSION}'"

CGO_ENABLED=0 go build -ldflags "${BUILD_FLAG}" .

echo "--------------------------"
ls -alh telebot
echo "--------------------------"

TAG=""
if [[ ! -z "$VERSION" ]]; then
    TAG=":${VERSION}"
fi

docker build \
    --build-arg http_proxy=${http_proxy:-} \
    --build-arg https_proxy=${https_proxy:-} \
    --build-arg HTTP_PROXY=${HTTP_PROXY:-} \
    --build-arg HTTPS_PROXY=${HTTPS_PROXY:-} \
    --build-arg no_proxy=${no_proxy:-} \
    --build-arg NO_PROXY=${NO_PROXY:-} \
    --network=host \
    -t hxstarrys/telebot${TAG} .
