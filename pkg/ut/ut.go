package ut

import (
	"strings"
)

func IsMobile(s string) bool {
	if len(s) != 11 {
		return false
	}

	return true
}

func IsURL(s string) bool {
	s = strings.TrimSpace(s)
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
