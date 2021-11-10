package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
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
			phoneNumber := strconv.Itoa(int(form.PhoneNumber))
			code, message := RegisterLogic(form, s)
			if code == FailedCode {

				// TODO 调用失败次数统计决定 DecisionType
				decisionType := Normal

				context.JSON(POSTFailedCode, gin.H{
					"Code":    FailedCode,
					"Message": message,
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": decisionType,
					},
				})
			} else {

				// TODO 生成 sessionID 及 失效时间
				sessionId := "123456"
				expireTime := service.SessionIdExpireTime

				s.InsertSessionId(phoneNumber, sessionId)
				context.JSON(POSTSuccessCode, gin.H{

					"Code":    SuccessCode,
					"Message": RegisterSuccess,
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

// RegisterLogic
// @Description 注册新用户逻辑
func RegisterLogic(data RegisterType, s service.Service) (int, string) {
	//测试验证码是否有效
	phoneNumber := strconv.Itoa(int(data.PhoneNumber))
	verifyCodeResult := s.GetVerifyCode(phoneNumber)
	switch {
	case verifyCodeResult == "nil": // 验证码不合法
		return FailedCode, VerifyCodeInvalid
	case verifyCodeResult != data.VerifyCode: // 验证码不正确
		return FailedCode, VerifyCodeError
	case s.QueryByUserName(data.UserName) != (service.UserTable{}): // 用户名已注册
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return FailedCode, UserNameAlreadyExists
	case s.QueryByPhoneNumber(strconv.Itoa(int(data.PhoneNumber))) != (service.UserTable{}): // 手机号已注册
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return FailedCode, PhoneNumberAlreadyExists
	default:
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		// TODO 加密密码
		PasswordAddSalt := data.Password
		Salt := "1234"

		user := service.UserTable{
			UserName:    data.UserName,
			Password:    PasswordAddSalt,
			PhoneNumber: strconv.Itoa(int(data.PhoneNumber)),
			Salt:        Salt,
		}
		if err := s.InsertUser(user); err != nil {
			fmt.Println(err)
		}
		return SuccessCode, RegisterSuccess
	}
}
