package decode

import (
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func BitVectorEncoding(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("fail to open CSV file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Failed to get file info: %v\n", err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	if err := binary.Read(file, binary.LittleEndian, &buffer); err != nil {
		return fmt.Errorf("failed to read bit vector data: %v", err)
	}

	outFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	switch datatype {
	case "int8":
		if err := bveInt8(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int16":
		if err := bveInt16(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int32":
		if err := bveInt32(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int64":
		if err := bveInt64(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	default:
		return errors.New("invalid dataType")
	}
	return nil
}

func bveInt8(buff []byte, writer *csv.Writer) error {
	n := int32(len(buff))
	length := int32(buff[0])<<24 | int32(buff[1])<<16 | int32(buff[2])<<8 | int32(buff[3])
	originalValues := make([]int8, length)
	bvLength := (length + 7) / 8
	for i := int32(4); i < n; i = i + 1 + bvLength {
		key := int8(buff[i])
		value := buff[i+1 : i+1+bvLength]
		for byteIndex, b := range value {
			for bitIndex := 0; bitIndex < 8; bitIndex++ {
				globalIndex := byteIndex*8 + bitIndex
				if (b & (1 << (7 - bitIndex))) != 0 {
					originalValues[globalIndex] = key
				}
			}
		}
	}
	// Write the decoded values to the CSV writer
	for _, value := range originalValues {
		if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}
	return nil
}

func bveInt16(buff []byte, writer *csv.Writer) error {
	n := int64(len(buff))
	length := int32(buff[0])<<24 | int32(buff[1])<<16 | int32(buff[2])<<8 | int32(buff[3])
	originalValues := make([]int16, length)
	bvLength := int64((length + 7) / 8)
	for i := int64(4); i < n; i = i + 2 + bvLength {
		key := int16(buff[i]) | int16(buff[i+1])<<8
		value := buff[i+2 : i+2+bvLength]
		for byteIndex, b := range value {
			for bitIndex := 0; bitIndex < 8; bitIndex++ {
				globalIndex := byteIndex*8 + bitIndex
				if (b & (1 << (7 - bitIndex))) != 0 {
					originalValues[globalIndex] = key
				}
			}
		}
	}
	// Write the decoded values to the CSV writer
	for _, value := range originalValues {
		if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}
	return nil
}

func bveInt32(buff []byte, writer *csv.Writer) error {
	n := int64(len(buff))
	length := int32(buff[0])<<24 | int32(buff[1])<<16 | int32(buff[2])<<8 | int32(buff[3])
	originalValues := make([]int32, length)
	bvLength := int64((length + 7) / 8)
	for i := int64(4); i < n; i = i + 4 + bvLength {
		key := int32(buff[i])<<24 | int32(buff[i+1])<<16 | int32(buff[i+2])<<8 | int32(buff[i+3])
		value := buff[i+4 : i+4+bvLength]
		for byteIndex, b := range value {
			for bitIndex := 0; bitIndex < 8; bitIndex++ {
				globalIndex := byteIndex*8 + bitIndex
				if (b & (1 << (7 - bitIndex))) != 0 {
					originalValues[globalIndex] = key
				}
			}
		}
	}
	// Write the decoded values to the CSV writer
	for _, value := range originalValues {
		if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}
	return nil
}

func bveInt64(buff []byte, writer *csv.Writer) error {
	n := int64(len(buff))
	length := int32(buff[0])<<24 | int32(buff[1])<<16 | int32(buff[2])<<8 | int32(buff[3])
	originalValues := make([]int64, length)
	bvLength := int64((length + 7) / 8)
	for i := int64(4); i < n; i = i + 8 + bvLength {
		key := int64(buff[i])<<56 | int64(buff[i+1])<<48 | int64(buff[i+2])<<40 | int64(buff[i+3])<<32 |
			int64(buff[i+4])<<24 | int64(buff[i+5])<<16 | int64(buff[i+6])<<8 | int64(buff[i+7])
		value := buff[i+8 : i+8+bvLength]
		for byteIndex, b := range value {
			for bitIndex := 0; bitIndex < 8; bitIndex++ {
				globalIndex := byteIndex*8 + bitIndex
				if (b & (1 << (7 - bitIndex))) != 0 {
					originalValues[globalIndex] = key
				}
			}
		}
	}
	// Write the decoded values to the CSV writer
	for _, value := range originalValues {
		if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}
	return nil
}
