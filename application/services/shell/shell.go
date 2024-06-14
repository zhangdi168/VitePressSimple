package shell

import (
	"bufio"
	"context"
	"errors"
	"fmt"

	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"wailstemplate/application/pkg/utils"
)

type NotifyMessage func(data NotifyShellData)
type ShellService struct {
	cmd       *exec.Cmd
	index     int
	isRunning bool //是否正在运行
	notify    NotifyMessage
	ctx       context.Context
	cancel    context.CancelFunc
	isAlone   bool
}

func NewShellService(index int, notify NotifyMessage) *ShellService {
	if notify != nil {
		return &ShellService{
			index:  index,
			notify: notify,
		}
	} else {
		return &ShellService{
			index: index,
			notify: func(data NotifyShellData) {
				fmt.Printf("[%d]%s%s", data.ShellIndex, utils.IfToString(data.IsError, "[error]", ""), data.Content)
			},
		}
	}
}

func (s *ShellService) notice(content string, isError bool) {
	if utils.IsWindows() && s.isAlone {
		//mylog.Info("cmd result:" + content)
		return
	}
	s.notify(NotifyShellData{
		ShellIndex: s.index,
		Content:    content,
		IsError:    isError,
	})
}

// StopCmd 停止命令运行(有问题，端口无法释放 暂时不用)
func (s *ShellService) StopCmd() string {
	if s.cmd == nil || s.cmd.Process == nil {
		return fmt.Sprintf("command is not running")
	}

	if utils.IsWindows() {
		return "windows 无法终止命令运行"
	}

	err := s.cmd.Process.Signal(os.Interrupt)
	if err != nil {

		s.notice("停止命令运行失败2:"+err.Error(), true)
		return err.Error()
	}
	//// 等待命令结束
	err = s.cmd.Wait()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// 命令可能因为接收到 SIGINT 而退出，检查退出状态码
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				// 根据需要处理退出状态码
				s.notice(fmt.Sprintf("%sCommand exited with status:%d", err.Error(), status.ExitStatus()), false)
			}
		}
		return err.Error()
	}
	s.notice("cmd stop \n", false)
	s.isRunning = false
	return ""
}

// printCmdOutput 捕获并打印命令的输出。
func (s *ShellService) printCmdOutput() {

	cmd := s.cmd
	// 将命令的标准输入设置为操作系统标准输入
	cmd.Stdin = os.Stdin

	// 使用等待组来等待所有goroutine完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 捕获并处理命令的标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		s.notice("捕获并处理命令的标准输出:"+err.Error(), true)
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		s.GetOutput(readout) // 在独立的goroutine中获取并处理标准输出
	}()

	// 捕获并处理命令的标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		s.notice("捕获并处理命令的标准错误:"+err.Error(), true)
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		s.GetOutput(readerr) // 在独立的goroutine中获取并处理标准错误
	}()

	// 执行命令
	err = cmd.Start()
	if err != nil {
		s.notice("执行命令:"+err.Error(), true)
		return
	}

	// 等待命令完成
	err = cmd.Wait()
	if err != nil {
		s.notice("cmd run error:"+err.Error(), true)
	}
	s.notice("\ncmd run complete!", false)
	s.isRunning = false
	wg.Wait() // 等待所有goroutine完成
	//
}

// GetOutput 从给定的reader中读取输出并打印。
// reader: 用于读取命令输出的bufio.Reader实例。
func (s *ShellService) GetOutput(reader *bufio.Reader) {
	var sumOutput string // 用于累计整个命令执行过程的所有输出
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes) // 读取命令的输出
		if err != nil {
			if err == io.EOF {
				break // 遇到文件结束EOF时退出循环
			}
			//fmt.Println(err)
			sumOutput += err.Error() // 将错误信息累加到sumOutput中
		}
		output := string(outputBytes[:n]) // 将读取到的字节转换为字符串
		// 立即打印输出，提供实时反馈

		s.notice(output, false)
		sumOutput += output // 将输出累加到sumOutput中
	}
	if sumOutput != "" {
		s.notice(sumOutput, true)
	}

}

// parseCmd 解析命令
func parseCmd(cmdString string) (string, []string) {
	cmdArr := strings.Split(cmdString, " ")
	var newCmdArr []string
	var name = "" //程序名 如 npm
	for i := 0; i < len(cmdArr); i++ {
		if cmdArr[i] != "" {
			if name == "" {
				name = cmdArr[i]
			} else {
				newCmdArr = append(newCmdArr, cmdArr[i])

			}
		}
	}
	return name, newCmdArr
}
