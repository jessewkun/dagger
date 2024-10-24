package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

// RandomNum 返回指定范围的随机数
func RandomNum(min int, max int) int {
	rand.Seed(uint64(time.Now().UnixNano()))
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}

// RandomElement 返回 map 中的随机元素
func RandomElement(m map[string]interface{}) (string, interface{}) {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	randomKey := keys[RandomNum(0, len(keys)-1)]
	return randomKey, m[randomKey]
}

// 定义字符集
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString 返回指定长度的随机字符串
func RandomString(n int) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
