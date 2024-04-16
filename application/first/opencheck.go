package first

import (
	"github.com/ncruces/zenity"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
)

// CheckCanOpen 返回true时可以打开
func CheckCanOpen() bool {
	isOpen := cfg.GetBool(keys.ConfigKeySysProgramIsOpen)
	if !isOpen {
		cfg.Set(keys.ConfigKeySysProgramIsOpen, "yes")
		return true
	}
	//canOpenMany := cfg.GetBool(keys.ConfigKeySysProgramCanOpenMany)
	err := zenity.Question("检测到程序已经打开，是否重复打开",
		zenity.Icon(zenity.QuestionIcon),
		zenity.OKLabel("重复打开"),
		zenity.CancelLabel("取消打开"))
	if err != nil {
		return false //不重复打开
	}
	return true
}
