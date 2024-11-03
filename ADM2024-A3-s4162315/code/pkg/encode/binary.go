package encode

import (
	"ADM2024/pkg/common"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Binary encode binary f
/* to save the storage use bit to create the bitVector instead of byte, run the binary_test.go file
rows := []string{
		"1", "1", "1", "2", "2", "2", "1", "1", "1", "2", "2", "2", "3", "3", "3",
	}
Value 1 bit vector: 11100011 10000000
Value 2 bit vector: 00011100 01110000
Value 3 bit vector: 00000000 00001110
*/
func Binary(datatype, filepath string) error {
	outputFilePath := filepath + ".bin"
	// read the csv file
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
	// generate the bit vector list
	rowLength := len(rows)
	bitVectorList := make([][]byte, 0)
	maxValue := -1
	for i, row := range rows {
		value, _ := strconv.Atoi(row[0])
		// expand the size of bitVectorList while loop the rows
		if value > maxValue {
			if value > maxValue {
				maxValue = value
				for len(bitVectorList) <= maxValue {
					bitVectorList = append(bitVectorList, make([]byte, (rowLength+7)/8))
				}
			}
		}
		bitVectorList[value][i/8] |= 1 << (8*(i/8+1) - i - 1)
	}
	// write the .bin file
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("fail to create binary file: %v", err)
	}
	defer outFile.Close()

	// write the bit packing with encoding/binary
	if err := binary.Write(outFile, binary.LittleEndian, int32(rowLength)); err != nil {
		return fmt.Errorf("failed to write number of csv rows: %v", err)
	}

	numVectors := int32(len(bitVectorList))
	if err := binary.Write(outFile, binary.LittleEndian, numVectors); err != nil {
		return fmt.Errorf("failed to write number of bit vectors: %v", err)
	}

	for _, bitVector := range bitVectorList {
		vectorSize := int32(len(bitVector))
		if err := binary.Write(outFile, binary.LittleEndian, vectorSize); err != nil {
			return fmt.Errorf("failed to write size of bit vector: %v", err)
		}
		if err := binary.Write(outFile, binary.LittleEndian, bitVector); err != nil {
			return fmt.Errorf("failed to write bit vector data: %v", err)
		}
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
