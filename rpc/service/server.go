package service

import (
	"github.com/micro/go-micro/v2/server"
)

type RegisterServerFunc func(srv server.Server) error

func RegisterServer(srv server.Server, registers ...RegisterServerFunc) error {
	err := srv.Init(
		// Graceful shutdown of a service using the server.Wait option
		// The server deregisters the service and waits for handlers to finish executing before exiting.
		server.Wait(nil),
	)
	if err != nil {
		return err
	}

	for _, r := range registers {
		err = r(srv)
		if err != nil {
			return err
		}
	}

	return nil
}
