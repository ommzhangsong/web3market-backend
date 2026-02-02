package service

import (
	"encoding/json"
	"fmt"
	"gin-campus-market/models"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// 1. 存储在线用户的连接：Key 改为 string (钱包地址)
var Clients = make(map[string]*websocket.Conn)
var Mux sync.Mutex

// 升级 HTTP 为 WebSocket
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 消息数据结构（接收前端 JSON）
type ChatMsg struct {
	ToUserID  string `json:"to_user_id"` // 接收方钱包地址
	Content   string `json:"content"`
	ProductID uint   `json:"product_id"`
}

// HandleChat：参数从 userID uint 改为 walletAddr string
func HandleChat(w http.ResponseWriter, r *http.Request, walletAddr string) {
	// 1. 升级协议
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WS Upgrade Error:", err)
		return
	}

	// 2. 记录用户在线
	Mux.Lock()
	Clients[walletAddr] = conn
	fmt.Printf("✅ 用户上线: %s, 当前在线人数: %d\n", walletAddr, len(Clients))
	Mux.Unlock()

	// 3. 断开连接后的处理
	defer func() {
		Mux.Lock()
		delete(Clients, walletAddr)
		fmt.Printf("❌ 用户下线: %s\n", walletAddr)
		Mux.Unlock()
		conn.Close()
	}()

	for {
		// 4. 读取前端发来的消息
		_, msgData, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var data ChatMsg
		if err := json.Unmarshal(msgData, &data); err != nil {
			fmt.Println("JSON Unmarshal Error:", err)
			continue
		}

		// --- 关键：持久化到数据库 ---
		// 此时 FromUserID 和 ToUserID 都是 string 类型
		msgRecord := models.Message{
			FromUserID: walletAddr,    // 发送者钱包
			ToUserID:   data.ToUserID, // 接收者钱包
			Content:    data.Content,
			ProductID:  data.ProductID,
			IsRead:     false,
		}

		// 保存到数据库
		if err := models.DB.Create(&msgRecord).Error; err != nil {
			fmt.Println("数据库保存失败:", err)
		}

		// 5. 实时转发消息
		Mux.Lock()
		// 根据目标钱包地址寻找连接
		if targetConn, ok := Clients[data.ToUserID]; ok {
			// 将包含 ID 和时间戳的完整记录发给对方
			response, _ := json.Marshal(msgRecord)
			targetConn.WriteMessage(websocket.TextMessage, response)
			fmt.Printf("🚀 消息已从 %s 转发至 %s\n", walletAddr, data.ToUserID)
		} else {
			fmt.Printf("😴 目标用户 %s 不在线，消息已存库\n", data.ToUserID)
		}
		Mux.Unlock()
	}
}
