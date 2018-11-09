package bcc

import (
	"testing"

	"fmt"

	"github.com/baidu/baiducloud-sdk-go/util"
)

func TestDescribeInstance(t *testing.T) {
	// ts := httptest.NewServer(InstancesHandler())
	// defer ts.Close()
	bccClient.SetDebug(true)
	// bccClient.Endpoint = ts.URL
	// ins, err := bccClient.DescribeInstance("i-YufwpQAe", nil)
	ins, err := bccClient.DescribeInstance("i-7VUJvwqR", nil)
	if err != nil {
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	if ins.InstanceName != "instance-luz2ef4l-1" {
		t.Error("name error!")
	}
}

func TestListInstances(t *testing.T) {
	// ts := httptest.NewServer(InstancesHandler())
	// defer ts.Close()
	// bccClient.Endpoint = ts.URL
	// bccClient.Endpoint = "bcc.bce-api.baidu.com"
	bccClient.SetDebug(true)
	list, err := bccClient.ListInstances(nil)

	if err != nil {
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	for _, ins := range list {
		fmt.Println(ins.VpcId)
		if ins.InstanceId != "i-IyWRtII7" {
			// t.Error("instanceId error")
		}
	}
}
