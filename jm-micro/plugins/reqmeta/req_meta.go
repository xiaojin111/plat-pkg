package reqmeta

import (
	"net/http"
	"strings"

	"github.com/micro/cli"
	"github.com/micro/micro/plugin"
)

type reqMetaPlugin struct {
	metadata map[string]string
}

func (p *reqMetaPlugin) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{
			Name:   "client_meta",
			Usage:  "A list of key-value pairs defining metadata. k1=v1",
			EnvVar: "CLIENT_META",
		},
	}
}

func (p *reqMetaPlugin) Commands() []cli.Command {
	return nil
}

func (p *reqMetaPlugin) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			for k, v := range p.metadata {
				r.Header.Set(k, v)
			}

			// serve the request
			h.ServeHTTP(rw, r)
		})
	}
}

func (p *reqMetaPlugin) Init(ctx *cli.Context) error {
	// Parse the server options
	metadata := make(map[string]string)
	for _, d := range ctx.StringSlice("client_meta") {
		var key, val string
		parts := strings.Split(d, "=")
		key = parts[0]
		if len(parts) > 1 {
			val = strings.Join(parts[1:], "=")
		}
		metadata[key] = val
	}

	p.metadata = metadata

	return nil
}

func (p *reqMetaPlugin) String() string {
	return "client-meta"
}

func NewPlugin() plugin.Plugin {
	return NewReqMetaPlugin()
}

func NewReqMetaPlugin() plugin.Plugin {
	// create plugin
	p := &reqMetaPlugin{}

	return p
}
