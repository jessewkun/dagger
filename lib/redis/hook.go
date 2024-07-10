package redis

import (
	"context"
	"dagger/lib/constant"
	"dagger/lib/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
 * @description: redis hook
 */

type daggerRedisHook struct{}

func (h *daggerRedisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	ctx = context.WithValue(ctx, constant.CTX_DAGGER_REDIS_START_TIME, time.Now())
	return ctx, nil
}

func (h *daggerRedisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	startTime := ctx.Value(constant.CTX_DAGGER_REDIS_START_TIME).(time.Time)
	duration := time.Since(startTime)
	logger.InfoWithField(ctx, TAGNAME, "AfterProcess", map[string]interface{}{
		"cmd":      cmd,
		"duration": duration,
	})
	return nil
}

func (h *daggerRedisHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	ctx = context.WithValue(ctx, constant.CTX_DAGGER_REDIS_START_TIME, time.Now())
	return ctx, nil
}

func (h *daggerRedisHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	startTime := ctx.Value(constant.CTX_DAGGER_REDIS_START_TIME).(time.Time)
	duration := time.Since(startTime)
	logger.InfoWithField(ctx, TAGNAME, "AfterProcessPipeline", map[string]interface{}{
		"cmds":     cmds,
		"duration": duration,
	})
	return nil
}
