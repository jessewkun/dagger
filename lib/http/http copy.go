package http

// import (
// 	"context"
// 	"dagger/lib/logger"
// 	"dagger/utils"
// 	"io"
// 	"net"
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"time"

// 	"github.com/go-resty/resty/v2"
// 	"github.com/spf13/cast"
// 	"github.com/spf13/viper"
// 	"github.com/valyala/fasthttp"
// )

// const TAGNAME = "DAGGER_HTTP"

// type HttpClient struct {
// 	Client  *resty.Client
// 	Timeout time.Duration
// }

// type HttpClientOptions struct {
// 	ReadTimeout                   time.Duration
// 	WriteTimeout                  time.Duration
// 	MaximumConnsPerHost           int
// 	NoDefaultUserAgentHeader      bool
// 	DisableHeaderNamesNormalizing bool
// 	DisablePathNormalizing        bool
// 	Dial                          func(addr string) (net.Conn, error)
// }

// // NewHttpClient create a new http client
// func NewHttpClient(t time.Duration) *HttpClient {
// 	return &HttpClient{
// 		Client:  resty.New(),
// 		Timeout: t,
// 	}
// }

// // WithOptions set http client options
// func (ho HttpClientOptions) WithOptions(key string, value interface{}) HttpClientOptions {
// 	switch key {
// 	case "ReadTimeout":
// 		ho.ReadTimeout = value.(time.Duration)
// 	case "WriteTimeout":
// 		ho.WriteTimeout = value.(time.Duration)
// 	case "MaximumConnsPerHost":
// 		ho.MaximumConnsPerHost = value.(int)
// 	case "NoDefaultUserAgentHeader":
// 		ho.NoDefaultUserAgentHeader = value.(bool)
// 	case "DisableHeaderNamesNormalizing":
// 		ho.DisableHeaderNamesNormalizing = value.(bool)
// 	case "DisablePathNormalizing":
// 		ho.DisablePathNormalizing = value.(bool)
// 	case "Dial":
// 		ho.Dial = value.(func(addr string) (net.Conn, error))
// 	}
// 	return ho
// }

// // defaultOptions default http client options
// func defaultOptions() HttpClientOptions {
// 	ho := HttpClientOptions{}
// 	ho.ReadTimeout = time.Duration(10 * time.Second)
// 	ho.WriteTimeout = time.Duration(10 * time.Second)
// 	ho.MaximumConnsPerHost = 10
// 	ho.NoDefaultUserAgentHeader = true
// 	ho.DisableHeaderNamesNormalizing = true
// 	ho.DisablePathNormalizing = true
// 	ho.Dial = func(addr string) (net.Conn, error) {
// 		return fasthttp.DialTimeout(addr, time.Duration(60)*time.Second)
// 	}
// 	return ho
// }

// // Post http request
// func (h HttpClient) Post(c context.Context, rawURL string, data io.Reader, headers map[string]string) (respData []byte, respHeader map[string]string, err error) {
// 	return h.Do(c, http.MethodPost, rawURL, data, headers)
// }

// // Get http request
// func (h HttpClient) Get(c context.Context, rawURL string, headers map[string]string) (respData []byte, respHeader map[string]string, err error) {
// 	return h.Do(c, http.MethodGet, rawURL, nil, headers)
// }

// // Set http client options
// func (h HttpClient) setHttpClientWithOption(option HttpClientOptions) {
// 	if utils.IsZeroValue(option) {
// 		option = defaultOptions()
// 	}
// 	h.Client.ReadTimeout = option.ReadTimeout
// 	h.Client.WriteTimeout = option.WriteTimeout
// 	h.Client.MaxConnsPerHost = option.MaximumConnsPerHost
// 	h.Client.NoDefaultUserAgentHeader = option.NoDefaultUserAgentHeader
// 	h.Client.DisableHeaderNamesNormalizing = option.DisableHeaderNamesNormalizing
// 	h.Client.DisablePathNormalizing = option.DisablePathNormalizing
// 	h.Client.Dial = option.Dial
// }

// // Do http request
// func (h HttpClient) Do(c context.Context, method, rawURL string, data io.Reader, headers map[string]string) (respData []byte, respHeader map[string]string, err error) {
// 	respHeader = make(map[string]string)
// 	if _, err := url.ParseRequestURI(rawURL); err != nil {
// 		return respData, respHeader, err
// 	}

// 	// h.setHttpClientWithOption(option)

// 	req := fasthttp.AcquireRequest()
// 	resp := fasthttp.AcquireResponse()
// 	defer func() {
// 		fasthttp.ReleaseResponse(resp)
// 		fasthttp.ReleaseRequest(req)
// 	}()

// 	method = strings.ToUpper(method)
// 	req.SetTimeout(h.Timeout)
// 	req.Header.SetMethod(method)
// 	req.SetRequestURI(rawURL)

// 	if transparentParameter := viper.GetStringSlice("log.transparent_parameter"); len(transparentParameter) > 0 {
// 		for _, parameter := range transparentParameter {
// 			if value := c.Value(parameter); value != nil {
// 				req.Header.Set(parameter, cast.ToString(value))
// 			}
// 		}
// 	}
// 	if len(headers) > 0 {
// 		for k, v := range headers {
// 			req.Header.Set(k, v)
// 		}
// 	}
// 	if method == http.MethodPost && data != nil {
// 		bytes, err := io.ReadAll(data)
// 		if err != nil {
// 			logger.ErrorWithMsg(c, TAGNAME, "HttpClient %s, method: %s, rawURL: %s, data: %s, headers: %s, option: %s, err: %s", h, method, rawURL, data, headers, option, err)
// 			return respData, respHeader, err
// 		}
// 		req.SetBody(bytes)
// 	}

// 	err = h.Client.Do(req, resp)
// 	respHeader["StatusCode"] = cast.ToString(resp.StatusCode())
// 	resp.Header.VisitAll(func(key, value []byte) {
// 		respHeader[string(key)] = string(value)
// 	})
// 	if err != nil {
// 		logger.ErrorWithMsg(c, TAGNAME, "HttpClient %s, method: %s, rawURL: %s, data: %s, headers: %s, option: %s, respHeader: %s, err: %s", h, method, rawURL, data, headers, option, respHeader, err)
// 		return respData, respHeader, err
// 	}

// 	if resp.Header.ContentLength() > 0 {
// 		respData = make([]byte, resp.Header.ContentLength())
// 		copy(respData, resp.Body())
// 	}
// 	if viper.GetBool("http.debug") {
// 		logger.Info(c, TAGNAME, "HttpClient %s, method: %s, rawURL: %s, data: %s, headers: %s, option: %s, respBody: %s, respHeader: %s", h, method, rawURL, data, headers, option, respData, respHeader)
// 	}

// 	return respData, respHeader, err
// }

// // HttpBuildQuery http build query
// func HttpBuildQuery(data map[string]interface{}) string {
// 	var uri url.URL

// 	q := uri.Query()
// 	for k, v := range data {
// 		q.Add(k, cast.ToString(v))
// 	}
// 	return q.Encode()
// }
