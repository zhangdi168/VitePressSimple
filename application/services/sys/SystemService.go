package service

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	runtime_ "runtime"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/mylog"
	"wailstemplate/application/pkg/mynet"
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

// ConfigSet 配置写入
func (s *SystemService) ConfigSet(k, v string) {
	cfg.Set(k, v)
}

// ConfigGet 配置写入
func (s *SystemService) ConfigGet(k string) any {
	return cfg.Get(k)
}

// ConfigGetBool 获取某个bool值
func (s *SystemService) ConfigGetBool(k string) bool {
	return cfg.GetBool(k)
}

// CheckUpdate 检查更新
func (s *SystemService) CheckUpdate() {
	////更新程序所在路径
	//updateExePath := cfg.GetString(constant.ConfigKeySysUpdateExePath)
	////当前程序名称
	//ProduceName := cfg.GetString(constant.ConfigKeySysProductName)
	//cmd := exec.Command(updateExePath)            //参数:【升级程序】HHUpdateApp.exe程序所在路径
	//cmd.Args = append(cmd.Args, ProduceName, "1") //参数1:【应用程序】的名词，例如：LOLClient；参数1:检查更新模式
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//err := cmd.Run()
	//if err != nil {
	//	mylog.Error("CheckUpdate" + err.Error())
	//}
}

// CheckNetConnect 检测网络连接是否正常
func (s *SystemService) CheckNetConnect() string {
	err := mynet.CheckConnection()
	if err != nil {
		return "网络连接不正常"
	}
	return ""
}

// SelectDir 选择一个文件夹
func (s *SystemService) SelectDir() string {
	ctx := wailshelper.GetCtx()
	if ctx == nil {
		mylog.Error("SelectDir() 获取ctx失败")
		return ""
	}
	filters := []runtime.FileFilter{
		{"Image Files", "*.jpg;*.png;*.*"},
	}
	opts := runtime.OpenDialogOptions{
		DefaultDirectory: fmt.Sprintf("C:\\"),
		DefaultFilename:  fmt.Sprintf("a"),
		Title:            fmt.Sprintf("请选择"),
		Filters:          filters,
	}
	res, err := runtime.OpenDirectoryDialog(ctx, opts)

	if err != nil {
		return err.Error()
	}

	return res
}
