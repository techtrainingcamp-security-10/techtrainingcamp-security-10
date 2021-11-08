package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/route/service"
	"techtrainingcamp-security-10/internal/utils"
)

// Register
// @Description 注册用户
// @Router /api/register [post]
func Register(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form constants.RegisterType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			phoneNumber := strconv.Itoa(int(form.PhoneNumber))
			code, message := RegisterLogic(form, s)
			if code == constants.FailedCode {

				decisionType := utils.CheckFailRecords(s, context.Request.RequestURI, context.Request.Method, form.UserName)

				context.JSON(constants.POSTFailedCode, gin.H{
					"Code":    constants.FailedCode,
					"Message": message,
					"Data": gin.H{
						"SessionID":    "",
						"ExpireTime":   "",
						"DecisionType": decisionType,
					},
				})
			} else {

				// TODO 手机验证码失效

				// TODO 生成 sessionID 及 失效时间
				sessionId := "123456"
				expireTime := service.SessionIdExpireTime

				utils.ClearFailRecords(s, context.Request.RequestURI, context.Request.Method, strconv.Itoa(int(form.PhoneNumber)))

				s.InsertSessionId(phoneNumber, sessionId)
				context.JSON(constants.POSTSuccessCode, gin.H{

					"Code":    constants.SuccessCode,
					"Message": constants.RegisterSuccess,
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

// RegisterLogic
// @Description 注册新用户逻辑
func RegisterLogic(data constants.RegisterType, s service.Service) (int, string) {
	//测试验证码是否有效
	verifyCodeResult := s.GetVerifyCode(strconv.Itoa(int(data.PhoneNumber)))
	switch {
	case verifyCodeResult == "nil": // 验证码不合法
		return constants.FailedCode, constants.VerifyCodeInvalid
	case verifyCodeResult != data.VerifyCode: // 验证码不正确
		return constants.FailedCode, constants.VerifyCodeError
	case s.QueryByUserName(data.UserName) != (service.UserTable{}): // 用户名已注册
		return constants.FailedCode, constants.UserNameAlreadyExists
	case s.QueryByPhoneNumber(strconv.Itoa(int(data.PhoneNumber))) != (service.UserTable{}): // 手机号已注册
		return constants.FailedCode, constants.PhoneNumberAlreadyExists
	default:

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
		return constants.SuccessCode, constants.RegisterSuccess
	}
}
