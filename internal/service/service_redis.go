package service

import (
	"github.com/gomodule/redigo/redis"

	"techtrainingcamp-security-10/internal/constants"
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

// DeleteVerifyCode
// @Description 登录注册成功后删除验证码
func (s *service) DeleteVerifyCode(phoneNumber string) bool {
	_, err := redis.Int(s.cache.Get().Do("del", phoneNumber+SplitChar+"VerifyCode"))
	if err != nil {
		return false
	}
	return true
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

// GetApiFailRecords
// @Description 获取某一 API 5s/1min 内调用失败记录；identifier 区别用户
func (s *service) GetApiFailRecords(identifier string) (int, int) {
	var err error
	data5s, err := redis.Int(s.cache.Get().Do("get", identifier+SplitChar+"5s"))
	if err != nil {
		return 0, 0
	}
	data1Min, err := redis.Int(s.cache.Get().Do("get", identifier+SplitChar+"60s"))
	if err != nil {
		return 0, 0
	}
	return data5s, data1Min
}

// SetApiFailRecords
// @Description 写入某一 API 5s/1min 内调用失败记录；identifier 区别用户
func (s *service) SetApiFailRecords(identifier string, records5s int, records1Min int) {
	var err error
	_, err = s.cache.Get().Do("set", identifier+SplitChar+"5s", records5s)
	if err != nil {
		return
	}
	_, err = s.cache.Get().Do("EXPIRE", identifier+SplitChar+"5s", 5)
	if err != nil {
		return
	}
	_, err = s.cache.Get().Do("set", identifier+SplitChar+"60s", records1Min)
	if err != nil {
		return
	}
	_, err = s.cache.Get().Do("EXPIRE", identifier+SplitChar+"60s", 60)
	if err != nil {
		return
	}

}

// GetUserLimitType
// @Description 查看某一用户名或手机号是否触发风控限制
func (s *service) GetUserLimitType(identifier string) int {
	conn := s.cache.Get()
	limitType, err1 := redis.Int(conn.Do("get", identifier+SplitChar+"LimitType"))
	_ = conn.Close()

	if err1 != nil {
		return constants.Normal
	} else {
		return limitType
	}
}

// SetUserLimitType
// @Description 设置某一用户名或手机号风控限制
func (s *service) SetUserLimitType(identifier string, limitType int) {
	conn := s.cache.Get()

	switch limitType {
	case constants.Normal:
		_, _ = conn.Do("del", identifier+SplitChar+"LimitType")
	case constants.FrequentLimit:
		_, _ = conn.Do("set", identifier+SplitChar+"LimitType", limitType)
		_, _ = conn.Do("EXPIRE", identifier+SplitChar+"LimitType", 10)
	case constants.Locked:
		_, _ = conn.Do("set", identifier+SplitChar+"LimitType", limitType)
	}

	_ = conn.Close()

}
