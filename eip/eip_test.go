package eip

import (
	"testing"
)

func TestCreateEip(t *testing.T) {
	bill := &Billing{
		PaymentTiming: "Postpaid",
		BillingMethod: "ByTraffic",
	}
	args := &CreateEipArgs{
		BandwidthInMbps: 998,
		Billing:         bill,
		Name:            "k8stestcgy",
	}
	ip, err := eipClient.CreateEip(args, nil)
	if err != nil {
		t.Error(err)
	}
	if ip != "180.181.3.133" {
		t.Error("ip error")
	}
}

var expectResizeEip = &ResizeEipArgs{
	BandwidthInMbps: 111,
	Ip:              "180.76.242.209",
}

func TestResizeEip(t *testing.T) {
	err := eipClient.ResizeEip(expectResizeEip, nil)
	if err != nil {
		t.Error(err)
	}
}

var expectBindEip = &BindEipArgs{
	Ip:           "180.76.247.62",
	InstanceType: "BCC",
	InstanceId:   "i-VAEyKKTh",
}
var expectUnbindEip = &EipArgs{
	Ip: "180.76.154.83",
}

func TestBindEip(t *testing.T) {
	err := eipClient.BindEip(expectBindEip, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUnbindEip(t *testing.T) {
	err := eipClient.UnbindEip(expectUnbindEip, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteEip(t *testing.T) {
	err := eipClient.DeleteEip(expectUnbindEip, nil)
	if err != nil {
		t.Error(err)
	}
}
func TestGetEips(t *testing.T) {
	eips, err := eipClient.GetEips(nil, nil)
	if err != nil {
		t.Error(err)
	}
	for _, eip := range eips {
		if eip.Eip != "180.181.3.133" && eip.Eip != "180.181.3.134" {
			t.Fatal("eip errpr")
		}
	}
}
