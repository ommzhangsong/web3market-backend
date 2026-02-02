package router

import (
	"gin-campus-market/controller"
	"gin-campus-market/middleware"
	"gin-campus-market/service" // 假设你的聊天逻辑写在 service 包里
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 中间件保持不变
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	api := r.Group("/api")
	{
		// === 商品公开接口 (无需登录) ===
		api.GET("/products", controller.GetProductList)       // 获取商品列表
		api.GET("/products/:id", controller.GetProductDetail) // 获取商品详情
		r.Static("/uploads", "./uploads")
		authGroup := api.Group("/auth")
		{
			// 登录接口
			authGroup.POST("/login", controller.LoginHandler)

			// === 需验证路由组 ===
			protected := authGroup.Group("")
			protected.Use(middleware.AuthMiddleware())
			{
				// 原有用户逻辑
				protected.GET("/me", controller.GetProfile)
				protected.POST("/send-code", controller.SendCodeHandler)
				protected.POST("/verify-email", controller.VerifyEmailHandler)

				// === 商品保护操作 ===
				protected.POST("/products", controller.CreateProduct)       // 发布商品
				protected.DELETE("/products/:id", controller.DeleteProduct) // 删除商品

				// === 聊天逻辑 (保持不变) ===
				protected.GET("/ws", func(c *gin.Context) {
					val, _ := c.Get("wallet_address")
					walletAddr, _ := val.(string)
					service.HandleChat(c.Writer, c.Request, walletAddr)
				})
				protected.GET("/chat/messages/:target_id", controller.GetChatHistory)
				protected.GET("/chat/sessions", controller.GetChatSessions)
				protected.POST("/chat/read", controller.MarkAsRead)
				protected.POST("/upload", controller.UploadFile)
				// 在 protected 路由组内添加
				protected.GET("/my-products", controller.GetMyProducts)
				protected.PUT("/products/:id", controller.UpdateProduct)
				protected.POST("/confirm-purchase", controller.ConfirmPurchase)
				protected.POST("/confirm-receipt", controller.ConfirmReceipt)
				protected.POST("/orders", controller.CreateOrder)
				protected.GET("/orders/my", controller.GetMyOrders)
			}
		}
	}
	return r
}
