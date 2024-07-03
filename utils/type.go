package utils

import (
	"reflect"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// IsOnlyChinese 判断字符串是否只包含中文
func IsOnlyChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	return count == utf8.RuneCountInString(str)
}

// IsOnlyNumber 判断字符串是否只包含数字
func IsOnlyNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}
	return false
}

// IsZeroValue 判断是否是零值
func IsZeroValue(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}
