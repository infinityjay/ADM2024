package encode

import (
	"ADM2024/pkg/common"
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunLengthEncoding(datatype, filepath string) error {
	outputFilePath := filepath + ".rle"
	// read the csv file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("fail to open CSV file: %v", err)
	}
	defer file.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()
	// encoding the rows
	if datatype == "string" {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			compressed := strEncoding(line)
			_, err := outputFile.WriteString(compressed + "\n")
			if err != nil {
				return fmt.Errorf("error writing to output file: %v", err)
			}
		}
	} else {
		reader := csv.NewReader(file)
		rows, err := reader.ReadAll()
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("fail to read CSV file: %v", err)
		}
		compressed := intEncoding(rows)
		_, err = outputFile.WriteString(compressed)
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}
	// get the compression ratio
	ratio, err := common.GetCompressionRatio(filepath, outputFilePath)
	if err != nil {
		return err
	}
	ratioStr := fmt.Sprintf("%.2f", ratio)
	fmt.Printf("The compression ratio is: %s\n", ratioStr)
	return nil
}

func strEncoding(row string) string {
	var encoded strings.Builder
	n := len(row)
	for i := 0; i < n; i++ {
		count := 1
		for i+1 < n && row[i] == row[i+1] {
			count++
			i++
		}
		encoded.WriteString(string(row[i]) + strconv.Itoa(count))
	}
	return encoded.String()
}

func intEncoding(rows [][]string) string {
	var encoded strings.Builder
	n := len(rows)

	for i := 0; i < n; i++ {
		count := 1
		for i+1 < n && rows[i][0] == rows[i+1][0] {
			count++
			i++
		}
		encoded.WriteString(rows[i][0])
		encoded.WriteString(",")
		encoded.WriteString(strconv.Itoa(count))
		encoded.WriteString("\n")
	}
	return encoded.String()
}
