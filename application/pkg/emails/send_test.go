// Package emails
// @Author: zhangdi
// @File: send_test
// @Version: 1.0.0
// @Date: 2023/6/25 15:17
package emails

import "testing"

func TestSend(t *testing.T) {
	err := SendEmail1("ideatools@qq.com", "验证码", "您的验证码为0092")
	if err != nil {
		panic(err)
	}
}
