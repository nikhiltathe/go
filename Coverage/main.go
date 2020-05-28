package main

import (
	"fmt"

	"github.com/go/Coverage/SVC1"
	"github.com/go/Coverage/SVC2"
)

func main() {
	one := SVC1.WhoAMI()
	two := SVC2.WhoAMI()
	fmt.Println(one, two)
}
