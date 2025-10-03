// https://leetcode.com/problems/maximum-subarray/description/

package main

import (
	"fmt"
	"math"
)

func main() {

	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	input = []int{-2, -3, 4, -1, -2, 1, 5, 3}
	ans := maxSubArray(input)
	fmt.Println(ans)
}

func maxSubArray(nums []int) int {

	max := math.MinInt32
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		sum += nums[i]

		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
