package route

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"techtrainingcamp-security-10/internal/api"
	"techtrainingcamp-security-10/internal/route/middleware"
	"techtrainingcamp-security-10/internal/route/service"
)

// NewRoute Restful 格式
func NewRoute(s service.Service) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(middleware.EnvCheck(s))
	{
		// 1. ApplyCode
		router.GET("/api/apply_code", api.ApplyCode(s))
		// 2. Register
		router.POST("/api/register", api.Register(s))
		// 3. Login
		router.POST("/api/login_uid", api.LoginByUID(s))
		router.POST("/api/login_phone", api.LoginByPhone(s))
		// 4. LogOut
		router.DELETE("/api/logout", api.LogOut(s))
	}
	return router, nil
}
