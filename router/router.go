package router

import (
	"gin-campus-market/controller"
	"gin-campus-market/middleware" // 引入中间件

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// 公开接口：登录
		auth := api.Group("/auth")
		{
			auth.POST("/login", controller.LoginHandler)
		}

		// 受保护接口：必须带 JWT Token 才能访问
		verified := api.Group("/auth")
		verified.Use(middleware.AuthMiddleware()) // 使用中间件进行拦截
		{
			verified.POST("/send-code", controller.SendCodeHandler)
			verified.POST("/verify-email", controller.VerifyEmailHandler)
		}
	}

	return r
}
