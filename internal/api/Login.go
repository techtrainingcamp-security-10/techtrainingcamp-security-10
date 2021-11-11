package api

import (
	"fmt"
	"strconv"
	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/route/service"
	"techtrainingcamp-security-10/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gofrs/uuid"
)

// LoginByUID
// @Description 用户名登录
// @Router /api/login_uid [post]
func LoginByUID(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form constants.LoginUIDType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			code, message, phoneNumber := LoginByUIDLogic(form.UserName, form.Password, s)

			if code == constants.FailedCode {

				decisionType := utils.CheckFailRecords(s, context.Request.RequestURI, context.Request.Method, form.UserName)

				context.JSON(constants.POSTFailedCode, gin.H{
					"Code":      constants.FailedCode,
					"Message":   message,
					"SessionID": "",
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": decisionType,
					},
				})
			} else {

				// 生成 sessionID 及 失效时间
				sessionId := getSessionId()
				expireTime := service.SessionIdExpireTime

				utils.ClearFailRecords(s, context.Request.RequestURI, context.Request.Method, form.UserName)

				s.InsertSessionId(phoneNumber, sessionId)
				context.JSON(constants.POSTSuccessCode, gin.H{
					"Code":      constants.SuccessCode,
					"Message":   constants.LoginSuccess,
					"SessionID": sessionId,
					"Data": gin.H{
						"SessionID":    sessionId,
						"ExpireTime":   expireTime,
						"DecisionType": constants.Normal,
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
	switch {
	case user == (service.UserTable{}): // 用户名不存在
		return constants.FailedCode, constants.UserNameNotRegister, ""
	case !user.Password.Verify(Password): // 密码错误
		return constants.FailedCode, constants.LoginFailed, ""
	default:
		return constants.SuccessCode, constants.LoginSuccess, user.PhoneNumber
	}
}

// LoginByPhone
// @Description 手机登录
// @Router /api/login_phone [post]
func LoginByPhone(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form constants.LoginPhoneType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			phoneNumber := strconv.Itoa(int(form.PhoneNumber))
			code, message := LoginByPhoneLogic(phoneNumber, form.VerifyCode, s)

			if code == constants.FailedCode {

				decisionType := utils.CheckFailRecords(s, context.Request.RequestURI, context.Request.Method, strconv.Itoa(int(form.PhoneNumber)))

				context.JSON(constants.POSTFailedCode, gin.H{
					"Code":      constants.FailedCode,
					"Message":   message,
					"SessionID": "",
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": decisionType,
					},
				})
			} else {
				// 生成 sessionID 及 失效时间
				sessionId := getSessionId()
				expireTime := service.SessionIdExpireTime

				utils.ClearFailRecords(s, context.Request.RequestURI, context.Request.Method, strconv.Itoa(int(form.PhoneNumber)))

				s.InsertSessionId(phoneNumber, sessionId)
				context.JSON(constants.POSTSuccessCode, gin.H{
					"Code":      constants.SuccessCode,
					"Message":   constants.LoginSuccess,
					"SessionID": sessionId,
					"Data": gin.H{
						"SessionID":    sessionId,
						"ExpireTime":   expireTime,
						"DecisionType": constants.Normal,
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
		return constants.FailedCode, constants.VerifyCodeInvalid
	case utils.IsVirtualPhoneNumber(phoneNumber): // 虚拟号段
		return constants.FailedCode, constants.PhoneNumberStateErr
	case verifyCodeResult != verifyCode: // 验证码不正确
		return constants.FailedCode, constants.VerifyCodeError
	case s.QueryByPhoneNumber(phoneNumber) == (service.UserTable{}): // 用户不存在
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return constants.FailedCode, constants.PhoneNumberNotRegister
	default:
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return constants.SuccessCode, constants.LoginSuccess
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
