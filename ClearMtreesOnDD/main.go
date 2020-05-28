package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var centOSHost = "ldpda089"
var vaultDD = "ldpda162.hop.lab.emc.com"

type linesData struct {
	lineVal string
}

func main() {
	// fmt.Println("Hello gopher !")

	// fmt.Println(len(os.Args))

	if len(os.Args) == 1 {
		showUsages()
	}
	strings.Join(os.Args, " ")
	filename := os.Args[1]

	//Read file
	allLines := readFile(filename)
	fmt.Println("Lines : ", len(allLines))

	op_filename := "output.txt"
	f, err := os.Create(op_filename)
	check(err)
	defer f.Close()

	//fmt.Printf("%+v", allLines)
	writeTofile(f, []string{
		"-------------------------------------------------------",
		"Run on DD",
		"-------------------------------------------------------",
	})
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	outputLines := nfsAdd(allLines)
	writeTofile(f, outputLines)

	writeTofile(f, []string{
		"-------------------------------------------------------",
		"Run on CentOS",
		"-------------------------------------------------------",
	})
	outputLines = mountAndCleanup(allLines)
	writeTofile(f, outputLines)

	writeTofile(f, []string{
		"-------------------------------------------------------",
		"Run on DD",
		"-------------------------------------------------------",
	})
	outputLines = nfsDel(allLines)
	writeTofile(f, outputLines)

	outputLines = mtreeDelete(allLines)
	writeTofile(f, outputLines)

	// for _, v := range outputLines {
	// 	fmt.Println(v)
	// }

}

func readFile(filename string) []linesData {
	data, err := ioutil.ReadFile(filename)
	var allLines []linesData

	if err != nil {
		fmt.Println("File reading error", err)
		return allLines
	}
	//fmt.Println("Contents of file:\n", string(data))
	lines := strings.Split(string(data), "\n")

	for _, v := range lines {
		//fmt.Println("Line : ", v)
		var repos []string
		if strings.Index(v, "cr-policy-") > 0 {
			// values := strings.Split(v, "\t")
			dataIndex := strings.Index(v, "/data/col1")
			repoLocation := strings.Index(v, "-repo")
			repoLocation += 5
			repo := v[dataIndex:repoLocation]
			// fmt.Println(" Policy Repo : ", repo)
			repos = append(repos, repo)
			data := linesData{lineVal: repo}
			allLines = append(allLines, data)
			//values := strings.Split(v, "/data/col1/")

			//for _, vv := range values {
			//		fmt.Println("VV : ", vv)
			//
			//}
		}
		//allLines = append(allLines, linesData{lineVal: values[0]})
	}
	//fmt.Printf("%+v", allLines)
	fmt.Println()
	return allLines
}

func showUsages() {
	fmt.Println("Usage:\n \t CleanupUtility.exe <filename.txt>")
}

func mtreeDelete(data []linesData) (commands []string) {
	for _, v := range data {
		command := "mtree delete " + v.lineVal + "\nyes"
		// fmt.Println(command)
		commands = append(commands, command)
	}
	return commands
}

func nfsAdd(data []linesData) (commands []string) {

	for _, v := range data {
		command := "nfs add " + v.lineVal + " " + centOSHost
		//fmt.Println(command)
		commands = append(commands, command)
	}
	return commands
}

func nfsDel(data []linesData) (commands []string) {
	for _, v := range data {
		command := "nfs del " + v.lineVal
		fmt.Println(command)
		commands = append(commands, command)
	}
	return commands
}

func mountAndCleanup(data []linesData) (commands []string) {
	dir := "/nikhil/DD"
	command := "cd " + dir
	// fmt.Println(command)
	commands = append(commands, command)

	folderName := "CleanDD"
	command = "mkdir " + folderName
	// fmt.Println(command)
	commands = append(commands, command)

	// command = "cd " + foldeName
	// fmt.Println(command)

	for _, v := range data {
		command := "mount " + vaultDD + ":" + v.lineVal + " " + folderName
		// fmt.Println(command)
		commands = append(commands, command)

		command = "ls -al " + folderName + "/*"
		// fmt.Println(command)
		commands = append(commands, command)

		command = "rm " + folderName + "/* -rf"
		// fmt.Println(command)
		commands = append(commands, command)

		command = "ls -al " + folderName + "/*"
		// fmt.Println(command)
		commands = append(commands, command)

		command = "umount " + folderName
		// fmt.Println(command)
		commands = append(commands, command)

	}
	return commands
}

func writeTofile(f *os.File, lines []string) {
	for _, v := range lines {
		_, err := f.WriteString(v + "\n")
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
