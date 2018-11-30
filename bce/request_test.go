package bce

import (
	"strconv"
	"testing"

	"github.com/baidu/baiducloud-sdk-go/util"
)

func TestAddHeadersForRequest(t *testing.T) {
	headers := map[string]string{"Host": "guoyao.me", "Content-Type": "text/plain"}
	req, err := NewRequest("GET", "guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("AddHeaders", err.Error(), "nil"))
	}

	req.AddHeaders(headers)

	if len(req.Header) != len(headers) {
		t.Error(util.FormatTest("AddHeaders", strconv.Itoa(len(req.Header)), strconv.Itoa(len(headers))))
	}
}

func TestAddHeaderForRequest(t *testing.T) {
	req, err := NewRequest("GET", "guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("addHeader", err.Error(), "nil"))
	}

	req.addHeader("Host", "guoyao.me")

	if len(req.Header) != 1 {
		t.Error(util.FormatTest("addHeader", strconv.Itoa(len(req.Header)), strconv.Itoa(1)))
	}
}

func TestSetHeaders(t *testing.T) {
	headers := map[string]string{"Host": "guoyao.me", "Content-Type": "text/plain"}
	req, err := NewRequest("GET", "guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("SetHeaders", err.Error(), "nil"))
	}

	req.Header = map[string][]string{
		"Host":            {"tocloud.org"},
		"Accept-Encoding": {"gzip, deflate"},
	}
	req.SetHeaders(headers)

	if len(req.Header) != 3 {
		t.Error(util.FormatTest("SetHeaders", strconv.Itoa(len(req.Header)), strconv.Itoa(3)))
	}

	if len(req.Header["Host"]) != 1 {
		t.Error(util.FormatTest("SetHeaders", strconv.Itoa(len(req.Header["Host"])), strconv.Itoa(1)))
	}
}

func TestSetHeader(t *testing.T) {
	req, err := NewRequest("GET", "guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("setHeader", err.Error(), "nil"))
	}

	req.Header = map[string][]string{
		"Host":            {"tocloud.org"},
		"Accept-Encoding": {"gzip, deflate"},
	}
	req.setHeader("Host", "guoyao.me")

	if len(req.Header) != 2 {
		t.Error(util.FormatTest("setHeader", strconv.Itoa(len(req.Header)), strconv.Itoa(2)))
	}

	if len(req.Header["Host"]) != 1 {
		t.Error(util.FormatTest("setHeaders", strconv.Itoa(len(req.Header["Host"])), strconv.Itoa(1)))
	}
}

func TestClearHeaders(t *testing.T) {
	req, err := NewRequest("GET", "guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("clearHeaders", err.Error(), "nil"))
	}

	req.Header = map[string][]string{
		"Host":            {"tocloud.org"},
		"Accept-Encoding": {"gzip, deflate"},
	}
	req.clearHeaders()

	if len(req.Header) != 0 {
		t.Error(util.FormatTest("clearHeaders", strconv.Itoa(len(req.Header)), strconv.Itoa(0)))
	}
}

func TestPrepareHeaders(t *testing.T) {
	req, err := NewRequest("GET", "http://guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("clearHeaders", err.Error(), "nil"))
	}

	signOption := &SignOption{
		Headers: map[string]string{"Content-Type": "text/plain"},
	}
	req.prepareHeaders(signOption)

	if !util.MapContains(signOption.Headers, func(key, value string) bool {
		return key == "host" && value == "guoyao.me"
	}) {
		t.Error(util.FormatTest("prepareHeaders", "no host", "host"))
	}

	req, err = NewRequest("GET", "http://guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("clearHeaders", err.Error(), "nil"))
	}

	req.Header = map[string][]string{
		"Host":            {"tocloud.org"},
		"Accept-Encoding": {"gzip, deflate"},
	}
	signOption = &SignOption{
		Headers: map[string]string{"Host": "tocloud.org", "Content-Type": "text/plain"},
	}
	req.prepareHeaders(signOption)

	if req.Header.Get("Host") != "guoyao.me" {
		t.Error(util.FormatTest("prepareHeaders", req.Header.Get("Host"), "guoyao.me"))
	}
}

func TestToCanonicalHeaderString(t *testing.T) {
	req, err := NewRequest("GET", "http://guoyao.me", nil)

	if err != nil {
		t.Error(util.FormatTest("toCanonicalHeaderString", err.Error(), "nil"))
	}

	req.Header = map[string][]string{
		"Host":            {"tocloud.org"},
		"Accept-Encoding": {"gzip, deflate"},
	}

	signOption := NewSignOption(
		"2015-04-27T08:23:49Z",
		ExpirationPeriodInSeconds,
		nil,
		nil,
	)

	canonicalString := req.toCanonicalHeaderString(signOption)
	expected := "host:tocloud.org"

	if canonicalString != expected {
		t.Error(util.FormatTest("toCanonicalHeaderString", canonicalString, expected))
	}

	signOption = NewSignOption(
		"2015-04-27T08:23:49Z",
		ExpirationPeriodInSeconds,
		nil,
		[]string{"Accept-Encoding"},
	)
	canonicalString = req.toCanonicalHeaderString(signOption)
	expected = "accept-encoding:gzip%2C%20deflate"

	if canonicalString != expected {
		t.Error(util.FormatTest("toCanonicalHeaderString", canonicalString, expected))
	}

	req.Header = map[string][]string{
		"Last-Modified":   {"2015-04-27T08:23:49Z"},
		"Accept-Encoding": {"gzip, deflate"},
	}
	signOption = NewSignOption(
		"2015-04-27T08:23:49Z",
		ExpirationPeriodInSeconds,
		nil,
		nil,
	)
	canonicalString = req.toCanonicalHeaderString(signOption)
	expected = ""

	if canonicalString != expected {
		t.Error(util.FormatTest("toCanonicalHeaderString", canonicalString, expected))
	}
}

func TestIsCanonicalHeader(t *testing.T) {
	expected := true
	result := isCanonicalHeader("content-type")

	if result != expected {
		t.Error(util.FormatTest("isCanonicalHeader", strconv.FormatBool(result), strconv.FormatBool(expected)))
	}

	expected = false
	result = isCanonicalHeader("Accept-Encoding")

	if result != expected {
		t.Error(util.FormatTest("isCanonicalHeader", strconv.FormatBool(result), strconv.FormatBool(expected)))
	}
}
