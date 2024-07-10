package redis

import (
	"context"
	"time"
)

// 注意判断 redis.Nil
func TestGet(ctx context.Context, key string) (string, error) {
	return mainDb().Get(ctx, key).Result()
}

func TestSet(ctx context.Context, key string, val string) error {
	return mainDb().Set(ctx, key, val, time.Second*10).Err()
}
