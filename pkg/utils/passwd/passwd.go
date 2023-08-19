package passwd

import (
	"math/rand"
	"strings"
	"time"

	"github.com/tjarratt/babble"
)

func GenRandomPasswd(length int, hasNum, hasSpec bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	if hasNum {
		chars = append(chars, []byte("0123456789")...)
	}
	if hasSpec {
		chars = append(chars, []byte(`~!@#$%^&*()-=_+./[]{}|/?`)...)
	}
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteByte(chars[r.Intn(len(chars))])
	}

	return b.String()
}

func GenRememberablePasswd(words int) string {
	babbler := babble.NewBabbler()
	babbler.Count = words
	return strings.ReplaceAll(babbler.Babble(), "'", "")
}
