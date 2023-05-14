package ec2

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DescribeInstancesCmd(t *testing.T) {
	status, err := DescribeInstances(context.TODO(), "ap-northeast-1")
	assert.Nil(t, err)
	assert.NotNil(t, status)
	if t.Failed() {
		return
	}
	for _, i := range status.Instances {
		d, _ := json.MarshalIndent(i, "", "  ")
		fmt.Printf("%v\n", string(d))
	}
}
