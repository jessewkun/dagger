package redis

import (
	"context"
	dlog "dagger/lib/logger"
	"dagger/utils"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

const TAGNAME = "DAGGER_REDIS"

var connList map[string]map[string]*redis.Client

func init() {
	connList = make(map[string]map[string]*redis.Client)
}

// InitRedis 初始化redis
func InitRedis(cfg map[string]Config) {
	for dbName, conf := range cfg {
		if err := redisConnect(dbName, conf); err != nil {
			dlog.ErrorWithMsg(context.Background(), TAGNAME, "connect to redis %s error %s", dbName, err)
			continue
		}
		dlog.Info(context.Background(), TAGNAME, "connect to redis %s succ", dbName)
	}
}

// redisConnect 连接 redis
func redisConnect(dbName string, conf Config) error {
	if len(conf.Nodes) < 1 {
		return errors.New(fmt.Sprintf("%s redis node is empty", dbName))
	}

	for _, node := range conf.Nodes {
		client := redis.NewClient(&redis.Options{
			Addr:               node,
			Password:           conf.Password,
			DB:                 conf.Db,
			PoolSize:           500,
			IdleTimeout:        time.Second,
			IdleCheckFrequency: 10 * time.Second,
			MinIdleConns:       3,
			MaxRetries:         3,
			DialTimeout:        2 * time.Second,
		})
		connList[dbName] = make(map[string]*redis.Client, 0)
		connList[dbName][node] = client
	}
	return nil
}

// GetConn 获得redis
func GetConn(dbIns string) *redis.Client {
	if len(connList) < 1 {
		return nil
	}
	if _, ok := connList[dbIns]; !ok {
		return nil
	}

	keys := make([]string, 0, len(connList[dbIns]))
	for key := range connList[dbIns] {
		keys = append(keys, key)
	}

	randomKey := keys[utils.RandomNum(0, len(keys)-1)]
	return connList[dbIns][randomKey]
}

// 探活服务
func Active() {
	for dbName, conns := range connList {
		for node, conn := range conns {
			if _, err := conn.Ping(context.Background()).Result(); err != nil {
				dlog.ErrorWithMsg(context.Background(), TAGNAME, "redis ping db %s node %s error %s", dbName, node, err)
			}

		}
	}
}
