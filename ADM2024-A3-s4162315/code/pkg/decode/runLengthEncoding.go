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

func RunLengthEncoding(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
	// Open the binary file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	outputFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer outputFile.Close()

	switch datatype {
	case "string":
		if err := rleStr(file, outputFile); err != nil {
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

		if err := rleInt8(buffer, writer); err != nil {
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

		if err := rleInt16(buffer, writer); err != nil {
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

		if err := rleInt32(buffer, writer); err != nil {
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

		if err := rleInt64(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	default:
		return errors.New("invalid dataType")
	}
	return nil
}

func rleStr(file *os.File, outputFile *os.File) error {
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading RLE file: %v", err)
	}
	i := 1
	frame := ""
	for scanner.Scan() {
		line := scanner.Text()
		if i%2 != 0 {
			frame = line
		} else {
			count, _ := strconv.Atoi(line)
			for j := 0; j < count; j++ {
				if _, err := outputFile.WriteString(frame + "\n"); err != nil {
					return fmt.Errorf("failed to write to CSV: %v", err)
				}
			}
		}
		i++
	}
	return nil
}

func rleInt8(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	for i := 0; i < n; i += 2 {
		value := int8(buff[i])
		count := buff[i+1]
		for j := uint8(0); j < count; j++ {
			if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
				return fmt.Errorf("failed to write to CSV: %v", err)
			}
		}
	}
	return nil
}

func rleInt16(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	for i := 0; i < n; i += 3 {
		value := int16(buff[i]) | int16(buff[i+1])<<8
		count := buff[i+2]
		for j := uint8(0); j < count; j++ {
			if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
				return fmt.Errorf("failed to write to CSV: %v", err)
			}
		}
	}
	return nil
}

func rleInt32(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	for i := 0; i < n; i += 5 {
		value := int32(buff[i])<<24 | int32(buff[i+1])<<16 | int32(buff[i+2])<<8 | int32(buff[i+3])
		count := buff[i+4]
		for j := uint8(0); j < count; j++ {
			if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
				return fmt.Errorf("failed to write to CSV: %v", err)
			}
		}
	}
	return nil
}

func rleInt64(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	for i := 0; i < n; i += 9 {
		value := int64(buff[i])<<56 | int64(buff[i+1])<<48 | int64(buff[i+2])<<40 | int64(buff[i+3])<<32 |
			int64(buff[i+4])<<24 | int64(buff[i+5])<<16 | int64(buff[i+6])<<8 | int64(buff[i+7])
		count := buff[i+8]
		for j := uint8(0); j < count; j++ {
			if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
				return fmt.Errorf("failed to write to CSV: %v", err)
			}
		}
	}
	return nil
}
