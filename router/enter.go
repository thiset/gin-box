package router

import (
	"gin-box/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	apiRouterGroup := r.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.settingsRouter()
	return r
}
