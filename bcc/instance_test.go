package bcc

import (
	"testing"
	"time"

	"github.com/baidu/baiducloud-sdk-go/billing"
)

func TestCreateInstance(t *testing.T) {
	bccClient := NewTestClient()
	args := &CreateInstanceArgs{
		Name:    "sdk-test",
		ImageID: "m-Sr1bsnee",
		Billing: billing.Billing{
			PaymentTiming: "Postpaid",
		},
		AdminPass:             "thisistestpw123!",
		CPUCount:              1,
		MemoryCapacityInGB:    1,
		PurchaseCount:         1,
		NetworkCapacityInMbps: 1,
	}
	instances, err := bccClient.CreateInstances(args, nil)
	if err != nil {
		t.Errorf("Failed to create image, err: %+v", err)
	}
	t.Logf("Created instances: %+v", instances)
	time.Sleep(5 * time.Second)
	ins, err := bccClient.DescribeInstance(instances[0], nil)
	if err != nil {
		t.Errorf("Describe instance err: %+v", err)
	}
	t.Logf("instance info: %+v", ins)
}

func TestDeleteIntsance(t *testing.T) {
	bccClient := NewTestClient()
	InstanceID := "i-EhnYCRaH"
	err := bccClient.DeleteInstance(InstanceID, nil)
	if err != nil {
		t.Errorf("Failed to delete instance: %+v, err: %+v", InstanceID, err)
	} else {
		t.Logf("delete instances %+v successfully", InstanceID)
	}
}
