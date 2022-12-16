#!/bin/bash

cd $(dirname $0)/../

if [[ ! -e "config.yaml" ]]; then
    cp config.yaml.example config.yaml
fi

docker build -t telebot .
