package main

import (
	"fmt"
	"sort"
	"sync"
)

func main_2() {
	arr := []int{10, 20, 3, 50, 500}

	index := 2
	// out := make(chan []int)
	// val := make(chan int)
	var out []int
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// defer close(out)
		// Sorter(index, arr, out, val, &wg)
		out = Sorter(index, arr, &wg)

	}()
	// fmt.Printf("\n%+T", out)

	wg.Wait()
	fmt.Println(out)

	// for _, v := range <-out {
	// 	fmt.Println(v)
	// }
	// v := <-val
	// fmt.Println("Value :", v)

	f1 := Frog{Type: "ABCD", Animal: Animal{Name: "XYZ"}}
	f1.Jump()
	a := Animal{Name: "saefe"}
	a.Jump()

	abc := GetAnimal("frog")
	abc.Jump()
	abc = GetAnimal("fox")
	abc.Jump()
}

func Sorter(index int, arr []int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	sort.Ints(arr)
	return arr
	// out <- arr
	// val <- arr[index-1]
	// fmt.Println(arr)
}

// func Sorter(index int, arr []int, out chan []int, val chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	sort.Ints(arr)
// 	out <- arr
// 	val <- arr[index-1]
// 	// fmt.Println(arr)
// }

type Frog struct {
	Animal
	Type string
}

type Fox struct {
	Animal
	Type string
}

type Animal struct {
	Name string
}

type actions interface {
	Jump()
}

func (a Animal) Jump() {
	fmt.Println("Animal jumping")
}

func (f Frog) Jump() {
	fmt.Println("Frog jumping")
}

func (f Fox) Jump() {
	fmt.Println("Fox jumping")
}

func GetAnimal(method string) actions {
	if method == "frog" {
		return Frog{}
	} else if method == "fox" {
		return Fox{}
	}
	return Animal{}
}
