package bcc

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/baidu/baiducloud-sdk-go/billing"
	"github.com/baidu/baiducloud-sdk-go/cds"
)

const (
	InstanceStatusRunning            string = "Running"
	InstanceStatusStarting           string = "Starting"
	InstanceStatusStopping           string = "Stopping"
	InstanceStatusStopped            string = "Stopped"
	InstanceStatusDeleted            string = "Deleted"
	InstanceStatusScaling            string = "Scaling"
	InstanceStatusExpired            string = "Expired"
	InstanceStatusError              string = "Error"
	InstanceStatusSnapshotProcessing string = "SnapshotProcessing"
	InstanceStatusImageProcessing    string = "ImageProcessing"
)

// Instance define instance model
type Instance struct {
	InstanceID            string `json:"id"`
	InstanceName          string `json:"name"`
	Description           string `json:"desc"`
	Status                string `json:"status"`
	PaymentTiming         string `json:"paymentTiming"`
	CreationTime          string `json:"createTime"`
	ExpireTime            string `json:"expireTime"`
	PublicIP              string `json:"publicIp"`
	InternalIP            string `json:"internalIp"`
	CPUCount              int    `json:"cpuCount"`
	GPUCount              int    `json:"gpuCount"`
	MemoryCapacityInGB    int    `json:"memoryCapacityInGB"`
	LocalDiskSizeInGB     int    `json:"localDiskSizeInGB"`
	ImageID               string `json:"imageId"`
	NetworkCapacityInMbps int    `json:"networkCapacityInMbps"`
	PlacementPolicy       string `json:"placementPolicy"`
	ZoneName              string `json:"zoneName"`
	SubnetID              string `json:"subnetId"`
	VpcID                 string `json:"vpcId"`
}

type EphemeralDisk struct {
	StorageType  cds.StorageType `json:"storageType,storageType"`
	SizeInGB     int             `json:"sizeInGB,omitempty"`
	FreeSizeInGB int             `json:"freeSizeInGB,omitempty"`
}

type CreateCdsModel struct {
	StorageType cds.StorageType `json:"storageType,storageType"`
	SnapshotID  string          `json:"snapshotId,omitempty"`
	CdsSizeInGB int             `json:"cdsSizeInGB,omitempty"`
}

type ListInstancesResponse struct {
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
	Instances   []Instance `json:"instances"`
}

// CreateInstanceArgs is args to create instances
// refers to https://cloud.baidu.com/doc/BCC/API.html#.E5.88.9B.E5.BB.BA.E5.AE.9E.E4.BE.8B
// TODO ssh public key still not supperted
type CreateInstanceArgs struct {
	ImageID               string           `json:"imageId"`
	Billing               billing.Billing  `json:"billing"`
	InstanceType          string           `json:"instanceType,omitempty"`
	CPUCount              int              `json:"cpuCount"`
	MemoryCapacityInGB    int              `json:"memoryCapacityInGB"`
	RootDiskSizeInGB      int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   int              `json:"rootDiskStorageType,omitempty"`
	LocalDiskSizeInGB     int              `json:"localDiskSizeInGB,omitempty"` // deprecated now
	EphemeralDisks        []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList         []CreateCdsModel `json:"createCdsList,omitempty"`
	NetworkCapacityInMbps int              `json:"networkCapacityInMbps,omitempty"`
	DedicatedHostID       int              `json:"dedicatedHostId,omitempty"`
	PurchaseCount         int              `json:"purchaseCount,omitempty"`
	Name                  string           `json:"name,omitempty"`
	AdminPass             string           `json:"adminPass,omitempty"`
	ZoneName              string           `json:"zoneName,omitempty"`
	SubnetID              string           `json:"subnetId,omitempty"`
	SecurityGroupID       string           `json:"securityGroupId,omitempty"`
	GPUCard               string           `json:"gpuCard,omitempty"`
	FPGACard              string           `json:"fpgaCard,omitempty"`
	CardCount             string           `json:"cardCount,omitempty"`
}

// CreateInstanceResponse is response of create instances
type CreateInstanceResponse struct {
	InstanceIDs []string `json:"instanceIds"`
}

type GetInstanceResponse struct {
	Ins Instance `json:"instance"`
}

func (args *CreateInstanceArgs) validate() error {
	if args == nil {
		return fmt.Errorf("CreateInstanceArgs cannot be nil")
	}
	if args.ImageID == "" {
		return fmt.Errorf("imageId cannot be empty")
	}
	if args.CPUCount <= 0 {
		return fmt.Errorf("cpuCount must be positive integer")
	}
	if args.MemoryCapacityInGB <= 0 {
		return fmt.Errorf("memoryCapacityInGB must be positive integer")
	}
	return nil
}

// CreateInstances create instances according to args
func (c *Client) CreateInstances(args *CreateInstanceArgs) ([]string, error) {
	instanceIDs := []string{}
	if err := args.validate(); err != nil {
		return instanceIDs, err
	}
	params := map[string]string{
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return instanceIDs, err
	}
	req, err := bce.NewRequest("POST", c.GetURL("v2/instance", params), bytes.NewBuffer(postContent))
	if err != nil {
		return instanceIDs, err
	}
	resp, err := c.SendRequest(req, nil)
	if err != nil {
		return instanceIDs, err
	}
	bodyContent, err := resp.GetBodyContent()
	if err != nil {
		return instanceIDs, err
	}
	var ciResp *CreateInstanceResponse
	if err := json.Unmarshal(bodyContent, &ciResp); err != nil {
		return instanceIDs, err
	}
	return ciResp.InstanceIDs, nil
}

// ListInstances gets all Instances.
func (c *Client) ListInstances(option *bce.SignOption) ([]Instance, error) {

	req, err := bce.NewRequest("GET", c.GetURL("v2/instance", nil), nil)

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

	var insList *ListInstancesResponse
	err = json.Unmarshal(bodyContent, &insList)

	if err != nil {
		return nil, err
	}

	return insList.Instances, nil
}

// DescribeInstance describe a instance
func (c *Client) DescribeInstance(instanceID string, option *bce.SignOption) (*Instance, error) {

	req, err := bce.NewRequest("GET", c.GetURL("v2/instance"+"/"+instanceID, nil), nil)

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

	var ins GetInstanceResponse
	err = json.Unmarshal(bodyContent, &ins)

	if err != nil {
		return nil, err
	}

	return &ins.Ins, nil
}
