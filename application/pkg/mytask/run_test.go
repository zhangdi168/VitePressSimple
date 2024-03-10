// Package mytask
// @Author: zhangdi
// @File: helper
// @Version: 1.0.0
// @Date: 2023/11/16 13:08
package mytask

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	task := NewTaskManager()

	//每隔10秒执行一次
	spec1 := FormatEverySpace(0, 0, 10)
	task.AddTask("every_10s", spec1, func() {
		//mylog.Task("定时任务被执行")
		fmt.Println("every_10s", time.Now().UTC().Format("2006-01-02 15:04:05"))
	})

	//每天的某点执行
	task.AddTask("every_day_15_15", FormatEveryDay(15, 31, 10), func() {
		//mylog.Task("【every_day_15_15】定时任务被执行")
		fmt.Println("每天的某点执行", time.Now().Format("2006-01-02 15:04:05"))
	})

	// 保持主程序运行
	select {}
}
