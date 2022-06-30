package reportexporter

type IReportExporter interface {
	SetComponents(comp []string)
	SetSubComponents(subComp []string)
	SetDuration(dur string)
	GenerateReport()
}

func GetReportExporter(reportType string) IReportExporter {
	switch reportType {
	case "csv":
		return CSVReportExporter{}
	case "pdf":
		return PDFReportExporter{}
	case "html":
		return HTMLReportExporter{}
	default:
		return nil
	}
}
