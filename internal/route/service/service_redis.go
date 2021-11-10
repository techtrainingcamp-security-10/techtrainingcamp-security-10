package service

import (
	"github.com/gomodule/redigo/redis"
)

// InsertVerifyCode
// @Description 插入验证码
func (s *service) InsertVerifyCode(phoneNumber string, verifyCode string) bool {
	_, err := s.cache.Get().Do("setex", phoneNumber+SplitChar+"VerifyCode", VerifyCodeExpireTime, verifyCode)
	if err != nil {
		return false
	}
	return true
}

// GetVerifyCode
// @Description 获取验证码
func (s *service) GetVerifyCode(phoneNumber string) string {
	data, err := redis.String(s.cache.Get().Do("get", phoneNumber+SplitChar+"VerifyCode"))
	if err != nil {
		return "nil"
	}
	return data
}

// TODO 登录注册成功后删除验证码

// InsertSessionId
// @Description 插入SessionId
func (s *service) InsertSessionId(phoneNumber string, sessionID string) bool {
	_, err := s.cache.Get().Do("setex", sessionID+SplitChar+"SessionId", SessionIdExpireTime, phoneNumber)
	if err != nil {
		return false
	}
	return true
}

// GetPhoneNumberBySessionId
// @Description 通过SessionId获取手机号
func (s *service) GetPhoneNumberBySessionId(sessionID string) string {
	data, err := redis.String(s.cache.Get().Do("get", sessionID+SplitChar+"SessionId"))
	if err != nil {
		return "nil"
	}
	return data
}

// DeleteSessionId
// @Description 删除SessionId记录
func (s *service) DeleteSessionId(sessionID string) bool {
	_, err := redis.Int(s.cache.Get().Do("del", sessionID+SplitChar+"SessionId"))
	if err != nil {
		return false
	}
	return true
}
