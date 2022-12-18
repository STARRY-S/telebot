package botcmd

import (
	"bytes"
	"fmt"

	"github.com/STARRY-S/telebot/utils"
	"gopkg.in/telebot.v3"
)

func AddUserCommands(bot *telebot.Bot) {
	bot.Handle("/hello", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		return c.Reply("Hi")
	})

	bot.Handle("/ping", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		return c.Reply("Pong")
	})

	bot.Handle("/start", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		return c.Reply(
			fmt.Sprintf("Hi @%s, glad to see you!", c.Chat().Username),
		)
	})
}

func GetUserHelpMessage() string {
	b := &bytes.Buffer{}
	fmt.Fprintln(b, "/hello Say hello")
	fmt.Fprintln(b, "/ping Ping")
	fmt.Fprintln(b, "/help Show this message")
	return b.String()
}
