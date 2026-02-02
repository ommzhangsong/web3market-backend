package controller

import (
	"fmt"
	"gin-campus-market/blockchain"
	"gin-campus-market/models"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetProductList 获取商品列表 (GET /api/products)
func GetProductList(c *gin.Context) {
	var products []models.Product

	// 1. 获取分页与筛选参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	category := c.Query("category")

	// 2. 构建查询
	query := models.DB.Model(&models.Product{}).Where("status = ?", 1) // 只看在售商品

	if category != "" {
		query = query.Where("category = ?", category)
	}

	// 3. 执行分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取商品列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": products,
	})
}

// GetProductDetail 获取商品详情 (GET /api/products/:id)
func GetProductDetail(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": product,
	})
}

// CreateProduct 发布商品 (POST /api/auth/products)
func CreateProduct(c *gin.Context) {
	// 1. 获取登录地址 (从 JWT 中解析出来的)
	val, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未获取到登录信息"})
		return
	}
	sellerAddr := val.(string)

	// 2. 绑定 JSON 到模型
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		// --- 加上这一行打印，你就能在黑窗口看到具体的解析错误了 ---
		fmt.Printf("JSON 绑定失败详情: %v\n", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	fmt.Println("✅ JSON 绑定成功，开始校验链上数据...")

	// --- 链上核实逻辑 (防止崩溃) ---
	if product.OnChainID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺失 on_chain_id"})
		return
	}

	onChainID := new(big.Int)
	onChainID, ok := onChainID.SetString(product.OnChainID, 10)
	if !ok || onChainID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的链上ID格式"})
		return
	}

	if blockchain.Instance == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "后端区块链连接未初始化"})
		return
	}

	// 调用合约查询
	// --- 找到这段代码进行修改 ---

	// 调用合约查询
	listing, err := blockchain.Instance.Listings(nil, onChainID)
	if err != nil {
		// 打印详细的 RPC 错误，比如是不是网络断了或 ID 太大
		fmt.Printf("❌ 合约查询异常: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "合约查询失败: " + err.Error()})
		return
	}

	// 重点！打印从链上拿到的原始数据
	fmt.Printf("🔍 链上数据详情: ID=%s, Seller=%s, IsActive=%v, Price=%s\n",
		product.OnChainID, listing.Seller.Hex(), listing.IsActive, listing.Price.String())

	if !listing.IsActive {
		// 如果这里打印了，说明链上确实没有这个 ID 的有效商品
		fmt.Printf("❌ 校验拦截：链上 ID %s 的 IsActive 为 false\n", product.OnChainID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "链上该商品已下架或 ID 错误"})
		return
	}

	// 统一地址格式对比
	contractSeller := listing.Seller.Hex()
	userWallet := common.HexToAddress(sellerAddr).Hex()
	if contractSeller != userWallet {
		// 这里在终端打印，方便你调试
		fmt.Printf("❌ 拒绝发布！合约卖家: %s, 登录用户: %s\n", contractSeller, userWallet)
		c.JSON(http.StatusForbidden, gin.H{
			"error":           "权限不足：你不是该链上商品的拥有者",
			"contract_seller": contractSeller,
		})
		return
	}

	// 3. 补全信息并存入数据库
	product.SellerAddr = userWallet
	product.Status = 1

	if err := models.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库存储失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "商品同步成功", "data": product})
}

// DeleteProduct 删除商品 (DELETE /api/auth/products/:id)
func DeleteProduct(c *gin.Context) {
	// 1. 获取当前登录用户的钱包地址
	val, _ := c.Get("wallet_address")
	myAddr := val.(string)

	// 2. 获取路径参数中的商品 ID
	id := c.Param("id")

	// 3. 先查出这个商品，判断是否属于当前用户
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品未找到"})
		return
	}

	// 4. 【关键】权限校验：只有卖家本人才能删除
	if product.SellerAddr != myAddr {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限删除他人的商品"})
		return
	}

	// 5. 执行删除 (gorm.Model 默认是软删除，会记录 DeletedAt)
	if err := models.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "商品已成功删除",
	})
}

// GetMyProducts 获取当前用户发布的商品 (GET /api/auth/my-products)
func GetMyProducts(c *gin.Context) {
	// 1. 从中间件获取当前登录用户的钱包地址
	val, exists := c.Get("wallet_address")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
		return
	}
	myAddr := val.(string)

	var products []models.Product

	// 2. 查询数据库：匹配卖家地址
	// 这里不需要过滤 status=1，因为用户在个人中心应该能看到自己“在售”、“下架”甚至“已售”的所有商品
	err := models.DB.Where("seller_addr = ?", myAddr).
		Order("created_at DESC").
		Find(&products).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取数据失败"})
		return
	}

	// 3. 返回数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": products,
	})
}

