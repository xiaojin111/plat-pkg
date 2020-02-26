package main

import (
	icmd "github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/micro/v2/cmd"
)

const (
	Name        = "jm-micro"
	Description = "A microservice runtime for Jinmu Platform (based on Go Micro)"
)

var (
	Version = "Not Defined"
)

func main() {
	app := icmd.App()
	cmd.Setup(app)

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
