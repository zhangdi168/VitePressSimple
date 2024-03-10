package mystatic

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
	"wails3templatevuets/app/mypkg/mylog"
)

func staticHandler(fs http.FileSystem) gin.HandlerFunc {
	return func(c *gin.Context) {
		//filePath := c.Param("filepath")
		http.FileServer(fs).ServeHTTP(c.Writer, c.Request)
	}
}

func StartStaticServer(port string, fsStatic embed.FS, topDirName string) {
	router := gin.Default()

	// 创建一个自定义FS，将资源路径映射到正确的位置
	resourceFS, _ := fs.Sub(fsStatic, topDirName)

	// 注册静态文件服务中间件，指定目录为cdn，并从resourceFS中提供文件
	//*filepath 是一个通配符参数，它会匹配任何在 /cdn 后面跟随的任意路径。这里的 filepath 是个变量名，你可以自定义这个名称（只要在处理函数中保持一致即可），它的值将包含实际请求的路径。
	dirs, err := listFirstLevelDirs(fsStatic, topDirName)
	if err != nil {
		mylog.Error("listFirstLevelDirs", err)
	} else {
		for _, dir := range dirs {
			//router.GET("/cdn/*filepath", staticHandler(http.FS(resourceFS)))
			router.GET("/"+dir+"/*filepath", staticHandler(http.FS(resourceFS)))
		}
	}

	// 启动服务器
	go func() {
		err = router.Run(":" + port)
		if err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()
}

// 获取embed.fs的所有第一级目录
func listFirstLevelDirs(fsStatic embed.FS, topDir string) ([]string, error) {
	var dirs []string

	// 获取根目录下的所有条目
	entries, err := fs.ReadDir(fsStatic, topDir)
	if err != nil {
		return nil, err
	}

	// 过滤出第一级目录
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}
