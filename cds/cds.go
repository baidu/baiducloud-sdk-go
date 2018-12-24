package cds

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baidu/baiducloud-sdk-go/bce"
)

type Volume struct {
	Id             string             `json:"id"`
	Name           string             `json:"name"`
	DiskSizeInGB   int                `json:"diskSizeInGB"`
	PaymentTiming  string             `json:"paymentTiming"`
	CreateTime     string             `json:"createTime"`
	ExpireTime     string             `json:"expireTime"`
	Status         VolumeStatus       `json:"status"`
	VolumeType     VolumeType         `json:"type"`
	StorageType    StorageType        `json:"storageType"`
	Desc           string             `json:"desc"`
	Attachments    []VolumeAttachment `json:"attachments"`
	ZoneName       string             `json:"zoneName"`
	IsSystemVolume bool               `json:"isSystemVolume"`
}

type VolumeStatus string

const (
	VOLUMESTATUS_CREATE             VolumeStatus = "Creating"
	VOLUMESTATUS_AVAILABLE          VolumeStatus = "Available"
	VOLUMESTATUS_ATTACHING          VolumeStatus = "Attaching"
	VOLUMESTATUS_NOTAVALIABLE       VolumeStatus = "NotAvailable"
	VOLUMESTATUS_INUSE              VolumeStatus = "InUse"
	VOLUMESTATUS_DETACHING          VolumeStatus = "Detaching"
	VOLUMESTATUS_DELETING           VolumeStatus = "Deleting"
	VOLUMESTATUS_DELETED            VolumeStatus = "Deleted"
	VOLUMESTATUS_SCALING            VolumeStatus = "Scaling"
	VOLUMESTATUS_EXPIRED            VolumeStatus = "Expired"
	VOLUMESTATUS_ERROR              VolumeStatus = "Error"
	VOLUMESTATUS_SNAPSHOTPROCESSING VolumeStatus = "SnapshotProcessing"
	VOLUMESTATUS_IMAGEPROCESSING    VolumeStatus = "ImageProcessing"
)

type VolumeType string

const (
	VOLUME_TYPE_SYSTEM    VolumeType = "System"
	VOLUME_TYPE_EPHEMERAL VolumeType = "Ephemeral"
	VOLUME_TYPE_CDS       VolumeType = "Cds"
)

type StorageType string

const (
	STORAGE_TYPE_STD1 StorageType = "std1"
	STORAGE_TYPE_HP1  StorageType = "hp1"
	STORAGE_TYPE_SATA StorageType = "sata"
	STORAGE_TYPE_SSD  StorageType = "ssd"
)

// VolumeAttachment define attach info
type VolumeAttachment struct {
	VolumeId   string `json:"volumeId"`
	InstanceId string `json:"instanceId"`
	// mount path
	Device string `json:"device"`
}

// CdsPreMountInfo define premount
type CdsPreMountInfo struct {
	MountPath string           `json:"mountPath"`
	CdsConfig []DiskSizeConfig `json:"cdsConfig"`
}

// DiskSizeConfig define distsize config
type DiskSizeConfig struct {
	Size         string `json:"size"`
	SnapshotID   string `json:"snapshotID"`
	SnapshotName string `json:"snapshotName"`
	VolumeType   string `json:"volumeType"`
	StorageType  string `json:"storageType"`
	LogicalZone  string `json:"logicalZone"`
}

type CreateVolumeArgs struct {
	PurchaseCount int          `json:"purchaseCount,omitempty"`
	CdsSizeInGB   int          `json:"cdsSizeInGB"`
	StorageType   StorageType  `json:"storageType"`
	Billing       *bce.Billing `json:"billing"`
	SnapshotId    string       `json:"snapshotId,omitempty"`
	ZoneName      string       `json:"zoneName,omitempty"`
}

type CreateVolumeResponse struct {
	VolumeIds []string `json:"volumeIds,omitempty"`
}

type GetVolumeListArgs struct {
	InstanceId string
	ZoneName   string
}

type GetVolumeListResponse struct {
	Volumes     []Volume `json:"volumes"`
	Marker      string   `json:"marker"`
	IsTruncated bool     `json:"isTruncated"`
	NextMarker  string   `json:"nextMarker"`
	MaxKeys     int      `json:"maxKeys"`
}

type DescribeVolumeResponse struct {
	Volume *Volume `json:"volume"`
}

// AttachCDSVolumeArgs describe attachcds args
type AttachVolumeArgs struct {
	VolumeId   string `json:"-"`
	InstanceId string `json:"instanceId"`
}
type AttachVolumeResponse struct {
	VolumeAttachment *VolumeAttachment `json:"volumeAttachment"`
}

