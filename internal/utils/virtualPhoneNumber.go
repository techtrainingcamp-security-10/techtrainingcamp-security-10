package utils

import "regexp"

var VirtualPhoneNumberSegments = []string{
	`^162\d{8}$`,
	`^165\d{8}$`,
	`^167\d{8}$`,
	`^171\d{8}$`,
	`^170\d{8}$`,
}

func IsVirtualPhoneNumber(s string) bool {
	for _, v := range VirtualPhoneNumberSegments {
		re := regexp.MustCompile(v)
		if res := re.Match([]byte(s)); res {
			if res {
				return true
			}
		}
	}
	return false
}
