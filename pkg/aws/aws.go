package aws

import (
	"context"
	"fmt"
	"regexp"

	"github.com/STARRY-S/telebot/pkg/aws/ec2"
	"github.com/STARRY-S/telebot/pkg/config"
	"github.com/sirupsen/logrus"
)

func GetEC2Status(
	ctx context.Context, regex *regexp.Regexp, outputStopped bool,
) (*ec2.EC2Status, error) {
	status := &ec2.EC2Status{}
	regions := config.AWSRegions()
	if len(regions) == 0 {
		logrus.Warnf("AWS region in config is empty")
	}
	for _, region := range regions {
		s, err := ec2.DescribeInstances(ctx, region)
		if err != nil {
			return nil, fmt.Errorf("GetEC2Status: %v", err)
		}
		if s.Instances != nil {
			for _, i := range s.Instances {
				if regex != nil && !regex.MatchString(i.Tags["Name"]) {
					// skip name not matched instance if regex is not nil
					logrus.Debugf("Skip regex not matched: %v", i)
					continue
				}
				if !outputStopped && !i.IsRunning() {
					// skip not running instance if outputStopped is false
					logrus.Debugf("Skip stopped: %v", i)
					continue
				}
				status.Instances = append(status.Instances, i)
			}
		}
	}
	return status, nil
}
