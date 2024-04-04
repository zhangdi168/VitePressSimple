package scripts

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/services"
	setting "wailstemplate/settings"
)

// 支持的平台
var platforms = []string{
	"windows/amd64",
	//"windows/arm64",
	"darwin/amd64",
	"darwin/arm64",
}

func TestBuildWindows(t *testing.T) {
	for _, platform := range platforms {
		if !strings.Contains(platform, "windows") {
			continue
		}
		s := services.NewShellService()
		s.RunCmd("wails build  -platform "+platform, "../")
		baseDir := filepath.Join("..", "build", "bin")
		zipPath := filepath.Join(baseDir, fmt.Sprintf("%s_%s.zip", strings.ReplaceAll(platform, "/", "_"), setting.Version))
		fileList := []string{
			filepath.Join(baseDir, fmt.Sprintf("%s.exe", setting.AppName)),
		}
		err := filehelper.CompressFilesToZip(zipPath, fileList)
		if err != nil {
			println(err.Error())
		} else {
			println("ok")
		}
	}

}
