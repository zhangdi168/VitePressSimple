package autoupdate

import (
	"fmt"
	"github.com/ncruces/zenity"
	"os"
	"os/exec"
)

func GetAppPathWindow() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return exe, nil
}

func UpWindowSelfApp(exeResourceFilePath string) (bool, string) {
	// window更新方法
	selfPathWindow, _ := GetAppPathWindow()

	//文件名 := filepath.Base(self路径Window)
	// 检查文件是否存在
	oldFileName := selfPathWindow + ".old.bak"
	if _, err := os.Stat(oldFileName); err == nil {
		// 删除文件
		os.Remove(oldFileName)
	}

	os.Rename(selfPathWindow, oldFileName)
	os.Rename(exeResourceFilePath, selfPathWindow)

	// 结束自身运行然后重启自己
	cmd := exec.Command(selfPathWindow)
	cmd.Args = append(cmd.Args, os.Args[1:]...)
	cmd.Start()

	os.Exit(0) // 此处退出当前进程
	return true, ""
}

func CheckUpdateWindow() {
	downloadFolderPath := GetDownloadPath()
	info := GetGitRepoLatestReleaseInfo()
	if info == nil {
		zenity.Info("网络原因无法获取更新信息")
		return
	}
	if info.Version == UpdateCfg.CurrVersion {
		err := zenity.Info("当前已经是最新版本")
		if err != nil {
			return
		}
		return
	}

	err := zenity.Question("软件有新版本可用，是否更新？\n当前版本："+
		UpdateCfg.CurrVersion+
		"\n最新版本："+info.Version,
		zenity.Title("更新提示"),
		zenity.Icon(zenity.QuestionIcon),
		zenity.OKLabel("更新"),
		zenity.CancelLabel("取消"))
	//ecore.E调试输出(err)
	println("更新", err)
	if err != nil {
		return
	}
	progress, _ := zenity.Progress(
		zenity.Title("软件更新"),
		zenity.MaxValue(100), // 设置最大进度值为100
	)

	progress.Text("正在下载...")

	err = DownloadCallbackProgress(info.WinDownloadURL, downloadFolderPath+"/"+UpdateCfg.AppName+".exe", func(进度 float64) {
		fmt.Println("正在下载...", 进度)
		progress.Text("正在下载..." + fmt.Sprintf("%.2f", 进度) + "%")
		progress.Value(int(进度))
	})
	if err != nil {
		fmt.Println("下载出错：", err)
		zenity.Info("下载错误，检查你的网络")
		progress.Close()
		return
	}
	progress.Text("下载完成 即将完成更新")
	if progress.Close() != nil {
		fmt.Println("点击了取消")
		return
	}
	fmt.Println("下载完成了")
	flag, s := UpWindowSelfApp(downloadFolderPath + "/" + UpdateCfg.AppName + ".exe")
	println(flag, s)
}
