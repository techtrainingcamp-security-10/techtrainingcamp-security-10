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
	InsertUser(user UserTable) error
	InsertVerifyCode(phoneNumber string, verifyCode string) bool
	GetVerifyCode(phoneNumber string) string
	InsertSessionId(phoneNumber string, sessionID string) bool
	GetPhoneNumberBySessionId(sessionID string) string
	DeleteSessionId(sessionID string) bool
}

func New(cache *redis.Pool, db *gorm.DB) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}