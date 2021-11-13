package api

import (
	"crypto/rand"
	"fmt"
	"strconv"

	// "time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/service"
	"techtrainingcamp-security-10/internal/utils"
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
			if !utils.IsNormalPhoneNumber(phoneNumber) {
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
				// rand.Seed(time.Now().UnixNano())
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

// 生成随机byte序列
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return b, nil
}

// RandomString returns a random string with a fixed length
// https://zhuanlan.zhihu.com/p/94684495

// 采用纯数字手机验证码
// var defaultLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var defaultLetters = []byte("0123456789")
var defaultMod = byte(len(defaultLetters))

// 使用crypto/rand
func RandomString(n int, allowedChars ...[]byte) string {
	var letters []byte
	var mod byte
	if len(allowedChars) == 0 {
		letters = defaultLetters
		mod = defaultMod
	} else {
		letters = allowedChars[0]
		mod = byte(len(allowedChars[0]))
	}
	b, _ := RandomBytes(n)
	for i := range b {
		b[i] = letters[b[i]%mod]
	}
	return string(b)
}

// 使用math/rand
// func RandomString(n int, allowedChars... []byte) string {
// 	var letters []byte
// 	if len(allowedChars) == 0 {
// 		letters = defaultLetters
// 	} else {
// 		letters = allowedChars[0]
// 	}
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}
// 	return string(b)
// }
