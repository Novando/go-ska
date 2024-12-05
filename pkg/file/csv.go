package file

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"reflect"
	"time"
)

func structToCsv(writer *csv.Writer, stringDto []interface{}) {
	// write CSV header
	header := make([]string, 0)
	valueOf := reflect.ValueOf(stringDto[0])
	typeOf := valueOf.Type()
	for i := 0; i < valueOf.NumField(); i++ {
		header = append(header, typeOf.Field(i).Name)
	}
	if err := writer.Write(header); err != nil {
		panic(err)
	}

	// write CSV body
	for _, datum := range stringDto {
		values := make([]string, 0)
		valueOf = reflect.ValueOf(datum)
		for i := 0; i < valueOf.NumField(); i++ {
			fieldVal := valueOf.Field(i).Interface()
			var strValue string
			switch v := fieldVal.(type) {
			case int, int8, int16, int32, int64:
				strValue = fmt.Sprintf("%d", v)
			case uint, uint8, uint16, uint32, uint64:
				strValue = fmt.Sprintf("%d", v)
			case float32, float64:
				strValue = fmt.Sprintf("%f", v)
			case time.Time:
				strValue = v.Format(time.RFC3339)
			default:
				strValue = fmt.Sprintf("%v", v)
			}
			values = append(values, strValue)
		}
		if err := writer.Write(values); err != nil {
			panic(err)
		}
	}
}

func StreamStructToCsv(stringDto []interface{}) bytes.Buffer {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	structToCsv(writer, stringDto)
	writer.Flush()
	return buf
}

func StreamStructToCsvHeadFoot(stringDto []interface{}, header, footer [][]string) bytes.Buffer {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Header
	for _, datum := range header {
		if err := writer.Write(datum); err != nil {
			panic(err)
		}
	}
	structToCsv(writer, stringDto)
	// Footer
	for _, datum := range footer {
		if err := writer.Write(datum); err != nil {
			panic(err)
		}
	}

	writer.Flush()
	return buf
}

func ToCsvQuoted(stringDto []interface{}, delimiter ...rune) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	if len(delimiter) > 0 {
		writer.Comma = delimiter[0]
	}
	for _, datum := range stringDto {
		if datum == nil {
			continue
		}
		var values []string
		valueOf := reflect.ValueOf(datum)
		for i := 0; i < valueOf.NumField(); i++ {
			fieldVal := valueOf.Field(i).Interface()
			var strValue string
			switch v := fieldVal.(type) {
			case int, int8, int16, int32, int64:
				strValue = fmt.Sprintf("%d", v)
			case uint, uint8, uint16, uint32, uint64:
				strValue = fmt.Sprintf("%d", v)
			case float32, float64:
				strValue = fmt.Sprintf("%f", v)
			case time.Time:
				strValue = v.Format(time.RFC3339)
			default:
				strValue = fmt.Sprintf("%v", v)
			}
			values = append(values, strValue)
		}
		if err := writer.Write(values); err != nil {
			return nil, err
		}
	}
	writer.Flush()
	return &buf, nil
}
