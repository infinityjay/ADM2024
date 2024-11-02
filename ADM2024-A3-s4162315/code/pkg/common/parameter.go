package common

import (
	"ADM2024/pkg/compression"
	"errors"
	"fmt"
)

func ValidateDataType(tech string, dataType string) error {
	integerTypes := map[string]bool{"int8": true, "int16": true, "int32": true, "int64": true}
	switch tech {
	case "bin", "for", "dif":
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

func Execute(mode, tech, datatype, filepath string) error {
	switch mode {
	case "en":
		return encode(tech, datatype, filepath)
	case "de":
		return decode(tech, datatype, filepath)
	default:
		return errors.New("invalid mode: must be 'en' for encode or 'de' for decode")
	}
}

func encode(tech, datatype, filepath string) error {
	switch tech {
	case "bin":
		return compression.EncodeBinary(datatype, filepath)
	case "rle":
		return compression.EncodeRLE(datatype, filepath)
	case "dic":
		return compression.EncodeDictionary(datatype, filepath)
	case "for":
		return compression.EncodeFrameOfReference(datatype, filepath)
	case "dif":
		return compression.EncodeDifferential(datatype, filepath)
	default:
		return errors.New("unsupported compression tech")
	}
}

func decode(tech, datatype, filepath string) error {
	switch tech {
	case "bin":
		return compression.DecodeBinary(datatype, filepath)
	case "rle":
		return compression.DecodeRLE(datatype, filepath)
	case "dic":
		return compression.DecodeDictionary(datatype, filepath)
	case "for":
		return compression.DecodeFrameOfReference(datatype, filepath)
	case "dif":
		return compression.DecodeDifferential(datatype, filepath)
	default:
		return errors.New("unsupported compression tech")
	}
}
