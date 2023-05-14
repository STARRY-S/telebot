package config

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiToken      string    `yaml:"apiToken"`
	Owner         string    `yaml:"owner"`
	OwnerID       string    `yaml:"ownerID"`
	Admins        []string  `yaml:"admins"`
	ExecWhiteList []string  `yaml:"execWhitelist"`
	AWS           AWSConfig `yaml:"aws"`
}

type AWSConfig struct {
	ReminderEnabled  bool   `yaml:"reminder"`
	ReminderTime     string `yaml:"reminderTime"`
	SendOnWorkingDay bool   `yaml:"sendOnWorkingDay"`
	SendOnWeekend    bool   `yaml:"sendOnWeekend"`

	Regions []string `yaml:"regions"`

	EC2 AWSEC2Config `yaml:"ec2"`

	EKS AWSEKSConfig `yaml:"eks,omitempty"`
}

type AWSEC2Config struct {
	InstanceNameRegex string `yaml:"instanceNameRegex"`
	// outputStoppedInstance
	StoppedInstance bool `yaml:"outputStoppedInstance"`
}

// TODO: Add EKS config
type AWSEKSConfig struct {
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
		logrus.Error("failed to get owner ID: %v", err)
		return 0
	}
	return int64(id)
}

func AWSReminderEnabled() bool {
	return config.AWS.ReminderEnabled
}

func AWSReminderSendOnWeekend() bool {
	return config.AWS.SendOnWeekend
}

func AWSReminderTime() (hour, minute int, err error) {
	spec := strings.Split(config.AWS.ReminderTime, ":")
	if len(spec) != 2 {
		return 0, 0, fmt.Errorf("invalid format, should be like '00:00'")
	}
	hour, err = strconv.Atoi(spec[0])
	if err != nil {
		return 0, 0, err
	}
	minute, err = strconv.Atoi(spec[1])
	if err != nil {
		return 0, 0, err
	}
	return hour, minute, nil
}

func AWSRegions() []string {
	region := make([]string, len(config.AWS.Regions))
	copy(region, config.AWS.Regions)
	return region
}

func EC2OutputStopped() bool {
	return config.AWS.EC2.StoppedInstance
}

var (
	instanceNameRegex *regexp.Regexp
)

func EC2InstanceNameRegex() *regexp.Regexp {
	if config.AWS.EC2.InstanceNameRegex == "" {
		return nil
	}

	if instanceNameRegex != nil {
		return instanceNameRegex
	}

	var err error
	instanceNameRegex, err = regexp.Compile(config.AWS.EC2.InstanceNameRegex)
	if err != nil {
		logrus.Error("EC2InstanceNameRegex: %v", err)
		return nil
	}
	return instanceNameRegex
}
