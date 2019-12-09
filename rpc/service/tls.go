package service

import (
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
)

func init() {
	if len(os.Getenv("NO_TLS_CLIENT")) == 0 {
		// 设置 Client 启用 TLS
		err := client.DefaultClient.Init(
			client.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			log.Fatalf("failed to set client TLS: %v", err)
		}

		log.Info("TLS for default client is enabled.")
	} else {
		log.Info("TLS for default client is disabled. INSECURE")
	}

	if len(os.Getenv("NO_TLS_SERVER")) == 0 {
		// 设置 Server 启用 TLS
		err := server.DefaultServer.Init(
			server.Transport(
				transport.NewTransport(transport.Secure(true)),
			),
		)
		if err != nil {
			log.Fatalf("failed to set server TLS: %v", err)
		}

		log.Info("TLS for default server is enabled.")
	} else {
		log.Info("TLS for default server is disabled. INSECURE")
	}
}
