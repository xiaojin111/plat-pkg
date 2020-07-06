package service

import (
	"fmt"

	"github.com/micro/cli/v2"
)

type options struct {
	// 微服务的命名空间
	Namespace string
	// 微服务的名称
	Name string

	// ProductVersion is current product version.
	ProductVersion string
	// GitCommit is the git commit short hash
	GitCommit string
	// GoVersion is go compiler version `go version`
	GoVersion string
	// BuildTime is go build time
	BuildTime string

	// Flags are CLI flags
	Flags []cli.Flag

	// CliPreAction 在标准 Action 之前调用
	CliPreAction func(c *cli.Context)

	// CliPostAction 在标准 Action 之后调用
	CliPostAction func(c *cli.Context)
}

// FQDN 返回微服务的全名
func (opts *options) FQDN() string {
	if opts == nil {
		return ""
	}

	return fmt.Sprintf("%s.%s", opts.Namespace, opts.Name)
}

// ServiceMetadata 返回微服务的 metadata
func (opts *options) ServiceMetadata() map[string]string {
	if opts == nil {
		return nil
	}

	return map[string]string{
		"ProductVersion": opts.ProductVersion,
		"GitCommit":      opts.GitCommit,
		"GoVersion":      opts.GoVersion,
		"BuildTime":      opts.BuildTime,
	}
}

func (opts *options) GetProductName() string {
	return opts.FQDN()
}

func (opts *options) GetProductVersion() string {
	if opts == nil {
		return ""
	}

	return opts.ProductVersion
}

func (opts *options) GetGitCommit() string {
	if opts == nil {
		return ""
	}

	return opts.GitCommit
}

func (opts *options) GetBuildTime() string {
	if opts == nil {
		return ""
	}

	return opts.BuildTime
}

func (opts *options) GetGoVersion() string {
	if opts == nil {
		return ""
	}

	return opts.GoVersion
}
