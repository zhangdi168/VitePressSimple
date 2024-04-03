package autoupdate

import (
	"fmt"
	"github.com/ncruces/zenity"
	"os/user"
	"path/filepath"
	"testing"
)

func TestSimple(t *testing.T) {
	CheckUpdateEntry()
}

func TestGetGithubRepoReleasesInfoAndChangelog(t *testing.T) {
	info := GetGitRepoLatestReleaseInfo()
	println(info)
}

func TestDownload(t *testing.T) {
	err := DownloadCallbackProgress("https://gitee.com/ranyuzhineng/qrpub-server-release/repository/blazearchive/1.0.zip?Expires=1705401575&Signature=SunOfUHT4G9PiCLL%2FI9L0YxTG%2BEBkL8dBwrHJWd3Flw%3D",
		"./GoEasyDesigner_MacOS.zip", func(progress float64) {
			fmt.Println("Downloading...", progress)
		})
	if err != nil {
		fmt.Println("Download error:", err)
	} else {
		fmt.Println("Download completed")
	}
}

func TestSystemVersion(t *testing.T) {
	fmt.Println("Is system Windows:", SysIsWindow())
	fmt.Println("Is system Linux:", SysIsLinux())
	fmt.Println("Is system Mac:", SysIsMac())
	fmt.Println("Is debug mode:", IsDebug())
}

func TestUnzip(t *testing.T) {
	archivePath := "/Users/ll/Downloads/GoEasyDesigner_MacOS.zip"
	UpdateSelfMacOSApp(archivePath, "GoEasyDesigner.app")
}
func TestUpdateProcessMacOS(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	downloadFolderPath := filepath.Join(usr.HomeDir, "Downloads")

	fmt.Println(downloadFolderPath)
	info := GetGithubRepoReleasesInfo()

	err = DownloadCallbackProgress(info.MacDownloadURL, downloadFolderPath+"/mactest.zip", func(progress float64) {
		fmt.Println("Downloading...", progress)
	})
	if err != nil {
		fmt.Println("Download error:", err)
		return
	}

	updateFlag, status := UpdateSelfMacOSApp(downloadFolderPath+"/mactest.zip", "qoq.app")
	fmt.Println(updateFlag, status)
}

func TestUpdateProcessWindows(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	downloadFolderPath := filepath.Join(usr.HomeDir, "Downloads")

	//fmt.Println(downloadFolderPath)
	//info := GetGithubRepoReleasesInfo()

	//err = DownloadWithProgressCallback(info.WinDownloadURL, downloadFolderPath+"/GoEasyDesigner.exe", func(progress float64) {
	//	fmt.Println("Downloading...", progress)
	//})
	//if err != nil {
	//	fmt.Println("Download error:", err)
	//	return
	//}

	updateFlag, status := UpWindowSelfApp(downloadFolderPath + "/GoEasyDesigner.exe")
	fmt.Println(updateFlag, status)
}

func TestFullUpdateFlow(t *testing.T) {
	downloadFolderDir := GetDownloadPath()
	info := GetGitRepoLatestReleaseInfo()
	fmt.Println(info.MacDownloadURL)
	fmt.Println(downloadFolderDir)
	if info.Version == UpdateCfg.CurrVersion {
		err := zenity.Info("Already on the latest version.")
		if err != nil {
			return
		}
		return
	}

	err := zenity.Question(fmt.Sprintf("A new software version is available. Would you like to autoupdate?\nCurrent version: %s\nLatest version: %s", UpdateCfg.CurrVersion, info.Version),
		zenity.Title("Update Prompt"),
		zenity.Icon(zenity.QuestionIcon),
		zenity.OKLabel("Update"),
		zenity.CancelLabel("Cancel"))

	fmt.Println("Update:", err)
	if err != nil {
		return
	}
	progress, _ := zenity.Progress(
		zenity.Title("Software Update"),
		zenity.MaxValue(100), // Set maximum progress value to 100
	)

	progress.Text("Downloading...")

	err = DownloadCallbackProgress(info.WinDownloadURL, "./win.zip", func(progressVal float64) {
		fmt.Println("Downloading...", fmt.Sprintf("%.2f", progressVal)+"%")
		progress.Text("Downloading..." + fmt.Sprintf("%.2f", progressVal) + "%")
		progress.Value(int(progressVal))
	})
	if err != nil {
		fmt.Println("Download error:", err)
		zenity.Info("Download error. Please check your network connection.")
		progress.Close()
		return
	}
	progress.Text("Download complete. Finishing autoupdate...")
	if progress.Close() != nil {
		fmt.Println("Cancelled by user")
		return
	}
	fmt.Println("Download completed")
	//updateFlag, status := UpdateSelfMacOSApp(downloadFolderDir+"/"+appName+"_MacOS.zip", appName+".app")
	//fmt.Println(updateFlag, status)
}
