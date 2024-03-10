package filehelper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FindFirstImage 遍历指定目录下的所有子文件夹，直到找到第一张图片并返回其路径。如果未找到任何图片，则返回空字符串
func FindFirstImage(dir string) string {
	// 使用 filepath.Walk 遍历指定目录及其所有子目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 如果在遍历过程中发生错误，返回该错误
		if err != nil {
			return err
		}

		// 如果当前路径是一个文件（而不是目录）
		if !info.IsDir() {
			// 获取文件扩展名并转换为小写
			ext := strings.ToLower(filepath.Ext(path))

			// 检查文件是否为图片（根据扩展名判断）
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".bmp" || ext == ".gif" {
				// 当找到第一张图片时，通过返回一个错误来停止遍历(错误的内容就是文件路径)
				return fmt.Errorf("%s", path)
			}
		}

		// 如果当前路径不是图片，则继续遍历
		return nil
	})

	// 检查遍历过程是否因找到图片而提前结束
	if err != nil {
		// 如果返回的错误消息表明找到了图片，则提取图片路径并返回
		if err.Error() != "" {
			return err.Error()
		}

		// 如果返回的错误不是由于找到图片引起的，则返回空字符串
		return ""
	}

	// 如果遍历完整个目录及其所有子目录都没有找到图片，则返回空字符串
	return ""
}

// ReadLatestImage 读取一个文件夹里最新的一张图片路径
func ReadLatestImage(dir string) (string, error) {
	//获取文件夹路径
	//dir := ternary.ToString(cameraIndex == 1, cfg.GetString(constant.ConfigKeyCameraQrcodeImageDir1), cfg.GetString(constant.ConfigKeyCameraQrcodeImageDir2))
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var newestFile os.FileInfo
	for _, file := range files {
		if !file.IsDir() && isImage(file.Name()) {
			if newestFile == nil || file.ModTime().After(newestFile.ModTime()) {
				newestFile = file
			}
		}
	}

	if newestFile != nil {
		return newestFile.Name(), nil
	} else {
		return "", fmt.Errorf("No resource found in the directory")
	}
}

func isImage(name string) bool {
	ext := filepath.Ext(name)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp": // 添加其他你想要支持的图片格式
		return true
	default:
		return false
	}
}
