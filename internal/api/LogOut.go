package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"techtrainingcamp-security-10/internal/route/service"
)

// LogOut
// @Description 登出或注销
// @Router /api/logout [delete]
func LogOut(s service.Service) gin.HandlerFunc {
	return func(context *gin.Context) {
		var form LogOutType
		err := context.ShouldBindBodyWith(&form, binding.JSON)
		if err == nil {
			code, message := LogOutLogic(form.SessionID, int(form.ActionType), s)
			if code == FailedCode {
				context.JSON(DELETESuccessCode, gin.H{
					"Code":    SuccessCode,
					"Message": message,
				})
			} else {
				context.JSON(DELETEFailedCode, gin.H{
					"Code":    FailedCode,
					"Message": LogOutFailed,
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
		return FailedCode, UserLoginStateInvalid
	}
	if result := s.DeleteSessionId(sessionID); result == false {
		return FailedCode, UserLoginStateInvalid
	}
	if actionType == 2 {
		s.DeleteUserByPhoneNumber(phoneNumber)
		return SuccessCode, CancellationSuccess
	}
	return SuccessCode, LogOutSuccess
}
