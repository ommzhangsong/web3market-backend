package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string  `json:"title" gorm:"index"`
	OnChainID   string  `json:"on_chain_id" gorm:"index"`
	OnChainTxID string  `json:"on_chain_tx_id" gorm:"index"` // 新增：交易在链上的 TransactionID
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"image_url"`
	Category    string  `json:"category" gorm:"index"`    // 比如：电子产品, 图书, 生活用品
	SellerAddr  string  `json:"seller_addr" gorm:"index"` // 关键：卖家的钱包地址
	Status      int     `json:"status" gorm:"default:1"`  // 1: 在售, 2: 售出, 3: 已下架
}
