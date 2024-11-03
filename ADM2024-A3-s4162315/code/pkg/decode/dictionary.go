package decode

import (
	"ADM2024/pkg/common"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Dictionary(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("fail to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("fail to read CSV file: %v", err)
	}

	outFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("fail to create binary file: %v", err)
	}
	defer outFile.Close()

	keyMap := make(map[string]string)
	mapEnd := true
	for _, row := range rows {
		if mapEnd && len(row) > 1 {
			keyMap[row[1]] = row[0]
			continue
		}
		if row[0] == common.EndOfMap {
			mapEnd = false
			continue
		}
		if !mapEnd {
			var rowString []string
			for _, item := range row {
				rowString = append(rowString, keyMap[item])
			}
			_, err := outFile.WriteString(strings.Join(rowString, ",") + "\n")
			if err != nil {
				return fmt.Errorf("error writing to output file: %v", err)
			}
		}
	}

	return nil
}
