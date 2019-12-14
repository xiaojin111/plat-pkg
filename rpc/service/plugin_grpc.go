// +build grpc

package service

import (
	"github.com/micro/go-micro/client"
	cli "github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/server"
	srv "github.com/micro/go-micro/server/grpc"

	"os"
)

func init() {
	if len(os.Getenv("NO_GRPC_CLIENT")) == 0 {
		// set the default client
		client.DefaultClient = cli.NewClient()
		log.Infoln("gRPC client enabled.")
	}

	if len(os.Getenv("NO_GRPC_SERVER")) == 0 {
		// set the default server
		server.DefaultServer = srv.NewServer()
		log.Infoln("gRPC server enabled.")
	}
}
