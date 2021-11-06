package service

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// QueryByUserName
// @Description 根据userName获取user信息
// @Type UserTable
func (s *service) QueryByUserName(userName string) UserTable {
	var user UserTable
	s.db.Where("user_name = ?", userName).First(&user)
	return user
}

// QueryByPhoneNumber
// @Description 根据phoneNumber获取user信息
// @Type UserTable
func (s *service) QueryByPhoneNumber(phoneNumber string) UserTable {
	var user UserTable
	s.db.Where("phone_number = ?", phoneNumber).First(&user)
	return user
}

// DeleteUserByPhoneNumber
// @Description 根据phoneNumber删除user信息
// @Type UserTable
func (s *service) DeleteUserByPhoneNumber(phoneNumber string) bool {
	var user UserTable
	s.db.Where("phone_number = ?", phoneNumber).Delete(&user)
	return true
}

// InsertUser
// @Description 向数据库新增user
func (s *service) InsertUser(verifyCode string, user UserTable) (int, string) {
	{ // TODO 拆分功能到注册api
		//测试验证码是否有效
		verifyCodeResult := s.GetVerifyCode(user.PhoneNumber)
		if verifyCodeResult == "nil" {
			return 1, VerifyCodeInvalid
		}
		if verifyCodeResult != verifyCode {
			return 1, VerifyCodeError
		}

		if s.QueryByUserName(user.UserName) != (UserTable{}) {
			return 1, UserNameAlreadyExists
		}
		if s.QueryByPhoneNumber(user.PhoneNumber) != (UserTable{}) {
			return 1, PhoneNumberAlreadyExists
		}
	}
	createUserResult := s.db.Create(user)
	if createUserResult.RowsAffected == 1 {
		return 0, RegisterSuccess
	} else {
		return 1, createUserResult.Error.Error()
	}
}

// LoginByUserName
// @Description 通过用户名登录
// @Return
// 0: 登录成功
// 1: 登录失败
func (s *service) LoginByUserName(userName string, password string) (int, string) {
	user := s.QueryByUserName(userName)
	if user == (UserTable{}) { // !为什么等于的时候登录失败
		return 1, LoginFailed
	}

	//TODO 哈希算法判断密码是否正确
	if password == user.Password {
		return 0, LoginSuccess
	} else {
		return 1, LoginFailed
	}
}

// LoginByPhoneNumber
// @Description 通过手机号登录
// @Return
// 0: 登录成功
// 1: 登录失败
func (s *service) LoginByPhoneNumber(phoneNumber string, verifyCode string) (int, string) {
	//测试验证码是否有效
	verifyCodeResult := s.GetVerifyCode(phoneNumber)
	if verifyCodeResult == "nil" {
		return 1, VerifyCodeInvalid
	}
	if verifyCodeResult != verifyCode {
		return 1, VerifyCodeError
	}

	user := s.QueryByPhoneNumber(phoneNumber)
	if user == (UserTable{}) {
		return 1, PhoneNumberNotRegister
	}
	return 0, LoginSuccess
}

// LogOutBySessionID
// @Description 登出或注销
// TODO 拆分功能到登出或注销api
func (s *service) LogOutBySessionID(sessionID string, actionType int) (int, string) {
	result, phoneNumber := s.DeleteSessionId(sessionID)
	if result == -1 || result == -2 {
		return 1, UserLoginStateInvalid
	}
	if actionType == 2 {
		s.DeleteUserByPhoneNumber(phoneNumber)
		return 0, CancellationSuccess
	}
	return 0, LogOutSuccess
}
