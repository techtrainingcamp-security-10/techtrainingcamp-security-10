package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

const (
	UserNameAlreadyExists    = "用户名已存在"
	PhoneNumberAlreadyExists = "电话号码已被注册"
	RegisterSuccess          = "注册成功"
	VerifyCodeInvalid        = "验证码无效"
	VerifyCodeError          = "验证码错误"
	LoginFailed              = "用户名或密码错误"
	LoginSuccess             = "登录成功"
	PhoneNumberNotRegister   = "手机号未注册"
	UserLoginStateInvalid    = "用户登录状态失效"
	LogOutSuccess            = "登出成功"
	CancellationSuccess      = "注销成功"
)

// QueryByUserName 根据userName获取user信息
func QueryByUserName(dbRead *gorm.DB, userName string) UserTable {
	var user UserTable
	dbRead.Where("user_name = ?", userName).First(&user)
	return user
}

// QueryByPhoneNumber 根据phoneNumber获取user信息
func QueryByPhoneNumber(dbRead *gorm.DB, phoneNumber string) UserTable {
	var user UserTable
	dbRead.Where("phone_number = ?", phoneNumber).First(&user)
	return user
}

//
func DeleteUserByPhoneNumber(dbRead *gorm.DB, phoneNumber string) bool {
	var user UserTable
	dbRead.Where("phone_number = ?", phoneNumber).Delete(&user)
	return true
}

// InsertUser 向数据库新增user
func InsertUser(cache *redis.Pool, dbRead *gorm.DB, verifyCode string, user UserTable) (int, string) {
	//测试验证码是否有效
	verifyCodeResult := GetVerifyCode(cache, user.PhoneNumber)
	if verifyCodeResult == "nil" {
		return 1, VerifyCodeInvalid
	}
	if verifyCodeResult != verifyCode {
		return 1, VerifyCodeError
	}

	if QueryByUserName(dbRead, user.UserName) != (UserTable{}) {
		return 1, UserNameAlreadyExists
	}
	if QueryByPhoneNumber(dbRead, user.PhoneNumber) != (UserTable{}) {
		return 1, PhoneNumberAlreadyExists
	}
	createUserResult := dbRead.Create(user)
	if createUserResult.RowsAffected == 1 {
		return 0, RegisterSuccess
	} else {
		return 1, createUserResult.Error.Error()
	}
}

// LoginByUserName 通过用户名登录
func LoginByUserName(dbRead *gorm.DB, userName string, password string) (int, string) {
	user := QueryByUserName(dbRead, userName)
	if user == (UserTable{}) {
		return 1, LoginFailed
	}

	//TODO 哈希算法判断密码是否正确
	if password == user.Password {
		return 0, LoginSuccess
	} else {
		return 1, LoginFailed
	}
}

// LoginByPhoneNumber 通过手机号登录
func LoginByPhoneNumber(cache *redis.Pool, dbRead *gorm.DB, phoneNumber string, verifyCode string) (int, string) {
	//测试验证码是否有效
	verifyCodeResult := GetVerifyCode(cache, phoneNumber)
	if verifyCodeResult == "nil" {
		return 1, VerifyCodeInvalid
	}
	if verifyCodeResult != verifyCode {
		return 1, VerifyCodeError
	}

	user := QueryByPhoneNumber(dbRead, phoneNumber)
	if user == (UserTable{}) {
		return 1, PhoneNumberNotRegister
	}
	return 0, LoginSuccess
}

// LogOutBySessionID 登出或注销
func LogOutBySessionID(cache *redis.Pool, dbRead *gorm.DB, sessionID string, actionType int) (int, string) {
	result, phoneNumber := DeleteSessionId(cache, sessionID)
	if result == -1 || result == -2 {
		return 1, UserLoginStateInvalid
	}
	if actionType == 2 {
		DeleteUserByPhoneNumber(dbRead, phoneNumber)
		return 0, CancellationSuccess
	}
	return 0, LogOutSuccess
}

// TestMysql
// @Description 测试mysql
// @Router /api/test_mysql [get]

func TestMysql(cache *redis.Pool, dbRead *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//code,message := InsertUser(cache,dbRead,"nil",UserTable{
		//	UserName:    "小李",
		//	Password:    "12345",
		//	PhoneNumber: "123445",
		//	Salt:        "abcdef",
		//})
		//code,message := LoginByUserName(dbRead,"小明","123456")
		//code,message := LoginByPhoneNumber(cache,dbRead,"12","123456")
		dbRead.AutoMigrate(&UserTable{})
		InsertSessionId(cache, "123", "123")
		code, message := LogOutBySessionID(cache, dbRead, "123", 2)
		context.JSON(200, gin.H{
			// "checkRev": form,
			"Code":    code,
			"Message": message,
			"Data": gin.H{
				"VerifyCode":   1234,
				"ExpireTime":   180,
				"DecisionType": 0,
			},
		})
	}
}
