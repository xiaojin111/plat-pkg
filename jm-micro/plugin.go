package main

import (
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/jwt"
	"github.com/micro/micro/api"
)

func init() {
	// 将 JWT 插件注入到 api 服务之中
	err := api.Register(jwt.NewJWT())
	die(err)
}
