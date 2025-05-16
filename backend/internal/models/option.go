package models

type Option struct {
	ID          int64  `json:"id"`
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
	AutoLoad    bool   `json:"auto_load"` // 是否自动加载
	Description string `json:"description"`
}

type UpdateOptionRequest struct {
	OptionValue string `json:"option_value" binding:"required"`
}

// 系统配置的键名常量
const (
	// 系统初始化标志
	OptionSystemInitialized = "system_initialized"
	// 用户注册设置
	OptionAllowRegistration = "allow_registration"
	// 系统名称
	OptionSystemName = "system_name"
	// 其他设置可以继续添加...
)

// 默认系统配置
var DefaultOptions = []Option{
	{
		OptionName:  OptionSystemInitialized,
		OptionValue: "false",
		AutoLoad:    true,
		Description: "系统是否已完成初始化",
	},
	{
		OptionName:  OptionAllowRegistration,
		OptionValue: "false",
		AutoLoad:    true,
		Description: "是否允许新用户注册",
	},
	{
		OptionName:  OptionSystemName,
		OptionValue: "后台管理系统",
		AutoLoad:    true,
		Description: "系统名称",
	},
} 