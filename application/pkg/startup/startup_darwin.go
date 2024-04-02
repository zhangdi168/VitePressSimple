// Package startup
// @Author: zhangdi
// @File: startup_drawin
// @Version: 1.0.0
// @Date: 2023/7/25 13:38
package startup

import (
	"fmt"
	"ideatools/application/pkgs/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func EnableAutoStart(appName string, appPath string) error {
	appPath = strings.ReplaceAll(appPath, ".exe", "")

	path := fmt.Sprintf("%s/Library/LaunchAgents/co.xiaod.%v.plist", utils.GetUserHomeDir(), appName)
	content := strings.ReplaceAll(macListFile, "{{AppName}}", appName)
	content = strings.ReplaceAll(content, "{{AppPath}}", appPath)
	content = strings.ReplaceAll(content, "{{AppDir}}", filepath.Dir(appPath))

	writer(true, path, content)
	return nil
}

func DisableAutoStart(appName string) error {
	path := fmt.Sprintf("%s/Library/LaunchAgents/co.xiaod.%v.plist", utils.GetUserHomeDir(), appName)
	content := strings.ReplaceAll(macListFile, "{{AppName}}", appName)
	writer(false, path, content)
	return nil
}

var macListFile = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>cn.xiaod.{{AppName}}</string>
	<key>ProgramArguments</key>
        <array>
        	<string>{{AppPath}}</string>
        </array>
	<key>RunAtLoad</key>
	<true/>
	<key>WorkingDirectory</key>
	<string>{{AppDir}}</string>
	<key>StandardErrorPath</key>
	<string>/tmp/down_tip.err</string>
	<key>StandardOutPath</key>
	<string>/tmp/down_tip.out</string>
</dict>
</plist>
`

//系统启动后会扫描下面这个目录的plist文件 ~/Library/LaunchAgents/
//cd ~/Library/LaunchAgents/
//xiaod@xiaoddeMac-mini LaunchAgents % ls -l
//-rwxr-xr-x@ 1 xiaod  staff  694  7 25 21:54 co.xiaod.ZhiJianChat.plist
//success example .plist file content

// <?xml version="1.0" encoding="UTF-8"?>
//
//	<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
//	<plist version="1.0">
//		<dict>
//			<key>Label</key>
//			<string>cn.xiaod.ZhiJianChat</string>
//			<key>ProgramArguments</key>
//			<array>
//				<string>/Users/xiaod/project/quickai/build/bin/ZhiJianChat.app/Contents/MacOS/ZhiJianChat</string>
//			</array>
//			<key>RunAtLoad</key>
//			<true/>
//			<key>WorkingDirectory</key>
//			<string>/Users/xiaod/project/quickai/build/bin/ZhiJianChat.app/Contents/MacOS</string>
//			<key>StandardErrorPath</key>
//			<string>/tmp/down_tip.err</string>
//			<key>StandardOutPath</key>
//			<string>/tmp/down_tip.out</string>
//		</dict>
//		</plist>
func writer(on bool, path, content string) error {
	var err error
	if on {
		stat, _ := os.Stat(path)
		if stat == nil {
			err = ioutil.WriteFile(path, []byte(content), os.ModePerm)
		}
	} else {
		err = os.Remove(path)
	}
	return err
}
