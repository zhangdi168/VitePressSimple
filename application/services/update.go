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

func (s *UpdateService) CheckUpdate() {
	autoupdate.SetUpdateConfig(autoupdate.UpdateConfig{
		AppName:                     setting.AppName,
		CurrVersion:                 setting.Version,
		GitType:                     autoupdate.GitTypeGitee,
		GitOwner:                    setting.GitAuthor,
		GitRepo:                     setting.GitRepo,
		WindowReleaseNameContainStr: setting.WindowZipContainStr,
		MacReleaseNameContainStr:    setting.MacosZipContainStr,
	})

	autoupdate.CheckUpdateEntry()
}
