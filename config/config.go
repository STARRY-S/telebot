package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"k8s.io/utils/strings/slices"
)

type Config struct {
	ApiToken      string   `yaml:"apiToken"`
	Owner         string   `yaml:"owner"`
	Admins        []string `yaml:"admins"`
	ExecWhiteList []string `yaml:"execWhitelist"`
}

const (
	CONFIG_FILE_NAME  = "config.yaml"
	APITOKEN_ENV_NAME = "TELEGRAM_APITOKEN"
)

var (
	config = Config{}
)

func Init() {
	data, err := os.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		logrus.Panic(err)
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		logrus.Panic(err)
	}

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

func Admins() []string {
	if config.Admins == nil {
		return make([]string, 0)
	}
	return slices.Clone(config.Admins)
}

func ExecWhiteListContains(s string) bool {
	if config.ExecWhiteList == nil {
		return false
	}
	return slices.Contains(config.ExecWhiteList, s)
}

func Owner() string {
	return config.Owner
}
