package cid

import (
	"net/http"

	"github.com/micro/cli/v2"
	"github.com/micro/micro/v2/plugin"

	rc "github.com/jinmukeji/plat-pkg/v2/rpc/cid"
)

type cidPlugin struct {
	headerKey       string // HTTP Request Header 中的 cid 使用的 key
	ignoreOriginCid bool   // 是否忽略原始 Request 中包含的 cid，为每次请求强制重新生成新 cid 并覆盖
}

func (p *cidPlugin) Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "cid_header_key",
			Usage:       "cid HTTP header key",
			EnvVars:     []string{"CID_HEADER_KEY"},
			Value:       rc.MetaCidKey,
			Destination: &(p.headerKey),
		},

		&cli.BoolFlag{
			Name:        "cid_ignore_origin",
			Usage:       "whether or not to ignore cid from origin request and generate new one to overwrite",
			EnvVars:     []string{"CID_IGNORE_ORIGIN"},
			Destination: &(p.ignoreOriginCid),
		},
	}
}

func (p *cidPlugin) Commands() []*cli.Command {
	return nil
}

func (p *cidPlugin) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			cid := r.Header.Get(p.headerKey)

			// Header 中提取的cid为空，或者开启强制生成cid
			if cid == "" || p.ignoreOriginCid {
				// 生成新的 cid，并注入到 Request Header 之中
				cid = rc.NewCid()
				r.Header.Set(p.headerKey, cid)
			}

			// 将 cid 写入到 Response Header 之中
			rw.Header().Add(p.headerKey, cid)

			// serve the request
			h.ServeHTTP(rw, r)
		})
	}
}

func (p *cidPlugin) Init(ctx *cli.Context) error {
	return nil
}

func (p *cidPlugin) String() string {
	return "cid"
}

func NewPlugin() plugin.Plugin {
	return NewCidPlugin()
}

func NewCidPlugin() plugin.Plugin {
	// create plugin
	p := &cidPlugin{}

	return p
}
