package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// HttpClient
type HttpClient struct {
	Client     *resty.Client // resty client
	Timeout    time.Duration `toml:"timeout"`      // 超时时间
	RetryCount int           `toml:"retry_count"`  // 重试次数
	isTraceLog bool          `toml:"is_trace_log"` // 是否记录trace log
}

func (h *HttpClient) String() string {
	return fmt.Sprintf("Client: %p, Timeout: %s, RetryCount: %d, isTraceLog: %v", h.Client, h.Timeout, h.RetryCount, h.isTraceLog)
}

// HttpResponse
type HttpResponse struct {
	Body       []byte          // http response body
	Header     http.Header     // http response header
	StatusCode int             // http response status code
	TraceInfo  resty.TraceInfo // http response trace info
}

func (h *HttpResponse) String() string {
	return fmt.Sprintf("Body: %s, Header: %v, StatusCode: %d, TraceInfo: %+v", h.Body, h.Header, h.StatusCode, h.TraceInfo)
}

type Config struct {
	TransparentParameter []string `toml:"transparent_parameter" mapstructure:"transparent_parameter"` // 透传参数，继承上下文中的参数
	isTraceLog           bool     `toml:"is_trace_log" mapstructure:"is_trace_log"`                   // 是否记录trace log
}
