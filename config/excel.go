package config

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ExcelizeSetup(path string) (*excelize.File, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return f, nil
}
