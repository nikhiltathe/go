//Timeout using Select
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Printf("You say %q\n", s)
		case <-timeout:
			fmt.Println("You sare too slow !")
			return
		}
	}
	fmt.Println("Done listening")
}

func boring(msg string) <-chan string { // Returns receive only chan
	c := make(chan string)
	go func() { // Launch go routine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c
}
