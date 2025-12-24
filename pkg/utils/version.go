package utils

import "fmt"

var (
	Version = "v0.1.0"
	Commit  = ""
)

func GetVersion() string {
	if Commit != "" {
		return fmt.Sprintf("Version %s - %s", Version, Commit)
	}
	return fmt.Sprintf("Version %s - HEAD", Version)
}
