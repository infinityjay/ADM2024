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

func Differential(datatype, filepath string) error {
	outputFilePath := filepath + ".dif"
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
		buff := difInt8(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}

	case "int16":
		buff := difInt16(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int32":
		buff := difInt32(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int64":
		buff := difInt64(rows)
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

// use 4 bit as offset size
func difInt8(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	previous, _ := strconv.Atoi(rows[0][0])
	packedData.WriteByte(byte(previous))
	var offsetList []int8
	for i := 1; i < n; i++ {
		value, _ := strconv.Atoi(rows[i][0])
		offset := value - previous
		previous = value
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

// 4 bit as offset size
func difInt16(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	previous, _ := strconv.Atoi(rows[0][0])

	// Write the frame as two bytes for int16 (correct order)
	packedData.Write([]byte{byte(previous), byte(previous >> 8)})
	var offsetList []int8

	for i := 1; i < n; i++ {
		value, _ := strconv.Atoi(rows[i][0])
		offset := int16(value) - int16(previous)
		previous = value

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

// use 8 bit offset
func difInt32(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	frameInt, _ := strconv.Atoi(rows[0][0])
	previous := int32(frameInt)
	packedData.Write([]byte{byte(previous >> 24), byte(previous >> 16), byte(previous >> 8), byte(previous)})
	var offsetList []int8
	for i := 1; i < n; i++ {
		valueInt, _ := strconv.Atoi(rows[i][0])
		value := int32(valueInt)
		offset := value - previous
		previous = value

		if (-128 <= offset && offset <= -2) || (0 <= offset && offset <= 127) {
			offsetList = append(offsetList, int8(offset))
		} else {
			if len(offsetList) != 0 {
				for len(offsetList)%4 != 0 {
					offsetList = append(offsetList, common.Bit8Separator)
				}
				for i := 0; i < len(offsetList); i += 4 {

					packedData.Write([]byte{
						uint8(offsetList[i]), uint8(offsetList[i+1]),
						uint8(offsetList[i+2]), uint8(offsetList[i+3]),
					})
				}
				offsetList = []int8{}
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
	if len(offsetList) != 0 {
		for len(offsetList)%4 != 0 {
			offsetList = append(offsetList, common.Bit8Separator)
		}
		for i := 0; i < len(offsetList); i += 4 {
			packedData.Write([]byte{
				uint8(offsetList[i]), uint8(offsetList[i+1]),
				uint8(offsetList[i+2]), uint8(offsetList[i+3]),
			})
		}
	}
	return packedData.Bytes()
}

func difInt64(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	previous, _ := strconv.ParseInt(rows[0][0], 10, 64)
	packedData.Write([]byte{
		byte(previous >> 56), byte(previous >> 48),
		byte(previous >> 40), byte(previous >> 32),
		byte(previous >> 24), byte(previous >> 16),
		byte(previous >> 8), byte(previous)})

	var offsetList []int8
	for i := 1; i < n; i++ {
		value, _ := strconv.ParseInt(rows[i][0], 10, 64)
		offset := value - previous
		previous = value

		if (-128 <= offset && offset <= -2) || (0 <= offset && offset <= 127) {
			offsetList = append(offsetList, int8(offset))
		} else {
			if len(offsetList) != 0 {
				for len(offsetList)%8 != 0 {
					offsetList = append(offsetList, common.Bit8Separator)
				}
				for i := 0; i < len(offsetList); i += 8 {
					packedData.Write([]byte{
						uint8(offsetList[i]), uint8(offsetList[i+1]),
						uint8(offsetList[i+2]), uint8(offsetList[i+3]),
						uint8(offsetList[i+4]), uint8(offsetList[i+5]),
						uint8(offsetList[i+6]), uint8(offsetList[i+7]),
					})
				}
				offsetList = []int8{}
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
	if len(offsetList) != 0 {
		for len(offsetList)%8 != 0 {
			offsetList = append(offsetList, common.Bit8Separator)
		}
		for i := 0; i < len(offsetList); i += 8 {
			packedData.Write([]byte{
				uint8(offsetList[i]), uint8(offsetList[i+1]),
				uint8(offsetList[i+2]), uint8(offsetList[i+3]),
				uint8(offsetList[i+4]), uint8(offsetList[i+5]),
				uint8(offsetList[i+6]), uint8(offsetList[i+7]),
			})
		}
	}
	return packedData.Bytes()
}
