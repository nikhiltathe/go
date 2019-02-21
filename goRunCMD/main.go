package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	// https: //tutorialedge.net/golang/executing-system-commands-with-golang/
	fmt.Println("Usage : \n go run main.go 3 a b c\n\n")
	argCount := os.Args[1]
	fmt.Println("Inputs :", os.Args[1])
	count, err := strconv.Atoi(argCount)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	newCommand := ""
	for i := 2; i < count+2; i++ {
		// fmt.Println("Reading for index ", i)
		newCommand += os.Args[i]
		if i+1 != count+2 {
			newCommand += " "
		}
	}

	var commandArgs []string
	for i := 3; i < count+2; i++ {
		// fmt.Println("Reading for index ", i)
		commandArgs = append(commandArgs, os.Args[i])
	}

	fmt.Println("New command is :", newCommand, commandArgs, "length is", len(newCommand))

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command(newCommand).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	// // let's try the pwd command herer
	// out, err = exec.Command("pwd").Output()
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
	// fmt.Println("Command Successfully Executed")
	// output = string(out[:])
	// fmt.Println(output)
}
