package xstrings

import (
	"crypto/rand"
)

var (
	numbers        = "0123456789"
	numberAndAlpha = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func Rand(n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = numberAndAlpha[b%byte(len(numberAndAlpha))]
	}
	return string(bytes), nil
}

//RandNNumber 随机生成N个整数
func RandNNumber(n int) (string, error) {
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = numbers[b%byte(len(numbers))]
	}
	return string(bytes), nil
}

//GenerateRandomBytes of n size
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
