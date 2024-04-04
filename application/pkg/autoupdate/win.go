package autoupdate

import (
	"fmt"
	"github.com/ncruces/zenity"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"wailstemplate/application/pkg/mylog"
)

// GetAppPathWindow 获取当前可执行文件的路径。
//
// 返回值:
// string: 可执行文件的完整路径。
// error: 如果获取过程中出现错误，则返回错误信息；否则返回nil。
func GetAppPathWindow() (string, error) {
	// 获取当前执行文件的路径
	exe, err := os.Executable()
	if err != nil {
		// 如果获取路径过程中出现错误，返回空字符串和错误信息
		return "", err
	}
	// 正常获取到路径后，返回该路径和nil作为错误信息
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
		err := os.Remove(oldFileName)
		if err != nil {
			mylog.Error("删除旧文件error ", err)
			return false, ""
		}
	}

	err := os.Rename(selfPathWindow, oldFileName)
	if err != nil {
		mylog.Error("修改旧文件路径error ", err)
		return false, err.Error()
	}
	//不能用重命名的方式移动，跨磁盘会报错
	err = CopyFile(exeResourceFilePath, selfPathWindow)
	if err != nil {
		mylog.Error("复制新文件路径error ", err)
		return false, err.Error()
	}

	// 结束自身运行然后重启自己
	cmd := exec.Command(selfPathWindow)
	cmd.Args = append(cmd.Args, os.Args[1:]...)
	err = cmd.Start()
	if err != nil {
		mylog.Error("重启程序error ", err)
		return false, ""
	}

	os.Exit(0) // 此处退出当前进程
	return true, ""
}

func checkUpdateWindow() {
	downloadDir := getDownloadDir()
	info := GetGitRepoLatestReleaseInfo()
	if info == nil {
		zenity.Info("Unable to retrieve update information due to network issues")
		return
	}
	if info.Version == UpdateCfg.CurrVersion {
		err := zenity.Info("Already on the latest version")
		if err != nil {
			return
		}
		return
	}
	fileName := filepath.Base(info.WinDownloadURL)
	saveFilePath := filepath.Join(downloadDir, fileName)
	err := zenity.Question("A new software version is available. Would you like to update?\nCurrent Version: "+
		UpdateCfg.CurrVersion+
		"\nLatest Version: "+info.Version,
		zenity.Title("Update Prompt"),
		zenity.Icon(zenity.QuestionIcon),
		zenity.OKLabel("Update"),
		zenity.CancelLabel("Cancel"))
	// ecore.E debug output(err)
	println("Update:", err)
	if err != nil {
		return
	}
	progress, _ := zenity.Progress(
		zenity.Title("Software Update"),
		zenity.MaxValue(100), // Set maximum progress value to 100
	)

	progress.Text("Downloading...")

	err = DownloadCallbackProgress(info.WinDownloadURL, saveFilePath, func(currProgress float64) {
		fmt.Println("Downloading...", currProgress)
		progress.Text("Downloading..." + fmt.Sprintf("%.2f", currProgress) + "%")
		progress.Value(int(currProgress))
	})
	if err != nil {
		fmt.Println("Download failed:", err)
		zenity.Info("Download error. Please check your network connection")
		progress.Close()
		return
	}
	progress.Text("Download complete. Updating soon...")
	if progress.Close() != nil {
		fmt.Println("Cancelled by user")
		return
	}
	fmt.Println("Download completed")
	// Unzip

	if strings.Contains(fileName, ".zip") {
		err := UnzipFile(filepath.Join(downloadDir, fileName), downloadDir)
		if err != nil {
			zenity.Info("UnzipFile error:" + err.Error())
			mylog.Err("解压更新包失败:", err)
			return
		}
		
	}
	// Perform the update operation
	flag, s := UpWindowSelfApp(downloadDir + "/" + UpdateCfg.AppName + ".exe")
	println(flag, s)
}
