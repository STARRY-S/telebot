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

## License

> MIT
