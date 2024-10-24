package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// HttpClient
type HttpClient struct {
	Client     *resty.Client  // resty client
	r          *resty.Request // resty request
	Timeout    time.Duration  // 超时时间
	RetryCount int            // 重试次数
	isTraceLog bool           // 是否记录trace log, 代码中手段控制单个请求是否记录trace log
	config     Config         // http config
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

// http config
type Config struct {
	TransparentParameter []string `toml:"transparent_parameter" mapstructure:"transparent_parameter"` // 透传参数，继承上下文中的参数
	IsTraceLog           bool     `toml:"is_trace_log" mapstructure:"is_trace_log"`                   // 是否记录全部请求的 trace log，配置文件控制，支持手动修改配置文件，自动生效
}

// post request
type PostRequest struct {
	URL     string            // 请求地址
	Data    interface{}       // 请求数据
	Headers map[string]string // 请求头
}

// upload request
type UploadRequest struct {
	URL       string            // 请求地址
	FileBytes []byte            // 文件字节
	Param     string            // 文件参数名
	FileName  string            // 文件名
	Data      map[string]string // 请求数据
	Headers   map[string]string // 请求头
}

// upload with file path request
type UploadWithFilePathRequest struct {
	URL      string            // 请求地址
	FileName string            // 文件名
	FilePath string            // 文件路径
	Param    string            // 文件参数名
	Data     map[string]string // 请求数据
	Headers  map[string]string // 请求头
}

// download request
type DownloadRequest struct {
	URL      string            // 请求地址
	FilePath string            // 文件路径
	Headers  map[string]string // 请求头
}

// get request
type GetRequest struct {
	URL     string            // 请求地址
	Headers map[string]string // 请求头
}

// get with query map request
type GetWithQueryMapRequest struct {
	URL      string            // 请求地址
	QueryMap map[string]string // 请求参数
	Headers  map[string]string // 请求头
}

// get with query request
type GetWithQueryStringRequest struct {
	URL     string            // 请求地址
	Query   string            // 请求参数
	Headers map[string]string // 请求头
}
