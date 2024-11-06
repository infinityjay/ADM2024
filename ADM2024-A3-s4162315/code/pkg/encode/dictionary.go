package encode

import (
	"ADM2024/pkg/common"
	"bytes"
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
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
	var encoded bytes.Buffer
	keyMap := make(map[int8]int8)
	uniqueKey := int8(1)
	var encodeRows []int8

	for _, row := range rows {
		valueInt, _ := strconv.Atoi(row[0])
		value := int8(valueInt)
		if val, ok := keyMap[value]; ok {
			encodeRows = append(encodeRows, val)
		} else {
			keyMap[value] = uniqueKey
			encodeRows = append(encodeRows, uniqueKey)
			uniqueKey++
		}
	}
	// encode map size
	encoded.WriteByte(byte(int8(len(keyMap))))
	// encode map
	for key, value := range keyMap {
		encoded.WriteByte(byte(key))
		encoded.WriteByte(byte(value))
	}
	// encode value
	for _, row := range encodeRows {
		encoded.WriteByte(byte(row))
	}
	return encoded.Bytes()
}

func dicInt16(rows [][]string) []byte {
	var encoded bytes.Buffer
	keyMap := make(map[int16]int16)
	uniqueKey := int16(1)
	var encodeRows []int16

	for _, row := range rows {
		valueInt, _ := strconv.Atoi(row[0])
		value := int16(valueInt)
		if val, ok := keyMap[value]; ok {
			encodeRows = append(encodeRows, val)
		} else {
			keyMap[value] = uniqueKey
			encodeRows = append(encodeRows, uniqueKey)
			uniqueKey++
		}
	}
	// encode map size
	lenMap := int16(len(keyMap))
	encoded.Write([]byte{byte(lenMap), byte(lenMap >> 8)})
	// encode map
	for key, value := range keyMap {
		encoded.Write([]byte{byte(key), byte(key >> 8)})
		encoded.Write([]byte{byte(value), byte(value >> 8)})
	}
	// encode value
	for _, row := range encodeRows {
		encoded.Write([]byte{byte(row), byte(row >> 8)})
	}
	return encoded.Bytes()
}

func dicInt32(rows [][]string) []byte {
	var encoded bytes.Buffer
	keyMap := make(map[int32]int32)
	uniqueKey := int32(1)
	var encodeRows []int32

	for _, row := range rows {
		valueInt, _ := strconv.Atoi(row[0])
		value := int32(valueInt)
		if val, ok := keyMap[value]; ok {
			encodeRows = append(encodeRows, val)
		} else {
			keyMap[value] = uniqueKey
			encodeRows = append(encodeRows, uniqueKey)
			uniqueKey++
		}
	}
	// encode map size
	lenMap := int32(len(keyMap))
	encoded.Write([]byte{
		byte(lenMap >> 24), byte(lenMap >> 16),
		byte(lenMap >> 8), byte(lenMap),
	})
	// encode map
	for key, value := range keyMap {
		encoded.Write([]byte{
			byte(key >> 24), byte(key >> 16),
			byte(key >> 8), byte(key),
		})
		encoded.Write([]byte{
			byte(value >> 24), byte(value >> 16),
			byte(value >> 8), byte(value),
		})
	}
	// encode value
	for _, row := range encodeRows {
		encoded.Write([]byte{
			byte(row >> 24), byte(row >> 16),
			byte(row >> 8), byte(row),
		})
	}
	return encoded.Bytes()
}

func dicInt64(rows [][]string) []byte {
	var encoded bytes.Buffer
	keyMap := make(map[int64]int64)
	uniqueKey := int64(1)
	var encodeRows []int64

	for _, row := range rows {
		value, _ := strconv.ParseInt(row[0], 10, 64)
		if val, ok := keyMap[value]; ok {
			encodeRows = append(encodeRows, val)
		} else {
			keyMap[value] = uniqueKey
			encodeRows = append(encodeRows, uniqueKey)
			uniqueKey++
		}
	}
	// encode map size
	lenMap := int64(len(keyMap))
	encoded.Write([]byte{
		byte(lenMap >> 56), byte(lenMap >> 48),
		byte(lenMap >> 40), byte(lenMap >> 32),
		byte(lenMap >> 24), byte(lenMap >> 16),
		byte(lenMap >> 8), byte(lenMap),
	})
	// encode map
	for key, value := range keyMap {
		encoded.Write([]byte{
			byte(key >> 56), byte(key >> 48),
			byte(key >> 40), byte(key >> 32),
			byte(key >> 24), byte(key >> 16),
			byte(key >> 8), byte(key),
		})
		encoded.Write([]byte{
			byte(value >> 56), byte(value >> 48),
			byte(value >> 40), byte(value >> 32),
			byte(value >> 24), byte(value >> 16),
			byte(value >> 8), byte(value),
		})
	}
	// encode value
	for _, row := range encodeRows {
		encoded.Write([]byte{
			byte(row >> 56), byte(row >> 48),
			byte(row >> 40), byte(row >> 32),
			byte(row >> 24), byte(row >> 16),
			byte(row >> 8), byte(row),
		})
	}
	return encoded.Bytes()
}
