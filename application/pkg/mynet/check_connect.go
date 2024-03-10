// Package mynet
// @Author: zhangdi
// @File: ping
// @Version: 1.0.0
// @Date: 2023/11/23 18:30
package mynet

import (
	"errors"
	"net"
	"net/url"
	"time"
)

// Ping
// 如果收到"network icmp unknown"的错误消息，操作系统不支持ICMP协议，可使用tcp协议
func Ping(host string, timeout time.Duration) error {
	conn, err := net.DialTimeout("icmp", host, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

// CheckConnection 通过TCP协议进行网络连接检查
func CheckConnection() error {
	host := "www.baidu.com" //可自行输入
	port := "80"
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

// CheckConnectionWithUrl 检测网络连接状态
func CheckConnectionWithUrl(host, port string) error {
	domain, err := getDomain(host)
	if err != nil {
		return errors.New("链接解析异常" + err.Error())
	}
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", domain+":"+port, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
func getDomain(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	domain := u.Hostname()

	return domain, nil
}
