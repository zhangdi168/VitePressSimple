package autoupdate

// UpdateCfg 更新配置的默认值
var UpdateCfg = UpdateConfig{
	AppName:                     "appName",
	CurrVersion:                 "v1.0.0",     //当前版本
	GitType:                     GitTypeGitee, //gitee || github
	GitOwner:                    "zhangdi168",
	GitRepo:                     "zhijian-chat",
	WindowReleaseNameContainStr: ".exe", //发行版名称包含该字符串则认为：windows系统
	MacReleaseNameContainStr:    ".zip", //发行版名称包含该字符串则认为：苹果系统
}

// SetUpdateConfig 设置当前程序更新信息 建议在检查更新前进行设置【外部调用】
func SetUpdateConfig(config UpdateConfig) {
	UpdateCfg.AppName = ifThenStr(config.AppName != "", config.AppName, "appName")
	UpdateCfg.CurrVersion = ifThenStr(config.CurrVersion != "", config.CurrVersion, "1.0.0")
	//更新匹配字符串
	UpdateCfg.WindowReleaseNameContainStr = ifThenStr(config.WindowReleaseNameContainStr != "", config.WindowReleaseNameContainStr, "1.0.0")
	UpdateCfg.MacReleaseNameContainStr = ifThenStr(config.MacReleaseNameContainStr != "", config.MacReleaseNameContainStr, "1.0.0")
	//git仓库信息
	UpdateCfg.GitType = GitTypeEnum(ifThenInt(config.CurrVersion != "", int(config.GitType), GitTypeGitee))
	UpdateCfg.GitOwner = ifThenStr(config.GitOwner != "", config.GitOwner, "1.0.0")
	UpdateCfg.GitRepo = ifThenStr(config.GitRepo != "", config.GitRepo, "1.0.0")

}

// GitTypeEnum git源
type GitTypeEnum int

const (
	GitTypeGitee  = iota
	GitTypeGithub = iota
)

type UpdateConfig struct {
	AppName                     string      //程序名称
	CurrVersion                 string      //当前版本
	GitType                     GitTypeEnum //git源
	GitOwner                    string      //仓库拥有者
	GitRepo                     string      //仓库名
	WindowReleaseNameContainStr string      //发行版名称包含该字符串则认为：windows系统
	MacReleaseNameContainStr    string      //发行版名称包含该字符串则认为：苹果系统
}
