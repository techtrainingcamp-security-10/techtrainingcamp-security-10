package api

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/gomodule/redigo/redis"
//	"github.com/jinzhu/gorm"
//)
//
//const (
//	VerifyCodeExpireTime = 3 * 60
//	SessionIdExpireTime  = 3 * 60 * 60
//	SplitChar            = "|"
//)
//
//// InsertVerifyCode 插入验证码
//func InsertVerifyCode(cache *redis.Pool, phoneNumber string, verifyCode string) bool {
//	_, err := cache.Get().Do("setex", phoneNumber+SplitChar+"VerifyCode", VerifyCodeExpireTime, verifyCode)
//	//fmt.Println(result)
//	if err != nil {
//		return false
//	}
//	return true
//}
//
//// GetVerifyCode 获取验证码
//func GetVerifyCode(cache *redis.Pool, phoneNumber string) string {
//	data, err := redis.String(cache.Get().Do("get", phoneNumber+SplitChar+"VerifyCode"))
//	//fmt.Println(data)
//	if err != nil {
//		return "nil"
//	}
//	return data
//}
//
//// InsertSessionId 插入SessionId
//func InsertSessionId(cache *redis.Pool, phoneNumber string, sessionID string) bool {
//	_, err := cache.Get().Do("setex", sessionID+SplitChar+"SessionId", SessionIdExpireTime, phoneNumber)
//	//fmt.Println(result)
//	if err != nil {
//		return false
//	}
//	return true
//}
//
//// DeleteSessionId 删除SessionId
//func DeleteSessionId(cache *redis.Pool, sessionID string) (int, string) {
//	phoneNumber, err1 := redis.String(cache.Get().Do("get", sessionID+SplitChar+"SessionId"))
//	if err1 != nil {
//		fmt.Println(err1)
//		return -1, ""
//	}
//	data, err2 := redis.Int(cache.Get().Do("del", sessionID+SplitChar+"SessionId"))
//	//fmt.Println(result)
//	if err2 != nil {
//		fmt.Println(err2)
//		return -2, ""
//	}
//	return data, phoneNumber
//}
//
//// TestRedis
//// @Description 测试redis
//// @Router /api/test_redis [get]
//func TestRedis(cache *redis.Pool, dbRead *gorm.DB) gin.HandlerFunc {
//	return func(context *gin.Context) {
//		if true {
//			context.JSON(200, gin.H{
//				// "checkRev": form,
//				"Code":    0,
//				"Message": "请求成功",
//				"Data": gin.H{
//					"VerifyCode":   GetVerifyCode(cache, "123456"),
//					"ExpireTime":   180,
//					"DecisionType": 0,
//				},
//			})
//		} else {
//			context.JSON(200, gin.H{
//				// "checkRev": form,
//				"Code":    1,
//				"Message": "请求失败",
//				"Data": gin.H{
//					"VerifyCode":   1234,
//					"ExpireTime":   180,
//					"DecisionType": 0,
//				},
//			})
//		}
//	}
//}
