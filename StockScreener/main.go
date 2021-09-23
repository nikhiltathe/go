package main

import (
	"fmt"

	"github.com/go/StockScreener/excelReader"
	"github.com/go/StockScreener/sorter"
)

func main() {

	fmt.Println("HI")

	excelReader.Read()
	sorter.Sort()
}
