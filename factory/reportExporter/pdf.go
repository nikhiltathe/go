package reportexporter

import "fmt"

type PDFReportExporter struct {
}

func (reportExporter PDFReportExporter) captureData() {
	fmt.Println("Getting Data for PDF export")
}

func (reportExporter PDFReportExporter) GenerateReport() {
	reportExporter.captureData()
	fmt.Println("Generated PDF report")
}

func (reportExporter PDFReportExporter) SetComponents(comp []string) {
	fmt.Println("Setting Components for PDF report as ", comp)
}

func (reportExporter PDFReportExporter) SetSubComponents(subComp []string) {
	fmt.Println("Setting sub components for PDF report as ", subComp)
}

func (reportExporter PDFReportExporter) SetDuration(dur string) {
	fmt.Println("Setting duration for PDF report ", dur)
}
