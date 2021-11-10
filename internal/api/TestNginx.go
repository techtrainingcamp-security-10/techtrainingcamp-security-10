package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

func TestNginx(cache *redis.Pool, dbRead *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			// "checkRev": form,
			"Code":    1,
			"Message": "请求成功",
			"port":    8080,
		})
	}
}
