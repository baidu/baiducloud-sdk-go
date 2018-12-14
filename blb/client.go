package blb

import (
	"github.com/baidu/baiducloud-sdk-go/bce"
)

// Endpoint contains all endpoints of Baidu Cloud BCC.
var Endpoint = map[string]string{
	"bj": "blb.bj.baidubce.com",
	"gz": "blb.gz.baidubce.com",
	"su": "blb.su.baidubce.com",
	"hk": "blb.hkg.baidubce.com",
	"bd": "blb.bd.baidubce.com",
}

// Client is the BLB client implemention for Baidu Cloud BLB API.
type Client struct {
	*bce.Client
}

func NewBLBClient(config *bce.Config) *Client {
	bceClient := bce.NewClient(config)
	return &Client{bceClient}
}

// GetURL generates the full URL of http request for Baidu Cloud BLB API.
func (c *Client) GetURL(version string, params map[string]string) string {
	host := c.Endpoint
	if host == "" {
		host = Endpoint[c.GetRegion()]
	}
	uriPath := version
	return c.Client.GetURL(host, uriPath, params)
}
