package controller

import (
	"gin-campus-market/common"
	"gin-campus-market/models"
	"time"

	"github.com/gin-gonic/gin"
)

// GetChatHistory 获取与特定用户的聊天记录
func GetChatHistory(c *gin.Context) {
	// 获取当前登录用户的钱包地址（从中间件存入的上下文获取）
	val, _ := c.Get("wallet_address")
	myWallet := val.(string)

	// 获取路径参数中的目标钱包地址
	targetWallet := c.Param("target_id")

	var messages []models.Message
	// 直接用钱包地址字符串进行查询
	err := models.DB.Where(
		"(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)",
		myWallet, targetWallet, targetWallet, myWallet,
	).Order("created_at asc").Find(&messages).Error

	if err != nil {
		common.Error(c, 500, "查询记录失败")
		return
	}

	common.Success(c, messages)
}

// SessionResp 结构体用于返回给前端
type SessionResp struct {
	TargetWallet string    `json:"target_wallet"`
	LastMessage  string    `json:"last_message"`
	LastTime     time.Time `json:"last_time"`
	UnreadCount  int64     `json:"unread_count"`
}

func GetChatSessions(c *gin.Context) {
	// 1. 获取当前用户钱包地址
	myWallet, _ := c.Get("wallet_address")
	addr := myWallet.(string)

	var sessions []SessionResp

	// 2. 执行原生 SQL 查询（这是最简单且性能最好的方式）
	// 逻辑：查询所有与我相关的消息，按对方地址分组，取最后一条消息和时间
	err := models.DB.Raw(`
		SELECT 
			target_wallet, 
			content as last_message, 
			created_at as last_time,
			(SELECT COUNT(*) FROM messages m2 WHERE m2.from_user_id = t.target_wallet AND m2.to_user_id = ? AND m2.is_read = 0) as unread_count
		FROM (
			SELECT 
				IF(from_user_id = ?, to_user_id, from_user_id) as target_wallet,
				content,
				created_at,
				ROW_NUMBER() OVER(PARTITION BY IF(from_user_id = ?, to_user_id, from_user_id) ORDER BY created_at DESC) as rn
			FROM messages
			WHERE from_user_id = ? OR to_user_id = ?
		) t
		WHERE t.rn = 1
		ORDER BY last_time DESC
	`, addr, addr, addr, addr, addr).Scan(&sessions).Error

	if err != nil {
		c.JSON(500, gin.H{"error": "获取列表失败"})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": sessions,
	})
}
func MarkAsRead(c *gin.Context) {
	myWallet, _ := c.Get("wallet_address")

	// 定义一个临时结构体来接收参数
	var input struct {
		TargetWallet string `json:"target_wallet" form:"target_wallet"`
	}

	// 使用 ShouldBind 自动识别 JSON 或 Form
	if err := c.ShouldBind(&input); err != nil || input.TargetWallet == "" {
		c.JSON(400, gin.H{"code": 400, "msg": "必须提供对方钱包地址 (target_wallet)"})
		return
	}

	// 执行更新：把对方发给我的未读消息改为已读
	result := models.DB.Model(&models.Message{}).
		Where("from_user_id = ? AND to_user_id = ? AND is_read = ?", input.TargetWallet, myWallet.(string), false).
		Update("is_read", true)

	if result.Error != nil {
		c.JSON(500, gin.H{"code": 500, "msg": "更新失败"})
		return
	}

	c.JSON(200, gin.H{
		"code":          200,
		"msg":           "已读成功",
		"affected_rows": result.RowsAffected, // 返回更新了多少条，方便调试
	})
}
