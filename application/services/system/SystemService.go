package system

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	runtime_ "runtime"
	"strings"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/pkg/mylog"
	"wailstemplate/application/pkg/mynet"
	"wailstemplate/application/pkg/utils"
	"wailstemplate/application/wailshelper"
)

type SystemService struct {
}

func NewSystemService() *SystemService {
	return &SystemService{}
}
func (s *SystemService) OpenFileBrowser(dir string) {
	var err error
	switch runtime_.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", dir).Run()
	case "darwin":
		err = exec.Command("open", dir).Run()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		fmt.Printf("failed to open file explorer: %v\n", err)
	}
}

// GetSystemType 获取系统类型
func (s *SystemService) GetSystemType() string {
	return runtime_.GOOS
}

// OpenURL 打开网址
func (s *SystemService) OpenURL(url string) {
	var err error
	switch runtime_.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Run()
	case "darwin":
		err = exec.Command("open", url).Run()
	case "linux":
		err = exec.Command("xdg-open", url).Run()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		fmt.Printf("failed to open URL: %v\n", err)
	}
}

// ConfigSet 配置写入
func (s *SystemService) ConfigSet(k, v string) {
	cfg.Set(k, v)
}

// GetSystemUserHomeDir 获取当前系统用户家目录
func (s *SystemService) GetSystemUserHomeDir() string {
	return utils.GetUserHomeDir()
}

// ConfigGet 配置写入
func (s *SystemService) ConfigGet(k string) any {
	return cfg.Get(k)
}

// ConfigGetBool 获取某个bool值
func (s *SystemService) ConfigGetBool(k string) bool {
	return cfg.GetBool(k)
}

// CheckNetConnect 检测网络连接是否正常
func (s *SystemService) CheckNetConnect() string {
	err := mynet.CheckConnection()
	if err != nil {
		return "网络连接不正常"
	}
	return ""
}

// PathExists 监测一个路径是否存在
func (s *SystemService) PathExists(path string) bool {
	return filehelper.FileExists(path)
}

// SelectDir 选择一个文件夹
func (s *SystemService) SelectDir(title string) string {
	ctx := wailshelper.GetCtx()
	if ctx == nil {
		mylog.Error("SelectDir() 获取ctx失败")
		return "获取ctx失败"
	}
	filters := []runtime.FileFilter{
		{"Image Files", "*.jpg;*.png;*.*"},
	}
	opts := runtime.OpenDialogOptions{
		//DefaultDirectory: fmt.Sprintf("C:\\"),
		DefaultFilename: fmt.Sprintf("a"),
		Title:           title,
		Filters:         filters,
	}
	res, err := runtime.OpenDirectoryDialog(ctx, opts)

	if err != nil {
		return err.Error()
	}

	return res
}

// SelectFile 选择文件
func (s *SystemService) SelectFile(title, baseDir string) string {
	ctx := wailshelper.GetCtx()
	if ctx == nil {
		mylog.Error("SelectFile() 获取ctx失败")
		return "获取ctx失败"
	}

	opts := runtime.OpenDialogOptions{
		DefaultDirectory: fmt.Sprintf(baseDir),
		DefaultFilename:  fmt.Sprintf("logo.png"),
		Title:            title,
	}
	res, err := runtime.OpenFileDialog(ctx, opts)

	if err != nil {
		return err.Error()
	}

	return res
}

// CopyPath 复制一个路径(传入文件或者目录均可)
func (s *SystemService) CopyPath(oriPath, targetPath string, isDir bool) string {
	if !isDir {
		//先创建目标目录
		filehelper.CreateDir(s.GetPathDir(targetPath))
		srcFile, err := os.Open(oriPath)
		if err != nil {
			return err.Error()
		}
		defer srcFile.Close()

		destFile, err := os.Create(targetPath)
		if err != nil {
			return err.Error()
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			return err.Error()
		}

		return ""
	} else { //复制目录
		filehelper.CreateDir(targetPath)
		err := filepath.Walk(oriPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativePath := strings.TrimPrefix(path, oriPath)
			targetFilePath := filepath.Join(targetPath, relativePath)

			if info.IsDir() {
				return os.MkdirAll(targetFilePath, info.Mode())
			}

			srcFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.Create(targetFilePath)
			if err != nil {
				return err
			}
			defer destFile.Close()

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err.Error()
		}
	}

	return ""
}

// GetPathFileName 获取一个文件路径的文件名部分
func (s *SystemService) GetPathFileName(path string) string {
	return filepath.Base(path)
}

// GetPathDir 获取一个文件的dir部分
func (s *SystemService) GetPathDir(path string) string {
	return filepath.Dir(path)
}

// GetPathExt 获取一个文件的后缀名
func (s *SystemService) GetPathExt(path string) string {
	return filepath.Ext(path)
}

// PathJoin 组装文件成系统能识别的路径
func (s *SystemService) PathJoin(elem []string) string {
	return filepath.Join(elem...)
}
