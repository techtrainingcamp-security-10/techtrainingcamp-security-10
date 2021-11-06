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

// InsertSessionId
// @Description 插入SessionId
func (s *service) InsertSessionId(phoneNumber string, sessionID string) bool {
	_, err := s.cache.Get().Do("setex", sessionID+SplitChar+"SessionId", SessionIdExpireTime, phoneNumber)
	if err != nil {
		return false
	}
	return true
}

// DeleteSessionId
// @Description 删除SessionId
func (s *service) DeleteSessionId(sessionID string) (int, string) {
	phoneNumber, err1 := redis.String(s.cache.Get().Do("get", sessionID+SplitChar+"SessionId"))
	if err1 != nil {
		return -1, ""
	}
	data, err2 := redis.Int(s.cache.Get().Do("del", sessionID+SplitChar+"SessionId"))
	if err2 != nil {
		return -2, ""
	}
	return data, phoneNumber
}
