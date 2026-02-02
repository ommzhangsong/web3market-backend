package models

import (
	"fmt"
	"gin-campus-market/config" // 引入你的配置包
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	c := config.AppConfig.Database
	// 使用 Sprintf 动态拼接配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	db.AutoMigrate(&User{})
	DB = db
	fmt.Println("✅ 数据库配置加载并连接成功")
	DB.AutoMigrate(
		&User{},    // 用户表
		&Message{}, // 聊天记录表（刚才聊到的）
		&Product{}, // 商品表（刚刚创建的）
		&Order{},
	)
}
