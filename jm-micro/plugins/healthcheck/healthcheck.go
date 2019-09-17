package healthcheck

import (
	"net/http"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
)

type healthCheck struct {
	path string
}

func (p *healthCheck) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "healthcheck_path",
			Usage:       "Health check URL path. Specified with leading slash e.g /_health",
			EnvVar:      "HEALTHCHECK_PATH",
			Value:       "/_health", // default path
			Destination: &(p.path),
		},
	}
}

func (p *healthCheck) Commands() []cli.Command {
	return nil
}

func (p *healthCheck) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(p.path) > 0 && r.URL.Path == p.path && r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK) // 200
				//nolint:errcheck
				w.Write([]byte("OK"))
				return
			}

			// serve request
			h.ServeHTTP(w, r)
		})
	}
}

func (p *healthCheck) Init(ctx *cli.Context) error {
	return nil
}

func (p *healthCheck) String() string {
	return "HealthCheck"
}

func NewPlugin() plugin.Plugin {
	return &healthCheck{}
}
