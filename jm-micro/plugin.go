package main

import (
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/cid"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/configloader"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/healthcheck"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/jwt"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/log"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/tls-client"

	// "github.com/jinmukeji/plat-pkg/jm-micro/plugins/whitelist"
	"github.com/micro/go-plugins/micro/cors"

	// "github.com/micro/go-plugins/micro/gzip"
	"github.com/micro/go-plugins/micro/metadata"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
	"github.com/micro/micro/proxy"
	"github.com/micro/micro/web"
)

func init() {
	// 全局插件
	err := plugin.Register(log.NewPlugin(Name))
	die(err)

	err = plugin.Register(configloader.NewPlugin())
	die(err)

	err = plugin.Register(metadata.NewPlugin())
	die(err)

	err = plugin.Register(cors.NewPlugin())
	die(err)

	// proxy 服务插件
	err = proxy.Register(tls.NewPlugin())
	die(err)

	// web 服务插件
	err = web.Register(tls.NewPlugin())
	die(err)

	// // Proxy
	// err = proxy.Register(tls.NewPlugin())
	// die(err)

	// // web 服务插件
	// err = web.Register(tls.NewPlugin())
	// die(err)

	// // api 服务插件

	// err = api.Register(tls.NewPlugin())
	// die(err)

	// micro gzip 插件存在 bug，当 response 数据量过小的时候，压缩后的数据丢失
	// err = api.Register(gzip.NewPlugin())
	// die(err)
	err = api.Register(healthcheck.NewPlugin())
	die(err)

	err = api.Register(cid.NewPlugin())
	die(err)

	err = api.Register(jwt.NewPlugin())
	die(err)

	// TODO: 白名单插件
	// err = api.Register(whitelist.NewRPCWhitelist("com.jinmuhealth.platform.srv.template-service1"))
	// die(err)
}
