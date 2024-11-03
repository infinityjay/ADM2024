package encode

import (
	"ADM2024/pkg/common"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Dictionary(datatype, filepath string) error {
	outputFilePath := filepath + ".dic"
	// read the csv file
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

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("fail to create binary file: %v", err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	keyMap := make(map[string]int)
	uniqueKeys := 1
	var encodeRows [][]string
	for _, row := range rows {
		encodedRow := make([]string, len(row))
		for j, item := range row {
			if val, ok := keyMap[item]; ok {
				encodedRow[j] = strconv.Itoa(val)
			} else {
				keyMap[item] = uniqueKeys
				encodedRow[j] = strconv.Itoa(uniqueKeys)
				uniqueKeys++
			}
		}
		encodeRows = append(encodeRows, encodedRow)
	}
	for key, value := range keyMap {
		if err := writer.Write([]string{key, strconv.Itoa(value)}); err != nil {
			return fmt.Errorf("Error writing dictionary: %v\n", err)
		}
	}
	writer.Flush()
	// add end of dictionary to
	if err := writer.Write([]string{common.EndOfMap}); err != nil {
		return fmt.Errorf("Error writing data: %v\n", err)
	}
	writer.Flush()
	for _, encodeRow := range encodeRows {
		if err := writer.Write(encodeRow); err != nil {
			return fmt.Errorf("Error writing data: %v\n", err)
		}
	}
	writer.Flush()

	// get the compression ratio
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}
