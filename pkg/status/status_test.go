package status

import (
	"os/exec"
	"testing"
)

func Test_GetStatus(t *testing.T) {
	if _, err := exec.LookPath("sensors"); err != nil {
		return
	}
	status, err := GetStatus()
	if err != nil {
		t.Error(err)
	}
	t.Logf("\n%v\n", status)
}
