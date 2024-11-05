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

		if firstOffset != common.Bit4Separator { // Check for separator
			// Interpret firstOffset as signed 4-bit integer
			if firstOffset >= 0x8 {
				firstOffset -= 0x10
			}
			originalValues = append(originalValues, int(frame+firstOffset))
		}

		if secondOffset != common.Bit4Separator { // Check for separator
			// Interpret secondOffset as signed 4-bit integer
			if secondOffset >= 0x8 {
				secondOffset -= 0x10
			}
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
	n := len(buff)
	if n < 2 {
		return fmt.Errorf("not enough data to decode int16 values")
	}

	// Read the frame from the first two bytes
	frame := int16(buff[0]) | int16(buff[1])<<8
	originalValues := []int16{frame} // Start with the frame value

	// Iterate through the buffer starting from the third byte
	for i := 2; i < n; {

		// Ensure we have a complete pair of bytes for packing
		if i+1 < n {
			if buff[i] == common.Int8Escape && buff[i+1] == common.Int8Escape { // Handle escape sequences
				i += 2
				if i < n {
					value := int16(buff[i]) | int16(buff[i+1])<<8
					originalValues = append(originalValues, value)
				}
				i += 2
				continue
			}
			packed := int16(buff[i]) | int16(buff[i+1])<<8
			// Decode the packed offsets
			for j := 0; j < 4; j++ {
				offset := int8((packed >> (12 - j*4)) & 0x0F) // Extract the j-th 4-bit offset
				if offset != common.Bit4Separator {           // Ignore separators
					// Interpret as signed 4-bit integer
					if offset >= 0x8 {
						offset -= 0x10
					}
					originalValues = append(originalValues, frame+int16(offset))
				}
			}
			i += 2 // Move to the next pair of bytes
		} else {
			break // Avoid reading out of bounds
		}
	}

	// Write all original values to the CSV
	for _, value := range originalValues {
		if err := writer.Write([]string{strconv.Itoa(int(value))}); err != nil {
			return fmt.Errorf("failed to write to CSV: %v", err)
		}
	}

	return nil
}

func forInt32(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	// Extract the frame from the first byte
	frame := int32(buff[0])<<24 | int32(buff[1])<<16 | int32(buff[2])<<8 | int32(buff[3])
	var originalValues []int32
	originalValues = append(originalValues, frame)

	// Process the buffer
	for i := 4; i < n; i += 4 {
		// Check for escape sequence (four consecutive common.Int8Escape bytes)
		if i+3 < n && buff[i] == common.Int8Escape && buff[i+1] == common.Int8Escape &&
			buff[i+2] == common.Int8Escape && buff[i+3] == common.Int8Escape {
			i += 4
			if i < n {
				value := int32(buff[i])<<24 | int32(buff[i+1])<<16 | int32(buff[i+2])<<8 | int32(buff[i+3])
				originalValues = append(originalValues, value)
			}
			continue
		}

		// Extract packed int32
		packed := int32(buff[i])<<24 | int32(buff[i+1])<<16 | int32(buff[i+2])<<8 | int32(buff[i+3])
		firstOffset := int16(packed >> 16)
		secondOffset := int16(packed & 0xFFFF)

		if firstOffset != common.Bit16Separator {
			originalValues = append(originalValues, frame+int32(firstOffset))
		}

		if secondOffset != common.Bit16Separator {
			originalValues = append(originalValues, frame+int32(secondOffset))
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

func forInt64(buff []byte, writer *csv.Writer) error {
	n := len(buff)
	// Extract the frame from the first byte
	frame := int64(buff[0])<<56 | int64(buff[1])<<48 | int64(buff[2])<<40 | int64(buff[3])<<32 |
		int64(buff[4])<<24 | int64(buff[5])<<16 | int64(buff[6])<<8 | int64(buff[7])
	var originalValues []int64
	originalValues = append(originalValues, frame)

	// Process the buffer
	for i := 8; i < n; i += 8 {
		// Check for escape sequence (four consecutive common.Int8Escape bytes)
		if i+7 < n && buff[i] == common.Int8Escape && buff[i+1] == common.Int8Escape &&
			buff[i+2] == common.Int8Escape && buff[i+3] == common.Int8Escape &&
			buff[i+4] == common.Int8Escape && buff[i+5] == common.Int8Escape &&
			buff[i+6] == common.Int8Escape && buff[i+7] == common.Int8Escape {
			i += 8
			if i < n {
				value := int64(buff[i])<<56 | int64(buff[i+1])<<48 | int64(buff[i+2])<<40 | int64(buff[i+3])<<32 |
					int64(buff[i+4])<<24 | int64(buff[i+5])<<16 | int64(buff[i+6])<<8 | int64(buff[i+7])
				originalValues = append(originalValues, value)
			}
			continue
		}

		// Extract packed int64
		packed := int64(buff[i])<<56 | int64(buff[i+1])<<48 | int64(buff[i+2])<<40 | int64(buff[i+3])<<32 |
			int64(buff[i+4])<<24 | int64(buff[i+5])<<16 | int64(buff[i+6])<<8 | int64(buff[i+7])
		firstOffset := int16(packed >> 48)
		secondOffset := int16(packed >> 32)
		thirdOffset := int16(packed >> 16)
		forthOffset := int16(packed & 0xFFFF)

		if firstOffset != common.Bit16Separator {
			originalValues = append(originalValues, frame+int64(firstOffset))
		}
		if secondOffset != common.Bit16Separator {
			originalValues = append(originalValues, frame+int64(secondOffset))
		}
		if thirdOffset != common.Bit16Separator {
			originalValues = append(originalValues, frame+int64(thirdOffset))
		}
		if forthOffset != common.Bit16Separator {
			originalValues = append(originalValues, frame+int64(forthOffset))
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
