package services

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
	"wailstemplate/application/pkg/vpsimpler"
)

type VitePress struct {
	//fs
	fs embed.FS
}

func NewVitePress(fs_ embed.FS) *VitePress {
	return &VitePress{fs: fs_}
}

func (s *VitePress) CreateProject() {

}

// CopyTemplateFile 复制模板文件到新建项目
func (s *VitePress) CopyTemplateFile(targetDir string) {
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
func (s *VitePress) GetNodeVersion() string {
	return vpsimpler.GetNodeVersion()
}

// GetNpmVersion 获取npm版本，如 8.19.2
func (s *VitePress) GetNpmVersion() string {
	return vpsimpler.GetNpmVersion()
}
