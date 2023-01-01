package botcmd

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"strings"

	"github.com/STARRY-S/telebot/utils"
	"github.com/STARRY-S/telebot/utils/passwd"
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

	bot.Handle("/sha256", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		if len(c.Args()) == 0 {
			return c.Reply("Usage: /sha256 <text>")
		}
		text := strings.TrimLeft(c.Text(), "/sha256 ")
		sum := sha256.Sum256([]byte(text))
		return c.Reply(
			fmt.Sprintf("`%x`", sum),
			telebot.ModeMarkdownV2,
		)
	})

	bot.Handle("/md5", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		if len(c.Args()) == 0 {
			return c.Reply("Usage: /md5 <text>")
		}
		text := strings.TrimLeft(c.Text(), "/md5 ")
		return c.Reply(
			fmt.Sprintf("`%x`", md5.Sum([]byte(text))),
			telebot.ModeMarkdownV2,
		)
	})

	bot.Handle("/base64", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		if len(c.Args()) == 0 {
			return c.Reply("Usage: /base64 <text>")
		}
		text := strings.TrimLeft(c.Text(), "/base64 ")
		str := base64.StdEncoding.EncodeToString([]byte(text))
		return c.Reply(
			fmt.Sprintf("`%s`", str),
			telebot.ModeMarkdownV2,
		)
	})

	bot.Handle("/decode_base64", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		if len(c.Args()) == 0 {
			return c.Reply("Usage: /decode_base64 <text>")
		}
		text := strings.TrimLeft(c.Text(), "/decode_base64 ")
		data, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			return c.Reply(fmt.Sprintf("failed: %s", err.Error()))
		}
		return c.Reply(
			fmt.Sprintf("`%s`", data),
			telebot.ModeMarkdownV2,
		)
	})

	bot.Handle("/genpasswd", func(c telebot.Context) error {
		if !isUser(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		cmdArgs := c.Args()
		var (
			passwdLength     int
			passwdWords      int
			passwdType       string
			passwdHasNum     bool
			passwdHasSpecial bool
		)

		cmd := flag.NewFlagSet("", flag.ContinueOnError)
		cmdOutput := bytes.NewBuffer(nil)
		cmd.SetOutput(cmdOutput)
		cmd.IntVar(
			&passwdLength,
			"len",
			16,
			"password length (random mode)")
		cmd.IntVar(
			&passwdWords,
			"words",
			4,
			"password words number (word mode)")
		cmd.StringVar(
			&passwdType,
			"type",
			"random",
			"password type ('random', 'word')")
		cmd.BoolVar(
			&passwdHasNum,
			"has-num",
			true,
			"password has number")
		cmd.BoolVar(
			&passwdHasSpecial,
			"has-special",
			true,
			"password has special characters")
		cmd.Parse(cmdArgs)
		if cmdOutput.String() != "" {
			return c.Reply(
				fmt.Sprintf("```\n%s\n```", cmdOutput.String()),
				telebot.ModeMarkdownV2,
			)
		}
		switch passwdType {
		case "random":
			s := passwd.GenRandomPasswd(
				passwdLength, passwdHasNum, passwdHasSpecial)
			return c.Reply(
				fmt.Sprintf("`%s`", s),
				telebot.ModeMarkdownV2,
			)
		case "word":
			s := passwd.GenRememberablePasswd(passwdWords)
			return c.Reply(
				fmt.Sprintf("`%s`", s),
				telebot.ModeMarkdownV2,
			)
		}

		cmd.Usage()
		return c.Reply(
			fmt.Sprintf("```\n%s\n```", cmdOutput.String()),
			telebot.ModeMarkdownV2,
		)
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
	fmt.Fprintln(b, "/sha256 Calculate sha256sum")
	fmt.Fprintln(b, "/md5 Calculate md5sum")
	fmt.Fprintln(b, "/base64 Calculate base64")
	fmt.Fprintln(b, "/decode_base64 Decode base64")
	fmt.Fprintln(b, "/genpasswd Generate password (-h to get more info)")
	fmt.Fprintln(b, "/help Show this message")
	return b.String()
}
