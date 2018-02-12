package opsutilsgo

import (
	"encoding/csv"
	"io"
	"os"
)

// CsvReader reads content from a csv file
func CsvReader(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var content [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return [][]string{}, err
		}
		content = append(content, record)
	}
	return content, nil
}

// CsvWriter writes content to a csv file
func CsvWriter(content [][]string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	writer := csv.NewWriter(file)
	for i := 0; i < len(content); i++ {
		writer.Write(content[i])
	}
	writer.Flush()
	return nil
}
