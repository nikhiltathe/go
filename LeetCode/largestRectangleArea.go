// https://leetcode.com/problems/largest-rectangle-in-histogram/
package main

import "fmt"

func main() {

	input := []int{2, 1, 5, 6, 2, 3}
	ans := largestRectangleArea(input)
	fmt.Println(ans)
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []int{}

	heights = append(heights, 0) // Append a zero height to flush out remaining bars in stack

	for i, h := range heights {
		// While current bar is smaller than the top of the stack,
		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			// pop stack and calculate area with popped bar as smallest
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := width * heights[idx]
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
