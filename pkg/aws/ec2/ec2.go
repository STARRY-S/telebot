package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type InstanceStateCode int32

const (
	InstanceStatePending      = InstanceStateCode(0)
	InstanceStateRunning      = InstanceStateCode(0x10)
	InstanceStateShuttingDown = InstanceStateCode(0x20)
	InstanceStateTerminated   = InstanceStateCode(0x30)
	InstanceStateStopping     = InstanceStateCode(0x40)
	InstanceStateStopped      = InstanceStateCode(0x50)
)

func (c *InstanceStateCode) IsRunning() bool {
	return *c == InstanceStateRunning
}

func (c *InstanceStateCode) IsStopped() bool {
	return *c == InstanceStateStopped
}

func (c *InstanceStateCode) IsTerminated() bool {
	return *c == InstanceStateTerminated
}

func DescribeInstances(
	ctx context.Context, region string,
) (*EC2Status, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("DescribeInstances: config: %w", err)
	}

	client := ec2.NewFromConfig(cfg)
	input := &ec2.DescribeInstancesInput{}
	result, err := client.DescribeInstances(ctx, input)
	if err != nil {
		return nil, fmt.Errorf(
			"DescribeInstances: failed to retrive EC2 information: %w", err)
	}

	ec2Status := &EC2Status{}
	if len(result.Reservations) == 0 {
		return ec2Status, nil
	}

	ec2Status.Instances = make(
		[]Instance, 0, len(result.Reservations[0].Instances))
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			running := time.Now().Sub(*i.LaunchTime)
			instance := Instance{
				ID:   *i.InstanceId,
				Type: string(i.InstanceType),
				Zone: "",
				RunningTime: fmt.Sprintf("%.2f",
					float64(running)/float64(time.Hour)),
				LaunchTime: *i.LaunchTime,
				Spot:       false,
				State:      "",
				Tags:       map[string]string{},
			}
			if i.Placement != nil {
				instance.Zone = *i.Placement.AvailabilityZone
			}
			if i.SpotInstanceRequestId != nil {
				instance.Spot = true
			}
			if i.State != nil {
				instance.State = string(i.State.Name)
				instance.stateCode = InstanceStateCode(*i.State.Code)
			}
			for _, t := range i.Tags {
				instance.Tags[*t.Key] = *t.Value
			}
			ec2Status.Instances = append(ec2Status.Instances, instance)
		}
	}
	return ec2Status, nil
}
