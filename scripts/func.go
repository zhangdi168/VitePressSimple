// Package script
// @Author: zhangdi
// @File: func
// @Version: 1.0.0
// @Date: 2023/9/13 10:16
package scripts

import (
	"path/filepath"
)

func FrontDir() string {
	return filepath.Join(ProjectDir(), "frontend")
}

func ProjectDir() string {
	return "../"
	//return setting.ServeDir
}
