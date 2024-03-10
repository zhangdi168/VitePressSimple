// Package pkg
// @Author: zhangdi
// @File: random
// @Version: 1.0.0
// @Date: 2023/9/11 17:12
package random

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(minLength, maxLength int) string {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxLength-minLength+1) + minLength

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
