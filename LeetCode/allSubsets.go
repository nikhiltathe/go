// https://leetcode.com/problems/subsets/description/
package main

import "fmt"

func main() {

	input := []int{1, 2, 3}
	ans := subsets(input)
	fmt.Println(ans)
}

func subsets(nums []int) [][]int {
	res := [][]int{{}} // start with empty subset

	for _, num := range nums {
		n := len(res)
		for i := 0; i < n; i++ {
			newSubset := append([]int{}, res[i]...) // copy current subset
			fmt.Println("i =", i, newSubset)
			newSubset = append(newSubset, num) // add current number
			fmt.Println(newSubset)
			res = append(res, newSubset)
		}
	}

	return res
}

// func subsets(nums []int) [][]int {
// 	ans := [][]int{}
// 	subset := []int{}
// 	backtrack(nums, 0, subset, &ans)
// 	return ans
// }

// func backtrack(nums []int, start int, subset []int, res *[][]int) {
// 	temp := make([]int, len(subset))
// 	copy(temp, subset)
// 	*res = append(*res, temp)

// 	for i := start; i < len(nums); i++ {
// 		subset = append(subset, nums[i])
// 		backtrack(nums, i+1, subset, res)
// 		subset = subset[:len(subset)-1]
// 	}
// }
