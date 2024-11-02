package compression

import (
	"ADM2024/pkg/common"
	"fmt"
)

//todo: print the compression ratio

func EncodeBinary(datatype, filepath string) error {

	// get the compression ratio
	outputFilePath := filepath + ".bin"
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}

func EncodeRLE(datatype, filepath string) error {

	// get the compression ratio
	outputFilePath := filepath + ".rle"
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}

func EncodeDictionary(datatype, filepath string) error {

	// get the compression ratio
	outputFilePath := filepath + ".dic"
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}

func EncodeFrameOfReference(datatype, filepath string) error {

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

func EncodeDifferential(datatype, filepath string) error {

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
