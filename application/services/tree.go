package services

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"wailstemplate/application/constant/keys"
	"wailstemplate/application/dto"
	"wailstemplate/application/pkg/cfg"
	"wailstemplate/application/pkg/filehelper"
	"wailstemplate/application/pkg/mylog"
)

type ArticleTreeData struct {
	currFilePath string //当前正在操作的文件
}

func NewTreeData() *ArticleTreeData {
	return &ArticleTreeData{}
}

// ParseTreeData 解析文件夹
func (s *ArticleTreeData) ParseTreeData(srcDoc string) []*dto.TreeNode {
	dir := cfg.GetString(keys.ConfigKeyProjectDir)
	if dir == "" {
		return nil
	}
	data, err := DirToTreeNode(filepath.Join(dir, srcDoc))
	if err != nil {
		mylog.Err("解析文件夹树", err)
		return nil
	}
	return data
}

// DirToTreeNode 将指定目录转换为树节点结构
//
// rootPath: 指定的目录路径
// 返回值: 树结构的根节点切片和可能发生的错误
func DirToTreeNode(rootPath string) ([]*dto.TreeNode, error) {
	// 定义递归构建树函数
	var buildTree func(path string, parent *dto.TreeNode) error

	buildTree = func(path string, parent *dto.TreeNode) error {
		// 读取目录内容
		files, err := os.ReadDir(path)
		if err != nil {
			return err
		}

		// 按照类型排序（目录在前，文件在后）
		sort.Slice(files, func(i, j int) bool {
			return files[i].IsDir() && !files[j].IsDir()
		})

		// 遍历目录中的每个条目

		for _, entry := range files {
			// 创建子节点
			childNode := &dto.TreeNode{
				Title: entry.Name(),
				Key:   filepath.Join(parent.Key, entry.Name()),
				//Class: "text-base",
			}

			parent.Children = append(parent.Children, childNode) // 尾部插入
			// 如果条目是目录，则递归构建子树
			if entry.IsDir() {
				//parent.Children = insertToFront(parent.Children, childNode) //文件夹 头部插入
				if err := buildTree(filepath.Join(path, entry.Name()), childNode); err != nil {
					return err
				}
			}
		}

		return nil
	}

	// 创建根节点
	root := &dto.TreeNode{
		Title: filepath.Base(rootPath),
		Key:   rootPath,
	}
	// 使用递归函数构建树
	if err := buildTree(rootPath, root); err != nil {
		return nil, err
	}

	// 返回根节点的子节点
	return root.Children, nil
}

// ReadFileContent 读取文件的内容
func (s *ArticleTreeData) ReadFileContent(filePath string) string {
	s.currFilePath = filePath
	content, err := filehelper.ReadContent(filePath)
	if err != nil {
		return ""
	}
	return content
}

// WriteFileContent 将文件保存到指定路径
func (s *ArticleTreeData) WriteFileContent(filePath string, content string) string {
	s.currFilePath = filePath
	err := filehelper.WriteContent(filePath, content)
	if err != nil {
		return err.Error()
	}
	return ""
}

// Rename 重命名
func (s *ArticleTreeData) Rename(oriPath string, newPath string) string {
	err := filehelper.RenamePath(oriPath, newPath)
	if err != nil {
		return err.Error()
	}
	return ""
}

// Move 移动文件
func (s *ArticleTreeData) Move(srcPath string, destDir string) string {
	err := filehelper.MovePath(srcPath, destDir)
	if err != nil {
		return err.Error()
	}
	return ""
}

// CreateDir 创建文件夹
func (s *ArticleTreeData) CreateDir(dir string) string {
	return filehelper.CreateDir(filepath.Join(dir))
}

// CreateFile  创建文件
func (s *ArticleTreeData) CreateFile(fullPath string) string {
	filehelper.CreateDir(filepath.Dir(fullPath))
	// 使用os.Create函数创建或打开并清空文件
	file, err := os.Create(filepath.Join(fullPath))
	if err != nil {
		return fmt.Sprintf("Error creating file:%s", err.Error())
	}
	defer file.Close() // 在函数结束时确保关闭文件
	// 写入内容到文件（如果需要的话）
	content := []byte("# Hello vitrpress <br /> > this is a  demo file created by vitepress simple!")
	if _, err := file.Write(content); err != nil {
		return "Error writing to file:" + err.Error()
	}
	return ""
}

// DeletePath 删除文件或文件夹
func (s *ArticleTreeData) DeletePath(filePath string) string {
	err := os.RemoveAll(filepath.Join(filePath))
	if err != nil {
		return err.Error()
	}
	return ""
}

// 从前面插入数据
func insertToFront(slice []*dto.TreeNode, node *dto.TreeNode) []*dto.TreeNode {
	// 创建一个新切片，长度为原切片长度加1
	newSlice := make([]*dto.TreeNode, len(slice)+1)

	// 将原切片内容从后往前拷贝到新切片，并将新值插入新切片的首位
	newSlice[0] = node
	copy(newSlice[1:], slice)

	return newSlice
}
