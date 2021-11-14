package resource

import (
	"github.com/jinzhu/gorm"

	"techtrainingcamp-security-10/internal/service"
)

var server resource

type resource struct {
	DbR     *gorm.DB
	DbW     *gorm.DB
	Redis   *Redis
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
	server.DbR.AutoMigrate(&service.UserTable{})
	// 2. cache
	redisOpts := &cfg.Redis
	server.Redis = NewRedis(redisOpts)
	server.Service = service.New(server.Redis.Conn, server.DbR)
	return &server, nil
}

func (*resource) Close() {
	// 资源释放
	return
}
