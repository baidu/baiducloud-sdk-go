package eip

import "github.com/baidu/baiducloud-sdk-go/bce"

// Endpoint contains all endpoints of Baidu Cloud BCC.
var Endpoint = map[string]string{
	"bj": "eip.bj.baidubce.com",
	"gz": "eip.gz.baidubce.com",
}

// Client is the bos client implemention for Baidu Cloud BOS API.
type Client struct {
	*bce.Client
}

func NewEIPClient(config *bce.Config) *Client {
	bceClient := bce.NewClient(config)
	return &Client{bceClient}
}

// GetURL generates the full URL of http request for Baidu Cloud BOS API.
func (c *Client) GetURL(version string, params map[string]string) string {
	host := c.Endpoint
	if host == "" {
		host = Endpoint[c.GetRegion()]
	}
	uriPath := version
	return c.Client.GetURL(host, uriPath, params)
}
