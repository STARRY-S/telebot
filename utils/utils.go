package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCommandFunc(cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%s, %w", stdout.String(), err)
	}

	return stdout.String(), nil
}
