package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/grpc"
	grpc2 "github.com/micro/go-micro/v2/codec/grpc"
	icmd "github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/micro/v2/cmd"
)

const (
	Name        = "jm-micro"
	Description = "A microservice runtime for Jinmu Platform (based on Go Micro)"
)

var (
	Version = "Not Defined"
)

func main() {
	app := icmd.App()
	cmd.Setup(app, func(options *micro.Options) {
		grpc2.MaxMessageSize = 1024 * 1024 * 100
		_ = options.Client.Init(
			grpc.MaxRecvMsgSize(1024*1024*100),
			grpc.MaxSendMsgSize(1024*1024*100),
		)
	})

	err := icmd.Init(
		icmd.Name(Name),
		icmd.Description(Description),
		icmd.Version(Version),
	)

	die(err)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
