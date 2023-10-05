## Telebot

This is my personal telebot, it has some common used functions by myself (password generator, base64 & SHA calculator, UID query, GIF downloader, etc...), API based on [go-telebot/telebot](https://github.com/go-telebot/telebot).

You can try this telebot via [Mikoto_bot](https://t.me/Mikoto_bot).

----

Available commands:

```text
/hello Say hello
/ping Ping
/sha256 Calculate sha256sum
/md5 Calculate md5sum
/base64 Calculate base64
/decode_base64 Decode base64
/genpasswd Generate password (-h to get more info)
/my_uid Get my telegram UID
/help Show this message

Administrators only:
/status Get system status (Private) (Admin)
/admins Get admins list (Private) (Admin)

Owner only:
/add_admin Register temporary admin user (Owner)
/del_admin Remove admin temporally (Owner)
/exec Run commands (Private) (Owner)
/restart Restart (kill) this bot (Owner)
```

Administrator users also have permission to convert Video (GIF Sticker) to GIF zip file.

## Usage

1. Build binary and create docker image:

    ```bash
    make build

    # You can specify the `HTTP_PROXY` and `HTTPS_PROXY` env if needed
    HTTP_RPOXY='http://127.0.0.1:8000' \
        HTTPS_PROXY='http://127.0.0.1:8000' \
        NO_PROXY='127.0.0.1' \
        make build
    ```

2. Setup the config file & Run telebot in container image:

    ```bash
    # Update config (Setup API tokens & ADMIN, Owner users)
    vim config.yaml

    # Run telebot in container
    make run

    # You can specify the `HTTP_PROXY` and `HTTPS_PROXY` env if needed
    HTTP_RPOXY='http://127.0.0.1:8000' \
        HTTPS_PROXY='http://127.0.0.1:8000' \
        NO_PROXY='127.0.0.1' \
        make run

    # View logs
    docker logs -f telebot
    ```

3. Destroy and clean-up resources:

    ```bash
    # kill & delete telebot container image
    make release
    ```

## License

> MIT
