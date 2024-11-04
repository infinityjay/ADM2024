package encode

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFrameOfReference(t *testing.T) {
	var packedData bytes.Buffer
	first := int16(3)   // Example positive offset
	second := int16(-4) // Example negative offset

	a := uint16(first)
	b := uint16(second)
	var packed int32

	// Pack the offsets into an int32
	packed |= int32(a) << 16 // left 16 bits
	packed |= int32(b)       // right 16 bits

	// Write the packed int32 to buffer, 4 bytes
	packedData.Write([]byte{
		byte(packed >> 24), byte(packed >> 16),
		byte(packed >> 8), byte(packed),
	})

	final := packedData.Bytes()
	value := int32(final[0])<<24 | int32(final[1])<<16 | int32(final[2])<<8 | int32(final[3])

	// Extract the 16-bit values from the packed int32 with proper masking
	firstOffset := int16((value >> 16) & 0xFFFF) // Mask to get the correct 16-bit signed value
	secondOffset := int16(value & 0xFFFF)        // Mask to get the correct 16-bit signed value

	fmt.Printf("firstOffset: %v, secondOffset: %v\n", firstOffset, secondOffset)
}
