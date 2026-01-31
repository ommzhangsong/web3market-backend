package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 对应你规范中的后端统一返回格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功返回：状态码 200
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}

// Error 失败返回：自定义状态码和信息
func Error(c *gin.Context, code int, msg string) {
	c.JSON(code, Response{
		Code:    code,
		Message: msg,
		Data:    nil, // 对应规范：失败时 data 为 null
	})
}
