package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	ServerAddress string `json:"server_address"`
	JWTSecret    string `json:"jwt_secret"`
	DataPath     string `json:"data_path"`
	DbName       string `json:"db_name"`
	DatabasePath string `json:"database_path"` // 完整的数据库路径
}

func LoadConfig() *Config {
	// 默认配置
	config := &Config{
		ServerAddress: ":8080",
		JWTSecret:    "M6pTFnXmYKpH2k",
		DataPath:     "data",
		DbName:       "app.db",
	}

	// 从环境变量加载配置
	if serverAddr := os.Getenv("SERVER_ADDRESS"); serverAddr != "" {
		config.ServerAddress = serverAddr
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		config.JWTSecret = jwtSecret
	}
	if dataPath := os.Getenv("DATA_PATH"); dataPath != "" {
		config.DataPath = dataPath
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DbName = dbName
	}

	// 组合完整的数据库路径
	config.DatabasePath = filepath.Join(config.DataPath, config.DbName)

	return config
} 