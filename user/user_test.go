package user

import "testing"

func Test_Register(t *testing.T) {
	if err := Register("test", LevelUser); err != nil {
		t.Error(err)
	}
	u, ok := users["test"]
	if !ok {
		t.Error("failed")
	}
	if u.Username != "test" || u.UserLevel != LevelUser {
		t.Error("failed")
	}
	if err := Register("test", LevelAdmin); err != nil {
		t.Error(err)
	}
	u, ok = users["test"]
	if !ok {
		t.Error("failed")
	}
	if u.Username != "test" || u.UserLevel != LevelAdmin {
		t.Error("failed")
	}
}
