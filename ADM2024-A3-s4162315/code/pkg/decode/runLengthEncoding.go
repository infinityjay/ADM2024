package decode

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunLengthEncoding(datatype, filepath string) error {
	outputCSVPath := filepath + ".csv"
	// Open the binary file
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open binary file: %v", err)
	}
	defer file.Close()

	outputFile, err := os.Create(outputCSVPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	if datatype == "string" {
		scanner := bufio.NewScanner(file)
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading RLE file: %v", err)
		}
		for scanner.Scan() {
			line := scanner.Text()
			decodedLine := strDecoding(line)
			_, err := outputFile.WriteString(decodedLine + "\n")
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
		decoded := intDecoding(rows)
		_, err = outputFile.WriteString(decoded)
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	return nil
}

func strDecoding(encoded string) string {
	var decoded strings.Builder
	n := len(encoded)
	for i := 0; i < n; {
		char := string(encoded[i])
		count, _ := strconv.Atoi(string(encoded[i+1]))
		for k := 0; k < count; k++ {
			decoded.WriteString(char)
		}
		i = i + 2
	}
	return decoded.String()
}

func intDecoding(rows [][]string) string {
	var decoded strings.Builder
	n := len(rows)
	for i := 0; i < n; i++ {
		value := rows[i][0]
		count, _ := strconv.Atoi(rows[i][1])
		for j := 0; j < count; j++ {
			decoded.WriteString(value)
			decoded.WriteString("\n")
		}
	}
	return decoded.String()
}
