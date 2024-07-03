package sys

import (
	"encoding/json"
	"errors"
)

type ApiResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *ApiResult) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

// NewApiResult create a new api result
func NewApiResult(code int, msg string, data interface{}) *ApiResult {
	return &ApiResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// SuccResp success response
func SuccResp(data interface{}) *ApiResult {
	return NewApiResult(0, "succ", data)
}

// ErrorResp error response
func ErrorResp(err error) *ApiResult {
	if !errors.As(err, &DaggerError{}) {
		err = newDefaultError(err)
	}
	e := err.(DaggerError)
	return NewApiResult(e.code, e.Error(), struct{}{})
}

// SystemErrorResp system error response
func SystemErrorResp() *ApiResult {
	return ErrorResp(SystemError)
}

// ParamErrorResp param error response
func ParamErrorResp() *ApiResult {
	return ErrorResp(ParamError)
}

// ForbiddenErrorResp forbidden error response
func ForbiddenErrorResp() *ApiResult {
	return ErrorResp(ForbiddenError)
}

// NotfoundErrorResp not found error response
func NotfoundErrorResp() *ApiResult {
	return ErrorResp(NotfoundError)
}
