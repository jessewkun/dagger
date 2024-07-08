package tools

import (
	"context"
	"dagger/lib/logger"
	"runtime/debug"
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (wg *WaitGroupWrapper) Wrap(c context.Context, f func()) {
	wg.Add(1)
	go func(c context.Context) {
		defer wg.Done()
		SafeGo(c, f)
	}(c)
}

// 仅仅是对函数本身做 Panic Recover，自身不会启动协程，需要在协程中调用
func SafeGo(c context.Context, fun func()) {
	defer func(ctx context.Context) {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			logger.ErrorWithMsg(c, "dagger_safego", "safeGo recover: %+v, stack is %s", err, stack)
		}
	}(c)
	fun()
}
