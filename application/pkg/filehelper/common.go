// Package filehelper
// @Author: zhangdi
// @File: common
// @Version: 1.0.0
// @Date: 2023/4/25 15:23
package filehelper

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// BatchDeleteFiles 用于批量删除给定路径下的文件
// 参数1 paths：一个字符串类型的slice，其中包含要删除的文件路径
// 参数2 delDirWhenEmpty 是否判断目录为空则删除目录，若为true则满足条件则删除该文件夹的上级文件夹，若为false则不进行处理
// 返回值1 count：已删除文件的数量
// 返回值2 error：如果出现错误，则返回相应的error对象，否则为nil
func BatchDeleteFiles(paths []string, delDirWhenEmpty bool) (int, error) {
	var count int // 定义一个计数器来记录已删除的文件数量

	for _, _path := range paths { // 对于每个指定的文件路径

		if !FileExists(_path) {
			continue
		}
		err := os.Remove(_path) // 使用os包的Remove函数删除当前文件
		if err != nil {
			info := err.Error()
			fmt.Print(info)
			return count, err // 如果删除失败，则返回计数器的值和对应的错误
		}
		count++ // 如果删除成功，则增加计数器的值

	}

	// 判断是否需要删除空文件夹
	if delDirWhenEmpty && len(paths) > 0 {
		DelEemtyDateDir(paths[len(paths)-1])
	}

	return count, nil // 如果所有操作都成功，则返回计数器的值和nil
}

func DelEemtyDateDir(_path string) error {
	dirPath := filepath.Dir(_path)
	ppdirPath := filepath.Dir(dirPath)
	// 判断上级文件夹是否存在递归子文件夹，不存在则删除该目录
	FilesOK, _ := os.ReadDir(path.Join(ppdirPath, "OK"))
	FilesNG, _ := os.ReadDir(path.Join(ppdirPath, "NG"))
	if len(FilesNG) <= 0 && len(FilesOK) <= 0 {
		//删除该日期文件
		err := os.RemoveAll(ppdirPath) // 使用os包的Remove函数删除该目录
		if err != nil {
			return err // 如果删除失败，则返回count和对应的错误
		}
	}
	return nil
}

// FileExists 判断文件或目录是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// CreateDir 创建目录-可多级
func CreateDir(dir string) string {
	if !FileExists(dir) { //如果文件夹不存在则创建
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err.Error()
		}
	}

	return ""
}

// CreateFile 创建文件，如果文件不存在则会先创建目录
func CreateFile(_path string) string {

	dir := filepath.Dir(_path)
	if !FileExists(dir) {
		resCreateDir := CreateDir(dir)
		if resCreateDir != "" {
			return resCreateDir
		}
	}
	file, err := os.Create(_path)
	if err != nil {
		return err.Error()
	}

	defer file.Close()

	return ""
}
