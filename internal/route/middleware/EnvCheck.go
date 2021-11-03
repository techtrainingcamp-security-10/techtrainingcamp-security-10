package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

type EnvironmentType struct {
	IP       string `json:"IP" binding:"required"`
	DeviceID string `json:"DeviceID" binding:"required"`
}

type RequestType struct {
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

func EnvCheck(cache *redis.Pool, dbRead *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 环境检测
		if envCheck(c) {
			c.Next()
		} else {
			c.JSON(500, gin.H{
				// "checkRev": form,
				"Code":    1,
				"Message": "请求失败",
				"Data": gin.H{
					"VerifyCode":   "",
					"ExpireTime":   "",
					"DecisionType": 3, // 1表示需要用户通过滑块验证，通过后才能注册，2表示需要用户过一段时间，才能重新注册，3表示这个用户不能注册
				},
			})
			c.Abort()
		}
	}
}

func envCheck(c *gin.Context) bool {
	var form RequestType
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err == nil {
		// 举个栗子
		if form.Environment.IP == "127.0.0.1" {
			return true
		}
	}
	return false
}
