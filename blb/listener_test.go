package blb

import (
	"fmt"
	"testing"

	"github.com/baidu/baiducloud-sdk-go/util"
)

func TestCreateTCPListener(t *testing.T) {
	args := &CreateTCPListenerArgs{
		LoadBalancerId: "lb-e5b33752",
		ListenerPort:   8088,
		BackendPort:    8080,
		Scheduler:      "LeastConnection",
	}
	err := blbClient.CreateTCPListener(args, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUDPListener(t *testing.T) {
	args := &CreateUDPListenerArgs{
		LoadBalancerId:    "lb-f5d263e5",
		ListenerPort:      8888,
		BackendPort:       8888,
		Scheduler:         "LeastConnection",
		HealthCheckString: "hello",
	}
	err := blbClient.CreateUDPListener(args, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateHTTPListener(t *testing.T) {
	args := &CreateHTTPListenerArgs{
		LoadBalancerId: "lb-f5d263e5",
		ListenerPort:   8899,
		BackendPort:    8899,
		Scheduler:      "LeastConnection",
	}
	err := blbClient.CreateHTTPListener(args, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDescribeTCPListener(t *testing.T) {
	args := &DescribeTCPListenerArgs{
		LoadBalancerId: "lb-e5b33752",
		ListenerPort:   8088,
	}
	list, err := blbClient.DescribeTCPListener(args, nil)

	if err != nil {
		t.Error(util.FormatTest("ListInstances", err.Error(), "nil"))
	}
	for _, blb := range list {
		fmt.Println(blb.ListenerPort)
	}
}

func TestDescribeUDPListener(t *testing.T) {
	args := &DescribeUDPListenerArgs{
		LoadBalancerId: "lb-07ab7a1d",
		// ListenerPort:   80,
	}
	list, err := blbClient.DescribeUDPListener(args, nil)

	if err != nil {
		t.Error(util.FormatTest("DescribeUDPListener", err.Error(), "nil"))
	}
	for _, blb := range list {
		fmt.Println(blb.ListenerPort)
	}
}
func TestUpdateTCPListener(t *testing.T) {
	args := &UpdateTCPListenerArgs{
		LoadBalancerId: "lb-e5b33752",
		ListenerPort:   8088,
		BackendPort:    999,
	}
	err := blbClient.UpdateTCPListener(args, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateUDPListener(t *testing.T) {
	args := &UpdateUDPListenerArgs{
		LoadBalancerId:    "lb-f5d263e5",
		ListenerPort:      8888,
		BackendPort:       8019,
		Scheduler:         "RoundRobin",
		HealthCheckString: "A",
	}
	err := blbClient.UpdateUDPListener(args, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteListeners(t *testing.T) {
	args := &DeleteListenersArgs{
		LoadBalancerId: "lb-e5b33752",
		PortList:       []int{8088},
	}
	err := blbClient.DeleteListeners(args, nil)
	if err != nil {
		t.Error(err)
	}
}
