package main

import (
	"fmt"

	"github.com/go/wrapper/wrapper"
)

func main() {
	obj := wrapper.GetWrapper()
	fmt.Printf("obj %+T\n", obj)
	obj.Say()
	obj2 := obj.Give(5)
	fmt.Printf("obj2 %+T\n", obj2)
	obj2.SayA1()
}
