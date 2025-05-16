package router

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// 创建处理器
	h := handlers.NewHandler(db, cfg)

	// 允许跨域
	r.Use(middleware.CORS())

	// 公开路由
	public := r.Group("/api")
	{
		public.POST("/login", h.Login)
		public.GET("/sysinfo", h.GetSysInfo)
		public.POST("/register", h.Register)
	}

	// 需要认证的路由
	auth := r.Group("/api")
	auth.Use(middleware.Auth(cfg.JWTSecret))
	{
		auth.GET("/user", h.GetUserInfo)
		auth.PUT("/user/password", h.UpdatePassword)
		auth.POST("/logout", h.Logout)
	}

	// 管理员路由
	admin := r.Group("/api/admin")
	admin.Use(middleware.Auth(cfg.JWTSecret), middleware.AdminOnly())
	{
		admin.GET("/users", h.ListUsers)
		admin.POST("/users", h.CreateUser)
		admin.PUT("/users/:id", h.UpdateUser)
		admin.DELETE("/users/:id", h.DeleteUser)
		admin.POST("/users/:id/reset-password", h.ResetUserPassword)
		admin.GET("/options", h.GetOptions)
		admin.GET("/options/:name", h.GetOption)
		admin.PUT("/options/:name", h.UpdateOption)
	}

	return r
} 