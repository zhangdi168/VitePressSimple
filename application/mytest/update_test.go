package mytest

import (
	"path/filepath"
	"testing"
	"wailstemplate/application/pkg/autoupdate"
	setting "wailstemplate/settings"
)

// 设置自动更新的配置，并检查是否有新版本可用。
// 如果有新版本，询问用户是否自动更新。
// 如果用户同意更新，则开始下载更新。
// 设置自动更新的配置信息
func TestAutoUpdate(t *testing.T) {
	autoupdate.SetUpdateConfig(autoupdate.UpdateConfig{
		AppName:                     setting.AppName,
		CurrVersion:                 setting.Version,
		GitType:                     autoupdate.GitTypeGithub,
		GitOwner:                    setting.GitAuthor,
		GitRepo:                     setting.GitRepo,
		WindowReleaseNameContainStr: setting.WindowZipContainStr,
		MacReleaseNameContainStr:    setting.MacosZipContainStr,
	})

	// 获取最新版本的信息
	autoupdate.CheckUpdateEntry()
}

// 测试解压压缩包
func TestUnzip(t *testing.T) {
	archivePath := "C:\\Users\\17788\\Downloads\\windows_1.0.1.zip"
	dir := "C:\\Users\\17788\\Downloads"
	autoupdate.UnzipFile(archivePath, dir)
	autoupdate.UpWindowSelfApp(filepath.Join(dir, setting.AppName+".exe"))
}
