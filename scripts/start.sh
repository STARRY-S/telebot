#!/bin/bash

cd $(dirname $0)/../

docker run -d \
    -e TELEGRAM_APITOKEN=$TELEGRAM_APITOKEN \
    -e HTTPS_PROXY=$HTTPS_PROXY \
    -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
    -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
    -v $(pwd)/config.yaml:/telebot/config.yaml \
    --restart=unless-stopped \
    --name=telebot \
    telebot
