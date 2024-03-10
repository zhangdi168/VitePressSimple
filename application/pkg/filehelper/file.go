package filehelper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateFileIfNotExists 判断文件不存在则进行创建
func CreateFileIfNotExists(filePath string) error {
	// 判断目录是否存在
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，进行创建
		if err := os.MkdirAll(dirPath, 0766); err != nil {
			return err
		}
	}

	// 判断文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，进行创建
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close() // 确保文件关闭
	}

	return nil
}

// ReadContent reads the content of a file and returns it as a string.
func ReadContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(byteValue), nil
}

// WriteContent writes data to a file.
func WriteContent(filePath string, data string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

// RenamePath 用于重命名文件或文件夹
func RenamePath(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// MovePath 移动文件或文件夹到新的文件夹
func MovePath(srcPath string, destDir string) error {
	// 检查源路径是否存在
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("source path does not exist: %w", err)
	}

	// 获取源路径的基本名称
	baseName := filepath.Base(srcPath)

	// 构建目标路径
	destPath := filepath.Join(destDir, baseName)

	// 对于文件或空目录，可以直接使用 os.Rename
	err := os.Rename(srcPath, destPath)
	if err == nil || os.IsExist(err) { // 如果目标已经存在且为文件或空目录，也可以接受
		return nil
	} else if !os.IsNotExist(err) { // 非“不存在”错误则返回
		return err
	}

	// 对于非空目录，需要递归移动
	return moveAll(srcPath, destPath)
}

// moveAll 递归移动非空目录内容
func moveAll(srcPath string, destPath string) error {
	// 创建目标目录
	err := os.MkdirAll(destPath, 0755)
	if err != nil {
		return err
	}

	directory, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer directory.Close()

	entries, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcEntryPath := filepath.Join(srcPath, entry.Name())
		destEntryPath := filepath.Join(destPath, entry.Name())

		if entry.IsDir() {
			err = MovePath(srcEntryPath, destEntryPath)
			if err != nil {
				return err
			}
		} else {
			// 移动文件
			err = os.Rename(srcEntryPath, destEntryPath)
			if err != nil {
				return err
			}
		}
	}

	// 现在目录为空，可以删除
	return os.Remove(srcPath)
}
