package decode

import (
	"bufio"
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Dictionary(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("fail to open CSV file: %v", err)
	}
	defer file.Close()

	if err != nil {
		return fmt.Errorf("fail to read CSV file: %v", err)
	}

	outputFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("fail to create binary file: %v", err)
	}
	defer outputFile.Close()

	switch datatype {
	case "string":
		if err := dicStr(file, outputFile); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int8":
		fileInfo, err := file.Stat()
		if err != nil {
			return fmt.Errorf("Failed to get file info: %v\n", err)
		}
		fileSize := fileInfo.Size()
		buffer := make([]byte, fileSize)
		if err := binary.Read(file, binary.LittleEndian, &buffer); err != nil {
			return fmt.Errorf("failed to read bit vector data: %v", err)
		}
		writer := csv.NewWriter(outputFile)
		defer writer.Flush()

		if err := dicInt8(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int16":
		fileInfo, err := file.Stat()
		if err != nil {
			return fmt.Errorf("Failed to get file info: %v\n", err)
		}
		fileSize := fileInfo.Size()
		buffer := make([]byte, fileSize)
		if err := binary.Read(file, binary.LittleEndian, &buffer); err != nil {
			return fmt.Errorf("failed to read bit vector data: %v", err)
		}
		writer := csv.NewWriter(outputFile)
		defer writer.Flush()

		if err := dicInt16(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int32":
		fileInfo, err := file.Stat()
		if err != nil {
			return fmt.Errorf("Failed to get file info: %v\n", err)
		}
		fileSize := fileInfo.Size()
		buffer := make([]byte, fileSize)
		if err := binary.Read(file, binary.LittleEndian, &buffer); err != nil {
			return fmt.Errorf("failed to read bit vector data: %v", err)
		}
		writer := csv.NewWriter(outputFile)
		defer writer.Flush()

		if err := dicInt32(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int64":
		fileInfo, err := file.Stat()
		if err != nil {
			return fmt.Errorf("Failed to get file info: %v\n", err)
		}
		fileSize := fileInfo.Size()
		buffer := make([]byte, fileSize)
		if err := binary.Read(file, binary.LittleEndian, &buffer); err != nil {
			return fmt.Errorf("failed to read bit vector data: %v", err)
		}
		writer := csv.NewWriter(outputFile)
		defer writer.Flush()

		if err := dicInt64(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	default:
		return errors.New("invalid dataType")
	}
	return nil
}

func dicStr(file *os.File, outputFile *os.File) error {
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading RLE file: %v", err)
	}
	i := 0
	mapSize := 0
	keyMap := make(map[string]string)
	preKey := ""
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			mapSize, _ = strconv.Atoi(line)
		} else if 0 < i && i <= mapSize*2 {
			if i%2 != 0 {
				preKey = line
			} else {
				keyMap[line] = preKey
			}
		} else {
			if val, ok := keyMap[line]; ok {
				if _, err := outputFile.WriteString(val + "\n"); err != nil {
					return fmt.Errorf("failed to write to CSV: %v", err)
				}
			} else {
				fmt.Printf("map not ok i: %v\n, line: %v\n", i, line)
				return fmt.Errorf("map is incomplete")
			}

		}
		i++
	}
	return nil
}

func dicInt8(buff []byte, writer *csv.Writer) error {
	return nil
}

func dicInt16(buff []byte, writer *csv.Writer) error {
	return nil
}

func dicInt32(buff []byte, writer *csv.Writer) error {
	return nil
}

func dicInt64(buff []byte, writer *csv.Writer) error {
	return nil
}
