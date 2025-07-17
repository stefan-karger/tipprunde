package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

var utf8BOM = []byte{0xEF, 0xBB, 0xBF}

func WriteStructsToCSV[T any](data []T, filePath string) error {

	if len(data) == 0 {
		return fmt.Errorf("no data to write to CSV")
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %w\n", err)
	}
	defer file.Close()

	if _, err := file.Write(utf8BOM); err != nil {
		return fmt.Errorf("Failed to write UTF-8 BOM: %w\n", err)
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = ';'
	defer func() {
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			log.Printf("Error flushing CSV writer for %s: %v", filePath, err)
		}
	}()

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
	if err := csvWriter.Write(headers); err != nil {
		return fmt.Errorf("failed to write CSV headers: %w", err)
	}

	fmt.Printf("Writing CSV headers: %q\n", headers)

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
			var cellValue string

			switch fieldVal.Kind() {
			case reflect.Struct:
				if t, ok := fieldVal.Interface().(time.Time); ok {
					if !t.IsZero() {
						cellValue = t.Format("02.01.2006")
					}
				} else {
					cellValue = fmt.Sprintf("%v", fieldVal.Interface())
				}
			case reflect.Ptr:
				if fieldVal.IsNil() {
					cellValue = "" // Empty string for nil pointers
				} else if t, ok := fieldVal.Interface().(*time.Time); ok {
					if !t.IsZero() {
						cellValue = t.Format("02.01.2006")
					}
				} else {
					cellValue = fmt.Sprintf("%v", fieldVal.Elem().Interface())
				}
			default:
				cellValue = fmt.Sprintf("%v", fieldVal.Interface())
			}
			row = append(row, cellValue)
		}

		fmt.Printf("Writing CSV row: %q\n", row) // Using %q for quoted string slice output

		if err := csvWriter.Write(row); err != nil {
			return fmt.Errorf("failed to write CSV row: %w", err)
		}
	}

	return nil
}
