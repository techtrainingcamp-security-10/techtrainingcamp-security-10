package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

// LogOut
// @Description 登出或注销
// @Router /api/logout [delete]
func LogOut(cache *redis.Pool, dbRead *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form LogOutType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			// 实现具体业务逻辑
			context.JSON(201, gin.H{
				//"checkRev": form,
				"Code":    0,
				"Message": "登出成功", // "注销成功"
			})
		} else {
			fmt.Println(err)
			context.JSON(400, gin.H{
				// "checkRev": form,
				"Code":    1,
				"Message": "登出失败", // "注销失败"
			})
		}
	}
}
