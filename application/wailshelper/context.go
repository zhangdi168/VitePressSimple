// Package wailshelper
// @Author: zhangdi
// @File: context
// @Version: 1.0.0
// @Date: 2023/5/19 18:54
package wailshelper

import (
	"context"
	"time"
)

var ctx *context.Context = nil
var IsSetCtx = false

// GetCtx 返回[副本]ctx
func GetCtx() context.Context {
	if !IsSetCtx || ctx == nil {
		//等待2s后继续读取
		time.Sleep(time.Millisecond * time.Duration(2000))
	}
	if ctx != nil {
		return *ctx
	}
	return nil
}

func SetCtx(ctx_ *context.Context) {
	IsSetCtx = true
	ctx = ctx_
}
