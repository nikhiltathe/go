package main

import (
	"fmt"
	// "strings"
)

func main() {
	// InputString := "Welcome to Pubmatic"
	// Reverse(InputString)
	mat1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	mat2 := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	result := mulityply(mat1, mat2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(result[i][j], " ")
		}
		fmt.Println()
	}
}

func mulityply(mat1 [][]int, mat2 [][]int) [][]int {

	res := make([][]int, 3)
	for i := 0; i < 3; i++ {
		row := make([]int, 3)
		res[i] = row
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				res[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}
	return res
}

// func Reverse(s string) {
// 	words := strings.Split(s, " ")
// 	len := len(words)

// 	for i := len - 1; i >= 0; i-- {
// 		fmt.Print(words[i], " ")
// 	}
// 	fmt.Println()
// }
/*
==============

[
 a1 a2 a3
 b1 b2 b3
 c1 c2 c3
]

[
	x1 x2 x3
	y1 y2 y3
	z1 z2 z3
]


[
	a1*x1+a2*y1+a3*z1 a1*x2+a2*y2+a3*z2 a1*x3+a2*y3+a3*z3
	b1*x1+b2*y1+b3*z1 b1*x2+b2*y2+b3*z2 b1*x3+b2*y3+b3*z3
	c1*x1+c2*y1+c3*z1 c1*x2+c2*y2+c3*z2 c1*x3+c2*y3+c3*z3
]
*/
