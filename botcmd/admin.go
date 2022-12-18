package botcmd

import (
	"bytes"
	"fmt"

	"github.com/STARRY-S/telebot/status"
	"github.com/STARRY-S/telebot/user"
	"github.com/STARRY-S/telebot/utils"
	"gopkg.in/telebot.v3"
)

func AddAdminCommands(bot *telebot.Bot) {
	bot.Handle("/status", func(c telebot.Context) error {
		// only available in private chat
		if !isPrivateChat(c) {
			return c.Reply(utils.ReplyOnlyPrivate)
		}
		if !isAdmin(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		status, err := status.GetStatus()
		if err != nil {
			return err
		}
		return c.Reply(
			fmt.Sprintf("```\n%s\n```", status),
			telebot.ModeMarkdownV2,
		)
	})

	bot.Handle("/admins", func(c telebot.Context) error {
		// only available in private chat
		if !isPrivateChat(c) {
			return c.Reply(utils.ReplyOnlyPrivate)
		}
		if !isAdmin(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}

		users := user.Users()
		if users == "" {
			return c.Reply(utils.ReplyFailed)
		}
		return c.Reply(
			fmt.Sprintf("```\n%s\n```", users),
			telebot.ModeMarkdownV2,
		)
	})
}

func GetAdminHelpMessage() string {
	b := &bytes.Buffer{}
	fmt.Fprintln(b, "/status Get system status (Private) (Admin)")
	fmt.Fprintln(b, "/admins Get admins list (Private) (Admin)")
	return b.String()
}
