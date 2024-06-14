//go:build linux || darwin

package shell

import (
	"context"
	"os/exec"
)

// RunCmd 运行一个cmd对象,返回cmd对象下标
func (s *ShellService) RunCmd(cmdString string, baseDir string, isAlone bool) string {
	s.isRunning = true
	s.isAlone = isAlone

	s.ctx, s.cancel = context.WithCancel(context.Background())
	name, newCmdArr := parseCmd(cmdString) //解析命令
	cmd := exec.CommandContext(s.ctx, name, newCmdArr...)
	cmd.Dir = baseDir
	s.cmd = cmd
	s.printCmdOutput()
	return ""
}

// RunCmdBySystem 调用系统自带的shell运行命令
func (s *ShellService) RunCmdBySystem(cmdString string, baseDir string) string {

	// 定义要执行的命令
	//var cmd *exec.Cmd
	s.cmd = exec.Command("sh", "-c", cmdString)
	s.cmd.Dir = baseDir

	// 启动命令
	err := s.cmd.Start()
	if err != nil {
		return err.Error()
	}
	return ""
}
