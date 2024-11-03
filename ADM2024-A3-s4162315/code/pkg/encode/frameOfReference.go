package encode

import (
	"ADM2024/pkg/common"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func FrameOfReference(datatype, filepath string) error {
	outputFilePath := filepath + ".for"
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("fail to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("fail to read CSV file: %v", err)
	}

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	frame, _ := strconv.Atoi(rows[0][0])
	offsetMax := 10

	if err := writer.Write(rows[0]); err != nil {
		return fmt.Errorf("failed to write to CSV file: %v", err)
	}
	writer.Flush()
	n := len(rows)
	for i := 1; i < n; i++ {
		value, _ := strconv.Atoi(rows[i][0])
		offset := strconv.Itoa(value - frame)
		if math.Abs(float64(value-frame)) < float64(offsetMax) {
			if err := writer.Write([]string{offset}); err != nil {
				return fmt.Errorf("failed to write to CSV file: %v", err)
			}
		} else {
			if err := writer.Write([]string{strconv.Itoa(common.Infinity)}); err != nil {
				return fmt.Errorf("failed to write to CSV file: %v", err)
			}
			if err := writer.Write([]string{rows[i][0]}); err != nil {
				return fmt.Errorf("failed to write to CSV file: %v", err)
			}
		}
	}

	// get the compression ratio
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}
