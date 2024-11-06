package encode

import (
	"ADM2024/pkg/common"
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
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

	switch datatype {
	case "string":
		if err := dicStr(rows, outFile); err != nil {
			return fmt.Errorf("failed to encode data: %v", err)
		}
	case "int8":
		buff := dicInt8(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}

	case "int16":
		buff := dicInt16(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int32":
		buff := dicInt32(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int64":
		buff := dicInt64(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	default:
		return errors.New("invalid dataType")
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

func dicStr(rows [][]string, file *os.File) error {
	var rowStrList []string
	keyMap := make(map[string]int)
	uniqueKey := 1
	var encodeRows []int
	for _, row := range rows {
		rowAsString := strings.Join(row, ",")
		rowStrList = append(rowStrList, rowAsString)
	}

	for _, rowStr := range rowStrList {
		if val, ok := keyMap[rowStr]; ok {
			encodeRows = append(encodeRows, val)
		} else {
			keyMap[rowStr] = uniqueKey
			encodeRows = append(encodeRows, uniqueKey)
			uniqueKey++
		}
	}
	// encode map size
	if _, err := file.WriteString(fmt.Sprintf("%d\n", len(keyMap))); err != nil {
		return fmt.Errorf("fail to write map size, %v", err)
	}
	// encode map
	for key, value := range keyMap {
		if _, err := file.WriteString(fmt.Sprintf("%s\n%d\n", key, value)); err != nil {
			return fmt.Errorf("fail to write map, %v", err)
		}
	}
	// encode value
	for _, row := range encodeRows {
		if _, err := file.WriteString(fmt.Sprintf("%d\n", row)); err != nil {
			return fmt.Errorf("fail to write value, %v", err)
		}
	}

	return nil
}

func dicInt8(rows [][]string) []byte {
	return nil
}

func dicInt16(rows [][]string) []byte {
	return nil
}

func dicInt32(rows [][]string) []byte {
	return nil
}

func dicInt64(rows [][]string) []byte {
	return nil
}
