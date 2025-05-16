package main

import (
	"log"
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/router"
)

func main() {
	// 初始化配置
	cfg := config.LoadConfig()

	// 初始化数据库
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 获取原生数据库连接以便关闭
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying database connection: %v", err)
	}
	defer sqlDB.Close()

	// 初始化路由
	r := router.SetupRouter(db, cfg)

	// 启动服务器
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 