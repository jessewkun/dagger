package test

import (
	"context"
	"testing"
	"time"

	dhttp "dagger/lib/http"
)

func TestHttpClient_Post(t *testing.T) {
	type fields struct {
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c   context.Context
		req dhttp.PostRequest
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		{"Test200", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.PostRequest{URL: "https://www.baidu.com", Data: nil, Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.PostRequest{URL: "https://www.baidu.com", Data: nil, Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.NewHttpClient(tt.fields.Timeout, tt.fields.RetryCount)
			h.SetIsTraceLog(true)
			_, err := h.Post(tt.args.c, tt.args.req)
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
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c   context.Context
		req dhttp.GetRequest
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		{"Test200", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.GetRequest{URL: "https://www.baidu.com", Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.GetRequest{URL: "https://www.baidu.com", Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.NewHttpClient(tt.fields.Timeout, tt.fields.RetryCount)
			_, err := h.Get(tt.args.c, tt.args.req)
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
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c   context.Context
		req dhttp.GetWithQueryMapRequest
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		{"Test200", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.GetWithQueryMapRequest{URL: "https://www.baidu.com", QueryMap: map[string]string{"a": "1"}, Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.GetWithQueryMapRequest{URL: "https://www.baidu.com", QueryMap: map[string]string{"a": "1"}, Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.NewHttpClient(tt.fields.Timeout, tt.fields.RetryCount)
			_, err := h.GetWithQueryMap(tt.args.c, tt.args.req)
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
		Timeout    time.Duration
		RetryCount int
	}
	type args struct {
		c   context.Context
		req dhttp.GetWithQueryStringRequest
	}
	ctx := context.Background()
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantRespData dhttp.HttpResponse
		wantStaus    int
	}{
		{"Test200", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.GetWithQueryStringRequest{URL: "https://www.baidu.com", Query: "a=1", Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 200}, 200},
		{"Test404", fields{Timeout: time.Duration(10 * time.Second), RetryCount: 1}, args{ctx, dhttp.GetWithQueryStringRequest{URL: "https://www.baidu.com", Query: "a=1", Headers: map[string]string{"trace_id": "1"}}}, dhttp.HttpResponse{StatusCode: 404}, 404},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := dhttp.NewHttpClient(tt.fields.Timeout, tt.fields.RetryCount)
			_, err := h.GetWithQueryString(tt.args.c, tt.args.req)
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
