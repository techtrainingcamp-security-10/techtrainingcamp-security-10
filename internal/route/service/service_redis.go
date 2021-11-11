package service

import (
	"github.com/gomodule/redigo/redis"
	"techtrainingcamp-security-10/internal/constants"
	"time"
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
// @Description 获取某一 API 调用失败记录；identifier 区别用户
func (s *service) GetApiFailRecords(apiRoute string, apiMethod string, identifier string) []int64 {
	conn := s.cache.Get()

	data, err := conn.Do("get", identifier+SplitChar+"FailRecords"+SplitChar+apiMethod+SplitChar+apiRoute)
	if err != nil || data == nil {
		return make([]int64, 0)
	}

	result := make([]int64, 0, len(data.([]uint8)))
	for _, record := range data.([]uint8) {
		result = append(result, int64(record))
	}

	_ = conn.Close()
	return result
}

// SetApiFailRecords
// @Description 写入某一 API 调用失败记录；identifier 区别用户
func (s *service) SetApiFailRecords(apiRoute string, apiMethod string, identifier string, records []int64) {
	conn := s.cache.Get()
	_, _ = conn.Do("set", identifier+SplitChar+"FailRecords"+SplitChar+apiMethod+SplitChar+apiRoute, records)
	_ = conn.Close()
}

// GetUserLimitType
// @Description 查看某一用户名或手机号是否触发风控限制
func (s *service) GetUserLimitType(identifier string) int {
	conn := s.cache.Get()
	limitType, err1 := redis.Int(conn.Do("get", identifier+SplitChar+"LimitType"))
	_ = conn.Close()

	if err1 != nil {
		return constants.Normal
	}

	if limitType == constants.FrequentLimit {
		limitExpired, err2 := redis.Int64(conn.Do("get", identifier+SplitChar+"LimitExpiredAt"))
		if err2 != nil || limitExpired < time.Now().Unix() {
			s.SetUserLimitType(identifier, constants.Normal)
			return constants.Normal
		}
	}

	return limitType
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
		_, _ = conn.Do("set", identifier+SplitChar+"LimitExpiredAt", time.Now().Unix() + 60 * 1000)
	case constants.Locked:
		_, _ = conn.Do("set", identifier+SplitChar+"LimitType", limitType)
	}

	_ = conn.Close()

}