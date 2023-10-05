package utils

import "fmt"

var (
	version   = "v0.1.0"
	gitCommit = ""
)

func GetVersion() string {
	if gitCommit != "" {
		return fmt.Sprintf("Version %s - %s", version, gitCommit)
	}
	return fmt.Sprintf("Version %s - HEAD", version)
}
