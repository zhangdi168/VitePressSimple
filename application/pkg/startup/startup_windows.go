// Package startup
// @Author: zhangdi
// @File: startup_windows
// @Version: 1.0.0
// @Date: 2023/7/25 13:38
package startup

import (
	"golang.org/x/sys/windows/registry"
	"path/filepath"
)

func EnableAutoStart(appName string, appPath string) error {
	//打开指定 Windows 注册表项（Registry Key）
	k, err := registry.OpenKey(
		registry.CURRENT_USER,                           //要打开的注册表项所在的根键（Root Key）registry.CURRENT_USER 表示当前用户的注册表项
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, //表示要打开的注册表项的完整路径。不能以 "\" 结尾
		registry.ALL_ACCESS,                             //表示打开注册表项时要求的访问权限,registry.ALL_ACCESS表示完全控制权限
	)
	if err != nil {
		return err
	}
	defer k.Close()

	//程序运行路径

	return k.SetStringValue(appName, appPath)
}

func DisableAutoStart(appName string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()
	return k.DeleteValue(appName)
}

// IsAutoStartEnabledWin IsAutoStartEnabledWin 判断是否成功加入开机自启列表的函数
func IsAutoStartEnabledWin(appName string, appPath string) (bool, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return false, err
	}

	// 获取myapp应用程序对应的值
	v, _, err := k.GetStringValue(appName)
	if err != nil {
		if err == registry.ErrNotExist {
			return false, nil
		}
		return false, err
	}
	defer k.Close()

	// 判断该值是否为应用程序执行路径，如果是则表示成功加入开机自启列表中
	return v == filepath.Join(appPath), nil
}
