package scripts

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"wailstemplate/application/constant/enums"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/services/shell"
	setting "wailstemplate/settings"
)

// 支持的平台
var platforms = []string{
	"windows/amd64",
	//"windows/arm64",
	"darwin/arm64",
	"darwin/amd64",
}

func TestBuildEntry(t *testing.T) {
	switch runtime.GOOS {
	case enums.SystemMac:
		buildMacos()
		break
	case enums.SystemLinux:
		break
	case enums.SystemWindows:
		buildWindows()
		break
	}

}
func buildWindows() {
	for _, platform := range platforms {
		if !strings.Contains(platform, enums.SystemWindows) {
			continue
		}
		s := shell.NewShellService(0, nil)
		s.RunCmd("wails build -platform "+platform, "../", false)
		baseDir := filepath.Join("..", "build", "bin")
		zipPath := filepath.Join(baseDir, fmt.Sprintf("vpsimple_%s_%s.zip", replaceSystemStr(platform), setting.Version))
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

func buildMacos() {
	for _, platform := range platforms {
		if !strings.Contains(platform, enums.SystemMac) {
			continue
		}
		s := shell.NewShellService(0, nil)
		s.RunCmd("wails build  -platform "+platform, "../", false)
		//baseDir := filepath.Join("..", "build", "bin")

		createDmg(platform)
		//zipPath := filepath.Join(baseDir, fmt.Sprintf("%s_%s.zip", replaceSystemStr(platform), setting.Version))
		//appPath := filepath.Join(baseDir, fmt.Sprintf("%sInstaller.dmg", setting.AppName))
		//if !filehelper.FileExists(appPath) {
		//	println(appPath + ": not exists")
		//	return
		//}
		//err := filehelper.CompressFilesToZip(zipPath, []string{appPath})
		//if err != nil {
		//	println(err.Error())
		//} else {
		//	println("ok")
		//}

	}

}

func replaceSystemStr(platform string) string {
	s1 := strings.ReplaceAll(platform, "/", "_")
	return strings.ReplaceAll(s1, "darwin", "mac")
}

//func TestCreateDmg(t *testing.T) {
//	createDmg(platforms[0])
//}

// 准备create-dmg命令参数
// 构建用于创建DMG文件的参数列表
// 参数:
//
//	scriptDir - 脚本目录的路径
//	sourceDir - 源文件目录的路径
//	volumeName - DMG卷的名称
//	volumeIcon - DMG卷的图标路径
//	appName - 应用程序的名称
//	dmgName - 生成的DMG文件的名称
func createDmg(platform string) {
	// 获取当前脚本所在目录
	scriptDir := filepath.Join("..", "build", "dmg")
	// 初始化源目录路径
	sourceDir := filepath.Join(scriptDir, "..", "bin")
	// 查找并设置应用程序路径
	appPath, err := findAppPath(sourceDir)
	if err != nil {
		panic(err)
	}
	appName := filepath.Base(appPath)
	// 提取卷宗名称（不包括.app后缀）
	//volumeName := strings.TrimSuffix(appName, ".app")
	// 查找卷宗图标路径
	volumeIcon, _ := findVolumeIcon(sourceDir)
	// 设置DMG文件名
	dmgName := fmt.Sprintf("vpsimple_%s_%s.dmg", replaceSystemStr(platform), setting.Version)
	// 如果DMG文件已存在，则删除
	removeExistingDmg(sourceDir, dmgName)
	msg := fmt.Sprintf("若提示'xx 损坏' 右键打开黑色图标修复")
	// 返回值:
	//   无
	args := []string{
		"--no-internet-enable",       // 禁用互联网启用
		"--hdiutil-quiet",            // 安静模式执行
		"--volname", setting.AppName, // 设置卷名称
		"--volicon", volumeIcon, // 设置卷图标
		"--background", filepath.Join(scriptDir, "background.tiff"), // 设置背景图像
		"--text-size", "12", // 设置文本大小
		"--window-pos", "400", "400", // 设置窗口位置
		"--window-size", "660", "450", // 设置窗口大小
		"--icon-size", "80", // 设置图标大小
		"--icon", appName, "180", "180", // 设置应用程序图标位置和大小
		"--hide-extension", appName, // 隐藏应用程序的文件扩展名
		"--app-drop-link", "480", "180", // 设置应用程序拖放链接的位置
		"--add-file", msg, filepath.Join(scriptDir, "fix-app"), "330", "290", // 添加修复文件及其位置和大小
		filepath.Join(sourceDir, dmgName), // 添加源文件目录中的DMG文件
		sourceDir,                         // 添加源文件目录
	}

	// 调用create-dmg工具创建DMG文件
	if err := runCreateDmg(scriptDir, args); err != nil {
		panic(err)
	}
}

func findAppPath(dir string) (string, error) {
	fs := afero.NewOsFs()
	matches, err := afero.Glob(fs, filepath.Join(dir, "*.app"))
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", fmt.Errorf("no app found in directory: %s", dir)
	}
	return matches[0], nil
}

func findVolumeIcon(dir string) (string, error) {
	fs := afero.NewOsFs()
	matches, err := afero.Glob(fs, filepath.Join(dir, "*.icns"))
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "", fmt.Errorf("no volume icon found in directory: %s", dir)
	}
	return matches[0], nil
}

func removeExistingDmg(dir, name string) {
	path := filepath.Join(dir, name)
	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(path); err != nil {
			panic(err)
		}
	}
}

func runCreateDmg(scriptDir string, args []string) error {
	createDmgPath := filepath.Join(scriptDir, "create-dmg")
	cmd := exec.Command(createDmgPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run create-dmg: %w\nOutput: %s", err, output)
	}
	return nil
}
