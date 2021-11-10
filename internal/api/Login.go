package api

import (
	"fmt"
	"strconv"
	"techtrainingcamp-security-10/internal/route/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofrs/uuid"
)

// LoginByUID
// @Description 用户名登录
// @Router /api/login_uid [post]
func LoginByUID(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form LoginUIDType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			code, message, phoneNumber := LoginByUIDLogic(form.UserName, form.Password, s)

			if code == FailedCode {

				// TODO 调用失败次数统计决定 DecisionType
				decisionType := Normal

				context.JSON(POSTFailedCode, gin.H{
					"Code":      FailedCode,
					"Message":   message,
					"SessionID": "",
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": decisionType,
					},
				})
			} else {

				// TODO 生成 sessionID 及 失效时间
				sessionId := getSessionId()
				expireTime := service.SessionIdExpireTime

				s.InsertSessionId(phoneNumber, sessionId)
				context.JSON(POSTSuccessCode, gin.H{
					"Code":      SuccessCode,
					"Message":   LoginSuccess,
					"SessionID": sessionId,
					"Data": gin.H{
						"SessionID":    sessionId,
						"ExpireTime":   expireTime,
						"DecisionType": Normal,
					},
				})
			}
		} else {
			fmt.Println(err)
		}
	}
}

// LoginByUIDLogic
// @Description 用户名登录验证逻辑
func LoginByUIDLogic(UserName string, Password string, s service.Service) (int, string, string) {
	user := s.QueryByUserName(UserName)

	// TODO 校验密码
	PasswordAddSalt := Password

	switch {
	case user == (service.UserTable{}): // 用户名不存在
		return FailedCode, UserNameNotRegister, ""
	case PasswordAddSalt != user.Password: // 密码错误
		return FailedCode, LoginFailed, ""
	default:
		return SuccessCode, LoginSuccess, user.PhoneNumber
	}
}

// LoginByPhone
// @Description 手机登录
// @Router /api/login_phone [post]
func LoginByPhone(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form LoginPhoneType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			phoneNumber := strconv.Itoa(int(form.PhoneNumber))
			code, message := LoginByPhoneLogic(phoneNumber, form.VerifyCode, s)

			if code == FailedCode {

				// TODO 调用失败次数统计决定 DecisionType
				decisionType := Normal

				context.JSON(POSTFailedCode, gin.H{
					"Code":      FailedCode,
					"Message":   message,
					"SessionID": "",
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": decisionType,
					},
				})
			} else {
				// TODO 生成 sessionID 及 失效时间
				sessionId := getSessionId()
				expireTime := service.SessionIdExpireTime

				s.InsertSessionId(phoneNumber, sessionId)
				context.JSON(POSTSuccessCode, gin.H{
					"Code":      SuccessCode,
					"Message":   LoginSuccess,
					"SessionID": sessionId,
					"Data": gin.H{
						"SessionID":    sessionId,
						"ExpireTime":   expireTime,
						"DecisionType": Normal,
					},
				})
			}
		} else {
			fmt.Println(err)
		}
	}
}

// LoginByPhoneLogic
// @Description 手机号登录验证逻辑
func LoginByPhoneLogic(phoneNumber string, verifyCode string, s service.Service) (int, string) {
	verifyCodeResult := s.GetVerifyCode(phoneNumber)
	switch {
	case verifyCodeResult == "nil": // 验证码不合法
		return FailedCode, VerifyCodeInvalid
	case verifyCodeResult != verifyCode: // 验证码不正确
		return FailedCode, VerifyCodeError
	case s.QueryByPhoneNumber(phoneNumber) == (service.UserTable{}): // 用户不存在
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return FailedCode, PhoneNumberNotRegister
	default:
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return SuccessCode, LoginSuccess
	}
}

// getSessionId
// @Description 生成随机UUID
func getSessionId() string {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("failed to generate UUID: %v\n", err)
		return ""
	}
	return id.String()
}
