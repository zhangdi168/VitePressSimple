// Package myimg
// @Author: zhangdi
// @File: base64save
// @Version: 1.0.0
// @Date: 2023/9/14 15:15
package myimg

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

func SaveBase64Image(base64String, fileName string) error {
	// 解码base64字符串
	imageBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}

	//文件夹不存在则先创建
	dirPath := filepath.Dir(fileName)
	_, err = os.Stat(dirPath)
	if err != nil {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 创建新文件
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件
	_, err = file.Write(imageBytes)
	if err != nil {
		return err
	}

	return nil
}

func GetImageBase64(path string) (string, error) {
	// 读取图片文件
	imageData, err := ioutil.ReadFile(path)
	if err != nil {

		return "", err
	}
	// 将图片文件编码为base64
	base64Data := base64.StdEncoding.EncodeToString(imageData)
	return base64Data, nil
}
