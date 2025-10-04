// https://leetcode.com/problems/permutations/description/
package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	fmt.Println(permute(input))
}

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, used, &res)
	return res
}

func backtrack(nums []int, curr []int, used []bool, res *[][]int) {

	if len(curr) == len(nums) {
		comb := make([]int, len(curr))
		copy(comb, curr)
		*res = append(*res, curr)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == true {
			continue
		}

		used[i] = true
		newPath := append(curr, nums[i])
		backtrack(nums, newPath, used, res)
		used[i] = false
	}
}
