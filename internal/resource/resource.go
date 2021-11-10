package resource

import (
	"github.com/jinzhu/gorm"
	"techtrainingcamp-security-10/internal/route/middleware"
	"techtrainingcamp-security-10/internal/route/service"
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
	var err error
	// 1. database
	dbReadOpts := &MySQLOpts{
		Address:  "127.0.0.1:3306",
		User:     "root",
		Password: "admin",
		Name:     "techtrainingcamp",
		// 连接信息
	}
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
	redisOpts := &RedisOpts{
		Host: "127.0.0.1:6379",
		// 连接信息
	}
	server.Redis = NewRedis(redisOpts)
	// 3. Middleware
	//server.Middles = middleware.NewMiddleware(server.Redis.Conn, server.DbR)
	// 4. Service
	server.Service = service.New(server.Redis.Conn, server.DbR)
	return &server, nil
}

func (*resource) Close() {
	// 资源释放
	return
}
