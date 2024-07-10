package redis

import (
	"context"
	"dagger/lib/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

type DaggerRedisHook struct{}

func (h *DaggerRedisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	ctx = context.WithValue(ctx, "dagger_redis_start_time", time.Now())
	return ctx, nil
}

func (h *DaggerRedisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	startTime := ctx.Value("dagger_redis_start_time").(time.Time)
	duration := time.Since(startTime)
	logger.InfoWithField(ctx, TAGNAME, "AfterProcess", map[string]interface{}{
		"cmd":      cmd,
		"duration": duration,
	})
	return nil
}

func (h *DaggerRedisHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	ctx = context.WithValue(ctx, "dagger_redis_start_time", time.Now())
	return ctx, nil
}

func (h *DaggerRedisHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	startTime := ctx.Value("dagger_redis_start_time").(time.Time)
	duration := time.Since(startTime)
	logger.InfoWithField(ctx, TAGNAME, "AfterProcessPipeline", map[string]interface{}{
		"cmd":      cmds,
		"duration": duration,
	})
	return nil
}
