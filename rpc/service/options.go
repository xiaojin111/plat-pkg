package service

import (
	"fmt"

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

	RegisterServerHook RegisterServerFunc
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
