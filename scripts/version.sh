#!/bin/bash

GITCOMMIT=${GITCOMMIT:-$(git rev-parse HEAD 2> /dev/null || echo -n head)}
VERSION=${VERSION:-$(git describe --tags 2>/dev/null || echo -n unknown)}

BUILD_FLAG="-s -w"

if [[ ! -z "${GITCOMMIT}" ]]; then
    BUILD_FLAG="${BUILD_FLAG:-} -X 'github.com/STARRY-S/telebot/pkg/utils.gitCommit=${GITCOMMIT}'"
fi
if [[ ! -z "${GITCOMMIT}" ]]; then
    BUILD_FLAG="${BUILD_FLAG:-} -X 'github.com/STARRY-S/telebot/pkg/utils.version=${VERSION}'"
fi
