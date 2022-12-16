#!/bin/bash

cd $(dirname $0)/../

docker run -d \
    -e TELEGRAM_APITOKEN=$TELEGRAM_APITOKEN \
    -e HTTPS_PROXY=$HTTPS_PROXY \
    -v $(pwd)/config.yaml:/telebot/config.yaml \
    --restart=always \
    --name=telebot \
    telebot
