package http

import (
	"bytes"
	"context"
	"dagger/lib/logger"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

const TAGNAME = "DAGGER_HTTP"

// NewHttpClient create a new http client
func NewHttpClient(t time.Duration, retryCount int) *HttpClient {
	h := &HttpClient{
		Client:     resty.New().SetTimeout(t),
		Timeout:    t,
		RetryCount: retryCount,
	}
	return h.setRetryCount().setTimeOut()
}

// SetTimeOut
func (h *HttpClient) setTimeOut() *HttpClient {
	h.Client.SetTimeout(h.Timeout)
	return h
}

// setRetryCount
func (h *HttpClient) setRetryCount() *HttpClient {
	h.Client.SetRetryCount(h.RetryCount)
	return h
}

// SetIsTraceLog
func (h *HttpClient) SetIsTraceLog(isTraceLog bool) *HttpClient {
	h.isTraceLog = isTraceLog
	return h
}

// isLogTraceInfo
func (h *HttpClient) isTraceInfo() bool {
	return viper.GetBool("http.is_trace_log") || h.isTraceLog
}

// setHeader
func (h *HttpClient) setHeader(c context.Context, headers map[string]string) *HttpClient {
	if transparentParameter := viper.GetStringSlice("log.transparent_parameter"); len(transparentParameter) > 0 {
		for _, parameter := range transparentParameter {
			if value := c.Value(parameter); value != nil {
				h.Client.R().SetHeader(parameter, cast.ToString(value))
			}
		}
	}
	if len(headers) > 0 {
		for k, v := range headers {
			h.Client.R().SetHeader(k, v)
		}
	}
	return h
}

// Post
func (h *HttpClient) Post(c context.Context, rawURL string, data interface{}, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().SetBody(data).Post(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, data: %s, headers: %s, err: %s", h, rawURL, data, headers, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, data: %s, headers: %s, respData: %+v", h, rawURL, data, headers, respData)
	}
	return respData, nil
}

// Upload
// fileBytes 可能太大，不记录日志
func (h *HttpClient) Upload(c context.Context, rawURL string, fileBytes []byte, param, fileName string, data map[string]string, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().SetFileReader(param, fileName, bytes.NewReader(fileBytes)).SetFormData(data).Post(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, param: %s, fileName: %s, data: %s, headers: %s, err: %s", h, rawURL, param, fileName, data, headers, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, param: %s, fileName: %s, data: %s, headers: %s, respData: %+v", h, rawURL, param, fileName, data, headers, respData)
	}
	return respData, nil
}

// UploadWithFilePath
func (h *HttpClient) UploadWithFilePath(c context.Context, rawURL string, filePath string, fileName string, data map[string]string, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().SetFile(fileName, filePath).SetFormData(data).Post(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, filePath: %s, fileName: %s, data: %s, headers: %s, err: %s", h, rawURL, filePath, fileName, data, headers, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, filePath: %s, fileName: %s, data: %s, headers: %s, respData: %+v", h, rawURL, filePath, fileName, data, headers, respData)
	}
	return respData, nil
}

// Download
func (h *HttpClient) Download(c context.Context, rawURL string, filePath string, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().SetOutput(filePath).Get(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, filePath: %s, headers: %s, err: %s", h, rawURL, filePath, headers, err)
		return
	}
	respData = &HttpResponse{
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, filePath: %s, headers: %s, respData: %+v", h, rawURL, filePath, headers, respData)
	}
	return
}

// Get
func (h *HttpClient) Get(c context.Context, rawURL string, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().Get(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, headers: %s, err: %s", h, rawURL, headers, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, headers: %s, respData: %+v", h, rawURL, headers, respData)
	}
	return respData, nil
}

// GetWithQueryMap
func (h *HttpClient) GetWithQueryMap(c context.Context, rawURL string, query map[string]string, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().SetQueryParams(query).Get(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, query: %s, headers: %s, err: %s", h, rawURL, query, headers, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, query: %s, headers: %s, respData: %+v", h, rawURL, query, headers, respData)
	}
	return respData, nil
}

// GetWithQueryString
func (h *HttpClient) GetWithQueryString(c context.Context, rawURL string, query string, headers map[string]string) (respData *HttpResponse, err error) {
	h.setHeader(c, headers)
	resp, err := h.Client.R().SetQueryString(query).Get(rawURL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, rawURL: %s, query: %s, headers: %s, err: %s", h, rawURL, query, headers, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, rawURL: %s, query: %s, headers: %s, respData: %+v", h, rawURL, query, headers, respData)
	}
	return respData, nil
}

// HttpBuildQuery http build query
func HttpBuildQuery(data map[string]interface{}) string {
	var uri url.URL

	q := uri.Query()
	for k, v := range data {
		q.Add(k, cast.ToString(v))
	}
	return q.Encode()
}
