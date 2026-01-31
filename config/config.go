package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
}

var AppConfig *Config

func InitConfig() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	AppConfig = &Config{}
	err = yaml.Unmarshal(file, AppConfig)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
}
