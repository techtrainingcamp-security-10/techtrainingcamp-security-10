// Package service 增删改查操作
package middleware

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"techtrainingcamp-security-10/internal/api"
)

type service struct {
	dbR   *gorm.DB
	dbW   *gorm.DB
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
	res := s.dbW.Create(Data)
	if res.Error != nil {
		return err
	}
	return nil
}

func New(cache *redis.Pool, dbR *gorm.DB, dbW *gorm.DB) Service {
	return &service{
		dbR:   dbR,
		dbW:   dbW,
		cache: cache,
	}
}

func (s *service) i() {}
