package main

import (
	mlog "github.com/jinmukeji/go-pkg/log"
	"github.com/jinmukeji/plat-pkg/rpc/service"

	"github.com/micro/go-micro/server"
)

const (
	// ServiceName 是本微服务的名称
	ServiceName = "minimal-example"
	// ServiceNamespace 是微服务的命名空间
	ServiceNamespace = "com.jinmuhealth.platform.srv"
)

var (
	log = mlog.StandardLogger()

	// Following values will be set during build.
	// Do NOT manually modify them.

	// ProductVersion is current product version.
	ProductVersion = "(n/a)"
	// GitCommit is the git commit short hash
	GitCommit = "(n/a)"
	// GoVersion is go compiler version `go version`
	GoVersion = "(n/a)"
	// BuildTime is go build time
	BuildTime = "(n/a)"
)

func main() {
	opts := &service.Options{
		Name:               ServiceName,
		Namespace:          ServiceNamespace,
		ProductVersion:     ProductVersion,
		GitCommit:          GitCommit,
		GoVersion:          GoVersion,
		BuildTime:          BuildTime,
		RegisterServerHook: register,
	}
	svc := service.CreateService(opts)

	// Run the service
	err := svc.Run()
	die(err)
}

func register(srv server.Server) error {
	// TODO: 注册 API 服务、设置订阅

	return nil
}

func die(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
