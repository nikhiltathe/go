package main

import (
	"fmt"
	"math"
)

func main() {

	nums := []int{10, 20, 30, 40, 50}
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			if i != j {
				dp[i][j] = -1
			}
		}
	}

	for i := n - 1; i > 0; i-- {
		for j := i; j < n; j++ {
			mini := math.MaxInt
			for k := i; k < j; k++ {
				// fmt.Println(i, j, k)
				fmt.Println("Products of Matrices ", i, " is ", nums[i-1], "*", nums[k], " X ", nums[k], "*", nums[j])
				steps := nums[i-1]*nums[k]*nums[j] + dp[i][k] +
					dp[k+1][j]
				if steps < mini {
					mini = steps
					dp[i][j] = mini
				}
			}
		}
	}
	fmt.Println("MCM is ", dp[1][n-1])
}

// func main() {

// 	nums := []int{10, 20, 30, 40, 50}
// 	n := len(nums)
// 	dp := make([][]int, n)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 		for j := range dp {
// 			dp[i][j] = -1
// 		}
// 	}
// 	fmt.Println("MCM is ", f(1, n-1, nums, dp))
// }

func f(i, j int, nums []int, dp [][]int) int {
	if i == j {
		return 0
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	mini := math.MaxInt
	for k := i; k < j; k++ {
		fmt.Println("Products of Matrices ", i, " is ", nums[i-1], "*", nums[k], " X ", nums[k], "*", nums[j])
		steps := nums[i-1]*nums[k]*nums[j] + f(i, k, nums, dp) +
			f(k+1, j, nums, dp)
		if steps < mini {
			mini = steps
			dp[i][j] = mini
		}
	}
	// dp[i][j] = mini
	return mini
}
