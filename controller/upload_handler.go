package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// 1. 从请求中获取文件
	file, err := c.FormFile("file") // 前端传参的 key 必须叫 "file"
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取文件失败"})
		return
	}

	// 2. 检查文件后缀（简单的安全校验）
	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" && ext != ".gif" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "只允许上传图片(jpg, png, gif)"})
		return
	}

	// 3. 生成唯一文件名，防止重名覆盖
	// 格式：时间戳 + 后缀
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// 4. 保存文件到本地目录
	dst := filepath.Join("./uploads", newFileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
		return
	}

	// 5. 返回图片的访问 URL
	// 注意：在实际生产中，这里的 IP 应该是你服务器的公网 IP 或域名
	fileURL := fmt.Sprintf("http://%s/uploads/%s", c.Request.Host, newFileName)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"url":  fileURL,
	})
}
