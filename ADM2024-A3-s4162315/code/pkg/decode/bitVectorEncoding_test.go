package decode

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	byteList := []byte{0b11100011, 0b10000000} // Example bytes: 11100011 10000000
	var indexes []int
	for byteIndex, b := range byteList {
		for bitIndex := 0; bitIndex < 8; bitIndex++ {
			globalIndex := byteIndex*8 + bitIndex
			if (b & (1 << (7 - bitIndex))) != 0 {
				indexes = append(indexes, globalIndex)
			}
		}
	}
	fmt.Println("Indexes of 1 bits:", indexes)
}
