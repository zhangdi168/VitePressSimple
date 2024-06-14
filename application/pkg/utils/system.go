package utils

import "runtime"

// IsMacOS 是否为苹果系统
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
