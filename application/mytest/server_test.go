package mytest

import (
	"path/filepath"
	"testing"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/utils"
	"wailstemplate/application/services"
	setting "wailstemplate/settings"
)

func TestServer(t *testing.T) {
	initData()
	StaticServerInstance := services.NewStaticServer()
	//StaticServerInstance.StartStaticServer("")
	StaticServerInstance.StartStaticServer("")
	select {}
}

func initData() {
	err := cfg.InitCfg(filepath.Join(utils.GetUserHomeDir(), setting.AppName, "configs.yaml"))
	if err != nil {
		println(err.Error())
		return
	}

	//cfg.SetDefault(constant.ConfigKeyXKUserCylinderTypeId, "qUUA+OwjWVrHDG7NmUy9EA==")
}
