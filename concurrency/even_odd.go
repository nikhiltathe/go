package main

import (
	"fmt"
	"time"
)

const maxNum = 20

// func printOdd(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for n := 1; n <= maxNum; n += 2 {
// 		if n == 1 {
// 			fmt.Println("Odd:", n)
// 		} else {
// 			fmt.Println("Odd:", <-ch)
// 		}

// 		next := n + 1
// 		if next <= maxNum {
// 			ch <- next
// 		}
// 	}
// }

// func printEven(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for n := 2; n <= maxNum; n += 2 {
// 		fmt.Println("Even:", <-ch)

// 		next := n + 1
// 		if next <= maxNum {
// 			ch <- next
// 		}
// 	}
// }

func printOdd(ch chan int) {
	for n := 1; n <= maxNum; n += 2 {
		if n == 1 {
			fmt.Println("Odd:", n)
		} else {
			fmt.Println("Odd:", <-ch)
		}

		next := n + 1
		if next <= maxNum {
			ch <- next
		}
	}
}

func printEven(ch chan int) {
	for n := 2; n <= maxNum; n += 2 {
		fmt.Println("Even:", <-ch)

		next := n + 1
		if next <= maxNum {
			ch <- next
		}
	}
}

func main() {
	ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(2)

	// go printOdd(ch, &wg)
	// go printEven(ch, &wg)

	go printOdd(ch)
	go printEven(ch)
	// wg.Wait()
	time.Sleep(1 * time.Second)
}
