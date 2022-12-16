package main

// // Juniper

// import "fmt"

// // True/ false
// // number of courses : 3
// // 2D arry [1,2] [2,3] > true
// // 2D arry [1,2] [2,3] [3,1] > false

// func main_3() {

// 	num := 3
// 	courseDeps := [][]int{[]int{1, 2}, []int{2, 3}}
// 	fmt.Println(num)
// 	fmt.Println(courseDeps)
// 	result := DetectCycle(num, courseDeps)
// 	fmt.Println("Possibility of course completion", !result)
// }

// type Elements struct {
// 	Curr       int
// 	LastCourse int
// }

// func DetectCycle(num int, courses [][]int) bool {

// 	queue := make([]Elements, 0)

// 	visited := make([]bool, num)

// 	for i := 1; i <= num; i++ {
// 		fmt.Println("visited :", visited)

// 		if !visited[i] {
// 			queue = append(queue, Elements{i, 0})
// 			fmt.Println("queue :", queue)

// 			for len(queue) > 0 {
// 				var curr Elements
// 				curr, queue = queue[0], queue[1:]
// 				// queue = queue[0 : len(queue)-1]
// 				fmt.Println("queue after pop :", queue)
// 				visited[curr.Curr-1] = true

// 				cycle := CheckCycleDFS(curr, courses[curr.Curr][0], courses, visited)
// 				if cycle {
// 					return true
// 				}

// 				queue = append(queue, Elements{courses[curr.Curr][0], i})
// 				fmt.Println("queue appended :", queue)

// 			}
// 		}
// 	}
// 	// O(N : M)

// 	return false
// }

// func CheckCycleDFS(currCourse Elements, num int, courses [][]int, visited []bool) bool {
// 	fmt.Println("visited in CheckCycleDFS :", visited)

// 	next := courses[currCourse.Curr][0]
// 	fmt.Println("next", next)

// 	if visited[next-1] == true && currCourse.LastCourse != num {
// 		return true
// 	}
// 	return false
// }
