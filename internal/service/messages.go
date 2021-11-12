package service

import "github.com/jinzhu/gorm"

type UserTable struct {
	gorm.Model
	UserName    string   `json:"UserName" binding:"required"`
	Password    Password `json:"Password" binding:"required"`
	PhoneNumber string   `json:"PhoneNumber" binding:"required"`
	Salt        string   `json:"Salt" binding:"required"`
}

const (
	VerifyCodeLength     = 6
	VerifyCodeExpireTime = 3 * 60
	SessionIdExpireTime  = 3 * 60 * 60
	SplitChar            = "|"
)
