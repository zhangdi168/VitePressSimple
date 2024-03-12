package vpsimpler

import (
	"fmt"
	"os/exec"
)

func NodejsIsInstall() bool {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		return false
	}
	fmt.Println(string(output))
	return true
}

// GetNodeVersion 获取nodejs版本
func GetNodeVersion() string {
	output, err := exec.Command("node", "-v").Output()
	if err != nil {
		return ""
	}
	return string(output)
}

// GetNpmVersion 获取npm版本
func GetNpmVersion() string {
	output, err := exec.Command("npm", "-v").Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func NpmIsInstall() bool {
	output, err := exec.Command("npm", "version").Output()
	if err != nil {
		return false
	}
	fmt.Println(string(output))
	return true
}
