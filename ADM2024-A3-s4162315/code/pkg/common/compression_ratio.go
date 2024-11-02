package common

import (
	"errors"
	"fmt"
	"os"
)

func GetCompressionRatio(originalFilePath string, compressedFilePath string) (float64, error) {
	originalSize, err := getFileSize(originalFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	compressedSize, err := getFileSize(compressedFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	if compressedSize == 0 {
		err := errors.New("error: Compressed file size is zero")
		return 0, err
	}
	compressionRatio := float64(originalSize) / float64(compressedSize)
	return compressionRatio, nil
}

func getFileSize(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
