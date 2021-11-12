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
				return
			}
			code, message := RegisterLogic(form, s)
			if code == constants.FailedCode {

				decisionType := utils.CheckFailRecords(s, form.UserName)

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

				// 生成 sessionID 及 失效时间
				sessionId := getSessionId()
				expireTime := service.SessionIdExpireTime

				utils.ClearFailRecords(s, phoneNumber)

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
	phoneNumber := strconv.Itoa(int(data.PhoneNumber))
	verifyCodeResult := s.GetVerifyCode(phoneNumber)
	switch {
	case verifyCodeResult == "nil": // 验证码不合法
		return constants.FailedCode, constants.VerifyCodeInvalid
	case !utils.IsNormalPhoneNumber(phoneNumber): // 虚拟号段
		return constants.FailedCode, constants.PhoneNumberStateErr
	case verifyCodeResult != data.VerifyCode: // 验证码不正确
		return constants.FailedCode, constants.VerifyCodeError
	case utils.SensitiveWordsFilter.Query(data.UserName) != 0: // 用户名含敏感词
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return constants.FailedCode, constants.UserNameErr
	case s.QueryByUserName(data.UserName) != (service.UserTable{}): // 用户名已注册
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return constants.FailedCode, constants.UserNameAlreadyExists
	case s.QueryByPhoneNumber(strconv.Itoa(int(data.PhoneNumber))) != (service.UserTable{}): // 手机号已注册
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		return constants.FailedCode, constants.PhoneNumberAlreadyExists
	default:
		// 手机验证码失效
		s.DeleteVerifyCode(phoneNumber)
		// 加密密码 加密后前 8 位是 salt 后 20 位是哈希后的值
		PasswordAddSalt := service.NewPassword(data.Password)
		fmt.Println(PasswordAddSalt)
		user := service.UserTable{
			UserName:    data.UserName,
			Password:    PasswordAddSalt,
			PhoneNumber: strconv.Itoa(int(data.PhoneNumber)),
		}
		if err := s.InsertUser(user); err != nil {
			fmt.Println(err)
		}
		return constants.SuccessCode, constants.RegisterSuccess
	}
}
