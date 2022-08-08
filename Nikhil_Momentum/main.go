package main

import (
	"fmt"
	"os"
	"strconv"

	m "github.com/go/Nikhil_Momentum/Model"
	model "github.com/go/Nikhil_Momentum/Model"
	excelize "github.com/xuri/excelize/v2"
)

var allData m.RawData

func main() {
	fmt.Println("Welcome ExcelNTTrade")

	portfolio := model.Porfolio{}
	fmt.Printf("My Portfolio %#v", portfolio)

	allData.AllRows = make([]m.StockPrices, 0)
	readRaw()

	for i := 0; i < len(allData.AllRows); i++ {
		// fmt.Printf("%#v", allData.AllRows[i].Symbol)
		fmt.Println(allData.AllRows[i].Symbol)
		for key, v := range allData.AllRows[i].Prices {
			fmt.Println("Week : ", key, " Price :", v[key])
		}
		fmt.Println()
	}

	GetPriceChanges(12)
}

func readRaw() {

	f, err := excelize.OpenFile("C:/Users/tathen/OneDrive - Dell Technologies/Nikhil/gofiles/src/github.com/go/Nikhil_Momentum/Data/Momentum back testing_0_104.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("NA->0", "A2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("NA->0")
	if err != nil {
		fmt.Println(err)
		return
	}

	nonPriceColums := 5
	nonStockRows := 3
	//Read 1st 10 rows
	for i, row := range rows {
		stock := m.StockPrices{}
		stock.Prices = make([]map[int]string, 0)

		if i >= nonStockRows {
			for j, colCell := range row {
				// fmt.Print( colCell, "\t")
				// fmt.Print(j, ":", colCell, "\t")

				switch j {
				case 0:
					stock.Index = colCell
				case 1:
					stock.Name = colCell
				case 2:
					stock.Industry = colCell
				case 3:
					stock.Symbol = colCell
				case 4:
					stock.MCap = colCell
				default:

					price := make(map[int]string)
					// weekNo := strconv.Itoa(j - nonPriceColums)
					weekNo := j - nonPriceColums
					price[weekNo] = colCell
					stock.Prices = append(stock.Prices, price)
					if j == 108 {
						allData.AllRows = append(allData.AllRows, stock)
					}
				}
			}
		}
		if i == 3 {
			break
		}
		fmt.Println()
	}
}

func GetPriceChanges(changePeriod int) m.RawData {
	var changePerctData m.RawData
	fmt.Printf("Finding price changes over %d period\n", changePeriod)
	changePerctData.AllRows = make([]m.StockPrices, 0)

	for i := 0; i < len(allData.AllRows); i++ {
		fmt.Println("Getting Price changes for ", allData.AllRows[i].Symbol)
		fmt.Printf("%#v", allData.AllRows[i].Prices)
		fmt.Println()
		fmt.Println()
		fmt.Println("0 ", allData.AllRows[i].Prices[0][0])
		fmt.Println("1 ", allData.AllRows[i].Prices[1][1])
		fmt.Println("2 ", allData.AllRows[i].Prices[2][2])
		fmt.Println("3 ", allData.AllRows[i].Prices[3][3])

		fmt.Println("len ", len(allData.AllRows[i].Prices))
		// for key, v := range allData.AllRows[i].Prices {
		for key := 0; key < len(allData.AllRows[i].Prices); key++ {
			v := allData.AllRows[i].Prices

			fmt.Println("Key :", key)
			if key-changePeriod >= 0 {

				newPrice := v[key][key]
				oldPrice := v[key-changePeriod][key-changePeriod]

				// fmt.Printf(" Price : %#v", oldPrice)
				// fmt.Printf(" Price : %#v", newPrice)

				newPriceVal, err := strconv.ParseFloat(newPrice, 32)
				if err != nil {
					fmt.Println("Failed to get new Price")
					os.Exit(2)
				}
				oldPriceVal, err := strconv.ParseFloat(oldPrice, 32)
				if err != nil {
					fmt.Println("Failed to get old Price")
					os.Exit(2)
				}
				// fmt.Println("Week : ", key-changePeriod, " Price :", oldPriceVal)
				// fmt.Println("Week : ", key, " Price :", newPriceVal)
				change := 100 * (newPriceVal - oldPriceVal) / oldPriceVal
				fmt.Println("", change, "%")
			}
		}
		fmt.Println()
	}

	return changePerctData
}
