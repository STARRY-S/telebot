package ec2

import "time"

type EC2Status struct {
	Instances []Instance
}

type Instance struct {
	ID          string
	Type        string
	Zone        string
	RunningTime string // Running time (hour)
	LaunchTime  time.Time
	Spot        bool
	State       string
	stateCode   InstanceStateCode
	Tags        map[string]string
}

func (i *Instance) IsRunning() bool {
	return i.stateCode.IsRunning()
}
