package xstrings

import "regexp"

var (
	emailRegExp = regexp.MustCompile(`^([a-zA-Z0-9._-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+`)
	phoneRegExp = regexp.MustCompile(`^(\+86)?1\d{10}$`)
)

func IsValidEmail(email string) bool {
	return emailRegExp.MatchString(email)
}

func IsValidPhone(PhoneNum string) bool {
	return phoneRegExp.MatchString(PhoneNum)
}
