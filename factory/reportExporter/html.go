package reportexporter

import "fmt"

type HTMLReportExporter struct {
}

func (reportExporter HTMLReportExporter) captureData() {
	fmt.Println("Getting Data for HTML export")
}

func (reportExporter HTMLReportExporter) GenerateReport() {
	reportExporter.captureData()
	fmt.Println("Generated HTML report")
}

func (reportExporter HTMLReportExporter) SetComponents(comp []string) {
	fmt.Println("Setting Components for HTML report as ", comp)
}

func (reportExporter HTMLReportExporter) SetSubComponents(subComp []string) {
	fmt.Println("Setting sub components for HTML report as ", subComp)
}

func (reportExporter HTMLReportExporter) SetDuration(dur string) {
	fmt.Println("Setting duration for HTML report ", dur)
}
