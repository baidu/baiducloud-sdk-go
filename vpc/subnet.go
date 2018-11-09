package vpc

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baidu/baiducloud-sdk-go/bce"
)

// Subnet define subnet of vpc
type Subnet struct {
	SubnetID    string `json:"subnetId"`
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VpcID       string `json:"vpcId"`
	SubnetType  string `json:"subnetType"`
	Description string `json:"description"`
}

// CreateSubnetArgs define args create a subnet
type CreateSubnetArgs struct {
	Name        string `json:"name"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	VpcID       string `json:"vpcId"`
	SubnetType  string `json:"subnetType,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateSubnetResponse define response of creating a subnet
type CreateSubnetResponse struct {
	SubnetID string `json:"subnetId"`
}

type ListSubnetResponse struct {
	Marker      string    `json:"marker"`
	IsTruncated bool      `json:"isTruncated"`
	NextMarker  string    `json:"nextMarker"`
	MaxKeys     int       `json:"maxKeys"`
	Subnets     []*Subnet `json:"subnets"`
}

type DescribeSubnetResponse struct {
	Subnet *Subnet `json:"subnet"`
}

// 在VPC中创建子网
func (c *Client) CreateSubnet(args *CreateSubnetArgs) (string, error) {
	if args == nil {
		return "", fmt.Errorf("CreateSubnet failed: CreateSubnetArgs is nil")
	}
	params := map[string]string{
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	req, err := bce.NewRequest("POST", c.GetURL("v1/vpc/subnet", params), bytes.NewBuffer(postContent))
	if err != nil {
		return "", err
	}
	resp, err := c.SendRequest(req, nil)
	if err != nil {
		return "", err
	}
	bodyContent, err := resp.GetBodyContent()
	if err != nil {
		return "", err
	}

	var createSubnetResponse *CreateSubnetResponse
	err = json.Unmarshal(bodyContent, &createSubnetResponse)
	if err != nil {
		return "", err
	}
	return createSubnetResponse.SubnetID, nil
}

// 查询指定VPC的所有子网列表信息
func (c *Client) ListSubnet(params map[string]string) ([]*Subnet, error) {
	req, err := bce.NewRequest("GET", c.GetURL("v1/vpc/subnet", params), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.SendRequest(req, nil)
	if err != nil {
		return nil, err
	}
	bodyContent, err := resp.GetBodyContent()
	if err != nil {
		return nil, err
	}

	var listSubnetResponse *ListSubnetResponse
	err = json.Unmarshal(bodyContent, &listSubnetResponse)
	if err != nil {
		return nil, err
	}
	return listSubnetResponse.Subnets, nil
}

// 查询指定子网的详细信息
func (c *Client) DescribeSubnet(subnetId string) (*Subnet, error) {
	if len(subnetId) == 0 {
		return nil, fmt.Errorf("DescribeSubnet failed, subnetId must not be empty")
	}
	req, err := bce.NewRequest("GET", c.GetURL("v1/vpc/subnet/"+subnetId, nil), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.SendRequest(req, nil)
	if err != nil {
		return nil, err
	}
	bodyContent, err := resp.GetBodyContent()
	if err != nil {
		return nil, err
	}

	var describeSubnetResponse *DescribeSubnetResponse
	err = json.Unmarshal(bodyContent, &describeSubnetResponse)
	if err != nil {
		return nil, err
	}
	return describeSubnetResponse.Subnet, nil
}
