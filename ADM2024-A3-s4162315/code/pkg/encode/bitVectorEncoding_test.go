package encode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBinary(t *testing.T) {
	rows := []string{
		"1", "1", "1", "2", "2", "2", "1", "1", "1", "2", "2", "2", "3", "3", "3",
	}

	// Initialize parameters
	rowLength := len(rows)
	maxValue := 3 // maximum value in the CSV
	bitVectorList := make([][]byte, maxValue)

	// Initialize bitVectorList for each unique value
	for i := 0; i < maxValue; i++ {
		bitVectorList[i] = make([]byte, (rowLength+7)/8)
	}

	// Populate the bit vector
	for i, row := range rows {
		var value int
		fmt.Sscan(row, &value)
		bitVectorList[value-1][i/8] |= 1 << (8*(i/8+1) - i - 1) // Set the i-th bit
	}

	// Print the result
	for i, bitVector := range bitVectorList {
		fmt.Printf("Value %d bit vector: ", i+1)
		for _, b := range bitVector {
			fmt.Printf("%08b ", b) // Print byte in binary format
		}
		fmt.Println()
	}
}

func TestBinary2(t *testing.T) {
	rows := []string{
		"1", "1", "1", "2", "2", "2", "1", "1",
		"1", "2", "2", "2", "3", "3", "3",
	}
	n := len(rows)
	valueMap := make(map[int][]byte)
	for i := 0; i < n; i++ {
		value, _ := strconv.Atoi(rows[i])
		if _, ok := valueMap[value]; ok {
			valueMap[value][i/8] |= 1 << (7 - i%8)
		} else {
			valueMap[value] = make([]byte, (n+7)/8)
			valueMap[value][i/8] |= 1 << (7 - i%8)
		}
	}
	for key, value := range valueMap {
		fmt.Printf("Value %d bit vector: ", key)
		for _, b := range value {
			fmt.Printf("%08b ", b) // Print byte in binary format
		}
		fmt.Println()
	}
}
