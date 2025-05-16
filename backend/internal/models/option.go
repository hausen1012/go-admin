package models

import (
	"time"
	"gorm.io/gorm"
)

type Option struct {
	ID              uint           `gorm:"primarykey"`
	OptionName      string         `gorm:"uniqueIndex;not null"`
	OptionValue     string         `gorm:"not null"`
	AutoLoad        bool           `gorm:"default:true"`
	Description     string
	ReturnToFrontend bool          `gorm:"default:true"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UpdateOptionRequest struct {
	OptionValue string `json:"option_value" binding:"required"`
}

// 系统选项常量
const (
	OptionSystemInitialized = "system_initialized"
	// 用户注册设置
	OptionAllowRegistration = "allow_registration"
	// 系统名称
	OptionSystemName = "system_name"
	// 其他设置可以继续添加...
)

// DefaultOptions 默认系统选项
var DefaultOptions = []Option{
	{
		OptionName:       OptionSystemInitialized,
		OptionValue:      "false",
		AutoLoad:         true,
		Description:      "系统是否已初始化",
		ReturnToFrontend: false,
	},
	{
		OptionName:       OptionAllowRegistration,
		OptionValue:      "false",
		AutoLoad:         true,
		Description:      "是否允许新用户注册",
		ReturnToFrontend: true,
	},
	{
		OptionName:       OptionSystemName,
		OptionValue:      "后台管理系统",
		AutoLoad:         true,
		Description:      "系统名称",
		ReturnToFrontend: true,
	},
} 