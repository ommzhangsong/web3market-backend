package middleware

import (
	"fmt"
	"gin-campus-market/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 尝试从 Header 获取
		authHeader := c.GetHeader("Authorization")
		tokenString := ""

		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = authHeader[7:]
		} else {
			// 2. 如果 Header 没有，尝试从 URL 参数 "token" 中获取
			// 这一步是专门给 WebSocket 准备的
			tokenString = c.Query("token")
		}

		// 3. 校验 Token 是否为空
		if tokenString == "" {
			c.JSON(401, gin.H{"code": 401, "msg": "未授权，请先登录"})
			c.Abort()
			return
		}
		fmt.Println("正在校验的 Token:", tokenString)
		// 4. 解析 Token (调用你之前的 common.ParseToken)
		claims, ok := common.ParseToken(tokenString)
		if !ok {
			c.JSON(401, gin.H{"code": 401, "msg": "Token 无效或已过期"})
			c.Abort()
			return
		}

		// 5. 存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("wallet_address", claims.WalletAddress)

		c.Next()
	}
}
