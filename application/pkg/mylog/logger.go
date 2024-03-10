package mylog

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// 日志等级常量定义
const (
	levelInfo  = "INFO"
	levelError = "ERROR"
	levelWarn  = "WARN"
	levelDebug = "DEBUG"
	//定时任务
	levelTask = "TASK"
)

type Logger struct {
	logPath string // 日志文件目录
}

var defaultLogger *Logger // 全局静态方法可直接调用

// 初始化单例默认日志对象
func init() {
	defaultLogger = &Logger{
		logPath: "./logs/",
	}
}

// SetLogPath 设置日志文件目录
func SetLogPath(logPath string) {
	//判断文件夹logPath是否存在，不存在则创建
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			log.Println("创建文件夹失败:", err)
			return
		}
		log.Println("文件夹不存在，已成功创建:", logPath)
	}

	defaultLogger.logPath = logPath + "/"
}

// Info 打印info级别日志
func Info(format string, args ...interface{}) {
	defaultLogger.writeLog(levelInfo, format, args...)
}

// Task 打印Task运行日志
func Task(format string, args ...interface{}) {
	defaultLogger.writeLog(levelTask, format, args...)
}

// Error 打印error级别日志
func Error(format string, args ...interface{}) {
	defaultLogger.writeLog(levelError, format, args...)
}

// Err 传入error 如果不为空则进行日志写入
func Err(message string, err error) {
	if err != nil {
		defaultLogger.writeLog(levelError, message+err.Error())
	}
}

// Warn 打印warn级别日志
func Warn(format string, args ...interface{}) {
	defaultLogger.writeLog(levelWarn, format, args...)
}

// Debug 打印Debug级别日志,仅在debug模式下生效
func Debug(format string, args ...interface{}) {
	defaultLogger.writeLog(levelWarn, format, args...)
}

// 写入日志到csv文件
func (l *Logger) writeLog(level, format string, args ...interface{}) {
	// 构建日志内容
	content := fmt.Sprintf(format, args...)

	// 获取当前时间并格式化输出时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 创建日志保存目录
	dir := fmt.Sprintf("%s/%s/", l.logPath, now[:10])
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("创建日志目录失败", err)
		return
	}

	// 构建日志文件路径
	filePath := fmt.Sprintf("%s%s.csv", dir, level)

	// 判断文件是否存在，不存在则新建并写入表头
	var isNewFile bool
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		isNewFile = true
	}
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("打开或创建日志文件失败", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	// 写入表头
	if isNewFile {
		header := []string{"类型", "日志等级", "写入时间", "日志内容"}
		err := writer.Write(header)
		if err != nil {
			fmt.Println("写入csv表头失败", err)
			return
		}
	}

	// 将日志信息写入csv文件
	record := []string{
		getType(content),
		level,
		now,
		content,
	}
	if err := writer.Write(record); err != nil {
		fmt.Println("写入csv日志信息失败", err)
		return
	}

	writer.Flush()

	fmt.Printf("[%s] [%s] %s\n", level, now, content)
}

// 根据日志内容判断类型（可能需要根据实际业务场景修改）
func getType(content string) string {
	if strings.Contains(content, "HTTP") || strings.Contains(content, "TCP") {
		return "SERVER"
	} else if strings.Contains(content, "DB") || strings.Contains(content, "SQL") {
		return "DATABASE"
	}
	return "SYSTEM"
}
