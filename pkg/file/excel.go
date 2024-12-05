package file

import (
	"bytes"
	"github.com/novando/go-ska/pkg/logger"
	"github.com/xuri/excelize/v2"
)

func structToXlsx(startRow int, writer *excelize.StreamWriter, stringDto [][]interface{}) (nextRow int, err error) {
	// write XLSX
	nextRow = startRow
	for y, datum := range stringDto {
		nextRow = startRow + y + 1
		cell, err1 := excelize.CoordinatesToCellName(1, startRow+y)
		if err1 = writer.SetRow(cell, datum); err1 != nil {
			err = err1
			return
		}
	}
	return
}

func StreamStructToXlsx(stringDto [][]interface{}, l ...*logger.Logger) (*bytes.Buffer, error) {
	log := logger.Call()
	if len(l) > 0 {
		log = l[0]
	}
	f := excelize.NewFile()
	defer f.Close()

	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	_, err = structToXlsx(1, sw, stringDto)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	sw.Flush()
	return f.WriteToBuffer()
}

func StreamStructToXlsxHeadFoot(body, header, footer [][]interface{}, log *logger.Logger) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	defer f.Close()

	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	nextRow, err := structToXlsx(1, sw, header)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	nextRow, err = structToXlsx(nextRow, sw, body)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	_, err = structToXlsx(nextRow, sw, footer)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	sw.Flush()
	return f.WriteToBuffer()
}
