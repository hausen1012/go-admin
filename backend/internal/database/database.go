package database

import (
	"os"
	"path/filepath"
	"backend/internal/config"
	"backend/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	// 确保数据目录存在
	if err := os.MkdirAll(filepath.Dir(cfg.DatabasePath), 0755); err != nil {
		return nil, err
	}

	// 打开数据库连接
	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库结构
	if err := db.AutoMigrate(&models.User{}, &models.Option{}); err != nil {
		return nil, err
	}

	// 检查系统是否已初始化
	var option models.Option
	result := db.Where("option_name = ?", models.OptionSystemInitialized).First(&option)
	
	if result.Error == gorm.ErrRecordNotFound || option.OptionValue != "true" {
		// 系统未初始化，执行初始化操作
		if err := initializeSystem(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func createDefaultAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&models.User{}).Where("username = ?", "admin").Count(&count)
	if count > 0 {
		return nil
	}

	// 创建默认管理员账户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := models.User{
		Username: "admin",
		Password: string(hashedPassword),
		IsAdmin:  true,
	}

	return db.Create(&admin).Error
}

func initDefaultOptions(db *gorm.DB) error {
	// 为每个默认选项检查是否存在，不存在则创建
	for _, option := range models.DefaultOptions {
		var count int64
		db.Model(&models.Option{}).Where("option_name = ?", option.OptionName).Count(&count)
		if count == 0 {
			if err := db.Create(&option).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// initializeSystem 执行系统初始化
func initializeSystem(db *gorm.DB) error {
	// 开始事务
	return db.Transaction(func(tx *gorm.DB) error {
		// 创建默认管理员账户
		if err := createDefaultAdmin(tx); err != nil {
			return err
		}

		// 初始化默认系统配置
		if err := initDefaultOptions(tx); err != nil {
			return err
		}

		// 设置系统初始化标志
		return tx.Model(&models.Option{}).
			Where("option_name = ?", models.OptionSystemInitialized).
			Update("option_value", "true").Error
	})
} 