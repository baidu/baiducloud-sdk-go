package blb

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baidu/baiducloud-sdk-go/bce"
)

type BackendServer struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight,omitempty"`
}

type BackendServerStatus struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight"`
	Status     string `json:"status"`
}

type AddBackendServersArgs struct {
	LoadBalancerId    string          `json:"-"`
	BackendServerList []BackendServer `json:"backendServerList"`
}

func (args *AddBackendServersArgs) validate() error {
	if args == nil {
		return fmt.Errorf("AddBackendServersArgs need args")
	}
	if args.LoadBalancerId == "" {
		return fmt.Errorf("AddBackendServersArgs need LoadBalancerId")
	}
	if args.BackendServerList == nil {
		return fmt.Errorf("UpdateUDPListener need BackendServerList")
	}

	return nil
}

func (c *Client) AddBackendServers(args *AddBackendServersArgs, option *bce.SignOption) error {
	err := args.validate()
	if err != nil {
		return err
	}
	params := map[string]string{
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return err
	}
	req, err := bce.NewRequest("POST", c.GetURL("v1/blb"+"/"+args.LoadBalancerId+"/backendserver", params), bytes.NewBuffer(postContent))
	if err != nil {
		return err
	}
	_, err = c.SendRequest(req, option)
	if err != nil {
		return err
	}
	return nil
}

type DescribeBackendServersArgs struct {
	LoadBalancerId string `json:"-"`
}

type DescribeBackendServersResponse struct {
	Marker            string          `json:"marker"`
	IsTruncated       bool            `json:"isTruncated"`
	NextMarker        string          `json:"nextMarker"`
	MaxKeys           int             `json:"maxKeys"`
	BackendServerList []BackendServer `json:"backendServerList"`
}

func (args *DescribeBackendServersArgs) validate() error {
	if args == nil {
		return fmt.Errorf("DescribeBackendServersArgs need args")
	}
	if args.LoadBalancerId == "" {
		return fmt.Errorf("DescribeBackendServersArgs need LoadBalancerId")
	}
	return nil
}

func (c *Client) DescribeBackendServers(args *DescribeBackendServersArgs, option *bce.SignOption) ([]BackendServer, error) {
	err := args.validate()
	if err != nil {
		return nil, err
	}
	req, err := bce.NewRequest("GET", c.GetURL("v1/blb"+"/"+args.LoadBalancerId+"/backendserver", nil), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.SendRequest(req, option)

	if err != nil {
		return nil, err
	}
	bodyContent, err := resp.GetBodyContent()

	if err != nil {
		return nil, err
	}
	var blbsResp *DescribeBackendServersResponse
	err = json.Unmarshal(bodyContent, &blbsResp)

	if err != nil {
		return nil, err
	}
	return blbsResp.BackendServerList, nil

}

type UpdateBackendServersArgs struct {
	LoadBalancerId    string          `json:"-"`
	BackendServerList []BackendServer `json:"backendServerList"`
}

func (args *UpdateBackendServersArgs) validate() error {
	if args == nil {
		return fmt.Errorf("UpdateBackendServersArgs need args")
	}
	if args.LoadBalancerId == "" {
		return fmt.Errorf("UpdateBackendServersArgs need LoadBalancerId")
	}
	if len(args.BackendServerList) == 0 {
		return fmt.Errorf("UpdateBackendServersArgs need BackendServerList")
	}
	return nil
}

// UpdateBackendServers update  BackendServers
func (c *Client) UpdateBackendServers(args *UpdateBackendServersArgs, option *bce.SignOption) error {
	err := args.validate()
	if err != nil {
		return err
	}
	params := map[string]string{
		"update":      "",
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return err
	}
	req, err := bce.NewRequest("PUT", c.GetURL("v1/blb"+"/"+args.LoadBalancerId+"/backendserver", params), bytes.NewBuffer(postContent))
	if err != nil {
		return err
	}
	_, err = c.SendRequest(req, option)
	if err != nil {
		return err
	}
	return nil
}

type RemoveBackendServersArgs struct {
	LoadBalancerId    string   `json:"-"`
	BackendServerList []string `json:"backendServerList"`
}

func (args *RemoveBackendServersArgs) validate() error {
	if args == nil {
		return fmt.Errorf("UpdateBackendServersArgs need args")
	}
	if args.LoadBalancerId == "" {
		return fmt.Errorf("UpdateBackendServersArgs need LoadBalancerId")
	}
	if len(args.BackendServerList) == 0 {
		return fmt.Errorf("UpdateBackendServersArgs need BackendServerList")
	}
	return nil
}

// RemoveBackendServers remove a BackendServers
func (c *Client) RemoveBackendServers(args *RemoveBackendServersArgs, option *bce.SignOption) error {
	err := args.validate()
	if err != nil {
		return err
	}
	params := map[string]string{
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return err
	}
	req, err := bce.NewRequest("PUT", c.GetURL("v1/blb"+"/"+args.LoadBalancerId+"/backendserver", params), bytes.NewBuffer(postContent))
	if err != nil {
		return err
	}
	_, err = c.SendRequest(req, option)
	if err != nil {
		return err
	}
	return nil
}