func (args *CreateVolumeArgs) validate() error {
	if args == nil {
		return fmt.Errorf("CreateVolumeArgs need args")
	}
	if args.StorageType == "" {
		return fmt.Errorf("CreateVolumeArgs need StorageType")
	}
	if args.Billing == nil {
		return fmt.Errorf("CreateVolumeArgs need Billing")
	}
	if args.CdsSizeInGB == 0 {
		return fmt.Errorf("CreateVolumeArgs need CdsSizeInGB")
	}
	return nil
}

// CreateVolumes create a volume
func (c *Client) CreateVolumes(args *CreateVolumeArgs, option *bce.SignOption) ([]string, error) {
	err := args.validate()
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := bce.NewRequest("POST", c.GetURL("v2/volume", params), bytes.NewBuffer(postContent))
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
	var blbsResp *CreateVolumeResponse
	err = json.Unmarshal(bodyContent, &blbsResp)

	if err != nil {
		return nil, err
	}
	return blbsResp.VolumeIds, nil

}

// DeleteVolume Delete a volume
func (c *Client) DeleteVolume(volumeId string, option *bce.SignOption) error {
	if volumeId == "" {
		return fmt.Errorf("DeleteVolume need a id")
	}
	params := map[string]string{
		"clientToken": c.GenerateClientToken(),
	}
	req, err := bce.NewRequest("DELETE", c.GetURL("v2/volume"+"/"+volumeId, params), nil)
	if err != nil {
		return err
	}
	_, err = c.SendRequest(req, option)
	if err != nil {
		return err
	}
	return nil
}

// GetVolumeList get all volumes
func (c *Client) GetVolumeList(args *GetVolumeListArgs, option *bce.SignOption) ([]Volume, error) {
	if args == nil {
		args = &GetVolumeListArgs{}
	}
	params := map[string]string{
		"zoneName":   args.ZoneName,
		"instanceId": args.InstanceId,
	}
	req, err := bce.NewRequest("GET", c.GetURL("v2/volume", params), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.SendRequest(req, option)
	bodyContent, err := resp.GetBodyContent()

	if err != nil {
		return nil, err
	}
	var blbsResp *GetVolumeListResponse
	err = json.Unmarshal(bodyContent, &blbsResp)

	if err != nil {
		return nil, err
	}
	return blbsResp.Volumes, nil
}

// DescribeVolume describe a volume
// More info see https://cloud.baidu.com/doc/BCC/API.html#.E6.9F.A5.E8.AF.A2.E7.A3.81.E7.9B.98.E8.AF.A6.E6.83.85
func (c *Client) DescribeVolume(id string, option *bce.SignOption) (*Volume, error) {
	if id == "" {
		return nil, fmt.Errorf("DescribeVolume need a id")
	}
	req, err := bce.NewRequest("GET", c.GetURL("v2/volume"+"/"+id, nil), nil)
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
	var ins DescribeVolumeResponse
	err = json.Unmarshal(bodyContent, &ins)

	if err != nil {
		return nil, err
	}
	return ins.Volume, nil
}

func (args *AttachVolumeArgs) validate() error {
	if args == nil {
		return fmt.Errorf("AttachCDSVolumeArgs need args")
	}
	if args.VolumeId == "" {
		return fmt.Errorf("AttachCDSVolumeArgs need VolumeId")
	}
	if args.InstanceId == "" {
		return fmt.Errorf("AttachCDSVolumeArgs need InstanceId")
	}
	return nil
}

// AttachCDSVolume attach a cds to vm
func (c *Client) AttachVolume(args *AttachVolumeArgs, option *bce.SignOption) (*VolumeAttachment, error) {
	err := args.validate()
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"attach":      "",
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	req, err := bce.NewRequest("PUT", c.GetURL("v2/volume"+"/"+args.VolumeId, params), bytes.NewBuffer(postContent))
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
	var blbsResp AttachVolumeResponse
	err = json.Unmarshal(bodyContent, &blbsResp)

	if err != nil {
		return nil, err
	}
	return blbsResp.VolumeAttachment, nil
}

// DetachCDSVolume detach a cds
// TODO: if a volume is detaching, need to wait
func (c *Client) DetachVolume(args *AttachVolumeArgs, option *bce.SignOption) error {
	err := args.validate()
	if err != nil {
		return err
	}
	params := map[string]string{
		"detach":      "",
		"clientToken": c.GenerateClientToken(),
	}
	postContent, err := json.Marshal(args)
	if err != nil {
		return err
	}
	req, err := bce.NewRequest("PUT", c.GetURL("v2/volume"+"/"+args.VolumeId, params), bytes.NewBuffer(postContent))
	if err != nil {
		return err
	}
	_, err = c.SendRequest(req, option)
	if err != nil {
		return err
	}
	return nil
}

// RollbackVolume rollback a volume
// TODO
func (c *Client) RollbackVolume() {

}

// PurchaseReservedVolume purchaseReserved a volume
// TODO
func (c *Client) PurchaseReservedVolume() {

}
