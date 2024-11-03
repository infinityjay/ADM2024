package encode

import (
	"ADM2024/pkg/common"
	"fmt"
)

func Differential(datatype, filepath string) error {

	// get the compression ratio
	outputFilePath := filepath + ".dif"
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}
