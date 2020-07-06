package web

import (
	"github.com/gin-gonic/gin"
	mlog "github.com/jinmukeji/go-pkg/v2/log"
)

// PrintRoutes 打印当前已经注册的路由
func PrintRoutes(eg *gin.Engine) {
	log := mlog.StandardLogger()
	rs := eg.Routes()
	for _, r := range rs {
		log.Infof("%s\t%s\t--> %s", r.Method, r.Path, r.Handler)
	}
}
