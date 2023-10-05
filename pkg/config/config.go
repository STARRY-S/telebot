package config

import (
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiToken      string   `yaml:"apiToken"`
	Owner         string   `yaml:"owner"`
	OwnerID       string   `yaml:"ownerID"`
	Admins        []string `yaml:"admins"`
	ExecWhiteList []string `yaml:"execWhitelist"`
	ExecTimeout   int      `yaml:"execTimeout"`
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
			logrus.Fatalf("%s env not set", APITOKEN_ENV_NAME)
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

func ExecTimeout() time.Duration {
	return time.Duration(config.ExecTimeout) * time.Second
}

func Owner() string {
	return config.Owner
}

func OwnerID() int64 {
	if config.OwnerID == "" {
		logrus.Error("failed to get owner ID: not set in config")
		return 0
	}
	id, err := strconv.Atoi(config.OwnerID)
	if err != nil {
		logrus.Errorf("failed to get owner ID: %v", err)
		return 0
	}
	return int64(id)
}
