package util

import (
	"bytes"
	"regexp"
)

func ValidatePhoneNumber(phoneNumber string) bool {
	r := regexp.MustCompile(`^\+998[0-9]{9}$`)
	return r.MatchString(phoneNumber)
}

func ValidateEmail(email string) bool {
	r := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return r.MatchString(email)
}

func GenerateString(keys ...string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(keys); i++ {
		buffer.WriteString(keys[i])
	}
	return buffer.String()
}
