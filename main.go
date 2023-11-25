package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {

	file, err := xlsx.OpenFile("gpv4.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	sheet := file.Sheets[0]

	for _, row := range sheet.Rows {
		// d := Data{
		// 	Name:  row.Cells[0].Value,
		// 	Email: row.Cells[1].Value,
		// }
		fmt.Println(row.Cells[0].Value)
		fmt.Println(row.Cells[1].Value)
		fmt.Println(row.Cells[2].Value)
		fmt.Println(row.Cells[3].Value)
		fmt.Println(row.Cells[4].Value)
		fmt.Println(row.Cells[5].Value)
		fmt.Println(row.Cells[6].Value)
		fmt.Println(row.Cells[7].Value)
		fmt.Println(row.Cells[8].Value)
		fmt.Println(row.Cells[9].Value)

		time.Sleep(1 * time.Second)

	}
}
