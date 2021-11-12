package utils

import (
	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/route/service"
)

func CheckFailRecords(s service.Service, identifier string) int {
	if limitType := s.GetUserLimitType(identifier); limitType != 0 {
		return limitType
	}

	recordsIn5s, recordsIn1Min := s.GetApiFailRecords(identifier)
	defer s.SetApiFailRecords(identifier, recordsIn5s+1, recordsIn1Min+1)

	// 五秒内
	if recordsIn5s >= 5 {
		s.SetUserLimitType(identifier, constants.FrequentLimit)
		return constants.FrequentLimit
	}

	// 一分钟内
	if recordsIn1Min <= 3 {
		return constants.SlideBar
	} else if recordsIn1Min <= 10 {
		s.SetUserLimitType(identifier, constants.FrequentLimit)
		return constants.FrequentLimit
	} else {
		s.SetUserLimitType(identifier, constants.Locked)
		return constants.Locked
	}
}

func ClearFailRecords(s service.Service, identifier string) {
	s.SetApiFailRecords(identifier, 0, 0)
	s.SetUserLimitType(identifier, 0)
}
