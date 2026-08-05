// Pattern - Fan in / multiplexer
package main

import (
	"fmt"
)

func main() {
	c := fanIn(boring(1), boring(2))
	for i := 0; i < 20; i++ {
		fmt.Printf("You say %d\n", <-c)
	}
	fmt.Println("Done listening")
}

func fanIn(input1, input2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg int) <-chan int { // Returns receive only chan
	c := make(chan int)
	go func() { // Launch go routine
		for i := msg; ; i++ {
			c <- i
			// time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			i++
		}

	}()
	return c
}
