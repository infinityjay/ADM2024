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

func RunLengthEncoding(datatype, filepath string) error {
	outputFilePath := filepath + ".rle"
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

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()
	// encoding the rows
	switch datatype {
	case "string":
		if err := rleStr(rows, outputFile); err != nil {
			return fmt.Errorf("failed to encode data: %v", err)
		}
	case "int8":
		buff := rleInt8(rows)
		if err := binary.Write(outputFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}

	case "int16":
		buff := rleInt16(rows)
		if err := binary.Write(outputFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int32":
		buff := rleInt32(rows)
		if err := binary.Write(outputFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int64":
		buff := rleInt64(rows)
		if err := binary.Write(outputFile, binary.LittleEndian, buff); err != nil {
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

func rleStr(rows [][]string, file *os.File) error {
	var rowStrList []string
	for _, row := range rows {
		rowAsString := strings.Join(row, ",")
		rowStrList = append(rowStrList, rowAsString)
	}
	n := len(rowStrList)
	for i := 0; i < n; i++ {
		count := uint8(1)
		for i+1 < n && rowStrList[i] == rowStrList[i+1] {
			count++
			i++
			if count >= 255 {
				break
			}
		}
		_, err := file.WriteString(fmt.Sprintf("%s\n%d\n", rowStrList[i], count))
		if err != nil {
			return fmt.Errorf("fail to write file, %v", err)
		}
	}
	return nil
}

func rleInt8(rows [][]string) []byte {
	var encoded bytes.Buffer
	n := len(rows)

	for i := 0; i < n; i++ {
		count := uint8(1)
		for i+1 < n && rows[i][0] == rows[i+1][0] {
			count++
			i++
			if count >= 255 {
				break
			}
		}
		value, _ := strconv.Atoi(rows[i][0])
		encoded.Write([]byte{byte(int8(value)), count})
	}
	return encoded.Bytes()
}

func rleInt16(rows [][]string) []byte {
	var encoded bytes.Buffer
	n := len(rows)

	for i := 0; i < n; i++ {
		count := uint8(1)
		for i+1 < n && rows[i][0] == rows[i+1][0] {
			count++
			i++
			if count >= 255 {
				break
			}
		}
		value, _ := strconv.Atoi(rows[i][0])
		packed := int16(value)
		encoded.Write([]byte{byte(packed), byte(packed >> 8), count})
	}
	return encoded.Bytes()
}

func rleInt32(rows [][]string) []byte {
	var encoded bytes.Buffer
	n := len(rows)

	for i := 0; i < n; i++ {
		count := uint8(1)
		for i+1 < n && rows[i][0] == rows[i+1][0] {
			count++
			i++
			if count >= 255 {
				break
			}
		}
		value, _ := strconv.Atoi(rows[i][0])
		packed := int32(value)
		encoded.Write([]byte{
			byte(packed >> 24), byte(packed >> 16),
			byte(packed >> 8), byte(packed),
			count,
		})
	}
	return encoded.Bytes()
}

func rleInt64(rows [][]string) []byte {
	var encoded bytes.Buffer
	n := len(rows)

	for i := 0; i < n; i++ {
		count := uint8(1)
		for i+1 < n && rows[i][0] == rows[i+1][0] {
			count++
			i++
			if count >= 255 {
				break
			}
		}
		value, _ := strconv.ParseInt(rows[i][0], 10, 64)
		encoded.Write([]byte{
			byte(value >> 56), byte(value >> 48),
			byte(value >> 40), byte(value >> 32),
			byte(value >> 24), byte(value >> 16),
			byte(value >> 8), byte(value),
			count,
		})
	}
	return encoded.Bytes()
}
