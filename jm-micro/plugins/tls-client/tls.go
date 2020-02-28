package tls

// FIXME: go-micro/v2 默认采用 gRPC 方式，TLS设定方式此处不再适用，考虑后续移除掉。

import (
	"net/http"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/transport"

	mlog "github.com/jinmukeji/go-pkg/v2/log"
)

var (
	log *mlog.Logger = mlog.StandardLogger()
)

type tlsPlugin struct {
	insecure bool
}

func (p *tlsPlugin) Flags() []cli.Flag {
	return []cli.Flag{
		// 日志相关
		&cli.BoolFlag{
			Name:        "no_tls_client",
			Usage:       "Disable TLS client",
			EnvVars:     []string{"NO_TLS_CLIENT"},
			Destination: &(p.insecure),
		},
	}
}

func (p *tlsPlugin) Commands() []*cli.Command {
	return nil
}

func (p *tlsPlugin) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		// 什么都不包装，透传
		return h
	}
}

func (p *tlsPlugin) Init(ctx *cli.Context) error {
	if !p.insecure {
		err := client.DefaultClient.Init(
			client.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			return err
		}

		log.Info("TLS is enabled for default client.")
	} else {
		log.Warn("TLS is disabled for default client. INSECURE")
	}

	return nil
}

func (p *tlsPlugin) String() string {
	return "tls-client"
}

func NewPlugin() plugin.Plugin {
	return NewTLS()
}

func NewTLS() plugin.Plugin {
	return &tlsPlugin{}
}
