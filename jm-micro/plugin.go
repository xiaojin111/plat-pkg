package main

import (
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/cid"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/jwt"
	"github.com/jinmukeji/plat-pkg/jm-micro/plugins/log"
	fstore "github.com/jinmukeji/plat-pkg/jwt/keystore/file"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
)

func init() {
	// 全局插件
	err := plugin.Register(log.NewPlugin(Name))
	die(err)

	// api 服务插件
	err = api.Register(cid.NewPlugin())
	die(err)

	s := fstore.NewFileStore()
	err = s.Load("../jwt/tools/testdata", "app-test1")
	if err != nil {
		panic(err)
	}
	err = api.Register(jwt.NewPlugin(s))
	die(err)
}
