package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiToken       string   `yaml:"apiToken"`
	AvailableUsers []string `yaml:"adminUsers"`
}

const (
	CONFIG_FILE_NAME  = "config.yaml"
	APITOKEN_ENV_NAME = "TELEGRAM_APITOKEN"
)

var (
	config = Config{}
)

func init() {
	data, err := os.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	logrus.Debugf("Config: %+v", config)
	if config.ApiToken == "" {
		logrus.Info("'apiToken' not set in config, reading it from env")
		config.ApiToken = os.Getenv(APITOKEN_ENV_NAME)
		if config.ApiToken == "" {
			logrus.Fatal(APITOKEN_ENV_NAME + "env not set")
		}
	}
}

func GetApiToken() string {
	return config.ApiToken
}

func IsAdmin(username string) bool {
	return slices.Contains(config.AvailableUsers, username)
}