// UpdateProduct 修改商品 (PUT /api/auth/products/:id)
func UpdateProduct(c *gin.Context) {
	// 1. 获取当前登录用户的钱包地址
	val, _ := c.Get("wallet_address")
	myAddr := val.(string)

	// 2. 获取路径参数中的 ID
	id := c.Param("id")

	// 3. 先查出这个商品，判断是否存在及权限
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品未找到"})
		return
	}

	// 4. 【关键】只有卖家本人才能修改
	if product.SellerAddr != myAddr {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限修改此商品"})
		return
	}

	// 5. 绑定前端传来的新数据
	// 我们定义一个临时结构体，只允许用户修改特定字段，防止通过 JSON 修改 ID 或 SellerAddr
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		ImageUrl    string  `json:"image_url"`
		Category    string  `json:"category"`
		Status      int     `json:"status"` // 允许用户自己上架或下架
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
		return
	}

	// 6. 更新数据库记录 (使用 Updates 会自动忽略空字段或只更新指定字段)
	updates := models.Product{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		ImageUrl:    input.ImageUrl,
		Category:    input.Category,
		Status:      input.Status,
	}

	if err := models.DB.Model(&product).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新商品失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改成功",
		"data":    product, // 返回更新后的商品
	})
}

// ConfirmPurchase 确认链上交易已创建 (POST /api/auth/confirm-purchase)
func ConfirmPurchase(c *gin.Context) {
	var input struct {
		ProductID   uint   `json:"product_id"`     // 后端数据库ID
		OnChainTxID string `json:"on_chain_tx_id"` // 合约生成的 transactionId
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	// --- 修改点 1: 严格检查 ID 转换 ---
	if input.OnChainTxID == "" {
		c.JSON(400, gin.H{"error": "链上交易 ID 不能为空"})
		return
	}

	txID := new(big.Int)
	txID, ok := txID.SetString(input.OnChainTxID, 10)
	if !ok || txID == nil {
		c.JSON(400, gin.H{"error": "链上交易 ID 格式错误"})
		return
	}

	// --- 修改点 2: 检查合约实例是否初始化 ---
	if blockchain.Instance == nil {
		c.JSON(500, gin.H{"error": "区块链服务未连接"})
		return
	}

	// 1. 从合约查询这笔交易的真实性
	// 现在 txID 保证不是 nil，不会再触发 reflect panic
	onChainTx, err := blockchain.Instance.Transactions(nil, txID)
	if err != nil {
		c.JSON(400, gin.H{"error": "查询链上交易失败: " + err.Error()})
		return
	}

	// 2. 验证：状态必须是 Locked (根据你的合约 enum: Created=0, Locked=1)
	if onChainTx.Status != 1 {
		c.JSON(400, gin.H{"error": "交易状态非法，链上显示未锁定"})
		return
	}

	// 3. 更新数据库
	var product models.Product
	if err := models.DB.First(&product, input.ProductID).Error; err != nil {
		c.JSON(404, gin.H{"error": "数据库中未找到该商品"})
		return
	}

	// 使用 Map 更新以确保字段被正确写入
	updates := map[string]interface{}{
		"status":         2, // 锁定状态
		"on_chain_tx_id": input.OnChainTxID,
	}

	if err := models.DB.Model(&product).Updates(updates).Error; err != nil {
		c.JSON(500, gin.H{"error": "更新数据库失败"})
		return
	}

	c.JSON(200, gin.H{"code": 200, "message": "交易已同步", "data": product})
}
func ConfirmReceipt(c *gin.Context) {
	var input struct {
		ProductID uint `json:"product_id"`
	}
	c.ShouldBindJSON(&input)

	var product models.Product
	models.DB.First(&product, input.ProductID)

	// 1. 获取对应的链上交易 ID
	txID, _ := new(big.Int).SetString(product.OnChainTxID, 10)

	// 2. 查合约状态
	onChainTx, _ := blockchain.Instance.Transactions(nil, txID)

	// 3. 验证状态：已完成应该是 Completed (索引为 3)
	if onChainTx.Status != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "买家尚未在链上点击确认收货"})
		return
	}

	// 4. 更新数据库状态为 3 (已售出/完成)
	models.DB.Model(&product).Update("status", 3)

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "交易圆满完成"})
}
