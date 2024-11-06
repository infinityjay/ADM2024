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
				return fmt.Errorf("map is incomplete")
			}

		}
		i++
	}
	return nil
}

func dicInt8(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	mapSize := int(buff[0])
	preKey := int8(0)
	keyMap := make(map[int8]int8)
	for i := 1; i < n; i++ {
		if i < mapSize*2+1 {
			if i%2 != 0 {
				preKey = int8(buff[i])
			} else {
				value := int8(buff[i])
				keyMap[value] = preKey
			}
		} else {
			value := int8(buff[i])
			if val, ok := keyMap[value]; ok {
				if err := writer.Write([]string{strconv.Itoa(int(val))}); err != nil {
					return fmt.Errorf("failed to write to CSV: %v", err)
				}
			} else {
				return fmt.Errorf("map is incomplete")
			}
		}
	}
	return nil
}

func dicInt16(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	mapSize := int16(buff[0]) | int16(buff[1])<<8

	keyMap := make(map[int16]int16)
	for i := 2; i < n; {
		if i < int(mapSize)*2*2+2 {
			key := int16(buff[i]) | int16(buff[i+1])<<8
			value := int16(buff[i+2]) | int16(buff[i+3])<<8
			keyMap[value] = key
			i += 4
		} else {
			value := int16(buff[i]) | int16(buff[i+1])<<8
			if val, ok := keyMap[value]; ok {
				if err := writer.Write([]string{strconv.Itoa(int(val))}); err != nil {
					return fmt.Errorf("failed to write to CSV: %v", err)
				}
			} else {
				return fmt.Errorf("map is incomplete")
			}
			i += 2
		}
	}
	return nil
}

func dicInt32(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	mapSize := int32(buff[0])<<24 | int32(buff[1])<<16 | int32(buff[2])<<8 | int32(buff[3])
	keyMap := make(map[int32]int32)
	for i := 4; i < n; {
		if i < int(mapSize)*2*4+4 {
			key := int32(buff[i])<<24 | int32(buff[i+1])<<16 | int32(buff[i+2])<<8 | int32(buff[i+3])
			value := int32(buff[i+4])<<24 | int32(buff[i+5])<<16 | int32(buff[i+6])<<8 | int32(buff[i+7])
			keyMap[value] = key
			i += 8
		} else {
			value := int32(buff[i])<<24 | int32(buff[i+1])<<16 | int32(buff[i+2])<<8 | int32(buff[i+3])
			if val, ok := keyMap[value]; ok {
				if err := writer.Write([]string{strconv.Itoa(int(val))}); err != nil {
					return fmt.Errorf("failed to write to CSV: %v", err)
				}
			} else {
				return fmt.Errorf("map is incomplete")
			}
			i += 4
		}
	}
	return nil
}

func dicInt64(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	mapSize := int64(buff[0])<<56 | int64(buff[1])<<48 | int64(buff[2])<<40 | int64(buff[3])<<32 |
		int64(buff[4])<<24 | int64(buff[5])<<16 | int64(buff[6])<<8 | int64(buff[7])
	keyMap := make(map[int64]int64)
	for i := 8; i < n; {
		if i < int(mapSize)*2*8+8 {
			key := int64(buff[i])<<56 | int64(buff[i+1])<<48 | int64(buff[i+2])<<40 | int64(buff[i+3])<<32 |
				int64(buff[i+4])<<24 | int64(buff[i+5])<<16 | int64(buff[i+6])<<8 | int64(buff[i+7])
			value := int64(buff[i+8])<<56 | int64(buff[i+9])<<48 | int64(buff[i+10])<<40 | int64(buff[i+11])<<32 |
				int64(buff[i+12])<<24 | int64(buff[i+13])<<16 | int64(buff[i+14])<<8 | int64(buff[i+15])
			keyMap[value] = key
			i += 16
		} else {
			value := int64(buff[i])<<56 | int64(buff[i+1])<<48 | int64(buff[i+2])<<40 | int64(buff[i+3])<<32 |
				int64(buff[i+4])<<24 | int64(buff[i+5])<<16 | int64(buff[i+6])<<8 | int64(buff[i+7])
			if val, ok := keyMap[value]; ok {
				if err := writer.Write([]string{strconv.FormatInt(val, 10)}); err != nil {
					return fmt.Errorf("failed to write to CSV: %v", err)
				}
			} else {
				return fmt.Errorf("map is incomplete")
			}
			i += 8
		}
	}
	return nil
}
