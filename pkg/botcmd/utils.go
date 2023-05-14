package botcmd

import (
	"github.com/STARRY-S/telebot/pkg/config"
	"github.com/STARRY-S/telebot/pkg/user"
	"gopkg.in/telebot.v3"
)

func isUser(c telebot.Context) bool {
	return user.Find(c.Chat().Username).IsUser()
}

func isAdmin(c telebot.Context) bool {
	return user.Find(c.Chat().Username).IsAdmin()
}

func isOwner(c telebot.Context) bool {
	return user.Find(c.Chat().Username).IsOwner() &&
		c.Chat().ID == config.OwnerID()
}

func isPrivateChat(c telebot.Context) bool {
	switch c.Chat().Type {
	case telebot.ChatPrivate:
		return true
	}
	return false
}

func isChannelChat(c telebot.Context) bool {
	switch c.Chat().Type {
	case telebot.ChatChannel:
		return true
	case telebot.ChatChannelPrivate:
		return true
	}
	return false
}

func isGroupChat(c telebot.Context) bool {
	switch c.Chat().Type {
	case telebot.ChatGroup:
		return true
	case telebot.ChatSuperGroup:
		return true
	}
	return false
}
