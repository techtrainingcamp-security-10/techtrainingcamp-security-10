package resource

import (
	"fmt"
	"techtrainingcamp-security-10/internal/route/middleware"

	"github.com/jinzhu/gorm"
)

var server resource

type resource struct {
	DbR     *gorm.DB
	DbW     *gorm.DB
	Redis   *Redis
	Middles middleware.Middleware
}

func NewServer() (*resource, error) {
	cfg, err := GetConfig()
	fmt.Printf("%v\n", cfg)
	if err != nil {
		return &server, err
	}
	// 1. database
	dbReadOpts := &cfg.Mysql
	server.DbR, err = NewDB(dbReadOpts)
	if err != nil {
		return nil, err
	}
	/*dbWriteOpts := &MySQLOpts{
		Address: "127.0.0.1:1234",
		// 连接信息
	}
	server.DbW, err = NewDB(dbWriteOpts)*/
	//if err != nil {
	//	return nil, err
	//}

	// 2. cache
	redisOpts := &cfg.Redis
	server.Redis = NewRedis(redisOpts)
	// 3. Middleware
	server.Middles = middleware.NewMiddleware(server.Redis.Conn, server.DbR)

	return &server, nil
}

func (*resource) Close() {
	// 资源释放
	return
}
