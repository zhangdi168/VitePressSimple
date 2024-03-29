package vpsimpler

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
)

type VpManager struct {
	//fs
	fs embed.FS
}

func NewVpManager(fs_ embed.FS) *VpManager {
	return &VpManager{fs: fs_}
}

func (s *VpManager) CreateProject(title, description, dir string) string {
	//1.检查node环境
	//if !NodejsIsInstall() {
	//	return "环境监测不通过：nodejs is not install"
	//}
	//if !NpmIsInstall() {
	//	return "环境监测不通过：npm is not install"
	//}

	//2.获取node版本
	if !strings.Contains(GetNodeVersion(), "v18") {
		return "环境监测不通过：nodejs 版本必须是v18及以上版本"
	}

	//3.检查目录
	if !filehelper.FileExists(dir) {
		return "目录检查不通过,目录不存在：" + dir
	}

	//4.复制模板文件到目标目录
	s.CopyTemplateFile(dir)

	//5.设置当前项目目录为新建的目录
	cfg.Set(keys.ConfigKeyProjectDir, dir)

	//6.检查是否复制成功：看看配置文件是否存在
	vpConfig := NewVpConfig()
	_, err := vpConfig.GetConfigPath()
	if err != nil {
		return err.Error()
	}

	//7.获取默认的配置文件内容替换标题描述
	defaultConfigContent := vpConfig.ReplaceDefaultConfigContent(title, description)

	//8.保存配置文件
	vpConfig.SaveConfig(defaultConfigContent)

	return ""
}

// CopyTemplateFile 复制模板文件到新建项目
func (s *VpManager) CopyTemplateFile(targetDir string) {
	err := fs.WalkDir(s.fs, "vitepress-template", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil // 跳过目录，仅处理文件
		}

		// 打开嵌入的文件
		fileData, err := s.fs.ReadFile(path)
		if err != nil {
			return err
		}

		// 计算目标路径（假设目标目录为 "./output"）
		targetPath := filepath.Join(targetDir + path[len("vitepress-template"):])
		// 创建包含所有上级目录的本地文件
		err = os.MkdirAll(filepath.Dir(targetPath), 0755)
		if err != nil {
			return err
		}

		// 创建本地文件用于写入嵌入的文件内容
		localFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer localFile.Close()

		// 将嵌入的文件内容写入到本地文件
		_, err = localFile.Write(fileData)
		if err != nil {
			return err
		}

		// 确保所有数据已刷新到磁盘
		err = localFile.Sync()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

// GetNodeVersion 获取nodejs版本，如 v18.14.2
func (s *VpManager) GetNodeVersion() string {
	return GetNodeVersion()
}

// GetNpmVersion 获取npm版本，如 8.19.2
func (s *VpManager) GetNpmVersion() string {
	return GetNpmVersion()
}
