// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

import "fmt"

func main() {

	// input := "abcabcbb"
	// input := "bbbbb"
	input := "pwwkew"
	ans := lengthOfLongestSubstring(input)
	fmt.Println(ans)
}

func lengthOfLongestSubstring(input string) int {

	if len(input) <= 1 {
		return 1
	}

	maxLength := 0
	mapCharvsIndex := make(map[string]int)
	left := 0
	right := 0

	for right < len(input)-1 {
		character := string(input[right])
		index, ok := mapCharvsIndex[character]
		if ok {
			left = index + 1
		}
		mapCharvsIndex[character] = right
		len := right - left + 1
		maxLength = max(maxLength, len)
		right++
	}

	return maxLength
}
