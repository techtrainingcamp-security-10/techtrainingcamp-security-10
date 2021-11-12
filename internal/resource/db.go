package resource

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLOpts struct {
	User     string `yml:"user" json:"user"`
	Password string `yml:"password" json:"password"`
	Address  string `yml:"address" json:"address"`
	Name     string `yml:"name" json:"name"`
}

func NewDB(opts *MySQLOpts) (*gorm.DB, error) {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		opts.User,
		opts.Password,
		opts.Address,
		opts.Name,
		true,
		"Local")
	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		return nil, err
	}
	return db, err
}