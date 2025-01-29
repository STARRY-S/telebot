FROM golang AS build

ARG GITCOMMIT
ARG VERSION

WORKDIR /app
# Pre-download go modules
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Build source
COPY . .
RUN ./scripts/build.sh

FROM archlinux

WORKDIR /telebot

# Install dependencies
RUN pacman --noconfirm -Syyu && \
    pacman --noconfirm -S lm_sensors words openssh ffmpeg && \
    pacman --noconfirm -Scc

COPY --from=build /app/build/telebot /usr/local/bin
COPY config.yaml.example ./config.yaml

CMD ["telebot"]
