package decode

import (
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func Binary(datatype, filepath string) error {
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
		if err := binInt8(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int16":
		if err := binInt16(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int32":
		if err := binInt32(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	case "int64":
		if err := binInt64(buffer, writer); err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
	default:
		return errors.New("invalid dataType")
	}
	return nil
}

func binInt8(buff []byte, writer *csv.Writer) error {
	return nil
}

func binInt16(buff []byte, writer *csv.Writer) error {
	return nil
}

func binInt32(buff []byte, writer *csv.Writer) error {
	return nil
}

func binInt64(buff []byte, writer *csv.Writer) error {
	return nil
}
