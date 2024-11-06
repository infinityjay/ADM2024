package encode

import (
	"ADM2024/pkg/common"
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

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
	return nil
}

func binInt16(rows [][]string) []byte {
	return nil
}

func binInt32(rows [][]string) []byte {
	return nil
}

func binInt64(rows [][]string) []byte {
	return nil
}
