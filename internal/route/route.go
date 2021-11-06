package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"techtrainingcamp-security-10/internal/api"
	"techtrainingcamp-security-10/internal/route/middleware"
)

// NewRoute Restful 格式
func NewRoute(cache *redis.Pool, dbRead *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()
	router.Use(middleware.EnvCheck(cache, dbRead))
	{
		// 1. ApplyCode
		router.GET("/api/apply_code", api.ApplyCode(cache, dbRead))
		// 2. Register
		router.POST("/api/register", api.Register(cache, dbRead))
		// 3. Login
		router.POST("/api/login_uid", api.LoginUID(cache, dbRead))
		router.POST("/api/login_phone", api.LoginPhone(cache, dbRead))
		// 4. LogOut
		router.DELETE("/api/logout", api.LogOut(cache, dbRead))

		//5. Test redis and mysql
		router.GET("/api/test_redis", api.TestRedis(cache, dbRead))
		router.GET("/api/test_mysql", api.TestMysql(cache, dbRead))
	}
	return router, nil
}
