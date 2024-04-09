package services

import (
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/autoupdate"
	"wailstemplate/application/pkg/cfg"
	setting "wailstemplate/settings"
)

type UpdateService struct {
}

func NewUpdateService() *UpdateService {
	return &UpdateService{}
}

func (s *UpdateService) UpdateNewVersion() {
	SetUpdateConfig()
	autoupdate.CheckUpdateEntry()
}
func (s *UpdateService) HasNewVersion() bool {
	SetUpdateConfig()
	info := autoupdate.GetGitRepoLatestReleaseInfo()
	if info == nil {
		return false
	}
	if info.Version != setting.Version {
		return true
	}
	return false
}

func SetUpdateConfig() {
	source := cfg.GetString(keys.ConfigKeySysUpdateSource)
	var gitType autoupdate.GitTypeEnum = autoupdate.GitTypeGithub
	if source == "gitee" {
		gitType = autoupdate.GitTypeGitee
	}
	autoupdate.SetUpdateConfig(autoupdate.UpdateConfig{
		AppName:                     setting.AppName,
		CurrVersion:                 setting.Version,
		GitType:                     gitType,
		GitOwner:                    setting.GitAuthor,
		GitRepo:                     setting.GitRepo,
		WindowReleaseNameContainStr: setting.WindowZipContainStr,
		MacReleaseNameContainStr:    setting.MacosZipContainStr,
	})
}
