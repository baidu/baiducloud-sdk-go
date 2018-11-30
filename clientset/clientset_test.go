package clientset

import (
	"testing"

	"github.com/baidu/baiducloud-sdk-go/bce"
)

var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

func TestNewFromConfig(t *testing.T) {
	cfg, err := bce.NewConfigFromFile("../aksk-test.json")
	if err != nil {
		t.Error(err)
	} else {
		cs, err := NewFromConfig(cfg)
		if err != nil {
			t.Error(err)
		} else {
			if cs.Bcc().AccessKeyID != credentials.AccessKeyID {
				t.Error("ak error")
			}
			if cs.Blb().AccessKeyID != credentials.AccessKeyID {
				t.Error("ak error")
			}
			if cs.Eip().AccessKeyID != credentials.AccessKeyID {
				t.Error("ak error")
			}
		}
	}
}
