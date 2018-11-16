package bcc

import (
	"os"

	"github.com/baidu/baiducloud-sdk-go/bce"
)

// Modify with your AccessKeyID and SecretAccessKey
var (
	TestAccessKeyID     = os.Getenv("AccessKeyId")
	TestSecretAccessKey = os.Getenv("SecretAccessKey")
	TestRegion          = os.Getenv("Region")
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		config := bce.NewConfigWithParams(TestAccessKeyID, TestSecretAccessKey, TestRegion)
		testClient = NewClient(config)
	}
	return testClient
}
