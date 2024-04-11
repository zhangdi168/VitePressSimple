package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
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

		if staticDir == "" {
			staticDir = s.GetStaticFullDir()
		}
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
		router.HandleFunc("/upload_image", s.handleImageUpload)

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

// GetStaticFullDir 获取静态资源目录
func (s *StaticServer) GetStaticFullDir() string {
	projectDir := cfg.GetStringDefault(keys.ConfigKeyProjectDir, "./")
	staticDirName := cfg.GetStringDefault(keys.ConfigKeySysProjectStaticDirName, "vpstatic")
	fullDir := filepath.Join(projectDir, staticDirName)
	filehelper.CreateDir(fullDir)
	return fullDir
}

// ConvertFullPathToUrl 将一个完整的文件路径转换成http url
func (s *StaticServer) ConvertFullPathToUrl(fullPath string) string {
	staticFullDir := s.GetStaticFullDir()
	path := strings.ReplaceAll(fullPath, staticFullDir, "")
	path = strings.ReplaceAll(path, "\\", "/")
	return s.GetStaticBaseUrl() + path
}

func (s *StaticServer) GetStaticImagesFullDir() string {
	fullDir := filepath.Join(s.GetStaticFullDir(), "images")
	filehelper.CreateDir(fullDir)
	return fullDir
}

func (s *StaticServer) GetProjectStaticName() string {
	return cfg.GetStringDefault(keys.ConfigKeySysProjectStaticDirName, "")
}

// UploadFile 上传图片
func (s *StaticServer) UploadFile() {
	s.serverMu.Lock()
	defer s.serverMu.Unlock()

}

const (
	// 图片上传目录

	// 最大允许的单个文件大小（以字节为单位）
	maxFileSize = 10 << 20 // 10 MB
)

// UploadResponseData 定义响应结构体
type UploadResponseData struct {
	ErrFiles []string          `json:"errFiles"`
	SuccMap  map[string]string `json:"succMap"`
}

type ImageUploadResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data UploadResponseData `json:"data"`
}

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400") // 缓存预检结果，避免频繁发送OPTIONS请求
}

func (s *StaticServer) handleImageUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		allowCors(w)
		w.WriteHeader(http.StatusOK)
		return
	}
	allowCors(w)
	// 检查Content-Type是否为multipart/form-data
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "multipart/form-data") {
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

	var response ImageUploadResponse
	response.Code = 0
	response.Msg = ""

	// 初始化数据结构体
	response.Data.ErrFiles = make([]string, 0)
	response.Data.SuccMap = make(map[string]string)
	errFiles := []string{}
	succMap := make(map[string]string)

	// 生成唯一的文件名（可包含原始扩展名）
	uniqFilename := generateUniqueFilename(header.Filename)

	// 计算上传路径
	uploadDir := s.GetStaticImagesFullDir()
	uploadPath := filepath.Join(uploadDir, time.Now().Format("20060102"), uniqFilename+ext)
	filehelper.CreateDir(filepath.Dir(uploadPath))
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

	//组装返回数据
	succMap[filepath.Base(uploadPath)] = fmt.Sprintf("%s", s.ConvertFullPathToUrl(uploadPath))
	response.Data.ErrFiles = errFiles
	response.Data.SuccMap = succMap
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, errors.Wrap(err, "failed to marshal JSON response").Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

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

	return uuid.New().String()
}
