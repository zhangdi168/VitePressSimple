package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
)

type ShellService struct {
	cmd       *exec.Cmd
	isRunning bool
}

func NewShellService() *ShellService {
	return &ShellService{}
}

// RunCmd 运行一个cmd对象,返回cmd对象下标
func (s *ShellService) RunCmd(cmdString string, baseDir string) string {
	//if s.isRunning {
	//	return "当前有一个命令正在运行，任务结束才可以运行新的命令"
	//}
	//s.isRunning = true
	name, newCmdArr := parseCmd(cmdString) //解析命令
	cmd := exec.Command(name, newCmdArr...)
	cmd.Dir = baseDir
	s.cmd = cmd
	s.printCmdOutput()

	return ""
}

// StopCmd 停止命令运行
func (s *ShellService) StopCmd() string {
	//if !s.isRunning {
	//	return "当前没有命令正在运行"
	//}
	err := s.cmd.Process.Signal(syscall.SIGINT)
	if err != nil {
		return err.Error()
	}
	//s.isRunning = false
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
		fmt.Println("INFO:", err)
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		GetOutput(readout) // 在独立的goroutine中获取并处理标准输出
	}()

	// 捕获并处理命令的标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		GetOutput(readerr) // 在独立的goroutine中获取并处理标准错误
	}()

	// 执行命令
	err = cmd.Run()
	if err != nil {
		return
	}
	wg.Wait() // 等待所有goroutine完成
	//s.isRunning = false
}

// GetOutput 从给定的reader中读取输出并打印。
// reader: 用于读取命令输出的bufio.Reader实例。
func GetOutput(reader *bufio.Reader) {
	var sumOutput string // 用于累计整个命令执行过程的所有输出
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes) // 读取命令的输出
		if err != nil {
			if err == io.EOF {
				break // 遇到文件结束EOF时退出循环
			}
			fmt.Println(err)
			sumOutput += err.Error() // 将错误信息累加到sumOutput中
		}
		output := string(outputBytes[:n]) // 将读取到的字节转换为字符串
		fmt.Print(output)                 // 立即打印输出，提供实时反馈
		sumOutput += output               // 将输出累加到sumOutput中
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
