package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"techtrainingcamp-security-10/internal/route/service"
)

// Register
// @Description 注册用户
// @Router /api/register [post]
func Register(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form RegisterType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			// 实现具体业务逻辑
			context.JSON(201, gin.H{
				// "checkRev": form,
				"Code":      0,
				"Message":   "注册成功",
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
				"Message": "相同的用户名已经被注册过了，请更换用户名试试",
				"Data": gin.H{
					"SessionID":    "",
					"ExpireTime":   "",
					"DecisionType": 1, // 1表示需要用户通过滑块验证，通过后才能注册，2表示需要用户过一段时间，才能重新注册，3表示这个用户不能注册
				},
			})
		}
	}
}
