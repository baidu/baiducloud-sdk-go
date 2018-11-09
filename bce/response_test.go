package bce

import (
	"net/http"
	"testing"

	"github.com/baidu/baiducloud-sdk-go/util"
)

func TestGetBodyContent(t *testing.T) {
	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)

	if err != nil {
		t.Error(util.FormatTest("GetBodyContent", err.Error(), "nil"))
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		t.Error(util.FormatTest("GetBodyContent", err.Error(), "nil"))
	}

	bceResponse := NewResponse(resp)
	bodyContent, err := bceResponse.GetBodyContent()

	if err != nil {
		t.Error(util.FormatTest("GetBodyContent", err.Error(), "nil"))
	}

	if bodyContent == nil {
		t.Error(util.FormatTest("GetBodyContent", "nil", "not nil"))
	} else if string(bodyContent) == "" {
		t.Error(util.FormatTest("GetBodyContent", "empty string", "none empty string"))
	}
}
