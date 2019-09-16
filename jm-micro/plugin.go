package main

import (
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/cid"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/configloader"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/jwt"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/log"
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/go-plugins/micro/gzip"
	"github.com/micro/go-plugins/micro/metadata"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
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

	// api 服务插件
	err = api.Register(gzip.NewPlugin())
	die(err)

	err = api.Register(cid.NewPlugin())
	die(err)

	err = api.Register(jwt.NewPlugin())
	die(err)
}
