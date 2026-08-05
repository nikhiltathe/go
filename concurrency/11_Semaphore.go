package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var sem = semaphore.NewWeighted(10) // Limit to 10 concurrent goroutines

func processTask(id int, wg *sync.WaitGroup) {
	sem.Acquire(context.Background(), 1) // Acquire a "slot"
	defer sem.Release(1)                 // Release the slot

	fmt.Printf("Worker %d processing task\n", id)
	// ... do some work ...
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d Done task\n", id)
	wg.Done()
}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go processTask(i, wg)
	}
	// ... wait for goroutines to finish ...
	wg.Wait()
}
