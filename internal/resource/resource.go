package resource

import (
	"techtrainingcamp-security-10/internal/route/middleware"
	"techtrainingcamp-security-10/internal/route/service"

	"github.com/jinzhu/gorm"
)

var server resource

type resource struct {
	DbR     *gorm.DB
	DbW     *gorm.DB
	Redis   *Redis
	Middles middleware.Middleware
	Service service.Service
}

func NewServer() (*resource, error) {
	cfg, err := GetConfig()
	if err != nil {
		return &server, err
	}
	// 1. database
	dbReadOpts := &cfg.Mysql
	server.DbR, err = NewDB(dbReadOpts)
	if err != nil {
		return nil, err
	}

	// 2. cache
	redisOpts := &cfg.Redis
	server.Redis = NewRedis(redisOpts)
	// 3. Middleware
	server.Middles = middleware.NewMiddleware(server.Redis.Conn, server.DbR)
	server.Service = service.New(server.Redis.Conn, server.DbR)
	return &server, nil
}

func (*resource) Close() {
	// 资源释放
	return
}
