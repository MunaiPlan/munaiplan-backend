package parser

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/xuri/excelize/v2"
)

const (
	excelFilePath = "data/catalog.xlsx"
)

type Catalog interface {
	GetCatalogName() string
	GetCatalogItemByID(id string) 
}


func writeCatalogToExcel(f *excelize.File, header *CatalogHeader, cache interface{}, sheetName string, itemStruct interface{}) error {
    index, _ := f.NewSheet(sheetName)
    f.SetActiveSheet(index)

	// Writing header information at the top of the sheet
	headerRow := 1
	writeHeaderInfoToExcel(f, sheetName, header, &headerRow)

	// Increment headerRow to leave a blank row between header info and items
	headerRow++

	itemType := reflect.TypeOf(itemStruct)
	headers := make([]string, itemType.NumField())

	for i := 0; i < itemType.NumField(); i++ {
		field := itemType.Field(i)
		tag := field.Tag.Get("xml")
		header := strings.Split(tag, ",")[0] // Use part before ',' in tag
		if header == "" {
			header = field.Name // Use field name as a fallback
		}
		headers[i] = header
		cell, _ := excelize.CoordinatesToCellName(i+1, headerRow)
		f.SetCellValue(sheetName, cell, headers[i])
	}

	headerRow++
	// Assuming all items are of the same type, get headers from the first item
    cacheValue := reflect.ValueOf(cache)
    if cacheValue.Kind() == reflect.Map {
        for _, key := range cacheValue.MapKeys() {
            item := cacheValue.MapIndex(key).Interface()
            writeItemToExcel(f, sheetName, headerRow, item)
            headerRow++
        }
    }

	


	// Save the Excel file
	if err := f.SaveAs(excelFilePath); err != nil {
		return err
	}

	i, _ := f.GetSheetIndex("Sheet1")
	f.SetActiveSheet(i)

	return nil
}

// Writes header information to the Excel file
func writeHeaderInfoToExcel(f *excelize.File, sheetName string, header *CatalogHeader, startingRow *int) {
	v := reflect.ValueOf(*header)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// Write field name
		cell, _ := excelize.CoordinatesToCellName(1, *startingRow)
		f.SetCellValue(sheetName, cell, t.Field(i).Name)

		// Write field value
		cell, _ = excelize.CoordinatesToCellName(2, *startingRow)
		f.SetCellValue(sheetName, cell, v.Field(i).Interface())

		*startingRow++
	}
}

// Dynamically writes item details to Excel based on reflection, similar to previous examples
func writeItemToExcel(f *excelize.File, sheetName string, rowIndex int, item interface{}) {
    v := reflect.Indirect(reflect.ValueOf(item)) // Handles both pointer and value types
    for i := 0; i < v.NumField(); i++ {
        // Convert field index to Excel column name (1-based index)
        colName, _ := excelize.ColumnNumberToName(i + 1)
        cell := fmt.Sprintf("%s%d", colName, rowIndex)
        f.SetCellValue(sheetName, cell, v.Field(i).Interface())
    }
}
