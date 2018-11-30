package cds

import (
	"github.com/baidu/baiducloud-sdk-go/bce"
)

// Endpoint contains all endpoints of Baidu Cloud BCC.
var Endpoint = map[string]string{
	"bj": "bcc.bj.baidubce.com",
	"gz": "bcc.gz.baidubce.com",
	"su": "bcc.su.baidubce.com",
	"bd": "bcc.bd.baidubce.com",
	"hk": "bcc.hkg.baidubce.com",
}

// Client is the bos client implemention for Baidu Cloud BOS API.
type Client struct {
	*bce.Client
}

func NewClient(config *bce.Config) *Client {
	bceClient := bce.NewClient(config)
	return &Client{bceClient}
}

// GetURL generates the full URL of http request for Baidu Cloud BOS API.
func (c *Client) GetURL(objectKey string, params map[string]string) string {
	host := c.Endpoint

	if host == "" {
		host = Endpoint[c.GetRegion()]
	}

	uriPath := objectKey

	return c.Client.GetURL(host, uriPath, params)
}
