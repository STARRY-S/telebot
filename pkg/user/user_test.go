package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	if err := Register("test", LevelAdmin); err != nil {
		t.Error(err)
		return
	}
	u, ok := users["test"]
	if !ok {
		t.Error("failed")
		return
	}
	assert.Equal(t, "test", u.Username)
	assert.Equal(t, LevelAdmin, u.UserLevel)
	if err := Register("test", LevelOwner); err != nil {
		t.Error(err)
	}
	u, ok = users["test"]
	if !ok {
		t.Error("failed")
		return
	}
	assert.Equal(t, "test", u.Username)
	assert.Equal(t, LevelOwner, u.UserLevel)
	// deregister
	if err := Register("test", LevelUnknow); err != nil {
		t.Error(err)
	}
	u, ok = users["test"]
	assert.False(t, ok)
	assert.Nil(t, u)

}
