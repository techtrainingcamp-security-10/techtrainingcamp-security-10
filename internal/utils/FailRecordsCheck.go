package utils

import (
	"techtrainingcamp-security-10/internal/constants"
	"techtrainingcamp-security-10/internal/route/service"
	"time"
)

func CheckFailRecords(s service.Service, apiRoute string, apiMethod string, identifier string) int {
	if limitType := s.GetUserLimitType(identifier); limitType != 0 {
		return limitType
	}

	records := s.GetApiFailRecords(apiRoute, apiMethod, identifier)
	nowTime := time.Now().Unix()
	records = append(records, nowTime)

	left := 0
	if len(records) >= 20 {
		left = len(records) - 20
	}

	defer s.SetApiFailRecords(apiRoute, apiMethod, identifier, records[left:])

	// 五秒内
	recordsIn5s := 0
	for _, record := range records {
		if record >= nowTime - 5 * 1000 {
			recordsIn5s ++
		}
	}

	if recordsIn5s >= 10 {
		s.SetUserLimitType(identifier, constants.Locked)
		return constants.Locked
	} else if recordsIn5s >= 3 {
		s.SetUserLimitType(identifier, constants.FrequentLimit)
		return constants.FrequentLimit
	}

	// 一分钟内
	recordsIn1Min := 0
	for _, record := range records {
		if record >= nowTime - 3 * 60 * 1000 {
			recordsIn1Min ++
		}
	}

	if recordsIn1Min > 3 {
		return constants.SlideBar
	}

	return constants.Normal
}

func ClearFailRecords(s service.Service, apiRoute string, apiMethod string, identifier string) {
	s.SetApiFailRecords(apiRoute, apiMethod, identifier, make([]int64, 0))
}