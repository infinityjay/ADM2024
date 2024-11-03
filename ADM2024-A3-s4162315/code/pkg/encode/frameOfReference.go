package encode

import (
	"ADM2024/pkg/common"
	"fmt"
)

func FrameOfReference(datatype, filepath string) error {

	// get the compression ratio
	outputFilePath := filepath + ".for"
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}
