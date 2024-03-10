// Package utils
// @Author: zhangdi
// @File: root
// @Version: 1.0.0
// @Date: 2023/5/6 16:51
package utils

import (
	"os"
	"path/filepath"
)

// GetRootDir 获取当前程序运行的根目录
func GetRootDir() string {
	//注意！！！go run 命令运行 Go 程序时，os.Executable() 函数返回的是
	//【一个临时文件的路径】，而不是程序最终生成的可执行文件的路径
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	root := filepath.Dir(ex)
	return root
}

// GetProgramName 获取当前运行程序的文件名如：a.exe
func GetProgramName() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Base(ex)
}
