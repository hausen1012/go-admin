package config

import (
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	JWTSecret    string `json:"jwt_secret"`
	DatabasePath string `json:"database_path"`
}

func LoadConfig() *Config {
	// 默认配置
	config := &Config{
		ServerAddress: ":8080",
		JWTSecret:    "M6pTFnXmYKpH2k",
		DatabasePath: "data/app.db",
	}

	// 从环境变量加载配置
	if serverAddr := os.Getenv("APP_SERVER_ADDRESS"); serverAddr != "" {
		config.ServerAddress = serverAddr
	}
	if jwtSecret := os.Getenv("APP_JWT_SECRET"); jwtSecret != "" {
		config.JWTSecret = jwtSecret
	}
	if dbPath := os.Getenv("APP_DATABASE_PATH"); dbPath != "" {
		config.DatabasePath = dbPath
	}

	return config
} 