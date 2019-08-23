package main

import (
	icmd "github.com/micro/go-micro/config/cmd"
	"github.com/micro/micro/cmd"
)

const (
	Name        = "jm-micro"
	Description = "A microservice runtime for Jinmu Platform (based on Go Micro)"
)

var (
	Version = "Not Defined"
)

func main() {
	cmd.Setup(icmd.App())

	err := icmd.Init(
		icmd.Name(Name),
		icmd.Description(Description),
		icmd.Version(Version),
	)
	die(err)
}

func die(err error) {
	if err != nil {
		panic(err)
	}
}
