package utils

import "net"

// 是否禁止ip
func IsBan(clientIP string, whiteList []string) bool {
	if len(whiteList) < 1 {
		return false
	}

	for _, row := range whiteList {
		if row == "*" {
			return false
		}

		if _, ipNet, err := net.ParseCIDR(row); err != nil {
			// ip地址不是掩码形式或者ip地址格式错误
			if clientIP == row {
				return false
			}
		} else {
			// 判断白名单ip是否包含客户端ip
			if ipNet.Contains(net.ParseIP(clientIP)) {
				return false
			}
		}
	}

	return true
}
