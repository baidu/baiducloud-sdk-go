package cds

import (
	"fmt"
	"testing"

	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/baidu/baiducloud-sdk-go/util"
)

var expectBill = &bce.Billing{
	PaymentTiming: "Postpaid",
	BillingMethod: "ByTraffic",
}
var expectCreateVolumeArgs = &CreateVolumeArgs{
	PurchaseCount: 1,
	Billing:       expectBill,
	StorageType:   STORAGE_TYPE_STD1,
	CdsSizeInGB:   10,
}

func TestCreateVolumes(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	_, err := bccClient.CreateVolumes(expectCreateVolumeArgs, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestGetVolumeList(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	list, err := bccClient.GetVolumeList(nil, nil)
	if err != nil {
		t.Error(err)
	}
	for _, v := range list {
		fmt.Println(v.Id)
	}
}

func TestDescribeVolume(t *testing.T) {
	// ts := httptest.NewServer(InstancesHandler())
	// defer ts.Close()
	// bccClient.Endpoint = ts.URL
	ins, err := bccClient.DescribeVolume("v-31wjHWIU", nil)
	if err != nil {
		t.Error(util.FormatTest("TestDescribeVolume", err.Error(), "nil"))
	}
	fmt.Println(ins.Id)
}

var expectAttach = &AttachVolumeArgs{
	VolumeId:   "v-JCvK3cpI",
	InstanceId: "i-NN0KeMyw",
}

func TestAttachVolume(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	bccClient.SetDebug(true)
	att, err := bccClient.AttachVolume(expectAttach, nil)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(att.Device)
	}

}

func TestDetachVolume(t *testing.T) {
	bccClient.SetDebug(true)
	err := bccClient.DetachVolume(expectAttach, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteVolume(t *testing.T) {
	bccClient.SetDebug(true)
	err := bccClient.DeleteVolume("v-JCvK3cpI", nil)
	if err != nil {
		t.Error(err)
	}
}
