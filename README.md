# BaiduCloud SDK for Go
[![GoDoc](https://godoc.org/github.com/baidu/baiducloud-sdk-go?status.svg)](https://godoc.org/github.com/baidu/baiducloud-sdk-go)
[![Build Status](https://travis-ci.org/baidu/baiducloud-sdk-go.svg?branch=master)](https://travis-ci.org/baidu/baiducloud-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/baidu/baiducloud-sdk-go)](https://goreportcard.com/report/github.com/baidu/baiducloud-sdk-go)

> This sdk is still under development and may be changed frequently.

baiudcloud-sdk-go is the official BaiudCloud SDK for the Go programming language, supporting convenient access to cloud resources, including [Baidu Cloud Compute (BCC)](https://cloud.baidu.com/doc/BCC/ProductDescription.html), [Baidu Load Balancer (BLB)](https://cloud.baidu.com/doc/BLB/ProductDescription.html), [Baidu Object Storage (BOS)](https://cloud.baidu.com/doc/BOS/ProductDescription.html), [Cloud Disk Service (CDS)](https://cloud.baidu.com/doc/CDS/ProductDescription.html), [Elastic IP (EIP)](https://cloud.baidu.com/doc/EIP/ProductDescription.html), [Virtual Private Cloud (VPC)](https://cloud.baidu.com/doc/VPC/ProductDescription.html), etc.

## Installing

Go >= 1.8 is required.

```bash
go get -u github.com/baidu/baiducloud-sdk-go/...
```

## Usage

This example shows a complete example to list all eips in some region.

```golang
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

	eipClient := eip.NewEIPClient(bceConf)
	eipClient.SetDebug(true)
	eips, err := eipClient.GetEips(nil, nil)
	if err != nil {
		panic(err)
	}
	for idx, ip := range eips {
		fmt.Printf("%d, %s\n", idx, ip.Eip)
	}
	return
}
```

## Contributors

- Guoyao Wu @guoyao
- Guoyan Chen @drinktee
- Yuxiao Song
- Hongbin Mao @hello2mao
- Yang Meng @m3ngyang
- Weidong Cai @cwdsuzhou

## License

This SDK is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0).