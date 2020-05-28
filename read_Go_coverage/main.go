package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var centOSHost = "ldpda089"
var vaultDD = "ldpda162.hop.lab.emc.com"

type linesData struct {
	lineVal string
}

func main() {
	if len(os.Args) != 3 {
		showUsages()
	}
	strings.Join(os.Args, " ")

	// fmt.Println(" Lines :", size)
	currentCoverage := getCoveragePercentage(os.Args[1])
	// fmt.Print("Current coverage : ", currentCoverage, "%\n")

	jenkinsCoverage := getCoverageFromJenkinsFile(os.Args[2])
	// fmt.Print("jenkins coverage : ", jenkinsCoverage, "%\n")

	difference := currentCoverage - jenkinsCoverage 
	if difference > 0.5 {
		fmt.Print("Congratualtions!!!! You have increased coverage from ", jenkinsCoverage, "% to ", currentCoverage,
			"% .Do consider updating JenkinsFile to have your efforts counted")
	} else if currentCoverage < jenkinsCoverage {
		fmt.Print("You have unit test coverage ", currentCoverage, "% compared to Jenkins coverage ", jenkinsCoverage,
			"%. Do Add more unit tests to get your PR checked in ")

	}
}

func showUsages() {
	fmt.Println("Usage:\n \t Counter.exe <filename.txt> <JenkinsFile>")
}

func getCoveragePercentage(filename string) (coverage float64) {

	//Read file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	size := len(lines)
	if size > 2 {

	} else {
		fmt.Println("Not enought information")
		os.Exit(1)
	}

	line := lines[size-1]
	// fmt.Println("Last Line :", line)
	key := "(statements)"
	pos := strings.Index(line, key)
	if pos > 0 {
		valueStr := line[pos+len(key) : len(line)-1]
		valueStr = strings.Trim(valueStr, " ")
		// fmt.Println(valueStr)
		var err error
		coverage, err = strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("Failed to get coverage ", err.Error())
			os.Exit(1)
		}
		// coverage = float64(coverageInt)
	}
	return coverage
}

func getCoverageFromJenkinsFile(filename string) (coverage float64) {

	//Read file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	// size := len(lines)

	for _, v := range lines {
		key := "lineCoverageTargets: '"
		pos := strings.Index(v, key)
		if pos > 0 {
			endStr := "',"
			end := strings.Index(v, endStr)
			// fmt.Println("Line is ", v)
			valueStr := v[pos+len(key) : end]

			// fmt.Println("All values ", valueStr)
			values := strings.Split(valueStr, ",")
			valuesSize := len(values)
			if valuesSize != 3 {
				fmt.Println("Can not find 3 values for ", key)
				os.Exit(1)
			}

			two := values[1]
			two = strings.Trim(two, " ")
			twoVal, err := strconv.ParseFloat(two, 64)
			if err != nil {
				fmt.Println("Failed to get coverage ", err.Error())
				os.Exit(1)
			}

			three := values[2]
			three = strings.Trim(three, " ")
			threeVal, err := strconv.ParseFloat(three, 64)
			if err != nil {
				fmt.Println("Failed to get coverage ", err.Error())
				os.Exit(1)
			}

			// fmt.Println(two)
			// fmt.Println(three)

			coverage = threeVal
			if twoVal < threeVal {
				coverage = twoVal
			}
		}
	}
	return coverage
}
