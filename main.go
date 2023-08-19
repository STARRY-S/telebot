package main

import (
	"flag"
	"os"
	"strings"
	"time"

	"github.com/STARRY-S/telebot/pkg/botcmd"
	"github.com/STARRY-S/telebot/pkg/config"
	"github.com/STARRY-S/telebot/pkg/user"
	"github.com/STARRY-S/telebot/pkg/utils"
	"github.com/sirupsen/logrus"
	telebot "gopkg.in/telebot.v3"
)

func main() {
	cmd := flag.NewFlagSet("", flag.ExitOnError)
	debugFlag := cmd.Bool("debug", false, "Enable debug output")
	versionFlag := cmd.Bool("version", false, "Show version")
	cmd.Parse(os.Args[1:])
	if *versionFlag {
		logrus.Info(utils.GetVersion())
		os.Exit(0)
	}
	if *debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug output enabled")
	}

	config.Init()
	bot, err := telebot.NewBot(telebot.Settings{
		OnError: func(err error, c telebot.Context) {
			logrus.Error(err)
			c.Reply(err.Error())
		},
		Token:  config.GetApiToken(),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
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

	logrus.Info("Successfully started telebot.")
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
