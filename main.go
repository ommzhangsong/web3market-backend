package main

import (
	"gin-campus-market/config"
	"gin-campus-market/models"
	"gin-campus-market/router"
)

func main() {
	// 1. 先读 YAML
	config.InitConfig()

	// 2. 再连数据库
	models.InitDB()

	// 3. 启动路由
	r := router.SetupRouter()
	r.Run(":" + config.AppConfig.Server.Port)
}
