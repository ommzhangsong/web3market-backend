package controller

import (
	"fmt"
	"gin-campus-market/common"
	"gin-campus-market/models"
	"math/rand"
	"net/smtp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	"gorm.io/gorm"
)

// LoginRequest 定义前端发送的参数
type LoginRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
}

// LoginHandler 处理模拟 Web3 登录
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	req.WalletAddress = strings.TrimSpace(req.WalletAddress)
	// 1. 参数校验
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "参数错误: 请提供钱包地址")
		return
	}

	var user models.User
	// 1. 查数据库
	result := models.DB.Where("wallet_address = ?", req.WalletAddress).First(&user)

	// 2. 如果报错了
	if result.Error != nil {
		// 只有当错误确实是 "Record Not Found" 时，才去执行创建
		if result.Error == gorm.ErrRecordNotFound {
			user = models.User{
				WalletAddress: req.WalletAddress,
				IsVerified:    false,
				CodeExpiresAt: time.Now(), // 加上我们刚才说的修复
			}
			// 执行创建
			if err := models.DB.Create(&user).Error; err != nil {
				common.Error(c, 500, "注册失败")
				return
			}
		} else {
			// 如果是其他数据库错误，报错返回
			common.Error(c, 500, "数据库异常")
			return
		}
	}

	// 3. 运行到这里，说明要么查到了，要么创建成功了
	// 直接生成 Token 即可

	token, err := common.ReleaseToken(user.ID, user.WalletAddress)
	if err != nil {
		common.Error(c, 500, "Token 生成失败")
		return
	}

	// 4. 返回符合你规范的统一格式响应
	common.Success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":             user.ID,
			"wallet_address": user.WalletAddress,
			"is_verified":    user.IsVerified,
		},
	})

}

// SendCodeHandler 发送验证码
func SendCodeHandler(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "请输入有效的校园邮箱")
		return
	}

	// 1. 生成 6 位随机验证码
	code := fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000))

	// 2. 获取当前用户（从 JWT 中间件存入的信息中取）
	walletAddress, _ := c.Get("wallet_address")

	// 3. 更新数据库：存入验证码和过期时间（5分钟有效）
	err := models.DB.Model(&models.User{}).Where("wallet_address = ?", walletAddress).Updates(map[string]interface{}{
		"email":             req.Email,
		"verification_code": code,
		"code_expires_at":   time.Now().Add(time.Minute * 5),
	}).Error

	if err != nil {
		common.Error(c, 500, "数据库更新失败")
		return
	}

	// 4. 模拟发送邮件（控制台打印）
	e := email.NewEmail()
	e.From = "1879566697@qq.com"
	e.To = []string{req.Email} // 前端传来的你的邮箱
	e.Subject = "校园二手市场验证码"
	e.Text = []byte(fmt.Sprintf("您的验证码是：%s，有效期5分钟。", code))

	// 设置 SMTP 服务器（以QQ为例，如果是网易请改写为 smtp.163.com）
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1879566697@qq.com", "xtoyfedxrgcqhdbf", "smtp.qq.com"))

	if err != nil {
		common.Error(c, 500, "邮件发送失败: "+err.Error())
		return
	}

	common.Success(c, "验证码已发送至您的真实邮箱")
}

// VerifyEmailHandler 校验验证码
func VerifyEmailHandler(c *gin.Context) {
	var req struct {
		Code string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Error(c, 400, "请输入验证码")
		return
	}

	walletAddress, _ := c.Get("wallet_address")
	var user models.User
	models.DB.Where("wallet_address = ?", walletAddress).First(&user)

	// 1. 判定是否过期
	if time.Now().After(user.CodeExpiresAt) {
		common.Error(c, 400, "验证码已过期")
		return
	}

	// 2. 判定验证码是否一致
	if user.VerificationCode != req.Code {
		common.Error(c, 400, "验证码错误")
		return
	}

	// 3. 验证通过，更新状态
	models.DB.Model(&user).Updates(map[string]interface{}{
		"is_verified":       true,
		"verification_code": "", // 验证后清空，防止重复使用
	})

	common.Success(c, "校园身份核验成功！")
}

// GetProfile 获取当前登录用户的详细信息
func GetProfile(c *gin.Context) {
	// 1. 从 Context 中获取中间件存入的 wallet_address
	// 这是在 middleware.AuthMiddleware 中 c.Set("wallet_address", ...) 存进去的
	walletAddress, exists := c.Get("wallet_address")
	if !exists {
		common.Error(c, 401, "未获取到登录信息")
		return
	}

	var user models.User
	// 2. 去数据库查一下最新的状态
	// 这里不需要 Unscoped，因为正常操作的用户不应该是删除状态
	if err := models.DB.Where("wallet_address = ?", walletAddress).First(&user).Error; err != nil {
		common.Error(c, 404, "用户不存在")
		return
	}

	// 3. 返回给前端
	common.Success(c, gin.H{
		"user": gin.H{
			"id":             user.ID,
			"wallet_address": user.WalletAddress,
			"email":          user.Email,
			"is_verified":    user.IsVerified,
		},
	})
}
