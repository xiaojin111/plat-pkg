package main

import (
	mlog "github.com/jinmukeji/go-pkg/log"
	"github.com/jinmukeji/plat-pkg/rpc/service"

	echosvc "github.com/jinmukeji/plat-pkg/rpc/service/minimal-example/handler"
	echopb "github.com/jinmukeji/proto/gen/micro/idl/examples/echo/v1"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
)

const (
	// ServiceName 是本微服务的名称
	ServiceName = "template-service"
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
		Name:                      ServiceName,
		Namespace:                 ServiceNamespace,
		ProductVersion:            ProductVersion,
		GitCommit:                 GitCommit,
		GoVersion:                 GoVersion,
		BuildTime:                 BuildTime,
		RegisterServer:            register,
		PreServerHandlerWrappers:  preHandlerWrappers(),
		PostServerHandlerWrappers: postHandlerWrappers(),
		ServiceOptions: []micro.Option{
			// 设置启用 TLS
			// micro 底层将同时设置 Sever 与 Client 启用 TLS
			micro.Transport(
				// create new transport
				transport.NewTransport(
					// set to automatically secure
					transport.Secure(true),
				),
			),
		},
	}
	svc := service.CreateService(opts)

	// Run the service
	err := svc.Run()
	die(err)
}

func register(srv server.Server) error {
	// TODO: 注册自定义 API 服务、设置订阅
	echoAPI := &echosvc.EchoAPIService{}
	if err := echopb.RegisterEchoAPIHandler(srv, echoAPI); err != nil {
		return err
	}

	return nil
}

func preHandlerWrappers() []server.HandlerWrapper {
	// TODO: 注册自定义 HandlerWrapper, 在标准 HandlerWrapper 之前注册
	return nil
}

func postHandlerWrappers() []server.HandlerWrapper {
	// TODO: 注册自定义 HandlerWrapper, 在标准 HandlerWrapper 之后注册

	return nil
}

func die(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
