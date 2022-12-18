package botcmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"

	"github.com/STARRY-S/telebot/config"
	"github.com/STARRY-S/telebot/user"
	"github.com/STARRY-S/telebot/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
)

const (
	ExecTimeout = time.Second * 3
)

func AddOwnerCommands(bot *telebot.Bot) {
	bot.Handle("/add_admin", func(c telebot.Context) error {
		if !isOwner(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /add_admin <user>")
		}
		if !user.Find(args[0]).IsAdmin() {
			err := user.Register(args[0], user.LevelAdmin)
			if err != nil {
				logrus.Errorf("add_admin failed: %v", err)
				return c.Reply(utils.ReplyFailed)
			}
		}
		return c.Reply(
			fmt.Sprintf("Add @%s to admin list temporally", args[0]),
		)
	})

	bot.Handle("/del_admin", func(c telebot.Context) error {
		if !isOwner(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /del_admin <user>")
		}
		if user.Find(args[0]).IsAdmin() {
			err := user.Register(args[0], user.LevelUnknow)
			if err != nil {
				logrus.Errorf("del_admin failed: %v", err)
				return c.Reply(utils.ReplyFailed)
			}
			return c.Reply(
				fmt.Sprintf("Remove @%s from admin list temporally", args[0]),
			)
		}

		return c.Reply(
			fmt.Sprintf("@%s is not in admin list", args[0]),
		)
	})

	bot.Handle("/exec", func(c telebot.Context) error {
		if !isOwner(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		if !isPrivateChat(c) {
			return c.Reply(utils.ReplyOnlyPrivate)
		}
		cmdArgs := c.Args()
		if len(cmdArgs) == 0 {
			return c.Reply("Usage: /exec [COMMAND] [OPTIONS]")
		}

		if !config.ExecWhiteListContains(cmdArgs[0]) {
			return c.Reply(
				fmt.Sprintf("`%s` is not allowed to execute", cmdArgs[0]),
				telebot.ModeMarkdownV2,
			)
		}

		cmdout := make(chan string)
		cmderr := make(chan string)
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		go func() {
			out := &bytes.Buffer{}
			cmd.Stdout = out
			cmd.Stderr = out
			if err := cmd.Start(); err != nil {
				cmderr <- err.Error()
				return
			}
			if err := cmd.Wait(); err != nil {
				cmderr <- err.Error()
				return
			}

			cmdout <- out.String()
		}()

		timer := time.NewTimer(ExecTimeout)
		select {
		case <-timer.C:
			// command executes timeout
			timer.Stop()
			if err := cmd.Process.Kill(); err != nil {
				logrus.Error("Failed to kill command: ", err)
				return c.Reply(fmt.Sprintf("Failed: execute timeout\n"+
					"failed to kill: \n```%s\n```", err.Error()),
					telebot.ModeMarkdownV2)
			}
			if err := cmd.Process.Release(); err != nil {
				logrus.Error("Failed to release command: ", err)
				return c.Reply(fmt.Sprintf("Failed: execute timeout\n"+
					"killed but failed to release: \n```%s\n```", err.Error()),
					telebot.ModeMarkdownV2)
			}
			return c.Reply("Failed: execute timeout, killed")
		case out := <-cmdout:
			// command executes successfully
			timer.Stop()
			if len(out) > 3000 {
				out = out[:3000] + "\n......"
			}
			return c.Reply(
				fmt.Sprintf("Done\n```\n%s\n```", out),
				telebot.ModeMarkdownV2,
			)
		case e := <-cmderr:
			// command executes failed
			timer.Stop()
			return c.Reply(
				fmt.Sprintf("Execute failed:\n```\n%s\n```", e),
				telebot.ModeMarkdownV2,
			)
		}
	})
}

func GetOwnerHelpMessage() string {
	b := &bytes.Buffer{}
	fmt.Fprintln(b, "/add_admin Register temporary admin user (Owner)")
	fmt.Fprintln(b, "/del_admin Remove admin temporally (Owner)")
	fmt.Fprintln(b, "/exec Run commands (Owner)")
	return b.String()
}
