package bce

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/baidu/baiducloud-sdk-go/util"
)

func TestError(t *testing.T) {
	bceError := &Error{
		StatusCode: 500,
		Code:       "StatusInternalServerError",
		Message:    "failed",
		RequestID:  "123",
	}

	result := bceError.Error()
	expected := "Error Message: \"failed\", Error Code: \"StatusInternalServerError\", Status Code: 500, Request Id: \"123\""

	if result != expected {
		t.Error(util.FormatTest("Error", result, expected))
	}
}

func TestBuildError(t *testing.T) {
	resp := &Response{BodyContent: []byte{}}
	err := buildError(resp)

	if _, ok := err.(*Error); ok {
		t.Error(util.FormatTest("buildError", "bceError", "error"))
	}

	bceError := &Error{
		StatusCode: 500,
		Code:       strconv.Itoa(http.StatusInternalServerError),
		Message:    "failed",
		RequestID:  "123",
	}
	byteArray, err := json.Marshal(bceError)

	if err != nil {
		t.Error(util.FormatTest("buildError", err.Error(), "nil"))
	}

	httpResponse := &http.Response{StatusCode: http.StatusInternalServerError}
	resp = &Response{BodyContent: byteArray, Response: httpResponse}
	err = buildError(resp)

	if _, ok := err.(*Error); !ok {
		t.Error(util.FormatTest("buildError", "error", "bceError"))
	}

	resp = &Response{BodyContent: []byte("Unknown Error"), Response: httpResponse}
	err = buildError(resp)

	if _, ok := err.(*Error); ok {
		t.Error(util.FormatTest("buildError", "bceError", "error"))
	}
}
