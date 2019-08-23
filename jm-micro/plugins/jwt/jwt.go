package jwt

import (
	"log"
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
)

type jwt struct{}

func (p *jwt) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "jwt_source",
			Usage:  "JWT Sources",
			EnvVar: "JWT_SOURCE",
		},
	}
}

func (p *jwt) Commands() []cli.Command {
	return nil
}

func (p *jwt) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("x-jwt")
			log.Printf("JWT Token: %s", token)

			// serve the request
			h.ServeHTTP(rw, r)
		})
	}
}

func (w *jwt) Init(ctx *cli.Context) error {
	return nil
}

func (p *jwt) String() string {
	return "jwt"
}

func NewPlugin() plugin.Plugin {
	return NewJWT()
}

func NewJWT() plugin.Plugin {
	// create plugin
	p := &jwt{}

	return p
}
