package services

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/mylog"
)

type StaticServer struct {
	server   *http.Server
	serverMu sync.Mutex // 保护server字段的互斥锁
}

func NewStaticServer() *StaticServer {
	return &StaticServer{server: nil}
}

// StartStaticServer 启动VitePress项目静态资源服务器
func (s *StaticServer) StartStaticServer(staticDir string) {
	go func() {
		s.serverMu.Lock()
		defer s.serverMu.Unlock()

		port := cfg.GetStringDefault(keys.ConfigKeySysStaticServerPort, "9874")
		if s.server != nil {
			if err := s.server.Shutdown(context.Background()); err != nil {
				mylog.Error(err.Error())
			}
			mylog.Info("切换静态资源路径并重启http服务目录：" + staticDir)
		}

		router := http.NewServeMux()
		router.Handle("/", http.FileServer(http.Dir(staticDir)))
		//图片上传接口
		router.HandleFunc("/upload_image", handleImageUpload)

		s.server = &http.Server{
			Addr:    ":" + port,
			Handler: router, // 你的路由处理程序
		}

		if err := s.server.ListenAndServe(); err != nil {
			mylog.Err("启动http服务失败", err)
		}
	}()
}

// GetStaticBaseUrl 获取静态资源基础路径带端口
func (s *StaticServer) GetStaticBaseUrl() string {
	port := cfg.GetStringDefault(keys.ConfigKeySysStaticServerPort, "9874")
	return "http://localhost:" + port
}

func (s *StaticServer) GetProjectBaseDir() string {
	return cfg.GetStringDefault(keys.ConfigKeyProjectDir, "")
}

// UploadFile 上传图片
func (s *StaticServer) UploadFile() {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()

}

const (
	// 图片上传目录
	uploadDir = "./images"
	// 最大允许的单个文件大小（以字节为单位）
	maxFileSize = 10 << 20 // 10 MB
)

func handleImageUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 检查Content-Type是否为multipart/form-data
	if r.Header.Get("Content-Type") != "multipart/form-data" {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return
	}

	// 创建一个multipart解析器
	r.ParseMultipartForm(maxFileSize)

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to retrieve file from request").Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 检查文件类型是否为允许的图片类型
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !isAllowedImageType(ext) {
		http.Error(w, fmt.Sprintf("Unsupported file type: %s", ext), http.StatusBadRequest)
		return
	}

	// 生成唯一的文件名（可包含原始扩展名）
	uniqFilename := generateUniqueFilename(header.Filename)

	// 计算上传路径
	uploadPath := filepath.Join(uploadDir, uniqFilename)

	// 创建文件并写入图片数据
	out, err := os.Create(uploadPath)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to create file on disk").Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to copy file data").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Image uploaded successfully: %s", uploadPath)
}

func isAllowedImageType(ext string) bool {
	allowedTypes := []string{".jpg", ".jpeg", ".png", ".gif"}
	for _, t := range allowedTypes {
		if ext == t {
			return true
		}
	}
	return false
}

func generateUniqueFilename(originalName string) string {
	// 这里可以实现生成唯一文件名的逻辑，例如结合时间戳、随机数等
	// 示例简化为仅保留原始文件名：
	// return originalName
	// 实际应用中应确保生成的文件名在服务器上是唯一的
	return ""
}
