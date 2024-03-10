package filehelper

import (
	"io"
	"os"
	"path/filepath"
)

// <editor-fold desc="Description">

// CopyFolder 复制文件夹
func CopyFolder(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := path[len(src):]

		if info.IsDir() {
			return os.MkdirAll(dst+relPath, info.Mode())
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(dst + relPath)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		return err
	})
}
