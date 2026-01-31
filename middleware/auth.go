package middleware

import (
	"gin-campus-market/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Header 中的 Authorization: Bearer {token}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			common.Error(c, 401, "未授权，请先登录")
			c.Abort() // 拦截请求，不再往下执行
			return
		}

		// 2. 提取并解析 Token
		tokenString := authHeader[7:] // 去掉 "Bearer "
		claims, ok := common.ParseToken(tokenString)
		if !ok {
			common.Error(c, 401, "Token 无效或已过期")
			c.Abort()
			return
		}

		// 3. 将用户信息存入上下文，方便后续接口使用
		c.Set("user_id", claims.UserID)
		c.Set("wallet_address", claims.WalletAddress)

		c.Next()
	}
}
