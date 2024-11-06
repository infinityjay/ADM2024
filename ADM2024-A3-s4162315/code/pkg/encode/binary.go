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

// Binary encode binary f
/* to save the storage use bit to create the bitVector instead of byte, see the binary_test.go file
rows := []string{
		"1", "1", "1", "2", "2", "2", "1", "1", "1", "2", "2", "2", "3", "3", "3",
	}
Value 1 bit vector: 11100011 10000000
Value 2 bit vector: 00011100 01110000
Value 3 bit vector: 00000000 00001110
*/
func Binary(datatype, filepath string) error {
	outputFilePath := filepath + ".bin"
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
		buff := binInt8(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}

	case "int16":
		buff := binInt16(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int32":
		buff := binInt32(rows)
		if err := binary.Write(outFile, binary.LittleEndian, buff); err != nil {
			return fmt.Errorf("failed to write data: %v", err)
		}
	case "int64":
		buff := binInt64(rows)
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

func binInt8(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	valueMap := make(map[int8][]byte)
	for i := 0; i < n; i++ {
		valueInt, _ := strconv.Atoi(rows[i][0])
		value := int8(valueInt)
		if _, ok := valueMap[value]; ok {
			valueMap[value][i/8] |= 1 << (7 - i%8)
		} else {
			valueMap[value] = make([]byte, (n+7)/8)
			valueMap[value][i/8] |= 1 << (7 - i%8)
		}
	}
	// write the length
	lenInt32 := int32(n)
	packedData.Write([]byte{byte(lenInt32 >> 24), byte(lenInt32 >> 16), byte(lenInt32 >> 8), byte(lenInt32)})
	for key, value := range valueMap {
		packedData.WriteByte(byte(key))
		for _, b := range value {
			packedData.WriteByte(b)
		}
	}

	return packedData.Bytes()
}

func binInt16(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	valueMap := make(map[int16][]byte)
	for i := 0; i < n; i++ {
		valueInt, _ := strconv.Atoi(rows[i][0])
		value := int16(valueInt)
		if _, ok := valueMap[value]; ok {
			valueMap[value][i/8] |= 1 << (7 - i%8)
		} else {
			valueMap[value] = make([]byte, (n+7)/8)
			valueMap[value][i/8] |= 1 << (7 - i%8)
		}
	}
	// write the length
	lenInt32 := int32(n)
	packedData.Write([]byte{byte(lenInt32 >> 24), byte(lenInt32 >> 16), byte(lenInt32 >> 8), byte(lenInt32)})
	for key, value := range valueMap {
		packedData.Write([]byte{byte(key), byte(key >> 8)})
		for _, b := range value {
			packedData.WriteByte(b)
		}
	}

	return packedData.Bytes()
}

func binInt32(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	valueMap := make(map[int32][]byte)
	for i := 0; i < n; i++ {
		valueInt, _ := strconv.Atoi(rows[i][0])
		value := int32(valueInt)
		if _, ok := valueMap[value]; ok {
			valueMap[value][i/8] |= 1 << (7 - i%8)
		} else {
			valueMap[value] = make([]byte, (n+7)/8)
			valueMap[value][i/8] |= 1 << (7 - i%8)
		}
	}
	// write the length
	lenInt32 := int32(n)
	packedData.Write([]byte{byte(lenInt32 >> 24), byte(lenInt32 >> 16), byte(lenInt32 >> 8), byte(lenInt32)})
	for key, value := range valueMap {
		packedData.Write([]byte{byte(key >> 24), byte(key >> 16), byte(key >> 8), byte(key)})
		for _, b := range value {
			packedData.WriteByte(b)
		}
	}

	return packedData.Bytes()
}

func binInt64(rows [][]string) []byte {
	var packedData bytes.Buffer
	n := len(rows)
	valueMap := make(map[int64][]byte)
	for i := 0; i < n; i++ {
		value, _ := strconv.ParseInt(rows[i][0], 10, 64)
		if _, ok := valueMap[value]; ok {
			valueMap[value][i/8] |= 1 << (7 - i%8)
		} else {
			valueMap[value] = make([]byte, (n+7)/8)
			valueMap[value][i/8] |= 1 << (7 - i%8)
		}
	}
	// write the length
	lenInt32 := int32(n)
	packedData.Write([]byte{byte(lenInt32 >> 24), byte(lenInt32 >> 16), byte(lenInt32 >> 8), byte(lenInt32)})
	for key, value := range valueMap {
		packedData.Write([]byte{
			byte(key >> 56), byte(key >> 48),
			byte(key >> 40), byte(key >> 32),
			byte(key >> 24), byte(key >> 16),
			byte(key >> 8), byte(key)})
		for _, b := range value {
			packedData.WriteByte(b)
		}
	}

	return packedData.Bytes()
}
