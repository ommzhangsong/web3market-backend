package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`            // 对应商品表的主键
	OnChainID  string  `json:"on_chain_id"`           // 链上的 Listing ID
	BuyerAddr  string  `json:"buyer_addr"`            // 买家钱包地址
	SellerAddr string  `json:"seller_addr"`           // 卖家钱包地址
	Price      float64 `json:"price"`                 // 购买时的价格
	TxHash     string  `json:"tx_hash" gorm:"unique"` // 交易哈希，防止重复记录
}
