package sys

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// NewApiResult create a new api result
func NewApiResult(c *gin.Context, code int, msg string, data interface{}) *ApiResult {
	resp := &ApiResult{
		Code:    code,
		Msg:     msg,
		Data:    data,
		TraceId: c.GetString("trace_id"),
	}
	c.Set(CTX_DAGGER_OUTPUT, resp)
	return resp
}

// SuccResp success response
func SuccResp(c *gin.Context, data interface{}) *ApiResult {
	if data == nil {
		data = struct{}{}
	}
	return NewApiResult(c, 0, "succ", data)
}

// ErrorResp error response
func ErrorResp(c *gin.Context, err error) *ApiResult {
	if !errors.As(err, &DaggerError{}) {
		err = newDefaultError(err)
	}
	e := err.(DaggerError)
	return NewApiResult(c, e.code, e.Error(), struct{}{})
}

// SystemErrorResp system error response
func SystemErrorResp(c *gin.Context) *ApiResult {
	return ErrorResp(c, SystemError)
}

// ParamErrorResp param error response
func ParamErrorResp(c *gin.Context) *ApiResult {
	return ErrorResp(c, ParamError)
}

// ForbiddenErrorResp forbidden error response
func ForbiddenErrorResp(c *gin.Context) *ApiResult {
	return ErrorResp(c, ForbiddenError)
}

// NotfoundErrorResp not found error response
func NotfoundErrorResp(c *gin.Context) *ApiResult {
	return ErrorResp(c, NotfoundError)
}

// RateLimiterErrorResp rate limiter error response
func RateLimiterErrorResp(c *gin.Context) *ApiResult {
	return ErrorResp(c, RateLimiterError)
}
