package constant

// CTX_DAGGER_OUTPUT
// 用于设置gin.Context中的key来存储返回结果，并在中间件中获取记录日志
const CTX_DAGGER_OUTPUT = "dagger_output"

// CTX_DAGGER_REDIS_START_TIME
// 用于在redis hook中设置context中的key来存储redis cmd开始时间
const CTX_DAGGER_REDIS_START_TIME = "dagger_redis_start_time"
