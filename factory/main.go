package main

import (
	"fmt"
	reportexporter "go/factory/reportExporter"
)

func main() {

	obj1 := reportexporter.GetReportExporter("pdf")
	obj1.SetDuration("2W")
	obj1.SetComponents([]string{"jobs"})
	obj1.SetComponents([]string{"failed", "Analyze"})
	obj1.GenerateReport()

	fmt.Println()
	obj2 := reportexporter.GetReportExporter("html")
	obj1.SetDuration("1W")
	obj1.SetComponents([]string{"capacity"})
	obj1.SetComponents([]string{"failed", "Analyze"})
	obj2.GenerateReport()
}
