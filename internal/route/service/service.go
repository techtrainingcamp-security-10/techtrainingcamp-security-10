// Package middleware 增删改查操作
package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type service struct {
	db    *gorm.DB
	cache *redis.Pool
}

var _ Service = (*service)(nil)

type Service interface {
	// 为了避免被其他包实现
	i()

	QueryByUserName(userName string) UserTable
	QueryByPhoneNumber(phoneNumber string) UserTable
	DeleteUserByPhoneNumber(phoneNumber string) bool
	InsertUser(verifyCode string, user UserTable) (int, string)
	LoginByUserName(userName string, password string) (int, string)
	LoginByPhoneNumber(phoneNumber string, verifyCode string) (int, string)
	LogOutBySessionID(sessionID string, actionType int) (int, string)
	InsertVerifyCode(phoneNumber string, verifyCode string) bool
	GetVerifyCode(phoneNumber string) string
	InsertSessionId(phoneNumber string, sessionID string) bool
	DeleteSessionId(sessionID string) (int, string)
}

func New(cache *redis.Pool, db *gorm.DB) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
