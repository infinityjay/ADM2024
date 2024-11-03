package decode

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Binary(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
	// Open the binary file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open binary file: %v", err)
	}
	defer file.Close()

	var numCSVRows int32
	if err := binary.Read(file, binary.LittleEndian, &numCSVRows); err != nil {
		return fmt.Errorf("failed to read number of bit vectors: %v", err)
	}
	var numVectors int32
	if err := binary.Read(file, binary.LittleEndian, &numVectors); err != nil {
		return fmt.Errorf("failed to read number of bit vectors: %v", err)
	}

	// Initialize bitVectorList
	bitVectorList := make([][]byte, numVectors)
	var vectorSize int32
	// Read each bit vector
	for i := int32(0); i < numVectors; i++ {

		if err := binary.Read(file, binary.LittleEndian, &vectorSize); err != nil {
			return fmt.Errorf("failed to read size of bit vector: %v", err)
		}
		bitVector := make([]byte, vectorSize)
		if err := binary.Read(file, binary.LittleEndian, &bitVector); err != nil {
			return fmt.Errorf("failed to read bit vector data: %v", err)
		}
		bitVectorList[i] = bitVector
	}

	// Decode the bit vectors into original values
	decodedValues := make([]string, numCSVRows)
	for value, bitVector := range bitVectorList {
		for byteIndex, b := range bitVector {
			for bitIndex := 0; bitIndex < 8; bitIndex++ {
				globalIndex := byteIndex*8 + bitIndex
				if (b & (1 << (7 - bitIndex))) != 0 {
					decodedValues[globalIndex] = strconv.Itoa(value)
				}
			}
		}
	}

	// Write the decoded values to a CSV file
	outFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// Write rows to CSV
	for _, value := range decodedValues {
		if err := writer.Write([]string{value}); err != nil {
			return fmt.Errorf("failed to write to CSV file: %v", err)
		}
	}

	return nil
}
