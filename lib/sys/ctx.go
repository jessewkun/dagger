package sys

import (
	"context"
)

// CopyCtx 复制新的 context
//
// 避免在gin框架中，http请求结束后，context被cancel，导致在请求中新开的 goroutine 中使用context时出现 ctx canceled 错误
func CopyCtx(ctx context.Context) context.Context {
	ctxUserID := ctx.Value(CTX_USERID)
	ctxTraceID := ctx.Value(CTX_TRACEID)

	newCtx := context.WithValue(context.Background(), CTX_USERID, ctxUserID)
	newCtx = context.WithValue(newCtx, CTX_TRACEID, ctxTraceID)

	return newCtx
}
