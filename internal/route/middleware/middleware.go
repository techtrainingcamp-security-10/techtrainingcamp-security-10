package middleware

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"techtrainingcamp-security-10/internal/route/service"
)

type middleware struct {
	cache *redis.Pool
	db    *gorm.DB
	// 具体功能
	service service.Service
}

type Middleware interface {
	// 为了避免被其他包实现
	i()

	// 实现功能接口
}

func NewMiddleware(cache *redis.Pool, db *gorm.DB) Middleware {
	return &middleware{
		cache:   cache,
		db:      db,
		service: service.New(cache, db),
	}
}

func (m *middleware) i() {}
