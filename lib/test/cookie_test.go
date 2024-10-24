package test

import (
	"net/http"
	"strings"
	"testing"
	"time"

	dhttp "dagger/lib/http"
)

// ResponseRecorder 记录 HTTP 响应
type ResponseRecorder struct {
	statusCode int
	headers    http.Header
	body       []byte
}

// NewResponseRecorder 创建一个新的 ResponseRecorder
func NewResponseRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		headers: http.Header{},
	}
}

// Header 返回 HTTP 头部
func (r *ResponseRecorder) Header() http.Header {
	return r.headers
}

// Write 将数据写入响应体
func (r *ResponseRecorder) Write(data []byte) (int, error) {
	r.body = append(r.body, data...)
	return len(data), nil
}

// WriteHeader 记录状态码
func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func TestSetCookie(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		maxAge   time.Duration
		path     string
		domain   string
		secure   bool
		httpOnly bool
	}{
		{"key1", "value1", 0 * time.Second, "/", "example.com", true, true},
		{"key2", "value2", 10 * time.Second, "/", "", true, false},
		{"key3", "value3", 20 * time.Second, "/test", "example.org", false, true},
		{"key4", "value4", 30 * time.Second, "/", "something.com", false, false},
	}
	for _, test := range tests {
		response := NewResponseRecorder()
		dhttp.SetCookie(response, test.name, test.value, test.maxAge, test.path, test.domain, test.secure, test.httpOnly)
		cookieStr := response.Header().Get("Set-Cookie")
		if !strings.Contains(cookieStr, test.name) {
			t.Errorf("SetCookie failed, name=%s, value=%s", test.name, test.value)
		}
	}
}
