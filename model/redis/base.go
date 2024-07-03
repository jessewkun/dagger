package redis

import (
	"dagger/lib/debug"
	redislib "dagger/lib/redis"

	"github.com/go-redis/redis/v8"
)

var debuglog = debug.InitDebug("redis")

func mainDb() *redis.Client {
	return redislib.GetConn("main")
}
