package decode

import (
	"ADM2024/pkg/common"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func FrameOfReference(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
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

	outFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()
	if err := writer.Write(rows[0]); err != nil {
		return fmt.Errorf("failed to write to CSV file: %v", err)
	}
	writer.Flush()

	frame, _ := strconv.Atoi(rows[0][0])
	n := len(rows)
	for i := 1; i < n; i++ {
		value, _ := strconv.Atoi(rows[i][0])
		if value == common.Infinity {
			valueNew, _ := strconv.Atoi(rows[i+1][0])
			if err := writer.Write([]string{strconv.Itoa(valueNew)}); err != nil {
				return fmt.Errorf("failed to write to CSV file: %v", err)
			}
			i++
		} else {
			valueNew := strconv.Itoa(value + frame)
			if err := writer.Write([]string{valueNew}); err != nil {
				return fmt.Errorf("failed to write to CSV file: %v", err)
			}
		}
	}

	return nil
}
