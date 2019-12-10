package service

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
)

func tlsCliFlags() []cli.Flag {
	return []cli.Flag{
		// TLS 相关
		cli.BoolFlag{
			Name:   "no_tls_client",
			Usage:  "Disable TLS client",
			EnvVar: "NO_TLS_CLIENT",
		},

		cli.BoolFlag{
			Name:   "no_tls_server",
			Usage:  "Disable TLS server",
			EnvVar: "NO_TLS_SERVER",
		},
	}
}

func setupTLS(c *cli.Context) error {
	// TLS Client
	if noTLSClient := c.Bool("no_tls_client"); noTLSClient {
		log.Warn("TLS is disabled for default client. INSECURE")
	} else {
		// 设置 Client 启用 TLS
		err := client.DefaultClient.Init(
			client.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			log.Fatalf("failed to enable client TLS: %v", err)
			return err
		}

		log.Info("TLS is enabled for default client.")
	}

	// TLS Server
	if noTLSServer := c.Bool("no_tls_server"); noTLSServer {
		log.Warn("TLS is disabled for default server. INSECURE")
	} else {
		// 设置 Server 启用 TLS
		err := server.DefaultServer.Init(
			server.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			log.Fatalf("failed to enable server TLS: %v", err)
			return err
		}

		log.Info("TLS is enabled for default server.")
	}

	return nil
}
