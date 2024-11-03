package encode

import (
	"fmt"
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
