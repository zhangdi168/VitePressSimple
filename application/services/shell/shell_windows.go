//go:build windows

package shell

import (
	"context"
	"golang.org/x/sys/windows"
	"os/exec"
	"syscall"
	"wailstemplate/application/pkg/utils"
)

// RunCmd 运行一个cmd对象,返回cmd对象下标
func (s *ShellService) RunCmd(cmdString string, baseDir string, isAlone bool) string {
	s.isRunning = true
	s.ctx, s.cancel = context.WithCancel(context.Background())
	name, newCmdArr := parseCmd(cmdString) //解析命令
	cmd := exec.CommandContext(s.ctx, name, newCmdArr...)
	cmd.Dir = baseDir
	s.cmd = cmd
	s.isAlone = isAlone
	if isAlone {
		s.cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: windows.
			DETACHED_PROCESS}
	}

	s.printCmdOutput()
	return ""
}

// RunCmdBySystem 调用系统自带的shell运行命令
func (s *ShellService) RunCmdBySystem(cmdString string, baseDir string) string {

	// 定义要执行的命令
	//var cmd *exec.Cmd
	if utils.IsWindows() {
		s.cmd = exec.Command("cmd", "/C", cmdString+" && pause")
		// 为了让命令提示符窗口可见，设置 CreationFlags
		s.cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: windows.CREATE_NEW_CONSOLE}
	} else {
		s.cmd = exec.Command("sh", "-c")

	}
	s.cmd.Dir = baseDir

	// 启动命令
	s.printCmdOutput()

	return ""
}

// sendCtrlBreak 向指定进程发送Ctrl+C控制事件。
// 这个函数尝试模拟用户向控制台进程发送Ctrl+C信号，主要用于测试或中断进程。
// 参数:
//
//	pid - 要发送控制事件的目标进程的ID。
//
// 返回值:
//
//	error - 如果函数执行失败，则返回错误；如果成功，则返回nil。
func sendCtrlBreak(pid int) error {
	// 加载kernel32.dll动态链接库，该库提供了生成控制事件的功能。
	dll, e := syscall.LoadDLL("kernel32.dll")
	if e != nil {
		return e
	}
	// 确保在函数返回前释放dll资源。
	defer dll.Release()

	// 在kernel32.dll中查找GenerateConsoleCtrlEvent函数。
	// 该函数用于生成控制事件，如Ctrl+C或Ctrl+BREAK。
	generateConsoleCtrlEvent, e := dll.FindProc("GenerateConsoleCtrlEvent")
	if e != nil {
		return e
	}

	// 调用GenerateConsoleCtrlEvent函数，发送Ctrl+C事件到指定的进程。
	// 使用Call方法直接调用DLL函数，传入事件类型和目标进程ID。
	r, _, e := generateConsoleCtrlEvent.Call(syscall.CTRL_C_EVENT, uintptr(pid))
	// 如果调用失败，r的值为0，此时返回错误。
	if r == 0 {
		return e
	}
	// 调用成功，返回nil。
	return nil
}
