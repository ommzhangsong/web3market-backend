package controller

import (
	"gin-campus-market/blockchain"
	"gin-campus-market/models"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	// 1. 获取买家地址 (从 JWT 中获取)
	val, _ := c.Get("wallet_address")
	buyerAddr := val.(string)

	var input struct {
		OnChainID string `json:"on_chain_id" binding:"required"`
		TxHash    string `json:"tx_hash" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "参数缺失"})
		return
	}

	// 2. 链上核实这笔交易
	// 注意：这里需要调用合约查询该 ID 的最新状态，确认 Seller 是否已变为买家，或者 IsActive 变为 false
	onChainID, _ := new(big.Int).SetString(input.OnChainID, 10)
	listing, err := blockchain.Instance.Listings(nil, onChainID)
	if err != nil {
		c.JSON(500, gin.H{"error": "链上数据校验失败"})
		return
	}

	// 3. 这里的校验逻辑取决于你的合约实现
	// 比如：如果合约里卖掉后 IsActive 会变 false，我们就检查这个
	if listing.IsActive {
		c.JSON(400, gin.H{"error": "该商品在链上尚未成交"})
		return
	}

	// 4. 从数据库找到对应的商品信息（为了拿价格和卖家地址）
	var product models.Product
	if err := models.DB.Where("on_chain_id = ?", input.OnChainID).First(&product).Error; err != nil {
		c.JSON(404, gin.H{"error": "后端未找到该商品记录"})
		return
	}

	// 5. 创建订单记录
	order := models.Order{
		ProductID:  product.ID,
		OnChainID:  input.OnChainID,
		BuyerAddr:  buyerAddr,
		SellerAddr: product.SellerAddr,
		Price:      product.Price,
		TxHash:     input.TxHash,
	}

	if err := models.DB.Create(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "订单存储失败或重复提交"})
		return
	}

	c.JSON(200, gin.H{"message": "购买记录同步成功", "order_id": order.ID})
}
func GetMyOrders(c *gin.Context) {
	// 1. 获取当前登录用户的钱包地址
	val, _ := c.Get("wallet_address")
	userAddr := val.(string)

	// 2. 定义一个结构体来接收查询结果（包含订单和商品基本信息）
	type OrderResult struct {
		ID        uint    `json:"id"`
		OnChainID string  `json:"on_chain_id"`
		TxHash    string  `json:"tx_hash"`
		CreatedAt string  `json:"created_at"`
		Title     string  `json:"title"`     // 来自 Product 表
		ImageUrl  string  `json:"image_url"` // 来自 Product 表
		Price     float64 `json:"price"`     // 购买时的成交价
	}

	var results []OrderResult

	// 3. 联表查询：orders 表 JOIN products 表
	// 通过 products.on_chain_id = orders.on_chain_id 进行关联
	err := models.DB.Table("orders").
		Select("orders.id, orders.on_chain_id, orders.tx_hash, orders.created_at, products.title, products.image_url, orders.price").
		Joins("left join products on products.on_chain_id = orders.on_chain_id").
		Where("orders.buyer_addr = ?", userAddr).
		Order("orders.created_at desc"). // 按时间倒序
		Scan(&results).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取购买记录失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": results,
	})
}
