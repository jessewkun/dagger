package sys

import "dagger/common"

// isDebug 是否是 debug 模式
// 开发环境
func IsDebug() bool {
	return common.Cfg.Mode == "debug"
}

// IsRelease 是否是 release 模式
// 生产环境
func IsRelease() bool {
	return common.Cfg.Mode == "release"
}

// IsTest 是否是 test 模式
// 测试环境
func IsTest() bool {
	return common.Cfg.Mode == "test"
}
