package user

import (
	"bytes"
	"fmt"
)

type Level int

const (
	// Unknow (common user) has the permission to interact with this bot,
	// common users are not stored in memory.
	LevelUnknow Level = iota

	// Admin has the permission to view the status of bot & system
	LevelAdmin

	// Owner has the permission to control the bot & system
	LevelOwner
)

type User struct {
	UserLevel Level
	Username  string
}

// users stores the data of owner and admin users,
// unknow users are not stored in this map.
var users map[string]*User = make(map[string]*User)

func (u *User) IsUser() bool {
	if u == nil {
		return false
	}
	return u.UserLevel >= LevelUnknow
}

func (u *User) IsAdmin() bool {
	if u == nil {
		return false
	}
	return u.UserLevel >= LevelAdmin
}

func (u *User) IsOwner() bool {
	if u == nil {
		return false
	}
	return u.UserLevel >= LevelOwner
}

func (u *User) Level() Level {
	if u == nil {
		return LevelUnknow
	}
	return u.UserLevel
}

func Register(name string, level Level) error {
	if name == "" {
		return fmt.Errorf("Register: invalid name")
	}
	switch level {
	case LevelUnknow:
		// de-register user
		delete(users, name)
		return nil
	case LevelAdmin:
	case LevelOwner:
	default:
		return fmt.Errorf("Register: invalid level")
	}

	u := User{
		UserLevel: Level(level),
		Username:  name,
	}
	users[name] = &u
	return nil
}

func Find(name string) *User {
	return users[name]
}

// Users get the formatted known user list (owner and admin)
func Users() string {
	buff := &bytes.Buffer{}
	num := 0
	for _, v := range users {
		fmt.Fprintf(buff, "--------------\n")
		fmt.Fprintf(buff, "Name: %s\n", v.Username)
		fmt.Fprintf(buff, "Level: %s\n", v.UserLevel.String())
		num++
	}
	if num > 0 {
		fmt.Fprintf(buff, "--------------\n")
	}
	return buff.String()
}

func (l Level) String() string {
	switch l {
	case LevelUnknow:
		return "User"
	case LevelAdmin:
		return "Admin"
	case LevelOwner:
		return "Owner"
	}
	return ""
}
