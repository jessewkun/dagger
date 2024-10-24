package utils

import (
	"reflect"
	"regexp"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// IsOnlyChinese 判断字符串是否只包含中文
func IsOnlyChinese(str string) bool {
	if len(str) < 1 {
		return false
	}
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
	if x == nil {
		return true
	}
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

// IsChinesePhoneNumber 判断字符串是否是中国手机号码
func IsChinesePhoneNumber(phone string) bool {
	// 定义中国手机号码的正则表达式
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
