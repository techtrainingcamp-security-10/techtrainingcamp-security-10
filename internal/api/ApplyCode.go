package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"strconv"
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
			// 基于时间戳的随机种子
			rand.Seed(time.Now().UnixNano())
			validCode := RandomString(service.VerifyCodeLength, defaultLetters)

			// TODO 对 PhoneNumber 判断风险

			s.InsertVerifyCode(strconv.Itoa(int(form.PhoneNumber)), validCode)
			context.JSON(GETSuccessCode, gin.H{
				"Code":    SuccessCode,
				"Message": RequestSuccess,
				"Data": gin.H{
					"VerifyCode":   validCode,
					"ExpireTime":   service.SessionIdExpireTime,
					"DecisionType": Normal,
				},
			})
		} else {
			fmt.Println(err)
		}
	}
}

// RandomString returns a random string with a fixed length
// https://zhuanlan.zhihu.com/p/94684495
var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
