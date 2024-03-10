// Package utils
// @Author: zhangdi
// @File: user
// @Version: 1.0.0
// @Date: 2023/5/16 17:33
package utils

import (
	"os/user"
)

// GetUserHomeDir 获取用户家目录
func GetUserHomeDir() string {
	currentUser, err := user.Current()
	if err != nil {
		return ""
	}
	// 返回格式形如: C:\Users\Administrator
	return currentUser.HomeDir
}
