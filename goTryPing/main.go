package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	// argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	// arg := os.Args[1]

	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

	out, _ := exec.Command("ping", argsWithoutProg[0], "-c 5", "-i 3", "-w 10").Output()
	// fmt.Println(err)
	//out, _ := exec.Command("ping", "192.168.0.111", "-c 5", "-i 3", "-w 10").Output()
	if strings.Contains(string(out), "Destination Host Unreachable") || strings.Contains(string(out), "Please check the name and try again.") {
		fmt.Println("TANGO DOWN")
	} else {
		fmt.Println("IT'S ALIVEEE")
	}

}
