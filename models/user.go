package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	WalletAddress    string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"wallet_address"` // 模拟 Web3 唯一标识
	Email            string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`                   // 校园邮箱
	IsVerified       bool           `gorm:"default:false" json:"is_verified"`                             // 是否核验通过
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"` // 软删除，不返回给前端
	VerificationCode string         `json:"-"`
	CodeExpiresAt    time.Time      `json:"-"`
}
