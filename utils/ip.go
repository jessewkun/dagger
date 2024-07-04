package utils

import (
	"fmt"
	"net"
)

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

// 获取本地ip
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		// 检查 IP 地址类型并跳过回环地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no IP address found")
}
