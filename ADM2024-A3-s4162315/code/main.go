package main

import (
	"ADM2024/pkg/common"
	"ADM2024/pkg/decode"
	"ADM2024/pkg/encode"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"os/exec"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Print(err)
		return
	}

}

func run() error {
	// get parameters from command line
	var mode string
	var tech string
	var datatype string
	var filepath string
	pflag.StringVar(&mode, "mode", "en", "'en' or 'de' to specify whether your program should encode the given data or decode data that your program has encoded")
	pflag.StringVar(&tech, "tech", "bin", "The compression technique to be used: bin, rle, dic, for, or dif")
	pflag.StringVar(&datatype, "datatype", "int8", "The data type of the input data: int8, int16, int32, int64, or string")
	pflag.StringVar(&filepath, "filepath", "./", "The name (or entire path) of the file to be en- or de-coded")
	pflag.Parse()

	if err := common.ValidateDataType(tech, datatype); err != nil {
		return err
	}

	// run the encode/decode and get execution time
	startTime := time.Now().UnixMilli()
	if err := execute(mode, tech, datatype, filepath); err != nil {
		return err
	}
	endTime := time.Now().UnixMilli()
	executionTime := endTime - startTime
	fmt.Printf("The execution time of %s file with %s technique is: %d ms.\n", mode, tech, executionTime)

	// decode: compare the difference of original csv and decoded csv files
	if mode == "de" {
		originalFilePath := filepath[:len(filepath)-4]
		decodedFilePath := filepath + ".csv"
		output, err := exec.Command("diff", originalFilePath, decodedFilePath).Output()
		if err != nil {
			return err
		}
		fmt.Printf("The difference between original csv file and decoded file:\n %s \n", output)
	}

	return nil
}

func execute(mode, tech, datatype, filepath string) error {
	switch mode {
	case "en":
		return encodeFunc(tech, datatype, filepath)
	case "de":
		return decodeFunc(tech, datatype, filepath)
	default:
		return errors.New("invalid mode: must be 'en' for encode or 'de' for decode")
	}
}

func encodeFunc(tech, datatype, filepath string) error {
	switch tech {
	case "bin":
		return encode.Binary(datatype, filepath)
	case "rle":
		return encode.RunLengthEncoding(datatype, filepath)
	case "dic":
		return encode.Dictionary(datatype, filepath)
	case "for":
		return encode.FrameOfReference(datatype, filepath)
	case "dif":
		return encode.Differential(datatype, filepath)
	default:
		return errors.New("unsupported compression tech")
	}
}

func decodeFunc(tech, datatype, filepath string) error {
	switch tech {
	case "bin":
		return decode.Binary(datatype, filepath)
	case "rle":
		return decode.RunLengthEncoding(datatype, filepath)
	case "dic":
		return decode.Dictionary(datatype, filepath)
	case "for":
		return decode.FrameOfReference(datatype, filepath)
	case "dif":
		return decode.Differential(datatype, filepath)
	default:
		return errors.New("unsupported compression tech")
	}
}
