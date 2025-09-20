package ut

import (
	"strings"
)

func IsMobile(s string) bool {
	return len(s) == 11
}

func IsURL(s string) bool {
	s = strings.TrimSpace(s)
	return strings.HasPrefix(s, "http://") || strings.HasPrefix(s, "https://")
}
