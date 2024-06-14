package main

import (
	"embed"
	"log"
	"os"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/first"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/utils"
	"wailstemplate/application/services"
	"wailstemplate/application/services/system"
	"wailstemplate/application/vitepress/vpsimpler"
	setting "wailstemplate/settings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

//go:embed all:resources
var resource embed.FS

//go:embed all:vitepress-template
var vpFs embed.FS

func main() {

	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	first.InitConfig()         //初始化配置文件
	first.InitDefaultConfig()  //初始化默认配置
	if !first.CheckCanOpen() { //重复打开检测
		os.Exit(0)
	}
	first.InitLog()      //初始化日志
	first.CheckStartup() //开机自启检查
	defer onCloseDo()
	//first.InitTask()   //初始化定时任务
	//静态文件服务器,如果需要请取消注释 localhost:9875/images/icon.pmh
	//mystatic.StartStaticServer("9874", resource, "resources")

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:  "VitePressSimple " + setting.Version,
		Width:  1080,
		Height: 720,
		//MinWidth:  900,
		//MinHeight: 600,
		//MaxWidth:          1200,
		//MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Maximised,
		CSSDragProperty:   utils.IfToString(utils.IsMacOS(), "--wails-draggable", ""),
		CSSDragValue:      utils.IfToString(utils.IsMacOS(), "drag", "drag-win"),
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			app,
			services.NewTreeData(),
			vpsimpler.NewVpManager(vpFs),
			vpsimpler.NewVpConfig(),
			system.NewSystemService(),
			services.NewShellManager(),
			services.NewUpdateService(), //在线更新
			services.NewStaticServer(),
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
		},

		// Mac platform specific options
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            true,
				UseToolbar:                 false,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Wails Template Vue",
				Message: "A Wails template based on Vue and Vue-Router",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func onCloseDo() {
	cfg.Set(keys.ConfigKeySysProgramIsOpen, "no") //打开状态设置成关闭
}
