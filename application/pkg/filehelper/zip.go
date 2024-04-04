package filehelper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CompressFilesToZip creates a ZIP archive at the specified output path,
// containing the files listed in filePaths. Each file is added to the archive
// using its original name and relative path.
func CompressFilesToZip(outputPath string, filePaths []string) error {
	// Create a new ZIP archive.
	zipFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create ZIP file '%s': %w", outputPath, err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, filePath := range filePaths {
		// Open the file to be added to the archive.
		fileToAdd, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file '%s': %w", filePath, err)
		}
		defer fileToAdd.Close()

		info, err := fileToAdd.Stat()
		if err != nil {
			return fmt.Errorf("failed to get file info for '%s': %w", filePath, err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("failed to create ZIP header for '%s': %w", filePath, err)
		}

		// Set the compression method and adjust the file name as needed.
		header.Method = zip.Deflate
		header.Name = filepath.Base(filePath)

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("failed to create writer for '%s' in ZIP: %w", filePath, err)
		}

		_, err = io.Copy(writer, fileToAdd)
		if err != nil {
			return fmt.Errorf("failed to copy file '%s' content to ZIP: %w", filePath, err)
		}
	}

	fmt.Printf("Successfully created ZIP archive: %s\n", outputPath)
	return nil
}

func demo() {
	filesToCompress := []string{
		"path/to/file1.txt",
		"path/to/subdir/file2.csv",
		"path/to/another_file.jpg",
	}

	outputPath := "compressed_files.zip"
	err := CompressFilesToZip(outputPath, filesToCompress)
	if err != nil {
		fmt.Printf("Error compressing files: %v\n", err)
	} else {
		fmt.Println("Files compressed successfully.")
	}
}
