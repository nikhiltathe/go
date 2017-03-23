package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	set := strings.Fields(s)
	fmt.Println("set is ",set, "and length is ", len(set))
	var st_map map[string]int = make(map[string]int)
	for i := range set {
		count := 0
		for j := range set {
			if set[i] == set[j] {
				count = count + 1
			}
		}
		st_map[set[i]] = count
	}
	return st_map
	//return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
