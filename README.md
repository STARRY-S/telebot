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

1. Build telebot in docker container.

    ```bash
    make image
    ```

2. Run telebot container image.

    ```bash
    # Update config (Setup API tokens & ADMIN, Owner users)
    cp config.yaml.example config.yaml
    vim config.yaml

    # Run telebot docker container image
    make run

    # You can specify the `HTTP_PROXY` and `HTTPS_PROXY` env if needed
    HTTP_RPOXY='http://127.0.0.1:8000' \
        HTTPS_PROXY='http://127.0.0.1:8000' \
        NO_PROXY='127.0.0.1' \
        make run

    # View logs
    docker logs -f telebot
    ```

## License

MIT License

Copyright (c) 2025 STARRY-S

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
