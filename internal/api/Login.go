package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"techtrainingcamp-security-10/internal/route/service"
)

// LoginUID
// @Description 用户名登录
// @Router /api/login_uid [post]
func LoginUID(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form LoginUIDType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			// 实现具体业务逻辑
			context.JSON(201, gin.H{
				// "checkRev": form,
				"Code":      0,
				"Message":   "登录成功",
				"SessionID": 123456,
				"Data": gin.H{
					"SessionID":    123456,
					"ExpireTime":   180,
					"DecisionType": 0,
				},
			})
		} else {
			fmt.Println(err)
			context.JSON(400, gin.H{
				// "checkRev": form,
				"Code":    1,
				"Message": "用户名或者密码不对",
				"Data": gin.H{
					"SessionID":    "",
					"ExpireTime":   "",
					"DecisionType": 1, // 1表示需要用户通过滑块验证，通过后才能注册，2表示需要用户过一段时间，才能重新注册，3表示这个用户不能注册
				},
			})
		}
	}
}

// LoginPhone
// @Description 手机登录
// @Router /api/login_phone [post]
func LoginPhone(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form LoginPhoneType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			// 实现具体业务逻辑
			context.JSON(200, gin.H{
				// "checkRev": form,
				"Code":      0,
				"Message":   "登录成功",
				"SessionID": 123456,
				"Data": gin.H{
					"SessionID":    123456,
					"ExpireTime":   180,
					"DecisionType": 0,
				},
			})
		} else {
			fmt.Println(err)
			context.JSON(200, gin.H{
				// "checkRev": form,
				"Code":    1,
				"Message": "用户名或者密码不对",
				"Data": gin.H{
					"SessionID":    "",
					"ExpireTime":   "",
					"DecisionType": 1, // 1表示需要用户通过滑块验证，通过后才能注册，2表示需要用户过一段时间，才能重新注册，3表示这个用户不能注册
				},
			})
		}
	}
}
