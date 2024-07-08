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
// http client 的使用不需要全局初始化，因此为了方便使用，这里的 config 信息直接从 viper 中获取，避免使用者传入 common.cfg 信息
func NewHttpClient(t time.Duration, retryCount int) *HttpClient {
	h := &HttpClient{
		Client:     resty.New().SetTimeout(t),
		Timeout:    t,
		RetryCount: retryCount,
		isTraceLog: false,
		config: Config{
			TransparentParameter: viper.GetStringSlice("http.transparent_parameter"),
			IsTraceLog:           viper.GetBool("http.is_trace_log"),
		},
	}
	h.Client = h.Client.SetTimeout(h.Timeout)
	h.Client = h.Client.SetRetryCount(h.RetryCount)
	return h
}

// SetIsTraceLog
func (h *HttpClient) SetIsTraceLog(isTraceLog bool) *HttpClient {
	h.isTraceLog = isTraceLog
	return h
}

// isLogTraceInfo
func (h *HttpClient) isTraceInfo() bool {
	return h.config.IsTraceLog || h.isTraceLog
}

// setHeader
func (h *HttpClient) setHeader(c context.Context, headers map[string]string) *HttpClient {
	if len(h.config.TransparentParameter) > 0 {
		for _, parameter := range h.config.TransparentParameter {
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
func (h *HttpClient) Post(c context.Context, req PostRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().SetBody(req.Data).Post(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
	}
	return respData, nil
}

// Upload
func (h *HttpClient) Upload(c context.Context, req UploadRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().SetFileReader(req.Param, req.FileName, bytes.NewReader(req.FileBytes)).SetFormData(req.Data).Post(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
	}
	return respData, nil
}

// UploadWithFilePath
func (h *HttpClient) UploadWithFilePath(c context.Context, req UploadWithFilePathRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().SetFile(req.FileName, req.FilePath).SetFormData(req.Data).Post(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
	}
	return respData, nil
}

// Download
func (h *HttpClient) Download(c context.Context, req DownloadRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().SetOutput(req.FilePath).Get(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
	}
	return
}

// Get
func (h *HttpClient) Get(c context.Context, req GetRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().Get(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
	}
	return respData, nil
}

// GetWithQueryMap
func (h *HttpClient) GetWithQueryMap(c context.Context, req GetWithQueryMapRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().SetQueryParams(req.QueryMap).Get(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
	}
	return respData, nil
}

// GetWithQueryString
func (h *HttpClient) GetWithQueryString(c context.Context, req GetWithQueryStringRequest) (respData *HttpResponse, err error) {
	h.setHeader(c, req.Headers)
	resp, err := h.Client.R().SetQueryString(req.Query).Get(req.URL)
	if err != nil {
		logger.ErrorWithMsg(c, TAGNAME, "HttpClient: %s, req: %+v, err: %s", h, req, err)
		return
	}
	respData = &HttpResponse{
		Body:       resp.Body(),
		Header:     resp.Header(),
		StatusCode: resp.StatusCode(),
		TraceInfo:  resp.Request.TraceInfo(),
	}
	if h.isTraceInfo() {
		logger.Info(c, TAGNAME, "HttpClient: %s, req: %+v, respData: %+v", h, req, respData)
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
