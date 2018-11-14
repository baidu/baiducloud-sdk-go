package main

import (
	"fmt"
	"time"

	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/baidu/baiducloud-sdk-go/eip"
)

func main() {
	cred, err := bce.NewCredentialsFromFile("example/main/aksk.json")
	if err != nil {
		panic(err)
	}
	bceConf := &bce.Config{
		Credentials: cred,
		Checksum:    true,
		Timeout:     5 * time.Second,
		Region:      "bj",
	}

	eipConf := eip.NewConfig(bceConf)
	eipClient := eip.NewEIPClient(eipConf)
	eipClient.SetDebug(true)
	eips, err := eipClient.GetEips(nil)
	if err != nil {
		panic(err)
	}
	for idx, ip := range eips {
		fmt.Printf("%d, %s\n", idx, ip.Eip)
	}
	return
}
