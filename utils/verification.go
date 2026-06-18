package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateVerificationCode generates a 6-digit verification code
func GenerateVerificationCode() string {
	const digits = "0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		code += string(digits[num.Int64()])
	}
	return code
}
