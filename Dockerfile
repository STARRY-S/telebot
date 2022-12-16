FROM archlinux

WORKDIR /telebot

# Configure Arch Linux CN repository
RUN echo "Server = https://mirrors.bfsu.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN echo "Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist
RUN pacman --noconfirm -Syyu
RUN pacman --noconfirm -S go lm_sensors

# Configure GOPROXY
RUN go env -w GO111MODULE=on
RUN go env -w  GOPROXY=https://goproxy.cn,direct

# pre-copy/cache go.mod for pre-downloading dependencies
# and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/telebot .

CMD ["telebot"]
