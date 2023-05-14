package main

import (
	"flag"
	"strings"
	"time"

	"github.com/STARRY-S/telebot/pkg/botcmd"
	"github.com/STARRY-S/telebot/pkg/config"
	"github.com/STARRY-S/telebot/pkg/reminder"
	"github.com/STARRY-S/telebot/pkg/user"
	"github.com/sirupsen/logrus"
	telebot "gopkg.in/telebot.v3"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Enable debug output")
	flag.Parse()
	if *debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug output enabled")
	}

	config.Init()
	pref := telebot.Settings{
		OnError: func(err error, c telebot.Context) {
			logrus.Error(err)
			c.Reply(err.Error())
		},
		Token:  config.GetApiToken(),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		errMsg := strings.Split(err.Error(), config.GetApiToken())
		if len(errMsg) > 1 {
			logrus.Fatalf("NewBot failed: %v", errMsg[1:])
		} else {
			logrus.Fatalf("NewBot failed: %v", err)
		}
	}

	for _, v := range config.Admins() {
		user.Register(v, user.LevelAdmin)
	}
	if owner := config.Owner(); owner != "" {
		id := config.OwnerID()
		if id == 0 {
			logrus.Fatal("failed to get owner UID")
		}
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

	reminder.Init(bot)

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
