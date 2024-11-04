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
	originalSizeStr := fmt.Sprintf("%.2f", originalSize)
	fmt.Printf("original size: %s MB\n", originalSizeStr)

	compressedSize, err := getFileSize(compressedFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	compressedSizeStr := fmt.Sprintf("%.2f", compressedSize)
	fmt.Printf("compressed size: %s MB\n", compressedSizeStr)

	if compressedSize == 0 {
		err := errors.New("error: Compressed file size is zero")
		return 0, err
	}
	compressionRatio := originalSize / compressedSize
	return compressionRatio, nil
}

func getFileSize(filePath string) (float64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	kbSize := float64(fileInfo.Size()) / 1024
	mbSize := kbSize / 1024
	return mbSize, nil
}
