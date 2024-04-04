package services

import (
	"wailstemplate/application/pkg/autoupdate"
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
	if info.Version != setting.Version {
		return true
	}
	return false
}

func SetUpdateConfig() {
	autoupdate.SetUpdateConfig(autoupdate.UpdateConfig{
		AppName:                     setting.AppName,
		CurrVersion:                 setting.Version,
		GitType:                     autoupdate.GitTypeGitee,
		GitOwner:                    setting.GitAuthor,
		GitRepo:                     setting.GitRepo,
		WindowReleaseNameContainStr: setting.WindowZipContainStr,
		MacReleaseNameContainStr:    setting.MacosZipContainStr,
	})
}
