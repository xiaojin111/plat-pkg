package main

import (
	icmd "github.com/micro/go-micro/config/cmd"
	"github.com/micro/micro/cmd"
)

const (
	name        = "jm-micro"
	description = "A microservice runtime for Jinmu Platform (based on Go Micro)"
	version     = "1.0.0"
)

func main() {
	cmd.Setup(icmd.App())

	err := icmd.Init(
		icmd.Name(name),
		icmd.Description(description),
		icmd.Version(version),
	)
	die(err)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
