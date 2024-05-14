package first

import (
	"wailstemplate/application/constant/cnts"
	"wailstemplate/application/constant/enums"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
)

func InitDefaultConfig() {
	cfg.SetDefault(keys.ConfigKeyProjectDir, "")
	cfg.SetDefault(keys.ConfigKeyIsStartup, "no")
	cfg.SetDefault(keys.ConfigKeySysUpdateSource, "github")
	cfg.SetDefault(keys.ConfigKeyHistoryProject, "[]")
	cfg.SetDefault(keys.ConfigKeyLayoutNavBgColor, "#ebebeb")
	cfg.SetDefault(keys.ConfigKeyChangeAutoSave, "no")
	cfg.SetDefault(keys.ConfigKeyFrontMatterSaveType, "json")

	//项目静态资源服务器的端口
	cfg.SetDefault(keys.ConfigKeySysStaticServerPort, "9874")
	cfg.SetDefault(keys.ConfigKeySysProjectStaticDirName, cnts.StaticBaseDir)

	cfg.SetDefault(keys.ConfigKeyVditorCdn, enums.VditorCdnZstatic)
	cfg.SetDefault(keys.ConfigKeyLang, "en")

	//cfg.SetDefault(keys.ConfigKeySysProgramCanOpenMany, "no")
	cfg.SetDefault(keys.ConfigKeySysProgramIsOpen, "no")

}
