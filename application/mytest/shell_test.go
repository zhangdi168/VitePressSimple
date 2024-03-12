package mytest

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"testing"
	"wailstemplate/application/pkg/vpsimpler"
)

func TestShell(t *testing.T) {
	makeCmd := exec.Command("ping", "baidu.com")
	PrintCmdOutput(makeCmd)
}

func TestCheckNpmInstall(t *testing.T) {
	print(vpsimpler.GetNpmVersion())
}

// PrintCmdOutput 打印命令执行后的输出。
// cmd: 指定要执行的命令。
func PrintCmdOutput(cmd *exec.Cmd) {
	// 将命令的标准输入设置为操作系统标准输入
	cmd.Stdin = os.Stdin

	// 使用等待组来等待所有goroutine完成
	var wg sync.WaitGroup
	wg.Add(2)

	// 捕获并处理命令的标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("INFO:", err)
		os.Exit(1)
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
		os.Exit(1)
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
