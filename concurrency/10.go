//Google search

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakesearch("web")
	Image = fakesearch("image")
	Video = fakesearch("video")
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elasped := time.Since(start)
	fmt.Println(results)
	fmt.Println(elasped)
}

type Search func(query string) Result

func fakesearch(kind string) Search {
	return func(query search) Result {
		time.Sleep(time.Duration(rand.Intn(100) * time.Millisecond))
		return Result(fmt.Sprintf("%s resukt for %q\n", kind, query))
	}
}
