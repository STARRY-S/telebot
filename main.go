package main

import (
	"strings"
	"time"

	"github.com/STARRY-S/telebot/botcmd"
	"github.com/STARRY-S/telebot/config"
	"github.com/STARRY-S/telebot/user"
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

	botcmd.AddUserCommands(bot)
	botcmd.AddAdminCommands(bot)
	botcmd.AddOwnerCommands(bot)

	bot.Handle("/help", func(c telebot.Context) error {
		return c.Reply(
			HelpMessage(user.Find(c.Chat().Username).Level()),
			telebot.ModeDefault,
		)
	})

	logrus.Info("Start telebot.")
	bot.Start()
}

func HelpMessage(level user.Level) string {
	var msg string = ""
	if level >= user.LevelOwner {
		msg = botcmd.GetOwnerHelpMessage() + msg
	}
	if level >= user.LevelAdmin {
		msg = botcmd.GetAdminHelpMessage() + msg
	}
	msg = botcmd.GetUserHelpMessage() + msg

	return msg
}
