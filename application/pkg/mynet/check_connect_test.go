// Package mynet
// @Author: zhangdi
// @File: ping_test
// @Version: 1.0.0
// @Date: 2023/11/23 18:32
package mynet

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	err := CheckConnection()
	if err != nil {
		fmt.Println("网络连接不正常：", err)
	} else {
		fmt.Println("网络连接正常")
	}
}

func TestPingUrl(t *testing.T) {
	err := CheckConnectionWithUrl("http://baidu.com/api2121121", "80")
	if err != nil {
		fmt.Println("网络连接不正常：", err)
	} else {
		fmt.Println("网络连接正常")
	}
}
