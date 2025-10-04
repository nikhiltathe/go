// https://leetcode.com/problems/combination-sum/
package main

import "fmt"

func main() {
	input := []int{2, 3, 6, 7}
	ans := combinationSum(input, 7)
	fmt.Println(ans)
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	backtrack(candidates, target, 0, []int{}, &ans)
	return ans
}

func backtrack(candidates []int, target int, index int, combination []int, ans *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(combination))
		copy(temp, combination)
		*ans = append(*ans, temp)
		return
	}
	for i := index; i < len(candidates); i++ {
		combination = append(combination, candidates[i])
		backtrack(candidates, target-candidates[i], i, combination, ans)
		combination = combination[:len(combination)-1]
	}
}
