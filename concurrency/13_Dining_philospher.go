package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var N int = 5

type Fork struct {
	Mutex sync.Mutex
}

type Philosopher struct {
	Num         int
	Left, Right *Fork
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {

		fmt.Println("Philosopher", p.Num, "wants to eat")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		p.Left.Mutex.Lock()
		p.Right.Mutex.Lock()

		fmt.Println("Philosopher", p.Num, "is eating...")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)

		p.Right.Mutex.Unlock()
		p.Left.Mutex.Unlock()

		fmt.Println("Philosopher", p.Num, "finishes eating [Tick]")
	}
}

func main() {
	fmt.Println("Dining philospher problem")

	// Init all forks
	forks := make([]*Fork, N)
	for i := 0; i < N; i++ {
		forks[i] = new(Fork)
	}

	// Init all philosphers
	philosophers := make([]*Philosopher, N)

	for i := 0; i < N; i++ {
		philosophers[i] = &Philosopher{i, forks[i], forks[(i+1)%N]}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(p *Philosopher) {
			p.eat(&wg)
		}(philosophers[i])
	}
	wg.Wait()
}

// Method 2
// The Dining Philosophers Problem is a classic concurrency problem that illustrates the challenges of resource allocation and deadlock prevention in multi-threaded environments. In Go, it can be implemented using goroutines and channels or mutexes.
// Problem Statement:
// Five philosophers are seated around a circular table, each with a plate of spaghetti. Between each pair of adjacent philosophers lies a single chopstick. To eat, a philosopher requires two chopsticks: one from their left and one from their right. Philosophers alternate between thinking and eating. The goal is to design a concurrent algorithm that allows all philosophers to eat without encountering deadlock or starvation. [1]
// Go Implementation using Mutexes:

// • Chopsticks as Mutexes: Each chopstick can be represented by a sync.Mutex.
// • Philosophers as Goroutines: Each philosopher can be a separate goroutine.
// • Picking up Forks: A philosopher attempts to acquire the mutexes for both their left and right chopsticks.
// • Eating: Once both chopsticks are acquired, the philosopher "eats" (simulated by a time.Sleep).
// • Putting down Forks: After eating, the philosopher releases both chopstick mutexes.

// Addressing Deadlock:
// A common deadlock scenario occurs if all philosophers simultaneously pick up their left chopstick and then wait indefinitely for their right chopstick, which is held by their neighbor. Solutions to prevent this include:

// • Asymmetric Acquisition: One philosopher picks up their right chopstick first, while others pick up their left first. This breaks the circular dependency.
// • Host/Waiter Solution: A central "host" or "waiter" controls access to the chopsticks, ensuring that no more than a certain number of philosophers (e.g., n-1) can hold chopsticks at any given time, preventing a full circular dependency.
// • Ordered Resource Hierarchy: Assign an order to the chopsticks and require philosophers to always acquire the lower-numbered chopstick before the higher-numbered one.

