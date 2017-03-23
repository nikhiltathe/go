package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(2.)
	fmt.Println("z is ",z)
	s := float64(0)
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}
	return s
}

func main() {
	fmt.Println(Sqrt(4))
}
