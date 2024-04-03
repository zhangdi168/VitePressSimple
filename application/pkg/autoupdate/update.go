package autoupdate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type ReleaseInfo struct {
	Version        string   `json:"版本号"`
	DownloadURLs   []string `json:"下载地址列表"`
	MacDownloadURL string   `json:"mac下载地址"`
	WinDownloadURL string   `json:"win下载地址"`
	Changelog      string   `json:"更新内容"`
	ReleaseTime    string   `json:"发布时间"`
}

// GetGitRepoLatestReleaseInfo 获取最新版本发行版的下载地址等信息
func GetGitRepoLatestReleaseInfo() *ReleaseInfo {
	if UpdateCfg.GitType == GitTypeGitee { //gitee
		return GetGiteeRepoReleasesInfo()
	}
	if UpdateCfg.GitType == GitTypeGithub { //github
		return GetGithubRepoReleasesInfo()
	}
	return nil
}

func GetGithubRepoReleasesInfo() *ReleaseInfo {
	//owner := "duolabmeng6"   // GitHub 仓库的所有者
	//repo := "GoEasyDesigner" // GitHub 仓库的名称
	//https://api.github.com/repos/duolabmeng6/projection_screen_tv/releases
	var urls []string
	_url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", UpdateCfg.CurrVersion, UpdateCfg.GitRepo)
	urls = append(urls, _url)
	_url = fmt.Sprintf("https://go.kenhong.com/releases_latest.json")
	urls = append(urls, _url)
	// 访问urls中的内容 如果第一个访问失败了 就继续下一个
	var resp *http.Response
	var err error
	for _, url := range urls {
		resp, err = http.Get(url)
		if err == nil {
			if resp.StatusCode != http.StatusOK {
				fmt.Println("请求失败:", resp.Status, url)
				continue
			}
			break
		}
		fmt.Println("请求失败:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil
	}
	//var err error
	//body := ecore.E读入文件("/Users/ll/Desktop/goproject/v3fanyi/mymodel/githubdata.json")

	var releases GithubJSONData

	err = json.Unmarshal(body, &releases)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		return nil
	}
	if len(releases) < 0 {
		return nil
	}

	//latestRelease := releases[0]
	//version := latestRelease.TagName
	//fmt.Println("最新的 Release 版本号:", version)
	//fmt.Println("最新的 Release 版本号:", latestRelease.Body)
	//// 循环获取 文件名和下载地址
	//for _, asset := range latestRelease.Assets {
	//	fmt.Println(asset.Name, asset.BrowserDownloadURL)
	//}

	latestRelease := releases[0]
	version := latestRelease.TagName
	downloadURLs := make([]string, 0)
	macDownloadURL := ""
	winDownloadURL := ""

	// 获取文件名和下载地址
	for _, asset := range latestRelease.Assets {
		downloadURLs = append(downloadURLs, asset.BrowserDownloadURL)
		//查找文本是否包含  macos.zip 和 .exe 不区分大小写
		if strings.Contains(strings.ToLower(asset.Name), "macos.zip") {
			macDownloadURL = asset.BrowserDownloadURL
		}
		if strings.Contains(strings.ToLower(asset.Name), ".exe") {
			winDownloadURL = asset.BrowserDownloadURL
		}

	}

	releaseInfo := &ReleaseInfo{
		Version:        version,
		DownloadURLs:   downloadURLs,
		MacDownloadURL: "https://ghproxy.com/" + macDownloadURL,
		WinDownloadURL: "https://ghproxy.com/" + winDownloadURL,
		Changelog:      latestRelease.Body,
		ReleaseTime:    latestRelease.CreatedAt.String(),
	}

	return releaseInfo
}

