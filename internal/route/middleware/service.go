// Package service 增删改查操作
package middleware

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"techtrainingcamp-security-10/internal/api"
)

type service struct {
	db    *gorm.DB
	cache *redis.Pool
}

var _ Service = (*service)(nil)

type Service interface {
	// 为了避免被其他包实现
	i()

	Create(Data *api.UserTable) (err error)

	// 增加其他功能
}

// Create
// @Description 增加一条记录
// @Type UserTable
func (s *service) Create(Data *api.UserTable) (err error) {
	// 建表
	// s.db.AutoMigrate(&api.RegisterType{})
	res := s.db.Create(Data)
	if res.Error != nil {
		return err
	}
	return nil
}

func New(cache *redis.Pool, db *gorm.DB) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
