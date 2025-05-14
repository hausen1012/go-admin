package handlers

import (
	"net/http"
	"backend/internal/models"
	"github.com/gin-gonic/gin"
)

// 获取所有系统配置
func (h *Handler) GetOptions(c *gin.Context) {
	rows, err := h.db.Query(`
		SELECT id, option_name, option_value, auto_load, description
		FROM options
		ORDER BY id ASC`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取系统配置失败"})
		return
	}
	defer rows.Close()

	var options []models.Option
	for rows.Next() {
		var option models.Option
		err := rows.Scan(
			&option.ID,
			&option.OptionName,
			&option.OptionValue,
			&option.AutoLoad,
			&option.Description,
		)
		if err != nil {
			continue
		}
		options = append(options, option)
	}

	c.JSON(http.StatusOK, options)
}

// 获取单个配置项
func (h *Handler) GetOption(c *gin.Context) {
	optionName := c.Param("name")
	
	var option models.Option
	err := h.db.QueryRow(`
		SELECT id, option_name, option_value, auto_load, description
		FROM options
		WHERE option_name = ?`,
		optionName,
	).Scan(
		&option.ID,
		&option.OptionName,
		&option.OptionValue,
		&option.AutoLoad,
		&option.Description,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "配置项不存在"})
		return
	}

	c.JSON(http.StatusOK, option)
}

// 更新配置项
func (h *Handler) UpdateOption(c *gin.Context) {
	optionName := c.Param("name")
	var req models.UpdateOptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	result, err := h.db.Exec(`
		UPDATE options
		SET option_value = ?, updated_at = CURRENT_TIMESTAMP
		WHERE option_name = ?`,
		req.OptionValue, optionName,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新配置失败"})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "配置项不存在"})
		return
	}

	// 如果更新的是注册设置，同步更新配置
	if optionName == models.OptionAllowRegistration {
		h.cfg.AllowRegistration = req.OptionValue == "true"
	}

	c.JSON(http.StatusOK, gin.H{"message": "配置更新成功"})
} 