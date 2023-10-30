package main

import (
	"gin-box/core"
	"gin-box/global"
	"gin-box/router"
)

func main() {
	// 读取配置文件
	core.InitCore()
	// 初始化Log
	global.Log = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	r := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("服务运行在:%s", addr)
	if err := r.Run(addr); err != nil {
		global.Log.Errorf("服务启动失败，失败原因：%s", err)
		return
	}
}
