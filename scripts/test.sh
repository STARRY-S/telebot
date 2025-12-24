#!/bin/bash

cd $(dirname $0)/../

set -exuo pipefail

go test -v -count=1 -timeout=5m ./...