// Example Structure (conceptual):
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type Chopstick struct {
// 	sync.Mutex
// }

// type Philosopher struct {
// 	id             int
// 	leftChopstick  *Chopstick
// 	rightChopstick *Chopstick
// }

// func (p *Philosopher) Eat() {
// 	// Implement logic to acquire chopsticks, eat, and release
// 	// Consider deadlock prevention strategy here (e.g., ordered acquisition or host)
// 	fmt.Printf("Philosopher %d is eating...\n", p.id)
// 	time.Sleep(time.Second) // Simulate eating
// }

// func main() {
// 	numPhilosophers := 5
// 	chopsticks := make([]*Chopstick, numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		chopsticks[i] = &Chopstick{}
// 	}

// 	philosophers := make([]*Philosopher, numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		philosophers[i] = &Philosopher{
// 			id:             i,
// 			leftChopstick:  chopsticks[i],
// 			rightChopstick: chopsticks[(i+1)%numPhilosophers],
// 		}
// 	}

// 	var wg sync.WaitGroup
// 	for _, p := range philosophers {
// 		wg.Add(1)
// 		go func(phil *Philosopher) {
// 			defer wg.Done()
// 			for i := 0; i < 3; i++ { // Simulate eating multiple times
// 				phil.Eat()
// 				time.Sleep(time.Millisecond * 2000) // Simulate thinking
// 			}
// 		}(p)
// 	}
// 	wg.Wait()
// }

// AI responses may include mistakes.

// [1] https://www.geeksforgeeks.org/operating-systems/dining-philosophers-problem/

// Method 3
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// // Number of philosophers and chopsticks
// const numPhilosophers = 5
// const maxEatingPhilosophers = 2 // Host limits concurrent eaters

// // Chopstick struct
// type Chopstick struct {
// 	mu sync.Mutex
// }

// // Philosopher struct
// type Philosopher struct {
// 	id                            int
// 	leftChopstick, rightChopstick *Chopstick
// 	host                          *Host
// 	eatCount                      int
// }

// // Host struct
// type Host struct {
// 	eatingChannel chan struct{} // Channel to control the number of eating philosophers
// 	wg            sync.WaitGroup
// }

// // initHost initializes the host with the eating channel
// func initHost() *Host {
// 	return &Host{
// 		eatingChannel: make(chan struct{}, maxEatingPhilosophers),
// 	}
// }

// // dine represents a philosopher's attempt to eat
// func (p *Philosopher) dine() {
// 	defer p.host.wg.Done()

// 	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
// 		// Think
// 		fmt.Printf("Philosopher %d is thinking.\n", p.id)
// 		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

// 		// Ask host for permission to eat
// 		fmt.Printf("Philosopher %d wants to eat.\n", p.id)
// 		p.host.eatingChannel <- struct{}{} // Request permission

// 		// Acquire chopsticks (protected by the host's permission)
// 		p.leftChopstick.mu.Lock()
// 		p.rightChopstick.mu.Lock()

// 		// Eat
// 		fmt.Printf("Philosopher %d is eating...\n", p.id)
// 		p.eatCount++
// 		// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 		time.Sleep(2 * time.Second)

// 		// Release chopsticks
// 		p.rightChopstick.mu.Unlock()
// 		p.leftChopstick.mu.Unlock()

// 		// Release permission to the host
// 		<-p.host.eatingChannel

// 		fmt.Printf("Philosopher %d finished eating %d time(s).\n", p.id, p.eatCount)
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	// Create chopsticks
// 	chopsticks := make([]*Chopstick, numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		chopsticks[i] = &Chopstick{}
// 	}

// 	// Create host
// 	host := initHost()

// 	// Create philosophers and assign chopsticks
// 	philosophers := make([]*Philosopher, numPhilosophers)
// 	for i := 0; i < numPhilosophers; i++ {
// 		philosophers[i] = &Philosopher{
// 			id:             i + 1,
// 			leftChopstick:  chopsticks[i],
// 			rightChopstick: chopsticks[(i+1)%numPhilosophers],
// 			host:           host,
// 		}
// 		host.wg.Add(1)
// 		go philosophers[i].dine()
// 	}

// 	// Wait for all philosophers to finish
// 	host.wg.Wait()

// 	fmt.Println("All philosophers have finished dining.")
// }

// Method 4 without Host
/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Number of philosophers and chopsticks
const numPhilosophers = 5
const maxEatingPhilosophers = 2 // Host limits concurrent eaters

// Chopstick struct
type Chopstick struct {
	mu sync.Mutex
}

// Philosopher struct
type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
	host                          *Host
	eatCount                      int
}

// Host struct
type Host struct {
	eatingChannel chan struct{} // Channel to control the number of eating philosophers
	wg            sync.WaitGroup
}

// initHost initializes the host with the eating channel
func initHost() *Host {
	return &Host{
		eatingChannel: make(chan struct{}, maxEatingPhilosophers),
	}
}

// dine represents a philosopher's attempt to eat
func (p *Philosopher) dine() {
	defer p.host.wg.Done()

	for i := 0; i < 3; i++ { // Each philosopher eats 3 times
		// Think
		fmt.Printf("Philosopher %d is thinking.\n", p.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Ask host for permission to eat
		fmt.Printf("Philosopher %d wants to eat.\n", p.id)
		// p.host.eatingChannel <- struct{}{} // Request permission

		// Acquire chopsticks (protected by the host's permission)
		p.leftChopstick.mu.Lock()
		p.rightChopstick.mu.Lock()

		// Eat
		fmt.Printf("Philosopher %d is eating.\n", p.id)
		p.eatCount++
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Release chopsticks
		p.rightChopstick.mu.Unlock()
		p.leftChopstick.mu.Unlock()

		// // Release permission to the host
		// <-p.host.eatingChannel

		fmt.Printf("Philosopher %d finished eating %d time(s).\n", p.id, p.eatCount)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Create chopsticks
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{}
	}

	// Create host
	host := initHost()

	// Create philosophers and assign chopsticks
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			host:           host,
		}
		host.wg.Add(1)
		go philosophers[i].dine()
	}

	// Wait for all philosophers to finish
	host.wg.Wait()

	fmt.Println("All philosophers have finished dining.")
}
*/
