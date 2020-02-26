package service

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/transport"
)

func tlsCliFlags() []cli.Flag {
	return []cli.Flag{
		// TLS 相关
		&cli.BoolFlag{
			Name:    "no_tls_client",
			Usage:   "Disable TLS client",
			EnvVars: []string{"NO_TLS_CLIENT"},
		},

		&cli.BoolFlag{
			Name:    "no_tls_server",
			Usage:   "Disable TLS server",
			EnvVars: []string{"NO_TLS_SERVER"},
		},
	}
}

func setupTLS(c *cli.Context) error {
	// TLS Client
	if noTLSClient := c.Bool("no_tls_client"); noTLSClient {
		log.Warn("TLS client is disabled. INSECURE")
	} else {
		// 设置 Client 启用 TLS
		err := client.DefaultClient.Init(
			client.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			log.Fatalf("failed to enable TLS client: %v", err)
			return err
		}

		log.Info("TLS client is enabled.")
	}

	// TLS Server
	if noTLSServer := c.Bool("no_tls_server"); noTLSServer {
		log.Warn("TLS server is disabled. INSECURE")
	} else {
		// 设置 Server 启用 TLS
		err := server.DefaultServer.Init(
			server.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			log.Fatalf("failed to enable TLS server: %v", err)
			return err
		}

		log.Info("TLS server is enabled.")
	}

	return nil
}
