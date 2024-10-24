package sys

// CTX_DAGGER_OUTPUT
// 用于设置gin.Context中的key来存储返回结果，并在中间件中获取记录日志
const CTX_DAGGER_OUTPUT = "dagger_output"

// CTX_USERID
// 用于设置gin.Context中的key来存储用户ID
const CTX_USERID = "user_id"

// CTX_TRACEID
// 用于设置gin.Context中的key来存储trace_id
const CTX_TRACEID = "trace_id"
