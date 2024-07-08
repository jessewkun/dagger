package test

// import (
// 	"context"
// 	dhttp "dagger/lib/http"
// 	"io"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/valyala/fasthttp"
// )

// func TestHttpClient_Get(t *testing.T) {
// 	type fields struct {
// 		client  *fasthttp.Client
// 		Timeout time.Duration
// 	}
// 	type args struct {
// 		c       context.Context
// 		rawURL  string
// 		headers map[string]string
// 		option  dhttp.HttpClientOptions
// 	}
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, "trace_id", uuid.New().String())
// 	tests := []struct {
// 		name           string
// 		fields         fields
// 		args           args
// 		wantRespData   []byte
// 		wantrespHeader map[string]string
// 		wantStaus      string
// 	}{
// 		{"Test200", fields{client: &fasthttp.Client{}, Timeout: time.Duration(10 * time.Second)}, args{ctx, "https://www.baidu.com", map[string]string{"trace_id": "1"}, dhttp.HttpClientOptions{}}, nil, nil, "200"},
// 		{"Test404", fields{client: &fasthttp.Client{}, Timeout: time.Duration(10 * time.Second)}, args{ctx, "https://www.baidu.com", map[string]string{"trace_id": "1"}, dhttp.HttpClientOptions{}}, nil, nil, "404"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := dhttp.HttpClient{
// 				Client:  tt.fields.client,
// 				Timeout: tt.fields.Timeout,
// 			}
// 			_, gotRespHeader, err := h.Get(tt.args.c, tt.args.rawURL, tt.args.headers, tt.args.option)
// 			if err != nil {
// 				t.Errorf("HttpClient.Get() error = %v", err)
// 				return
// 			}
// 			if gotRespHeader["StatusCode"] != tt.wantStaus {
// 				t.Errorf("HttpClient.Get() StatusCode = %v, wantStaus = %v", gotRespHeader["StatusCode"], tt.wantStaus)
// 				return
// 			}
// 		})
// 	}
// }

// func TestHttpClient_Post(t *testing.T) {
// 	type fields struct {
// 		Client  *fasthttp.Client
// 		Timeout time.Duration
// 	}
// 	type args struct {
// 		c       context.Context
// 		rawURL  string
// 		data    io.Reader
// 		headers map[string]string
// 		option  dhttp.HttpClientOptions
// 	}
// 	ctx := context.Background()
// 	ctx = context.WithValue(ctx, "trace_id", uuid.New().String())
// 	tests := []struct {
// 		name           string
// 		fields         fields
// 		args           args
// 		wantRespData   []byte
// 		wantRespHeader map[string]string
// 		wantStaus      string
// 	}{
// 		{"Test200", fields{Client: &fasthttp.Client{}, Timeout: time.Duration(10 * time.Second)}, args{ctx, "https://www.baidu.com", nil, map[string]string{"trace_id": "1"}, dhttp.HttpClientOptions{}}, nil, nil, "200"},
// 		{"Test404", fields{Client: &fasthttp.Client{}, Timeout: time.Duration(10 * time.Second)}, args{ctx, "https://www.baidu.com", nil, map[string]string{"trace_id": "1"}, dhttp.HttpClientOptions{}}, nil, nil, "404"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := dhttp.HttpClient{
// 				Client:  tt.fields.Client,
// 				Timeout: tt.fields.Timeout,
// 			}
// 			_, gotRespHeader, err := h.Post(tt.args.c, tt.args.rawURL, tt.args.data, tt.args.headers, tt.args.option)
// 			if err != nil {
// 				t.Errorf("HttpClient.Post() error = %v", err)
// 				return
// 			}
// 			if gotRespHeader["StatusCode"] != tt.wantStaus {
// 				t.Errorf("HttpClient.Post() StatusCode = %v, wantStaus = %v", gotRespHeader["StatusCode"], tt.wantStaus)
// 				return
// 			}
// 		})
// 	}
// }
