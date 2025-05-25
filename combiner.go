package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	dateFormat = "02-Jan-2006" // Date format used in the CSV file names
	headerLine = "Stock Name,Symbol,Exch,LTP,Chg%,Sector Name,Industry Name,Type,Close Within 52 Week High Zone,Stock Outperforming Benchmark Index in both 1 Week and 3 Month,Stock Outperforming Sectoral Index in both 1 Week and 3 Month,Strongly Outperforming Benchmark Index (55 Days),Strongly Outperforming Sector Index (55 Days),Both SRS and ARS Above Zero,High Return on Equity,Consistent Sales Growth Ratio (Annual Report),Consistent Cash From Operations,Increase in Institution shareholding,Symbol with Comma for External Upload"
)

// fileInfo stores the file path and its parsed date
type fileInfo struct {
	path string
	date time.Time
}

// ByDate implements sort.Interface for []fileInfo based on the date field.
type ByDate []fileInfo

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].date.Before(a[j].date) }

func main() {
	folderPaths := []string{
		"C:\\Users\\tathen\\OneDrive - Dell Technologies\\Nikhil\\Personal\\Share market\\stockedge\\eagle",
		"C:\\Users\\tathen\\OneDrive - Dell Technologies\\Nikhil\\Personal\\Share market\\stockedge\\ATH_Strong",
		"C:\\Users\\tathen\\OneDrive - Dell Technologies\\Nikhil\\Personal\\Share market\\stockedge\\Strong_FIPA_Value",
		"C:\\Users\\tathen\\OneDrive - Dell Technologies\\Nikhil\\Personal\\Share market\\stockedge\\HighVol_52W",
		"C:\\Users\\tathen\\OneDrive - Dell Technologies\\Nikhil\\Personal\\Share market\\stockedge\\HOLS",
		"C:\\Users\\tathen\\OneDrive - Dell Technologies\\Nikhil\\Personal\\Share market\\stockedge\\Momentum",
	}

	for _, folderPath := range folderPaths {

		files, err := ioutil.ReadDir(folderPath)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}

		csvFiles := []fileInfo{}
		re := regexp.MustCompile(`_(\d{2}-\w{3}-\d{4})\.csv`)

		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
				matches := re.FindStringSubmatch(file.Name())
				if len(matches) == 2 {
					date, err := time.Parse(dateFormat, matches[1])
					if err == nil {
						csvFiles = append(csvFiles, fileInfo{path: filepath.Join(folderPath, file.Name()), date: date})
					}
				}
			}
		}

		if len(csvFiles) == 0 {
			fmt.Println("No eaglemomentum CSV files found in the directory.")
			return
		}

		// Sort files by date to find the latest
		sort.Sort(ByDate(csvFiles))

		latestFile := csvFiles[len(csvFiles)-1]
		fmt.Printf("Latest file: %s (Date: %s)\n", latestFile.path, latestFile.date.Format(dateFormat))

		// Read content of the latest file
		latestFileContent, err := ioutil.ReadFile(latestFile.path)
		if err != nil {
			fmt.Printf("Error reading latest file %s: %v\n", latestFile.path, err)
			return
		}

		// Create a new target file or overwrite the existing latest file
		outputFile, err := os.Create(latestFile.path)
		if err != nil {
			fmt.Printf("Error creating/opening output file %s: %v\n", latestFile.path, err)
			return
		}
		defer outputFile.Close()

		writer := bufio.NewWriter(outputFile)

		// Write content of the latest file to the output file
		_, err = writer.WriteString(string(latestFileContent))
		if err != nil {
			fmt.Printf("Error writing latest file content: %v\n", err)
			return
		}

		// Add a blank line after the content of the latest file
		_, err = writer.WriteString("\n")
		if err != nil {
			fmt.Printf("Error writing blank line: %v\n", err)
			return
		}

		// Process other CSV files
		for _, file := range csvFiles {
			if file.path == latestFile.path {
				continue // Skip the latest file as its content is already copied
			}

			fmt.Printf("Processing file: %s\n", file.path)
			content, err := ioutil.ReadFile(file.path)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", file.path, err)
				continue
			}

			lines := strings.Split(string(content), "\n")
			headerFound := false
			for _, line := range lines {
				if strings.HasPrefix(line, "Stock Name") {
					headerFound = true
					_, err = writer.WriteString(line + "\n")
					if err != nil {
						fmt.Printf("Error writing header line: %v\n", err)
						break
					}
					continue
				}
				if headerFound {
					_, err = writer.WriteString(line + "\n")
					if err != nil {
						fmt.Printf("Error writing line from %s: %v\n", file.path, err)
						break
					}
				}
			}
			if headerFound {
				// Add a blank line after each appended file's content
				_, err = writer.WriteString("\n")
				if err != nil {
					fmt.Printf("Error writing blank line after %s: %v\n", file.path, err)
				}
			}
		}

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Error flushing writer: %v\n", err)
			return
		}

		fmt.Println("Data consolidation complete.")

		err = writer.Flush()
		if err != nil {
			fmt.Printf("Error flushing writer: %v\n", err)
			return
		}
	}
	fmt.Println("Data consolidation complete.")

}
