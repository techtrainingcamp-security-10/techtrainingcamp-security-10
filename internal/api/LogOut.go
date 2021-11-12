package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/route/service"
)

// LogOut
// @Description 登出或注销
// @Router /api/logout [delete]
func LogOut(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form constants.LogOutType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			code, message := LogOutLogic(form.SessionID, int(form.ActionType), s)
			if code == constants.FailedCode {
				context.JSON(constants.DELETESuccessCode, gin.H{
					"Code":    constants.FailedCode,
					"Message": constants.LogOutFailed,
				})
			} else {
				context.JSON(constants.DELETEFailedCode, gin.H{
					"Code":    constants.SuccessCode,
					"Message": message,
				})
			}
		} else {
			fmt.Println(err)
		}
	}
}

// LogOutLogic
// @Description 登出或注销逻辑
func LogOutLogic(sessionID string, actionType int, s service.Service) (int, string) {
	phoneNumber := s.GetPhoneNumberBySessionId(sessionID)
	if phoneNumber == "nil" {
		return constants.FailedCode, constants.UserLoginStateInvalid
	}
	if result := s.DeleteSessionId(sessionID); result == false {
		return constants.FailedCode, constants.UserLoginStateInvalid
	}
	if actionType == 2 {
		s.DeleteUserByPhoneNumber(phoneNumber)
		return constants.SuccessCode, constants.CancellationSuccess
	}
	return constants.SuccessCode, constants.LogOutSuccess
}
