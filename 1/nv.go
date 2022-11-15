package main

import "fmt"

/*
 * Given a sorted integer array Arr, two integers N and X,
 * return the N closest integers to X from the input array Arr.
 * An integer a is closer to x than an integer b if |a - x| < |b - x|,
 * OR |a - x| == |b - x| and a < b.
 */
func findClosest(arr []int, n int, x int) []int {
	if n <= 0 || n > len(arr) {
		return []int{}
	} else if n == len(arr) {
		return arr
	}

	out := make([]int, 0)
	size := len(arr)
	start := 0
	end := size - 1
	mid := (start + end) / 2
	for start <= end {
		fmt.Println(start, " ", end)
		start_diff := mod(arr[start], x)
		end_diff := mod(arr[end], x)
		fmt.Println("diff", start_diff, " ", end_diff)

		if start_diff < end_diff && mid != start {
			end = mid // mid - 1
			mid = (start + end) / 2
		} else if mid != start {
			start = mid // mid  + 1
			mid = (start + end) / 2
		} else {
			fmt.Println("All checks done")
			break
		}
		// mid = (start+end)/2
	}

	// mid is closes to x now
	fmt.Println("Mid index is :", mid)
	out = append(out, arr[mid])
	min := mid

	for i := 1; i < n && len(out) < n; i++ {
		if min-i >= 0 {
			out = append(out, arr[mid-i])
			min = mid - i
		}
		if len(out) < n {
			out = append(out, arr[mid+i])
		}
	}
	//  fmt.Println("Mid index is :",mid)

	// TODO: implement the solution here

	return arr[min : n-min]
}

func mod(a, b int) int {
	val := a - b
	if val < 0 {
		return 0 - val
	}
	return val
}

func main() {
	/*
	   // Arr, N, X
	    { in: [[1, 2, 3, 4, 5], 4, 3], out: [1, 2, 3, 4] },
	    { in: [[1, 2, 3, 4, 5], 4, -1], out: [1, 2, 3, 4] },
	    { in: [[1, 2, 3, 4, 5], 4, 10], out: [2, 3, 4, 5] },
	    { in: [[1, 2, 3, 4, 5], 1, 10], out: [5] },
	    { in: [[1, 2, 3, 4, 5], 1, 3], out: [3] },
	    { in: [[-2, -1, 0, 1, 2], 2, 1], out: [0, 1] },
	  ];
	*/
	//res := findClosest([]int{1, 2, 3, 4, 5}, 4, -1)
	res := findClosest([]int{-5, -1}, 1, -1)
	fmt.Println(res)
}

/*
N: 4, X : 3

arr: 1 2 3 4 5
diff 2 1 0 1 2

start : 1,2
mid   : 3,0
end   : 5,2

start : 1,2
mid   : 2,1
end   : 3,0
3->4


*/
