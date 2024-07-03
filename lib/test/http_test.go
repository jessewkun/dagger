package test

import (
	"context"
	"testing"
	"time"

	dhttp "dagger/lib/http"

	"github.com/go-resty/resty/v2"
)

func TestHttpClient_Post(t *testing.T) {
	type fields struct {
		Client     *resty.Client
		Timeout    time.Duration
		RetryCount int
		debug      bool
	}
	type args struct {
		c       context.Context
		rawURL  string
		data    interface{}
		headers map[string]string
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		// TODO: Add test cases.
		{"Test200", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", nil, map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", nil, map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &dhttp.HttpClient{
				Client:     tt.fields.Client,
				Timeout:    tt.fields.Timeout,
				RetryCount: tt.fields.RetryCount,
			}
			h.SetDebug(true)
			_, err := h.Post(tt.args.c, tt.args.rawURL, tt.args.data, tt.args.headers)
			if err != nil {
				t.Errorf("HttpClient.Post() error = %v", err)
				return
			}
			if tt.wantRespData.StatusCode != tt.wantStaus {
				t.Errorf("HttpClient.Get() StatusCode = %v, wantStaus = %v", tt.wantRespData.StatusCode, tt.wantStaus)
			}
		})
	}
}

func TestHttpClient_Get(t *testing.T) {
	type fields struct {
		Client     *resty.Client
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c       context.Context
		rawURL  string
		headers map[string]string
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		// TODO: Add test cases.
		{"Test200", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.HttpClient{
				Client:     tt.fields.Client,
				Timeout:    tt.fields.Timeout,
				RetryCount: tt.fields.RetryCount,
			}
			_, err := h.Get(tt.args.c, tt.args.rawURL, tt.args.headers)
			if err != nil {
				t.Errorf("HttpClient.Get() error = %v", err)
				return
			}
			if tt.wantRespData.StatusCode != tt.wantStaus {
				t.Errorf("HttpClient.Get() StatusCode = %v, wantStaus = %v", tt.wantRespData.StatusCode, tt.wantStaus)
			}
		})
	}
}

func TestHttpClient_GetWithQueryMap(t *testing.T) {
	type fields struct {
		Client     *resty.Client
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c       context.Context
		rawURL  string
		query   map[string]string
		headers map[string]string
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		// TODO: Add test cases.
		{"Test200", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", map[string]string{"a": "1"}, map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", map[string]string{"a": "1"}, map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.HttpClient{
				Client:     tt.fields.Client,
				Timeout:    tt.fields.Timeout,
				RetryCount: tt.fields.RetryCount,
			}
			_, err := h.GetWithQueryMap(tt.args.c, tt.args.rawURL, tt.args.query, tt.args.headers)
			if err != nil {
				t.Errorf("HttpClient.GetWithQueryMap() error = %v", err)
				return
			}
			if tt.wantRespData.StatusCode != tt.wantStaus {
				t.Errorf("HttpClient.GetWithQueryMap() StatusCode = %v, wantStaus = %v", tt.wantRespData.StatusCode, tt.wantStaus)
			}
		})
	}
}

func TestHttpClient_GetWithQueryString(t *testing.T) {
	type fields struct {
		Client     *resty.Client
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c       context.Context
		rawURL  string
		query   string
		headers map[string]string
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		// TODO: Add test cases.
		{"Test200", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", "a=1", map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Client: resty.New(), Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, "https://www.baidu.com", "a=1", map[string]string{"trace_id": "1"}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.HttpClient{
				Client:     tt.fields.Client,
				Timeout:    tt.fields.Timeout,
				RetryCount: tt.fields.RetryCount,
			}
			_, err := h.GetWithQueryString(tt.args.c, tt.args.rawURL, tt.args.query, tt.args.headers)
			if err != nil {
				t.Errorf("HttpClient.GetWithQueryString() error = %v", err)
				return
			}
			if tt.wantRespData.StatusCode != tt.wantStaus {
				t.Errorf("HttpClient.GetWithQueryString() StatusCode = %v, wantStaus = %v", tt.wantRespData.StatusCode, tt.wantStaus)
			}
		})
	}
}
