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
// @Type UserTable
func (s *service) InsertUser(user UserTable) error {
	createUserResult := s.db.Create(user)
	if createUserResult.RowsAffected == 1 {
		return nil
	} else {
		return createUserResult.Error
	}
}
