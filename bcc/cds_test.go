package bcc

import (
	"fmt"
	"testing"

	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/baidu/baiducloud-sdk-go/util"
)

func TestDeleteVolume(t *testing.T) {
	// ts := httptest.NewServer(InstancesHandler())
	// defer ts.Close()
	// bccClient.SetDebug(true)
	// bccClient.Endpoint = ts.URL
	err := bccClient.DeleteVolume("v-MK288vVC")
	if err != nil {
		t.Error(util.FormatTest("DeleteVolume", err.Error(), "nil"))
	}
}

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
	_, err := bccClient.CreateVolumes(expectCreateVolumeArgs)
	if err != nil {
		t.Error(err)
	}
}

func TestGetVolumeList(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	list, err := bccClient.GetVolumeList(nil)
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
	ins, err := bccClient.DescribeVolume("v-31wjHWIU")
	if err != nil {
		t.Error(util.FormatTest("TestDescribeVolume", err.Error(), "nil"))
	}
	fmt.Println(ins.Id)
}

var expectAttach = &AttachCDSVolumeArgs{
	VolumeId:   "v-JCvK3cpI",
	InstanceId: "i-NN0KeMyw",
}

func TestAttachCDSVolume(t *testing.T) {
	// ts := httptest.NewServer(EipHandler())
	// defer ts.Close()
	// eipClient.Endpoint = ts.URL
	bccClient.SetDebug(true)
	att, err := bccClient.AttachCDSVolume(expectAttach)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(att.Device)
	}

}

func TestDetachCDSVolume(t *testing.T) {
	bccClient.SetDebug(true)
	err := bccClient.DetachCDSVolume(expectAttach)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteCDSVolume(t *testing.T) {
	bccClient.SetDebug(true)
	err := bccClient.DeleteCDS("v-JCvK3cpI")
	if err != nil {
		t.Error(err)
	}
}
