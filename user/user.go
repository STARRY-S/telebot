package user

import (
	"bytes"
	"fmt"
)

type Level int

const (
	LevelUnknow Level = iota
	LevelUser
	LevelAdmin
	LevelOwner
)

const (
	CacheFolder = "cache"
	CacheFile   = "user-cache.json"
)

type User struct {
	UserLevel Level
	Username  string
}

var users map[string]*User = make(map[string]*User)

func (u *User) IsUser() bool {
	if u == nil {
		return false
	}
	return u.UserLevel >= LevelUser
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
	case LevelUser:
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

func FindUser(name string) *User {
	return users[name]
}

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
	case LevelUser:
		return "User"
	case LevelAdmin:
		return "Admin"
	case LevelOwner:
		return "Owner"
	}
	return "Unknow"
}
