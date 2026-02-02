package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model

	FromUserID string `json:"from_user_id" gorm:"type:varchar(64);index"`
	ToUserID   string `json:"to_user_id" gorm:"type:varchar(64);index"`
	Content    string `json:"content"`
	IsRead     bool   `json:"is_read" gorm:"default:false"`
	ProductID  uint   `json:"product_id" gorm:"index"`
}
