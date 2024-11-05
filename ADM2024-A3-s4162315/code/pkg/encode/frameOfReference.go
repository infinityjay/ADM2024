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
)

func FrameOfReference(datatype, filepath string) error {
	outputFilePath := filepath + ".for"
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("fail to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("fail to read CSV file: %v", err)
	}

	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	switch datatype {
	case "int8":
		buff := forInt8(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}

	case "int16":
		buff := forInt16(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int32":
		buff := forInt32(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int64":
		buff := forInt64(rows)
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

// use 4 bit offset, and (1111 1111) as separator, thus -1 is deleted, the offset range [-8, -2], [0, 7]
// can also add (1111) to odd number 4 bit list to pack them into int8
func forInt8(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	frame, _ := strconv.Atoi(rows[0][0])
	packedData.WriteByte(uint8(frame))
	var offsetList []int8
	for i := 1; i < n; i++ {
		value, _ := strconv.Atoi(rows[i][0])
		offset := value - frame
		if (-8 <= offset && offset <= -2) || (0 <= offset && offset <= 7) {
			offsetList = append(offsetList, int8(offset))
		} else {
			if len(offsetList) != 0 {
				// pack the 4 bit to 8 bit
				if len(offsetList)%2 != 0 {
					offsetList = append(offsetList, common.Bit4Separator)
				}
				for i := 0; i < len(offsetList); i += 2 {
					var packed int8
					packed |= (offsetList[i] & common.Bit4Separator) << 4 // set to left 4 bit
					packed |= offsetList[i+1] & common.Bit4Separator      // set to right 4 bits

					packedData.WriteByte(byte(packed)) // Write the packed byte to buffer
				}
				offsetList = []int8{}
			}
			packedData.WriteByte(common.Int8Escape)
			packedData.WriteByte(uint8(value))
		}
	}
	if len(offsetList) > 0 {
		if len(offsetList)%2 != 0 {
			offsetList = append(offsetList, common.Bit4Separator)
		}
		for i := 0; i < len(offsetList); i += 2 {
			var packed int8
			packed |= (offsetList[i] & common.Bit4Separator) << 4 // set to left 4 bit
			packed |= offsetList[i+1] & common.Bit4Separator      // set to right 4 bits

			packedData.WriteByte(byte(packed)) // Write the packed byte to buffer
		}
	}
	return packedData.Bytes()
}

// use 4 bit offset, thus -1 is deleted, the offset range [-8, -2], [0, 7]
func forInt16(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	frame, _ := strconv.Atoi(rows[0][0])

	// Write the frame as two bytes for int16 (correct order)
	packedData.Write([]byte{byte(frame), byte(frame >> 8)})
	var offsetList []int8

	for i := 1; i < n; i++ {
		value, _ := strconv.Atoi(rows[i][0])
		offset := int16(value) - int16(frame)

		if (-8 <= offset && offset <= -2) || (0 <= offset && offset <= 7) {
			offsetList = append(offsetList, int8(offset))
		} else {
			if len(offsetList) > 0 {
				// Pack the 4-bit offsets into 16-bit values
				for j := 0; j < len(offsetList); j += 4 {
					var packed int16
					// Ensure not to exceed the bounds of offsetList
					for k := 0; k < 4; k++ {
						if j+k < len(offsetList) {
							packed |= int16(offsetList[j+k]&0x0F) << (12 - (k * 4)) // Pack 4 bits
						} else {
							packed |= common.Bit4Separator << (12 - (k * 4)) // Fill with separator if not enough offsets
						}
					}

					// Write the packed 16-bit value as two bytes
					packedData.Write([]byte{byte(packed), byte(packed >> 8)})
				}
				offsetList = []int8{}
			}
			// two (11111111) to indicate the escape
			packedData.WriteByte(common.Int8Escape)
			packedData.WriteByte(common.Int8Escape)
			valueInt16 := int16(value)
			packedData.Write([]byte{byte(valueInt16), byte(valueInt16 >> 8)})
		}
	}

	// Final packing for any remaining offsets in offsetList
	if len(offsetList) > 0 {
		for j := 0; j < len(offsetList); j += 4 {
			var packed int16
			for k := 0; k < 4; k++ {
				if j+k < len(offsetList) {
					packed |= int16(offsetList[j+k]&0x0F) << (12 - (k * 4)) // Pack 4 bits
				} else {
					packed |= common.Bit4Separator << (12 - (k * 4)) // Fill with separator if not enough offsets
				}
			}

			// Write the packed 16-bit value as two bytes
			packedData.Write([]byte{byte(packed), byte(packed >> 8)})
		}
	}
	return packedData.Bytes()
}

// use 16 bit offset, thus -1 is deleted and use -1 as separator
func forInt32(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	frameInt, _ := strconv.Atoi(rows[0][0])
	frame := int32(frameInt)
	packedData.Write([]byte{byte(frame >> 24), byte(frame >> 16), byte(frame >> 8), byte(frame)})
	var offsetList []int16
	for i := 1; i < n; i++ {
		valueInt, _ := strconv.Atoi(rows[i][0])
		value := int32(valueInt)
		offset := value - frame

		if (-32768 <= offset && offset <= -2) || (0 <= offset && offset <= 32767) {
			offsetList = append(offsetList, int16(offset))
		} else {
			if len(offsetList) != 0 {
				if len(offsetList)%2 != 0 {
					offsetList = append(offsetList, common.Bit16Separator)
				}
				for i := 0; i < len(offsetList); i += 2 {
					var packed int32
					packed |= int32(uint16(offsetList[i])) << 16 // left 16 bits
					packed |= int32(uint16(offsetList[i+1]))     // right 16 bits

					// Write the packed int32 to buffer, 4 bytes
					packedData.Write([]byte{
						byte(packed >> 24), byte(packed >> 16),
						byte(packed >> 8), byte(packed),
					})
				}
				offsetList = []int16{}
			}
			// four (11111111) to indicate the escape
			packedData.WriteByte(common.Int8Escape)
			packedData.WriteByte(common.Int8Escape)
			packedData.WriteByte(common.Int8Escape)
			packedData.WriteByte(common.Int8Escape)
			packedData.Write([]byte{
				byte(value >> 24), byte(value >> 16),
				byte(value >> 8), byte(value),
			})
		}
	}
	// Final check if any offsets are left unprocessed
	if len(offsetList) > 0 {
		if len(offsetList)%2 != 0 {
			offsetList = append(offsetList, common.Bit16Separator)
		}
		for i := 0; i < len(offsetList); i += 2 {
			var packed int32
			packed |= int32(uint16(offsetList[i])) << 16
			packed |= int32(uint16(offsetList[i+1]))

			packedData.Write([]byte{
				byte(packed >> 24), byte(packed >> 16),
				byte(packed >> 8), byte(packed),
			})
		}
	}
	return packedData.Bytes()
}

func forInt64(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	frame, _ := strconv.ParseInt(rows[0][0], 10, 64)
	packedData.Write([]byte{
		byte(frame >> 56), byte(frame >> 48),
		byte(frame >> 40), byte(frame >> 32),
		byte(frame >> 24), byte(frame >> 16),
		byte(frame >> 8), byte(frame)})

	var offsetList []int16
	for i := 1; i < n; i++ {
		value, _ := strconv.ParseInt(rows[i][0], 10, 64)
		offset := value - frame

		if (-32768 <= offset && offset <= -2) || (0 <= offset && offset <= 32767) {
			offsetList = append(offsetList, int16(offset))
		} else {
			if len(offsetList) != 0 {
				for len(offsetList)%4 != 0 {
					offsetList = append(offsetList, common.Bit16Separator)
				}
				for i := 0; i < len(offsetList); i += 4 {
					var packed int64
					packed |= int64(uint16(offsetList[i])) << 48
					packed |= int64(uint16(offsetList[i+1])) << 32
					packed |= int64(uint16(offsetList[i+2])) << 16
					packed |= int64(uint16(offsetList[i+3]))

					packedData.Write([]byte{
						byte(packed >> 56), byte(packed >> 48),
						byte(packed >> 40), byte(packed >> 32),
						byte(packed >> 24), byte(packed >> 16),
						byte(packed >> 8), byte(packed)})
				}
				offsetList = []int16{}
			}
			// eight (11111111) to indicate the escape
			for i := 0; i < 8; i++ {
				packedData.WriteByte(common.Int8Escape)
			}
			packedData.Write([]byte{
				byte(value >> 56), byte(value >> 48),
				byte(value >> 40), byte(value >> 32),
				byte(value >> 24), byte(value >> 16),
				byte(value >> 8), byte(value),
			})
		}
	}
	// Final check if any offsets are left unprocessed
	if len(offsetList) > 0 {
		for len(offsetList)%4 != 0 {
			offsetList = append(offsetList, common.Bit16Separator)
		}
		for i := 0; i < len(offsetList); i += 4 {
			var packed int64
			packed |= int64(uint16(offsetList[i])) << 48
			packed |= int64(uint16(offsetList[i+1])) << 32
			packed |= int64(uint16(offsetList[i+2])) << 16
			packed |= int64(uint16(offsetList[i+3]))

			packedData.Write([]byte{
				byte(packed >> 56), byte(packed >> 48),
				byte(packed >> 40), byte(packed >> 32),
				byte(packed >> 24), byte(packed >> 16),
				byte(packed >> 8), byte(packed)})
		}
	}
	return packedData.Bytes()
}
