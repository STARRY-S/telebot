package reminder

import (
	"context"
	"fmt"
	"time"

	"github.com/STARRY-S/telebot/pkg/aws"
	"github.com/STARRY-S/telebot/pkg/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"gopkg.in/yaml.v3"
)

func Init(bot *telebot.Bot) {
	if !config.AWSReminderEnabled() {
		return
	}

	go reminderFunc(bot)
}

func reminderFunc(bot *telebot.Bot) {
	var sent bool = false
	for {
		time.Sleep(time.Second * 5)
		w := time.Now().Weekday()
		if w == time.Saturday || w == time.Sunday {
			// skip reminder on weekend
			if !config.AWSReminderSendOnWeekend() {
				continue
			}
		}

		hour := time.Now().Hour()
		minute := time.Now().Minute()
		h, m, err := config.AWSReminderTime()
		if err != nil {
			logrus.Error("failed to get ReminderTime from config: %v", err)
			return
		}

		if hour == h && minute == m && !sent {
			logrus.Infof("Reminder: query AWS EC2 status...")
			sent = true
			chat, err := bot.ChatByID(config.OwnerID())
			if err != nil {
				logrus.Error(err)
				continue
			}
			status, err := aws.GetEC2Status(
				context.Background(),
				config.EC2InstanceNameRegex(),
				config.EC2OutputStopped())
			if err != nil {
				logrus.Error("%v", err)
			}
			if len(status.Instances) == 0 {
				logrus.Infof("No instances found")
				continue
			}
			d, err := yaml.Marshal(status)
			if err != nil {
				logrus.Error("yaml.Marshal: %v", err)
			}
			_, err = bot.Send(chat, fmt.Sprintf(
				"AWS EC2 Reminder: \n<code>%v</code>", string(d)),
				telebot.ModeHTML)
			if err != nil {
				logrus.Error(err)
				continue
			}
			logrus.Infof("Sent AWS EC2 status")
		} else if minute != m && sent {
			logrus.Debugf("Reset reminder sent status")
			sent = false
		}
	}
}
