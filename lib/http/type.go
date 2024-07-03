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
	Timeout    time.Duration // timeout
	RetryCount int           // retry count
	debug      bool          // debug
}

func (h *HttpClient) String() string {
	return fmt.Sprintf("Client: %p, Timeout: %s, RetryCount: %d, debug: %v", h.Client, h.Timeout, h.RetryCount, h.debug)
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
	TransparentParameter []string `toml:"transparent_parameter"` // 透传参数，继承上下文中的参数
}
