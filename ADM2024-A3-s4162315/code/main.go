package main

import (
	"ADM2024/pkg/common"
	"fmt"
	"github.com/spf13/pflag"
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

	startTime := time.Now().Second()
	if err := common.Execute(mode, tech, datatype, filepath); err != nil {
		return err
	}
	endTime := time.Now().Second()
	executionTime := endTime - startTime
	fmt.Printf("The execution time of %s file with %s technique is: %d s.\n", mode, tech, executionTime)

	return nil
}
