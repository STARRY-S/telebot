package botcmd

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/STARRY-S/telebot/pkg/status"
	"github.com/STARRY-S/telebot/pkg/user"
	"github.com/STARRY-S/telebot/pkg/utils"
	"github.com/sirupsen/logrus"
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

	bot.Handle(telebot.OnMedia, func(c telebot.Context) error {
		if !isPrivateChat(c) {
			return c.Reply(utils.ReplyOnlyPrivate)
		}
		if !isAdmin(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		doc, err := convertVideoMediaToGIF(bot, c.Message())
		if err != nil {
			logrus.Error(err)
			return c.Reply(fmt.Sprintf("<pre>%s</pre>", err), telebot.ModeHTML)
		}
		defer os.Remove(doc.File.FileLocal)
		return c.Reply(doc)
	})

	bot.Handle(telebot.OnSticker, func(c telebot.Context) error {
		if !isPrivateChat(c) {
			return c.Reply(utils.ReplyOnlyPrivate)
		}
		if !isAdmin(c) {
			return c.Reply(utils.ReplyPermissionDenied)
		}
		doc, err := convertVideoMediaToGIF(bot, c.Message())
		if err != nil {
			logrus.Error(err)
			return c.Reply(fmt.Sprintf("<pre>%s</pre>", err), telebot.ModeHTML)
		}
		defer os.Remove(doc.File.FileLocal)
		return c.Reply(doc)
	})
}

func convertVideoMediaToGIF(
	bot *telebot.Bot, message *telebot.Message,
) (*telebot.Document, error) {
	media := message.Media()
	switch media.MediaType() {
	case "video":
	case "animation":
	case "sticker":
		if !message.Sticker.Video {
			return nil, fmt.Errorf(
				"unsupported image sticker, only video sticker supported")
		}
	default:
		return nil, fmt.Errorf(
			"unsupported type: %v, only video and animation supported",
			media.MediaType())
	}

	if media.MediaFile().FileSize > 1024*1024*10 {
		return nil, fmt.Errorf(
			"failed to convert file: size too large: %v",
			media.MediaFile().FileSize)
	}

	logrus.Debugf("file size: %.2fK",
		float32(media.MediaFile().FileSize)/float32(1024))
	r, err := bot.File(media.MediaFile())
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %v", err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	f, err := os.CreateTemp("", "media")
	if err != nil {
		return nil, fmt.Errorf("failed to create tmp file: %v", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		return nil, fmt.Errorf("failed to write tmp file: %v", err)
	}

	cmd := exec.Command(
		"bash", "-c",
		fmt.Sprintf("ffmpeg -i %s %s.gif", f.Name(), f.Name()),
	)
	out := &bytes.Buffer{}
	cmd.Stdout = out
	cmd.Stderr = out
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("%s\n%s", out.String(), err.Error())
	}
	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("%s\n%s", out.String(), err.Error())
	}
	logrus.Debugf("ffmpeg output: %v", out.String())
	defer os.Remove(f.Name() + ".gif")

	gif, err := os.Open(f.Name() + ".gif")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	z, err := os.Create(f.Name() + ".zip")
	if err != nil {
		return nil, fmt.Errorf("failed to create zip file: %v", err)
	}
	zw := zip.NewWriter(z)
	zgif, err := zw.Create(path.Base(f.Name() + ".gif"))
	if err != nil {
		return nil, fmt.Errorf("failed to create file in zip archive: %v", err)
	}
	_, err = io.Copy(zgif, gif)
	if err != nil {
		return nil, fmt.Errorf("failed to write file in zip archive: %v", err)
	}
	zw.Close()

	return &telebot.Document{
		File:     telebot.FromDisk(f.Name() + ".zip"),
		FileName: path.Base(f.Name()) + ".zip",
	}, nil
}

func GetAdminHelpMessage() string {
	b := &bytes.Buffer{}
	fmt.Fprintln(b, "/status Get system status (Private) (Admin)")
	fmt.Fprintln(b, "/admins Get admins list (Private) (Admin)")
	return b.String()
}
