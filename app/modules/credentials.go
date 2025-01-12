package modules

import (
	"regexp"
)

func ValidUserName(name string) bool {
	for _, char := range name {
		if !(char < 127 && char > 32) {
			return false
		}
	}
	return true
}

func ValidEmail(email string) bool {
	valid := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return valid.MatchString(email)
}

func ValidPassword(password string) bool {
	return true
}
