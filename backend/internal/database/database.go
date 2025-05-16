package database

import (
	"database/sql"
	"os"
	"path/filepath"
	"backend/internal/config"
	"backend/internal/models"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	// 确保数据目录存在
	if err := os.MkdirAll(filepath.Dir(cfg.DatabasePath), 0755); err != nil {
		return nil, err
	}

	// 打开数据库连接
	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		return nil, err
	}

	// 创建表
	if err := createTables(db); err != nil {
		return nil, err
	}

	// 检查系统是否已初始化
	var initialized bool
	err = db.QueryRow("SELECT option_value = 'true' FROM options WHERE option_name = ?", 
		models.OptionSystemInitialized).Scan(&initialized)
	
	if err == sql.ErrNoRows || !initialized {
		// 系统未初始化，执行初始化操作
		if err := initializeSystem(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL,
			is_admin BOOLEAN NOT NULL DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS options (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			option_name TEXT UNIQUE NOT NULL,
			option_value TEXT NOT NULL,
			auto_load BOOLEAN NOT NULL DEFAULT 1,
			description TEXT,
			return_to_frontend BOOLEAN NOT NULL DEFAULT 1,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func createDefaultAdmin(db *sql.DB) error {
	// 检查是否已存在管理员账户
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	// 创建默认管理员账户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password, is_admin) VALUES (?, ?, ?)",
		"admin", string(hashedPassword), true)
	return err
}

func initDefaultOptions(db *sql.DB) error {
	// 为每个默认选项检查是否存在，不存在则创建
	for _, option := range models.DefaultOptions {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM options WHERE option_name = ?", option.OptionName).Scan(&count)
		if err != nil {
			return err
		}

		if count == 0 {
			_, err = db.Exec(`
				INSERT INTO options (option_name, option_value, auto_load, description, return_to_frontend)
				VALUES (?, ?, ?, ?, ?)`,
				option.OptionName, option.OptionValue, option.AutoLoad, option.Description, option.ReturnToFrontend)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// initializeSystem 执行系统初始化
func initializeSystem(db *sql.DB) error {
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 创建默认管理员账户
	if err := createDefaultAdmin(db); err != nil {
		return err
	}

	// 初始化默认系统配置
	if err := initDefaultOptions(db); err != nil {
		return err
	}

	// 设置系统初始化标志
	_, err = db.Exec("UPDATE options SET option_value = 'true' WHERE option_name = ?",
		models.OptionSystemInitialized)
	if err != nil {
		return err
	}

	// 提交事务
	return tx.Commit()
} 