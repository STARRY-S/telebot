FROM archlinux

# Configure Arch Linux CN repository & Install dependencies
RUN echo "Server = https://mirrors.bfsu.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist && \
    echo "Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist && \
    pacman --noconfirm -Syyu && \
    pacman --noconfirm -S lm_sensors words && \
    pacman --noconfirm -Scc

WORKDIR /telebot
COPY . .
# Install telebot
RUN cp config.yaml.example config.yaml && \
    mv telebot /usr/local/bin/

CMD ["telebot"]