func GetGiteeRepoReleasesInfo() *ReleaseInfo {
	//owner := "duolabmeng6"   // GitHub 仓库的所有者
	//repo := "GoEasyDesigner" // GitHub 仓库的名称
	//https://api.github.com/repos/duolabmeng6/projection_screen_tv/releases
	var urls []string
	_url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/releases", UpdateCfg.GitOwner, UpdateCfg.GitRepo)
	urls = append(urls, _url)
	//_url = fmt.Sprintf("https://go.kenhong.com/releases_latest.json")
	//urls = append(urls, _url)
	// 访问urls中的内容 如果第一个访问失败了 就继续下一个
	var resp *http.Response
	var err error
	for _, url := range urls {
		resp, err = http.Get(url)
		if err == nil {
			if resp.StatusCode != http.StatusOK {
				fmt.Println("请求失败:", resp.Status, url)
				continue
			}
			break
		}
		fmt.Println("请求失败:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return nil
	}
	//var err error
	//body := ecore.E读入文件("/Users/ll/Desktop/goproject/v3fanyi/mymodel/githubdata.json")

	var releases GiteeRelease

	err = json.Unmarshal(body, &releases)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		return nil
	}
	if len(releases) < 0 {
		return nil
	}

	//latestRelease := releases[0]
	//version := latestRelease.TagName
	//fmt.Println("最新的 Release 版本号:", version)
	//fmt.Println("最新的 Release 版本号:", latestRelease.Body)
	//// 循环获取 文件名和下载地址
	//for _, asset := range latestRelease.Assets {
	//	fmt.Println(asset.Name, asset.BrowserDownloadURL)
	//}

	//gitee 发型版列表排序是创建顺序来的 越新的越后面
	latestRelease := releases[len(releases)-1]
	version := latestRelease.TagName
	downloadURLs := make([]string, 0)
	macDownloadURL := ""
	winDownloadURL := ""

	// 获取文件名和下载地址
	for _, asset := range latestRelease.Assets {
		downloadURLs = append(downloadURLs, asset.BrowserDownloadURL)
		//查找文本是否包含  macos.zip 和 .exe 不区分大小写
		if strings.Contains(strings.ToLower(asset.Name), UpdateCfg.MacReleaseNameContainStr) {
			macDownloadURL = asset.BrowserDownloadURL
		}
		if strings.Contains(strings.ToLower(asset.Name), UpdateCfg.WindowReleaseNameContainStr) {
			winDownloadURL = asset.BrowserDownloadURL
		}

	}

	releaseInfo := &ReleaseInfo{
		Version:        version,
		DownloadURLs:   downloadURLs,
		MacDownloadURL: macDownloadURL,
		WinDownloadURL: winDownloadURL,
		Changelog:      latestRelease.Body,
		ReleaseTime:    latestRelease.CreatedAt.String(),
	}

	return releaseInfo
}

// DownloadCallbackProgress 下载进度回调
func DownloadCallbackProgress(downloadUrl string, savePath string, fc func(progress float64)) error {
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer out.Close()

	DownloadFileSize := resp.ContentLength
	currDownloadSize := int64(0)
	buf := make([]byte, 1024)

	progressTicker := time.Tick(time.Millisecond * 50) //降低刷新的频率

	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			_, err := out.Write(buf[:n])
			if err != nil {
				return err
			}

			currDownloadSize += int64(n)
			progress_ := float64(currDownloadSize) / float64(DownloadFileSize) * 100

			select {
			case <-progressTicker:
				if fc != nil {
					fc(progress_) // 调用进度回调函数
				}
			default:
				// 不触发进度回调
			}
		}

		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	fc(100) // 调用进度回调函数

	return nil
}

// IsDebug
// 判断当前的环境是否为编译后的程序
// 真为调试模式
// 假为编译后的程序
func IsDebug() bool {
	debugMode := false
	if os.Getenv("DEBUG") != "" {
		debugMode = true
	}

	return debugMode
}

// CheckUpdateEntry 检查更新的入口函数
func CheckUpdateEntry() {
	if SysIsMac() {
		CheckUpdateMac()
	}
	if SysIsWindow() {
		CheckUpdateWindow()
	}
}
