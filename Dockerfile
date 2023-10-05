FROM archlinux

ENV http_proxy=${http_proxy} https_proxy=${https_proxy} no_proxy=${no_proxy}

# Install dependencies
RUN pacman --noconfirm -Syyu && \
    pacman --noconfirm -S lm_sensors words openssh pacman-contrib && \
    pacman --noconfirm -Scc

WORKDIR /telebot
COPY config.yaml.example telebot ./

# Install telebot
RUN cp config.yaml.example config.yaml && \
    mv telebot /usr/local/bin/

CMD ["telebot"]
