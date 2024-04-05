package autoupdate

import (
	"fmt"
	"github.com/ncruces/zenity"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/pkg/mylog"
)

func checkUpdateMac() {
	downloadFolderPath := getDownloadDir()
	info := GetGitRepoLatestReleaseInfo()
	println(info.MacDownloadURL)
	println(downloadFolderPath)
	if info.Version == UpdateCfg.CurrVersion {
		err := zenity.Info("当前已经是最新版本")
		if err != nil {
			return
		}
		return
	}
	fileName := filepath.Base(info.MacDownloadURL) //xx.dmg
	saveDownloadFilePath := filepath.Join(downloadFolderPath, fileName)
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
	if !filehelper.FileExists(saveDownloadFilePath) {
		progress, _ := zenity.Progress(
			zenity.Title("软件更新"),
			zenity.MaxValue(100), // 设置最大进度值为100
		)

		progress.Text("正在下载...")

		err = DownloadCallbackProgress(info.MacDownloadURL, saveDownloadFilePath, func(currProgress float64) {
			fmt.Println("正在下载...", currProgress)
			progress.Text("正在下载..." + fmt.Sprintf("%.2f", currProgress) + "%")
			progress.Value(int(currProgress))
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
	}

	//打开dmg文件
	if filepath.Ext(saveDownloadFilePath) == ".dmg" {
		cmd := exec.Command("open", saveDownloadFilePath)
		err := cmd.Run()
		if err != nil {
			mylog.Error("Failed to open DMG using 'open' command: %v", err)
		}
	} else {
		mylog.Error(saveDownloadFilePath + "不是dmg文件")
	}

	//flag, s := UpdateSelfMacOSApp(downloadFolderPath+"/"+UpdateCfg.AppName+"_MacOS.zip", UpdateCfg.AppName+".app")
	//println(flag, s)
}

func GetMacOsAppPath() string {
	// 如果不处于编译状态反馈空
	compiledPath, err := os.Executable()
	//println("compiledPath", compiledPath)
	if err != nil {
		return ""
	}
	// 调试模式下的路径
	//compiledPath = "/Users/ll/Desktop/goproject/v3fanyi/bin/qoq.app/Contents/MacOS/qoq"
	// compiledPath 取父目录
	compiledPath = compiledPath[:strings.LastIndex(compiledPath, "/")]

	appDir := compiledPath[:strings.LastIndex(compiledPath, "/")]
	appDir = appDir[:strings.LastIndex(appDir, "/")]
	FatherDir := compiledPath[strings.LastIndex(compiledPath, "/")+1:]
	if FatherDir == "MacOS" {
		return appDir
	} else {
		return ""
	}
}

func UpdateSelfMacOSApp(ResourceCompressionPackage string, appName string) (bool, string) {
	//ResourceCompressionPackage = "/Users/ll/Downloads/GoEasyDesigner_MacOS.zip"
	//appName = "GoEasyDesigner.app"
	//println("ResourceCompressionPackage", ResourceCompressionPackage)
	//println("appName", appName)

	MacOsAppPath := GetMacOsAppPath()
	//fmt.Println("MacOsAppPath", MacOsAppPath)
	//MacOsAppPath := "/Users/ll/Downloads/" + appName
	if MacOsAppPath != "" {
		appFatherDir := MacOsAppPath[:strings.LastIndex(MacOsAppPath, "/")]
		fmt.Printf("ResourceCompressionPackage %s appFatherDir%s MacOsAppPath%s \r\n", ResourceCompressionPackage, appFatherDir, MacOsAppPath)
		if MacOsAppPath != "" {
			tarResult := Unzip(ResourceCompressionPackage, appFatherDir, []string{appName + "/Contents/"})
			println("tarResult", tarResult)
			// 解压完成后删除压缩包
			//os.Remove(ResourceCompressionPackage)
			MacOsAppPath = filepath.Join(appFatherDir, appName)
			// QApplication.quit()
			appName = appName[:strings.LastIndex(appName, ".")]
			runCmd := fmt.Sprintf("killall %s && open -n -a %s", appName, MacOsAppPath)
			println("runCmd", runCmd)
			cmd := exec.Command("sh", "-c", runCmd)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
				return false, "killall失败 请自己重启应用"
			}
			return true, MacOsAppPath
		}
	} else {
		fmt.Println("非MacOS编译环境")
		return false, "非MacOS编译环境"
	}
	fmt.Println("应用路径为空")
	return false, "应用路径为空"
}
