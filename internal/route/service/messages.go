package service

import "techtrainingcamp-security-10/internal/utils"

type UserTable struct {
	UserName    string         `json:"UserName" binding:"required"`
	Password    utils.Password `json:"Password" binding:"required"`
	PhoneNumber string         `json:"PhoneNumber" binding:"required"`
	Salt        string         `json:"Salt" binding:"required"`
}

const (
	VerifyCodeLength     = 6
	VerifyCodeExpireTime = 3 * 60
	SessionIdExpireTime  = 3 * 60 * 60
	SplitChar            = "|"
)
