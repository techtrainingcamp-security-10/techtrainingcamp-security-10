package middleware

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type middleware struct {
	cache *redis.Pool
	dbR   *gorm.DB
	dbW   *gorm.DB
	// 具体功能
	service Service
}

type Middleware interface {
	// 为了避免被其他包实现
	i()

	// 实现功能接口
}

func NewMiddleware(cache *redis.Pool, dbR *gorm.DB, dbW *gorm.DB) Middleware {
	return &middleware{
		cache:   cache,
		dbR:     dbR,
		dbW:     dbW,
		service: New(cache, dbR, dbW),
	}
}

func (m *middleware) i() {}
