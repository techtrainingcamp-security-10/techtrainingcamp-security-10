package api

import (
	"fmt"
	"techtrainingcamp-security-10/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"strconv"
	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/route/service"
)

// ApplyCode
// @Description 获取验证码
// @Router /api/apply_code [get]
func ApplyCode(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form constants.ApplyCodeType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			phoneNumber := strconv.Itoa(int(form.PhoneNumber))
			// 判断是否被频控或封禁
			if limitType := s.GetUserLimitType(phoneNumber); limitType != 0 {
				var message string
				switch limitType {
				case 2:
					message = constants.FrequencyLimit
				case 3:
					message = constants.Lock
				}
				context.JSON(constants.POSTFailedCode, gin.H{
					"Code":      constants.FailedCode,
					"Message":   message,
					"SessionID": "",
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": limitType,
					},
				})
			}
			// 对 PhoneNumber 判断风险
			if utils.IsVirtualPhoneNumber(phoneNumber) {
				context.JSON(constants.GETSuccessCode, gin.H{
					"Code":    constants.FailedCode,
					"Message": constants.PhoneNumberStateErr,
					"Data": gin.H{
						"VerifyCode":   "",
						"ExpireTime":   "",
						"DecisionType": constants.Normal,
					},
				})
			} else {
				// 基于时间戳的随机种子
				rand.Seed(time.Now().UnixNano())
				validCode := RandomString(service.VerifyCodeLength, defaultLetters)
				s.InsertVerifyCode(phoneNumber, validCode)
				context.JSON(constants.GETSuccessCode, gin.H{
					"Code":    constants.SuccessCode,
					"Message": constants.RequestSuccess,
					"Data": gin.H{
						"VerifyCode":   validCode,
						"ExpireTime":   service.SessionIdExpireTime,
						"DecisionType": constants.Normal,
					},
				})
			}
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
