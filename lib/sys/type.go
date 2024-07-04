package sys

import "encoding/json"

// ApiResult 接口返回数据结构
type ApiResult struct {
	Code    int         `json:"code"`     // 接口错误码，0表示成功，非0表示失败
	Msg     string      `json:"msg"`      // 错误信息
	Data    interface{} `json:"data"`     // 返回数据
	TraceId string      `json:"trace_id"` // 请求唯一标识
}

func (r *ApiResult) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

// DaggerError 自定义错误
type DaggerError struct {
	code int
	err  error
}

// Error 实现error接口
func (e DaggerError) Error() string {
	return e.err.Error()
}

// Unwrap 实现errors.Unwrap接口
func (e DaggerError) Unwrap() error {
	return e.err
}
