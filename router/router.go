package router

import (
	"gin-campus-market/controller"
	"gin-campus-market/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 中间件：处理跨域问题，解决前端的 OPTIONS 404 错误
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 允许前端所有来源访问，如果你知道前端的具体端口（如 http://localhost:3000），建议替换掉 "*"
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 如果是预检请求 (OPTIONS)，直接返回 204 (No Content) 即可
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 1. 注册跨域中间件（必须放在最前面）
	r.Use(Cors())

	// 基础测试接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		// --- 公开路由 ---
		auth := api.Group("/auth")
		{
			// 登录不需要 Token
			auth.POST("/login", controller.LoginHandler)
		}

		// --- 需要身份验证的路由 ---
		verified := api.Group("/auth")
		verified.Use(middleware.AuthMiddleware())
		{

			// 发送验证码
			verified.POST("/send-code", controller.SendCodeHandler)
			// 核验验证码
			verified.POST("/verify-email", controller.VerifyEmailHandler)
		}
	}

	return r
}
