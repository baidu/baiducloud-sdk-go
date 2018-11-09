package clientset

import (
	"fmt"

	"github.com/baidu/baiducloud-sdk-go/bcc"
	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/baidu/baiducloud-sdk-go/blb"
	"github.com/baidu/baiducloud-sdk-go/eip"
	"github.com/baidu/baiducloud-sdk-go/vpc"
)

// Interface contains all methods of clients
type Interface interface {
	Bcc() *bcc.Client
	Blb() *blb.Client
	Eip() *eip.Client
	Vpc() *vpc.Client
}

// Clientset contains the clients for groups.
type Clientset struct {
	BccClient *bcc.Client
	BlbClient *blb.Client
	EipClient *eip.Client
	VpcClient *vpc.Client
}

// Bcc retrieves the BccClient
func (c *Clientset) Bcc() *bcc.Client {
	if c == nil {
		return nil
	}
	return c.BccClient
}

// Blb retrieves the BccClient
func (c *Clientset) Blb() *blb.Client {
	if c == nil {
		return nil
	}
	return c.BlbClient
}

// Eip retrieves the BccClient
func (c *Clientset) Eip() *eip.Client {
	if c == nil {
		return nil
	}
	return c.EipClient
}

// Vpc retrieves the VpcClient
func (c *Clientset) Vpc() *vpc.Client {
	if c == nil {
		return nil
	}
	return c.VpcClient
}

// NewFromConfig create a new Clientset for the given config.
func NewFromConfig(cfg *bce.Config) (*Clientset, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	var cs Clientset
	bccConfig := bcc.NewConfig(cfg)
	blbConfig := blb.NewConfig(cfg)
	eipConfig := eip.NewConfig(cfg)
	vpcConfig := vpc.NewConfig(cfg)
	cs.BccClient = bcc.NewClient(bccConfig)
	cs.BlbClient = blb.NewBLBClient(blbConfig)
	cs.EipClient = eip.NewEIPClient(eipConfig)
	cs.VpcClient = vpc.NewVPCClient(vpcConfig)
	return &cs, nil
}
