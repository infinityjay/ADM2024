package common

import (
	"fmt"
)

const (
	EndOfMap             = "END_OF_DICTIONARY"
	Int8Escape           = 0xFF
	Bit16Separator int16 = -1
	Bit8Separator  int8  = -1
	Bit4Separator        = 0x0F
)

func ValidateDataType(tech string, dataType string) error {
	integerTypes := map[string]bool{"int8": true, "int16": true, "int32": true, "int64": true}
	switch tech {
	case "bin", "for", "dif", "bve":
		if !integerTypes[dataType] {
			return fmt.Errorf("compression type '%s' supports integer types only", tech)
		}
	case "rle", "dic":
		// These support both integers and strings, so no validation needed
	default:
		return fmt.Errorf("unsupported compression type: %s", tech)
	}
	return nil
}
