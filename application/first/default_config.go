package first

import (
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
)

func InitDefaultConfig() {
	cfg.SetDefault(keys.ConfigKeyProjectDir, "")
	cfg.SetDefault(keys.ConfigKeyIsStartup, "no")
	cfg.SetDefault(keys.ConfigKeySysUpdateSource, "github")
	cfg.SetDefault(keys.ConfigKeyHistoryProject, "[]")
	cfg.SetDefault(keys.ConfigKeyLayoutNavBgColor, "#ebebeb")
	cfg.SetDefault(keys.ConfigKeyChangeAutoSave, "yes")
	cfg.SetDefault(keys.ConfigKeyFrontMatterSaveType, "json")

}
