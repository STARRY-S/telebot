#!/bin/bash

set -euo pipefail

cd $(dirname $0)/../

docker kill telebot &> /dev/null || true
docker rm telebot &> /dev/null || true

docker run -dit \
    -v $(pwd)/config.yaml:/telebot/config.yaml \
    --name=telebot \
    --restart=always \
	ghcr.io/starry-s/telebot:latest
