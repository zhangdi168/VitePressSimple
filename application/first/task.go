package first

import (
	"wailstemplate/application/pkg/mytask"
)

// Manager 定时任务管理器
var Manager *mytask.TaskManager

// 定时任务在这里添加
func taskAdd() {

	//使用示例如下：
	//spec := mytask.FormatEverySpace(0, 5, 0)
	//_, err := Manager.AddTask("同步统计数据到db", spec, SyncToDb)
	//if err != nil {
	//	mylog.Error("初始化定时任务[同步统计数据到db]error:" + err.Error())
	//	return
	//}
}

// SyncToDb 定时同步
func SyncToDb() {

}
