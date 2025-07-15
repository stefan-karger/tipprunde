package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect" // Used for generic struct processing
)

// WriteStructsToCSV is a generic function to write a slice of structs to a CSV file.
// It uses reflection to dynamically get headers and values.
func WriteStructsToCSV[T any](data []T, filePath string) error {
	if len(data) == 0 {
		return fmt.Errorf("no data to write to CSV")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Get headers from the first struct using reflection and 'csv' tags
	headers := []string{}
	val := reflect.ValueOf(data[0])
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		csvTag := field.Tag.Get("csv")
		if csvTag != "" {
			headers = append(headers, csvTag)
		} else {
			// If no 'csv' tag, use the field name itself
			headers = append(headers, field.Name)
		}
	}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	// Write data rows
	for _, item := range data {
		row := []string{}
		val := reflect.ValueOf(item)
		typ := val.Type()

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			// Ensure it's an exported field
			if field.PkgPath != "" { // PkgPath is empty for exported fields
				continue
			}

			fieldVal := val.Field(i)
			row = append(row, fmt.Sprintf("%v", fieldVal.Interface()))
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	return nil
}
