package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"techtrainingcamp-security-10/internal/route/service"
)

// ApplyCode
// @Description 获取验证码
// @Router /api/apply_code [get]
func ApplyCode(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form ApplyCodeType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			// 实现具体业务逻辑
			context.JSON(200, gin.H{
				// "checkRev": form,
				"Code":    0,
				"Message": "请求成功",
				"Data": gin.H{
					"VerifyCode":   1234,
					"ExpireTime":   180,
					"DecisionType": 0,
				},
			})
		} else {
			fmt.Println(err)
			context.JSON(500, gin.H{
				// "checkRev": form,
				"Code":    1,
				"Message": "请求失败",
				"Data": gin.H{
					"VerifyCode":   "",
					"ExpireTime":   "",
					"DecisionType": 1, // 1表示需要用户通过滑块验证，通过后才能注册，2表示需要用户过一段时间，才能重新注册，3表示这个用户不能注册
				},
			})
		}
	}
}
