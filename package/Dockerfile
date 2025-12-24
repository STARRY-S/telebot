FROM golang:1.25 AS build

WORKDIR /telebot

# Pre-download dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

# Build the binary
ENV GOCACHE=/root/.cache/go-build
ARG TAG=v0.1.0-alpha
ARG COMMIT=unknown
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 \
    go build \
        -buildmode=pie \
        -ldflags="-extldflags='-static' -s -w -X github.com/STARRY-S/telebot/pkg/utils.Version=${TAG} -X github.com/STARRY-S/telebot/pkg/utils.Commit=${COMMIT}" \
        -o /usr/local/bin/telebot \
        .
RUN telebot -version

FROM archlinux

RUN pacman --noconfirm -Syyu && \
    pacman --noconfirm -S lm_sensors words openssh ffmpeg && \
    pacman --noconfirm -Scc

WORKDIR /telebot
COPY config.yaml.example ./config.yaml
COPY --from=build /usr/local/bin/telebot /usr/local/bin

ENTRYPOINT [ "telebot" ]
