package services

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"wailstemplate/application/pkg/mylog"
	"wailstemplate/application/wailshelper"
)

// ShellManager 终端管理器
type ShellManager struct {
	ShellList []*ShellService
}

type NotifyShellData struct {
	ShellIndex int    `json:"shell_index,omitempty"`
	Content    string `json:"content,omitempty"`
	IsError    bool   `json:"is_error,omitempty"`
}

func NewShellManager() *ShellManager {
	return &ShellManager{}
}

// CreateShellAndRunCmd 创建一个终端，同时运行命令
func (s *ShellManager) CreateShellAndRunCmd(baseDir, cmd string, isAlone bool) int {
	shellIndex := s.CreateShell()
	res := s.RunCmd(baseDir, cmd, shellIndex, isAlone)
	if res != "" {
		mylog.Error("运行终端命令失败：" + res)
		return -1
	}
	return shellIndex
}

// CreateShell 新建一个终端
func (s *ShellManager) CreateShell() int {
	shellIndex := len(s.ShellList)
	//回调函数
	handle := func(data NotifyShellData) {
		if wailshelper.IsSetCtx {
			ctx := wailshelper.GetCtx()
			if ctx != nil {
				runtime.EventsEmit(ctx, "shell", data)
			}
		} else {
			fmt.Printf("shell:%d %s\n", shellIndex, data.Content)
		}
	}
	s.ShellList = append(s.ShellList, NewShellService(shellIndex, handle))
	return shellIndex
}

// RunCmd 运行一个终端命令
func (s *ShellManager) RunCmd(baseDir, cmd string, shellIndex int, isAlone bool) string {
	//baseDir := cfg.GetString(keys.ConfigKeyShellBaseDir)
	if baseDir == "" {
		return fmt.Sprintf("请先设置终端基础目录")
	}
	if shellIndex >= len(s.ShellList) {
		return fmt.Sprintf("shellIndex超出范围")
	}
	if s.ShellList[shellIndex].isRunning {
		return fmt.Sprintf("shell%d 正在运行", shellIndex)
	}
	go s.ShellList[shellIndex].RunCmd(cmd, baseDir, isAlone)
	return ""
}

// RunCmdBySystem 运行一个终端命令
func (s *ShellManager) RunCmdBySystem(baseDir, cmd string, shellIndex int) string {
	//baseDir := cfg.GetString(keys.ConfigKeyShellBaseDir)
	if baseDir == "" {
		return fmt.Sprintf("请先设置终端基础目录")
	}
	if shellIndex >= len(s.ShellList) {
		return fmt.Sprintf("shellIndex超出范围")
	}
	if s.ShellList[shellIndex].isRunning {
		return fmt.Sprintf("shell%d 正在运行", shellIndex)
	}
	go s.ShellList[shellIndex].RunCmdBySystem(cmd, baseDir)
	return ""
}

// StopShell 停止正在运行的终端()
func (s *ShellManager) StopShell(shellIndex int) string {
	if len(s.ShellList) < shellIndex {
		return "shellIndex不存在"
	}
	return s.ShellList[shellIndex].StopCmd()
}

func (s *ShellManager) Export(data NotifyShellData) {

}
