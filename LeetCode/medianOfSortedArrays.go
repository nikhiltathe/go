// https://leetcode.com/problems/median-of-two-sorted-arrays/
package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	n1, n2 := len(nums1), len(nums2)
	low, high := 0, n1

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := (n1+n2+1)/2 - mid1

		L1, L2 := math.Inf(-1), math.Inf(-1)
		if mid1 > 0 {
			L1 = float64(nums1[mid1-1])
		}
		if mid2 > 0 {
			L2 = float64(nums2[mid2-1])
		}

		R1, R2 := math.Inf(1), math.Inf(1)
		if mid1 < n1 {
			R1 = float64(nums1[mid1])
		}
		if mid2 < n2 {
			R2 = float64(nums2[mid2])
		}

		if L1 <= R2 && L2 <= R1 {
			if (n1+n2)%2 == 0 {
				return (math.Max(L1, L2) + math.Min(R1, R2)) / 2.0
			}
			return math.Max(L1, L2)
		} else if L1 > R2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}
	return 0.0
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("Median is: %.2f\n", findMedianSortedArrays(nums1, nums2))

	nums3 := []int{1, 2}
	nums4 := []int{3, 4}
	fmt.Printf("Median is: %.2f\n", findMedianSortedArrays(nums3, nums4))
}
