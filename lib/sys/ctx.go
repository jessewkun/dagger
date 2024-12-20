package sys

import (
	"context"
)

const CtxUserIDKey = "user_id"
const CtxTraceIDKey = "trace_id"

// CopyCtx 复制新的 context
//
// 避免在gin框架中，http请求结束后，context被cancel，导致在请求中新开的 goroutine 中使用context时出现 ctx canceled 错误
func CopyCtx(ctx context.Context) context.Context {
	ctxUserID := ctx.Value(CtxUserIDKey)
	ctxTraceID := ctx.Value(CtxTraceIDKey)

	newCtx := context.WithValue(context.Background(), CtxUserIDKey, ctxUserID)
	newCtx = context.WithValue(newCtx, CtxTraceIDKey, ctxTraceID)

	return newCtx
}
