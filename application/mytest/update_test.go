package mytest

import (
	"testing"
	"wailstemplate/application/pkg/autoupdate"
	setting "wailstemplate/settings"
)

func TestAutoUpdate(t *testing.T) {
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
