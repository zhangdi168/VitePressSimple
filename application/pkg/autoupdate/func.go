package autoupdate

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"wailstemplate/application/pkg/filehelper"
)

// SysIsWindow 当前系统是否为windows
func SysIsWindow() bool {
	return runtime.GOOS == "windows"
}

// SysIsMac 当前系统是否为mac
func SysIsMac() bool {
	return runtime.GOOS == "darwin"
}

// SysIsLinux 当前系统是否为linux
func SysIsLinux() bool {
	return runtime.GOOS == "linux"
}

// UnzipFile 解压缩单个ZIP文件到指定目录
func UnzipFile(zipFilePath string, destDir string) error {
	// 打开ZIP文件
	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return fmt.Errorf("failed to open ZIP file: %w", err)
	}
	defer reader.Close()

	// 遍历ZIP文件中的每个条目
	for _, file := range reader.File {
		// 获取条目的路径信息
		filePath := filepath.Join(destDir, file.Name)

		// 检查是否为目录，如果是则创建该目录
		if file.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		// 创建目标文件并打开写入流
		targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("failed to create target file: %w", err)
		}
		defer targetFile.Close()

		// 打开ZIP条目的读取流
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open ZIP entry for reading: %w", err)
		}
		defer rc.Close()

		// 将ZIP条目的内容复制到目标文件
		_, err = io.Copy(targetFile, rc)
		if err != nil {
			return fmt.Errorf("failed to copy data from ZIP entry to target file: %w", err)
		}
	}

	return nil
}

// Unzip 解压zip压缩包（MACOS 用到）
func Unzip(compressedPackagePath string, unzipDir string, unzipDirPre []string) bool {
	// 保持权限和软连接解压
	// unzipDirPre 例如 ["my_app.app/Contents/"] 不填则全部解压

	file, err := zip.OpenReader(compressedPackagePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	for _, info := range file.File {
		// 检查目标文件路径是否在允许解压路径前缀中
		if len(unzipDirPre) > 0 {
			allowUnzip := false
			for _, path_ := range unzipDirPre {
				if strings.HasPrefix(info.Name, path_) {
					allowUnzip = true
				}
			}
			if !allowUnzip {
				continue
			}
		}

		fileName := info.Name

		targetFilePath := filepath.Join(unzipDir, fileName)

		permission := info.Mode().Perm()
		if info.Mode()&os.ModeSymlink != 0 { // 检查是否为软连接
			softLinkPath, err := readLink(info)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 检查目标文件路径是否存在，如果存在就删除，防止创建失败
			if _, err := os.Lstat(targetFilePath); err == nil {
				os.Remove(targetFilePath)
			}
			err = os.Symlink(softLinkPath, targetFilePath)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// 删除文件重新解压
			if _, err := os.Lstat(targetFilePath); err == nil {
				// 检查是否为文件
				if !info.Mode().IsDir() {
					os.Remove(targetFilePath)
				}
			}
			err = extractFile(info, targetFilePath)
			if err != nil {
				fmt.Println(err)
			}

			// 修改文件权限
			if _, err := os.Lstat(targetFilePath); err == nil {
				err = os.Chmod(targetFilePath, permission)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	return true
}

func readLink(info *zip.File) (string, error) {
	file, err := info.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	linkName := make([]byte, info.FileInfo().Size())
	_, err = io.ReadFull(file, linkName)
	if err != nil {
		return "", err
	}

	return string(linkName), nil
}

// 提取文件
func extractFile(info *zip.File, targetFilePath string) error {
	file, err := info.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	if info.FileInfo().IsDir() {
		err = os.MkdirAll(targetFilePath, info.Mode())
		if err != nil {
			return err
		}
	} else {
		writer, err := os.OpenFile(targetFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, info.Mode())
		if err != nil {
			return err
		}
		defer writer.Close()

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}
	}

	return nil
}

// getDownloadDir 获取用户的【下载】文件夹
func getDownloadDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("pull:", err)
		return ""
	}
	downloadFolderPath := filepath.Join(usr.HomeDir, UpdateCfg.AppName, "download")
	filehelper.CreateDir(downloadFolderPath)
	return downloadFolderPath
}

// 模拟三元运算符 返回字符串
func ifThenStr(flag bool, resTrue, resFalse string) string {
	if flag {
		return resTrue
	}
	return resFalse
}

// 模拟三元运算符 返回int
func ifThenInt(flag bool, resTrue, resFalse int) int {
	if flag {
		return resTrue
	}
	return resFalse
}

// CopyFile copies the contents of the source file to the destination file.
// It returns an error if any occurs during the copy process.
func CopyFile(srcFilePath, dstFilePath string) error {
	// Open the source file
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return fmt.Errorf("failed to open source file '%s': %w", srcFilePath, err)
	}
	defer srcFile.Close()

	// Create or truncate the destination file
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		return fmt.Errorf("failed to create destination file '%s': %w", dstFilePath, err)
	}
	defer dstFile.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents from '%s' to '%s': %w", srcFilePath, dstFilePath, err)
	}

	return nil
}

func GetCpuArch() string {
	// 获取CPU架构
	cpuArch := runtime.GOARCH
	switch cpuArch {
	case "arm64":
		return "arm64"
	case "amd64":
		return "amd64"
	default:
		return "amd64"
	}
	return "amd64"
}
