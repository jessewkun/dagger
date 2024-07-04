package sys

import (
	"errors"
)

// 默认业务错误码
const DEFAULT_ERROR_CODE = 10000

// 系统错误
var SystemError = DaggerError{1000, errors.New("系统错误，请稍后重试")}

// 参数错误
var ParamError = DaggerError{1001, errors.New("参数错误")}

// 权限错误
var ForbiddenError = DaggerError{1002, errors.New("Permission denied")}

// 未找到
var NotfoundError = DaggerError{1003, errors.New("Not found")}

// NewDaggerError 创建自定义错误
// 业务自定义错误码必须大于10000，小于10000的错误码为系统错误码，10000为默认业务错误码
func NewDaggerError(code int, err error) DaggerError {
	if code < DEFAULT_ERROR_CODE+1 {
		panic("error code must greater than 10000")
	}
	return DaggerError{code, err}
}

// newDefaultError 创建默认错误
func newDefaultError(err error) DaggerError {
	return DaggerError{DEFAULT_ERROR_CODE, err}
}
