package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// sharedResource is the data that readers and writers access
	sharedResource int = 0
	// rwMutex protects the sharedResource
	rwMutex sync.RWMutex
)

// Reader goroutine
func reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire read lock
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	fmt.Printf("Reader %d reading: %d\n", id, sharedResource)
	// Simulate reading time
	time.Sleep(1 * time.Second)
}

// Writer goroutine
func writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Acquire write lock
	rwMutex.Lock()
	defer rwMutex.Unlock()

	sharedResource++
	fmt.Printf("Writer %d writing...: %d\n", id, sharedResource)
	// Simulate writing time
	time.Sleep(2 * time.Second)
	fmt.Printf("Writer %d writing Done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Create a few readers and writers
	numReaders := 50
	numWriters := 2

	// Launch reader goroutines
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go reader(i, &wg)
	}

	// Launch writer goroutines
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go writer(i, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All readers and writers finished.")
}
