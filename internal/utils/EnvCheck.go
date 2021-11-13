package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/phuslu/iploc"
	"math"
	"net"

	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/service"
)

type EnvironmentType struct {
	IP       string `json:"IP" binding:"required"`
	DeviceID string `json:"DeviceID" binding:"required"`
}

type RequestType struct {
	Environment EnvironmentType `json:"Environment" binding:"required"`
}

func EnvCheck(s service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch envCheck(c, s) {
		case constants.Normal:
			c.Next()
		case constants.NoEnvCode:
			c.JSON(500, gin.H{
				"Code":    constants.FailedCode,
				"Message": constants.NoEnv,
			})
			c.Abort()
		case constants.NotChinaIPCode:
			c.JSON(500, gin.H{
				"Code":    constants.FailedCode,
				"Message": constants.NotChinaIP,
			})
			c.Abort()
		case constants.FrequentLimit:
			c.JSON(500, gin.H{
				"Code":    constants.FailedCode,
				"Message": constants.FrequencyLimit,
			})
			c.Abort()
		case constants.Locked:
			c.JSON(500, gin.H{
				"Code":    constants.FailedCode,
				"Message": constants.Lock,
			})
			c.Abort()
		}
	}
}

// envCheck
// @Description ip、环境风险检测
// @Return
// 0: 正常
// -1：缺少环境信息
// 1: 非国内IP
// 2: 同环境 1 分钟内请求 15 次
// 3: 同环境 1 分钟内请求 30 次
func envCheck(c *gin.Context, s service.Service) int {
	var form RequestType
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err == nil {
		var loc = iploc.Country(net.ParseIP(form.Environment.IP))
		switch {
		case string(loc) != "CN": // 非国内IP
			return constants.NotChinaIPCode
		case s.GetUserLimitType(form.Environment.IP) == constants.FrequentLimit:
			return constants.FrequentLimit
		case s.GetUserLimitType(form.Environment.IP) == constants.Locked:
			return constants.Locked
		default:
			cntIP := s.GetIRequests(form.Environment.IP)
			cntID := s.GetIRequests(form.Environment.DeviceID)
			cntTotal := int(math.Max(float64(cntIP), float64(cntID)))

			defer s.SetIRequests(form.Environment.IP, cntTotal+1)
			// 同环境 1 分钟内请求 15 次 拦截
			switch {
			case cntTotal > 30:
				s.SetUserLimitType(form.Environment.IP, constants.Locked)
				return constants.Locked
			case cntTotal > 15:
				s.SetUserLimitType(form.Environment.IP, constants.FrequentLimit)
				return constants.FrequentLimit
			default:
				return constants.Normal
			}
		}
	} else {
		return constants.NoEnvCode
	}
}
