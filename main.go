package main

import (
	"time"

	"github.com/STARRY-S/telebot/config"
	"github.com/STARRY-S/telebot/status"
	"github.com/sirupsen/logrus"
	telebot "gopkg.in/telebot.v3"
)

func main() {
	pref := telebot.Settings{
		Token:  config.GetApiToken(),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	bot.Handle("/hello", func(c telebot.Context) error {
		return c.Send("Hi!")
	})

	bot.Handle("/ping", func(c telebot.Context) error {
		return c.Reply("Pong!")
	})

	bot.Handle("/status", func(c telebot.Context) error {
		if !config.IsAdmin(c.Chat().Username) {
			return nil
		}
		// only available in private chat
		switch c.Chat().Type {
		case telebot.ChatChannelPrivate:
		case telebot.ChatPrivate:
		default:
			return nil
		}

		status, err := status.GetStatus()
		if err != nil {
			return err
		}
		return c.Reply(status)
	})

	bot.Handle("/help", func(c telebot.Context) error {
		return c.Reply(getHelpMessage())
	})

	logrus.Info("Start telebot.")
	bot.Start()
}

func getHelpMessage() string {
	return `
/status Get status of system (admin only).
/hello Say hello.
/ping Ping.
/help Show this message.`
}
