package usermangment

import (
	"errors"
	"regexp"
)

func ValidUserName(name string) error {
	for _, char := range name {
		if !(char < 127 && char > 32) {
			return errors.New("username contains invalid characters")
		}
	}
	return nil
}

func ValidEmail(email string) error {
	valid := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if valid.MatchString(email) {
		return nil
	}
	return errors.New("invalid email")
}

func ValidPassword(pass, confirm string) error {
	if pass != confirm {
		return errors.New("incorrect confirm password")
	}
	valid := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	if valid.MatchString(pass) {
		return nil
	}
	return errors.New("invalid password")
}
