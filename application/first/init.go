package first

import (
	"fmt"
	"path"
	"path/filepath"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/db/dbinit"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/pkg/mylog"
	"wailstemplate/application/pkg/mytask"
	"wailstemplate/application/pkg/startup"
	"wailstemplate/application/pkg/utils"
	setting "wailstemplate/settings"
)

// InitDb 初始化sqlite数据库连接
func InitDb() {

	dbinit.InitDb(filepath.Join(utils.GetUserHomeDir(), setting.AppName, "db"),
		setting.EntityAutoMigrateList)

}

// InitConfig 初始化配置文件
func InitConfig() {
	yamlPath := filepath.Join(utils.GetUserHomeDir(), setting.AppName, "configs.yaml")
	err := cfg.InitCfg(yamlPath)
	if err != nil {
		mylog.Error(err.Error())
		panic(err.Error())
	}
}

// CheckStartup 检查开机自启
func CheckStartup() {
	err := startup.SetAutoStart(setting.AppName, cfg.GetBool(keys.ConfigKeyIsStartup))
	if err != nil {
		mylog.Error(fmt.Sprintf("设置开机自启失败:%s", err.Error()))
	}

}

// InitLog 初始化日志系统
func InitLog() {
	LogDir := path.Join(utils.GetUserHomeDir(), setting.AppName, "logs")
	filehelper.CreateDir(LogDir)
	mylog.SetLogPath(LogDir)
}

// InitTask 初始化定时任务
func InitTask() {
	Manager = mytask.NewTaskManager()
	taskAdd() //添加定时任务
}
