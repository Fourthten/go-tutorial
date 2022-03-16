package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type M map[string]interface{}

func main() {
	xlsx, err := excelize.OpenFile("./file/file1.xlsx")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	sheetName := "Sheet One"
	rows := make([]M, 0)
	for i := 2; i < 5; i++ {
		row := M{
			"Name":   xlsx.GetCellValue(sheetName, fmt.Sprintf("A%d", i)),
			"Gender": xlsx.GetCellValue(sheetName, fmt.Sprintf("B%d", i)),
			"Age":    xlsx.GetCellValue(sheetName, fmt.Sprintf("C%d", i)),
		}
		rows = append(rows, row)
	}
	fmt.Printf("%v \n", rows)
}
