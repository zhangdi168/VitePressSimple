package first

import (
	"path"
	"path/filepath"
	"wailstemplate/application/db/dbinit"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/pkg/mylog"
	"wailstemplate/application/pkg/mytask"
	setting "wailstemplate/settings"
)

// InitDb 初始化sqlite数据库连接
func InitDb() {

	dbinit.InitDb(filepath.Join("data"),
		setting.EntityAutoMigrateList)

}

// InitConfig 初始化配置文件
func InitConfig() {
	yamlPath := filepath.Join("data", "configs.yaml")
	err := cfg.InitCfg(yamlPath)
	if err != nil {
		mylog.Error(err.Error())
		panic(err.Error())
	}
}

// InitLog 初始化日志系统
func InitLog() {
	LogDir := path.Join("data", "logs")
	filehelper.CreateDir(LogDir)
	mylog.SetLogPath(LogDir)
}

// InitTask 初始化定时任务
func InitTask() {
	Manager = mytask.NewTaskManager()
	taskAdd() //添加定时任务
}
