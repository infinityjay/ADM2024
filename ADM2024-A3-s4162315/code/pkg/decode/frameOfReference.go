package decode

import (
	"ADM2024/pkg/common"
	"encoding/binary"
	"encoding/csv"
	"errors"
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
		if err := forInt8(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int16":
		if err := forInt16(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int32":
		if err := forInt32(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int64":
		if err := forInt64(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	default:
		return errors.New("invalid dataType")
	}
	return nil
}

// use 4 bit offset, and (1111 1111) as separator, thus -1 is deleted, the offset range [-8, -2], [0, 7]
// can also add (1111) to odd number 4 bit list to pack them into int8
func forInt8(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	frame := int8(buff[0])
	var originalValues []int
	originalValues = append(originalValues, int(frame))
	for i := 1; i < n; i++ {
		currentByte := buff[i]
		if currentByte == common.Int8Escape { // Handle escape sequences
			i++
			if i < n {
				value := int(buff[i]) // The next byte is a literal value
				originalValues = append(originalValues, value)
			}
			continue
		}
		firstOffset := int8((currentByte >> 4) & 0x0F) // Get the first 4 bits
		secondOffset := int8(currentByte & 0x0F)       // Get the second 4 bits

		// Interpret firstOffset as signed 4-bit integer
		if firstOffset >= 0x8 {
			firstOffset -= 0x10
		}
		if firstOffset != common.Bit4Separator { // Check for separator
			originalValues = append(originalValues, int(frame+firstOffset))
		}

		if secondOffset != common.Bit4Separator { // Check for separator
			originalValues = append(originalValues, int(frame+secondOffset))
		}
	}
	for _, value := range originalValues {
		if err := writer.Write([]string{strconv.Itoa(value)}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}
	return nil
}

func forInt16(buff []byte, writer *csv.Writer) error {
	return nil
}

func forInt32(buff []byte, writer *csv.Writer) error {
	return nil
}

func forInt64(buff []byte, writer *csv.Writer) error {
	return nil
}
