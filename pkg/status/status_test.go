package status

import "testing"

func Test_GetStatus(t *testing.T) {
	status, err := GetStatus()
	if err != nil {
		t.Error(err)
	}
	t.Logf("\n" + status)
}
