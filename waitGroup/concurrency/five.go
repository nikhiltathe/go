//Generator pattern
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-joe)
		fmt.Printf("You say %q\n", <-ann)
	}
	fmt.Println("Done listening")
}

func boring(msg string) <-chan string { // Returns receive only chan
	c := make(chan string)
	go func() {   // Launch go routine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c
}
