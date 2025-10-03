// https://leetcode.com/problems/product-of-array-except-self/
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	ans := productExceptSelf(nums)
	fmt.Println(ans)
}

func productExceptSelf(nums []int) []int {
	// 1. Create an empty answer array
	ans := make([]int, len(nums))
	postFix := make([]int, len(nums))

	product := 1
	for i := len(nums) - 1; i >= 0; i-- {
		postFix[i] = product
		product = product * nums[i]
	}
	fmt.Println(postFix)

	product = 1
	for i := 0; i < len(nums); i++ {
		ans[i] = postFix[i] * product
		product = product * nums[i]
	}

	return ans
}
