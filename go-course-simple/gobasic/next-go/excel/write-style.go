package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type M map[string]interface{}

var data = []M{
	M{"Name": "Ajung", "Gender": "male", "Age": 21},
	M{"Name": "Setia", "Gender": "female", "Age": 20},
	M{"Name": "Jung", "Gender": "male", "Age": 15},
}

func main() {
	xlsx := excelize.NewFile()
	sheetName := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)
	xlsx.SetCellValue(sheetName, "A1", "Name")
	xlsx.SetCellValue(sheetName, "B1", "Gender")
	xlsx.SetCellValue(sheetName, "C1", "Age")
	err := xlsx.AutoFilter(sheetName, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	for i, each := range data {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), each["Name"])
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), each["Gender"])
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), each["Age"])
	}

	sheet2Name := "Sheet two"
	sheetIndex := xlsx.NewSheet(sheet2Name)
	xlsx.SetActiveSheet(sheetIndex)

	xlsx.SetCellValue(sheet2Name, "A1", "Name")
	xlsx.MergeCell(sheet2Name, "A1", "B1")

	style, err := xlsx.NewStyle(`{"font":{"bold":true,"size":36},"fill":{"type":"pattern","color":["#E0EBF5"],"pattern":1}}`)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	xlsx.SetCellStyle(sheet2Name, "A1", "A1", style)

	err = xlsx.SaveAs("./file/file2.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
