package first

import (
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
)

func InitDefaultConfig() {
	cfg.SetDefault(keys.ConfigKeyProjectDir, "")
	cfg.SetDefault(keys.ConfigKeyHistoryProject, "[]")

}
