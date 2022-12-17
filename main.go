package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/STARRY-S/telebot/config"
	"github.com/STARRY-S/telebot/status"
	"github.com/STARRY-S/telebot/user"
	"github.com/STARRY-S/telebot/utils"
	"github.com/sirupsen/logrus"
	telebot "gopkg.in/telebot.v3"
)

func main() {
	config.Init()
	pref := telebot.Settings{
		Token:  config.GetApiToken(),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		errMsg := strings.Split(err.Error(), config.GetApiToken())
		if len(errMsg) > 1 {
			logrus.Fatal(errMsg[1:])
		} else {
			logrus.Fatal(err)
		}
	}

	for _, v := range config.Admins() {
		user.Register(v, user.LevelAdmin)
	}
	if owner := config.Owner(); owner != "" {
		user.Register(owner, user.LevelOwner)
	}

	bot.Handle("/hello", func(c telebot.Context) error {
		return c.Reply("Hi!")
	})

	bot.Handle("/ping", func(c telebot.Context) error {
		return c.Reply("Pong!")
	})

	bot.Handle("/status", func(c telebot.Context) error {
		// only available in private chat
		switch c.Chat().Type {
		case telebot.ChatChannelPrivate:
		case telebot.ChatPrivate:
		default:
			return nil
		}

		if !user.FindUser(c.Chat().Username).IsAdmin() {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		status, err := status.GetStatus()
		if err != nil {
			return err
		}
		return c.Reply(status)
	})

	bot.Handle("/add_admin", func(c telebot.Context) error {
		if !user.FindUser(c.Chat().Username).IsOwner() {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /addAdmin <user>")
		}
		if !user.FindUser(args[0]).IsAdmin() {
			err := user.Register(args[0], user.LevelAdmin)
			if err != nil {
				logrus.Errorf("addAdmin failed: %v", err)
				return c.Reply("failed")
			}
		}
		return nil
	})

	bot.Handle("/users", func(c telebot.Context) error {
		// only available in private chat
		switch c.Chat().Type {
		case telebot.ChatChannelPrivate:
		case telebot.ChatPrivate:
		default:
			return nil
		}
		if !user.FindUser(c.Chat().Username).IsOwner() {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		users := user.Users()
		if users == "" {
			return c.Reply("failed")
		}
		return c.Reply(users)
	})

	bot.Handle("/help", func(c telebot.Context) error {
		return c.Reply(HelpMessage(user.FindUser(c.Chat().Username).Level()))
	})

	bot.Handle("/start", func(c telebot.Context) error {
		if !user.FindUser(c.Chat().Username).IsUser() {
			err := user.Register(c.Chat().Username, user.LevelAdmin)
			if err != nil {
				logrus.Errorf("/start failed: %v", err)
				return nil
			}
			return c.Reply("Registered")
		}

		return c.Reply("Already registered")
	})

	logrus.Info("Start telebot.")
	bot.Start()
}

func HelpMessage(level user.Level) string {
	buff := &bytes.Buffer{}
	if level >= user.LevelOwner {
		fmt.Fprintln(buff, "/users Get users")
		fmt.Fprintln(buff, "/add_admin Register admin user")
	}
	if level >= user.LevelAdmin {
		fmt.Fprintln(buff, "/status Get system status")
	}
	if level >= user.LevelUser {
	}
	fmt.Fprintln(buff, "/hello Say hello")
	fmt.Fprintln(buff, "/ping Ping")
	fmt.Fprintln(buff, "/help Show this message")

	return buff.String()
}
