#!/usr/bin/env bash

set -euo pipefail

cd $(dirname $0)/../
source scripts/version.sh

TAG="${VERSION:-latest}"

# Cleanup exist telebot instance.
docker kill telebot &> /dev/null || true
docker rm telebot &> /dev/null || true

docker run -d \
    -e TELEGRAM_APITOKEN=${TELEGRAM_APITOKEN:-} \
    -e HTTP_PROXY=${HTTP_PROXY:-} \
    -e http_proxy=${http_proxy:-} \
    -e HTTPS_PROXY=${HTTPS_PROXY:-} \
    -e https_proxy=${https_proxy:-} \
    -e NO_PROXY=${NO_PROXY:-} \
    -e no_proxy=${no_proxy:-} \
    -v $(pwd)/config.yaml:/telebot/config.yaml \
    --network=host \
    --restart=always \
    --name=telebot \
    hxstarrys/telebot:${TAG}
