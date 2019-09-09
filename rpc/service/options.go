package service

import (
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

type RegisterServerFunc func(srv server.Server) error

type Options struct {
	// 微服务的名称
	Name string
	// 微服务的命名空间
	Namespace string

	// ProductVersion is current product version.
	ProductVersion string
	// GitCommit is the git commit short hash
	GitCommit string
	// GoVersion is go compiler version `go version`
	GoVersion string
	// BuildTime is go build time
	BuildTime string

	// 自定义HandlerWrapper，在标准 HandlerWrapper 之前注册
	PreServerHandlerWrappers []server.HandlerWrapper

	// 自定义HandlerWrapper，在标准 HandlerWrapper 之后注册
	PostServerHandlerWrappers []server.HandlerWrapper

	// 注册 micro.Server
	RegisterServer RegisterServerFunc

	// 自定义 Client Wrapper，在标准 Wrapper 之前注册
	PreClientWrappers []client.Wrapper

	// 自定义 Client Wrapper，在标准 Wrapper 之前注册
	PostClientWrappers []client.Wrapper
}

// ServiceFQDN 返回微服务的全名
func (opts *Options) FQDN() string {
	return fmt.Sprintf("%s.%s", opts.Namespace, opts.Name)
}

// ServiceMetadata 返回微服务的 metadata
func (opts *Options) ServiceMetadata() map[string]string {
	return map[string]string{
		"ProductVersion": opts.ProductVersion,
		"GitCommit":      opts.GitCommit,
		"GoVersion":      opts.GoVersion,
		"BuildTime":      opts.BuildTime,
	}
}
