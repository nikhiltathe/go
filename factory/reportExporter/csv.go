package reportexporter

import "fmt"

type CSVReportExporter struct {
}

func (reportExporter CSVReportExporter) captureData() {
	fmt.Println("Getting Data for CSV export")
}

func (reportExporter CSVReportExporter) GenerateReport() {
	reportExporter.captureData()
	fmt.Println("Generated CSV report")
}

func (reportExporter CSVReportExporter) SetComponents(comp []string) {
	fmt.Println("Setting Components for CSV report as ", comp)
}

func (reportExporter CSVReportExporter) SetSubComponents(subComp []string) {
	fmt.Println("Setting sub components for CSV report as ", subComp)
}

func (reportExporter CSVReportExporter) SetDuration(dur string) {
	fmt.Println("Setting duration for CSV report ", dur)
}
