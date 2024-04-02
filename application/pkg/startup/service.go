// Package startup
// @Author: zhangdi
// @File: service
// @Version: 1.0.0
// @Date: 2023/4/26 16:51
package startup

import (
	"fmt"
	"path/filepath"
	"runtime"
	"wailstemplate/application/pkg/utils"
)

// SetAutoStart 开机自启设置
// enable 是否启用
func SetAutoStart(appName string, enable bool) error {

	switch runtime.GOOS {
	case "windows":
	case "darwin":
		appPath := filepath.Join(utils.GetRootDir(), appName+".exe")
		if enable {
			return EnableAutoStart(appName, appPath)
		} else {
			return DisableAutoStart(appName)
		}
		break
	default:
		return fmt.Errorf("平台不支持")
	}

	return nil
}
