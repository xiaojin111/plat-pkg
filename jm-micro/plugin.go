package main

import (
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/jwt"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/log"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
)

func init() {
	// 全局插件
	err := plugin.Register(log.NewPlugin(Name))
	die(err)

	// api 服务插件
	err = api.Register(jwt.NewPlugin())
	die(err)
}
