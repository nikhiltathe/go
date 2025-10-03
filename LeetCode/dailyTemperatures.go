// LeetCode 739. Daily input
package main

import "fmt"

func main() {
	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	//input := []int{30, 40, 50, 60}
	ans := dailyTemperatures(input)
	fmt.Println(ans)
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	stack := []int{} // stores indices of warmer days

	for i := n - 1; i >= 0; i-- {
		// Pop indices with temperatures less than or equal to current temperature
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		// If stack not empty, top of stack is next warmer day
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}

		// Push current day index onto stack
		stack = append(stack, i)
	}

	return res
}
