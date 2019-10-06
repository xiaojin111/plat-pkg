package tls

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/transport"

	mlog "github.com/jinmukeji/go-pkg/log"
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
		cli.BoolFlag{
			Name:        "insecure",
			Usage:       "insecure without TLS client",
			EnvVar:      "INSECURE",
			Destination: &(p.insecure),
		},
	}
}

func (p *tlsPlugin) Commands() []cli.Command {
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
		client.DefaultClient.Init(
			client.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)

		log.Info("TLS is enabled.")
	} else {
		log.Warn("TLS is disabled. INSECURE")
	}

	return nil
}

func (p *tlsPlugin) String() string {
	return "tls"
}

func NewPlugin() plugin.Plugin {
	return NewTLS()
}

func NewTLS() plugin.Plugin {
	return &tlsPlugin{}
}
