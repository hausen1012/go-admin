package config

import (
	"os"
)

type Config struct {
	ServerAddress     string `json:"server_address"`
	JWTSecret        string `json:"jwt_secret"`
	AllowRegistration bool   `json:"allow_registration"`
	DatabasePath     string `json:"database_path"`
}

func LoadConfig() *Config {
	// 默认配置
	config := &Config{
		ServerAddress:     ":8080",
		JWTSecret:        "your-secret-key",
		AllowRegistration: true,
		DatabasePath:     "data/app.db",
	}

	// 从环境变量加载配置
	if serverAddr := os.Getenv("APP_SERVER_ADDRESS"); serverAddr != "" {
		config.ServerAddress = serverAddr
	}
	if jwtSecret := os.Getenv("APP_JWT_SECRET"); jwtSecret != "" {
		config.JWTSecret = jwtSecret
	}
	if allowReg := os.Getenv("APP_ALLOW_REGISTRATION"); allowReg != "" {
		config.AllowRegistration = allowReg == "true"
	}
	if dbPath := os.Getenv("APP_DATABASE_PATH"); dbPath != "" {
		config.DatabasePath = dbPath
	}

	return config
} 