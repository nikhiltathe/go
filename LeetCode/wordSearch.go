// https://www.youtube.com/watch?v=pfiQ_PS1g8E

package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word1 := "ABCCED"
	word2 := "SEE"
	word3 := "ABCB"

	fmt.Println("Searching for:", word1, "->", exist(board, word1))
	fmt.Println("Searching for:", word2, "->", exist(board, word2))
	fmt.Println("Searching for:", word3, "->", exist(board, word3))
}

// standalone dfs function
func dfs(board [][]byte, word string, i, j, k int) bool {
	m, n := len(board), len(board[0])
	if k == len(word) {
		return true
	}
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
		return false
	}
	temp := board[i][j]
	board[i][j] = '#' // mark as visited
	found := dfs(board, word, i+1, j, k+1) ||
		dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) ||
		dfs(board, word, i, j-1, k+1)
	board[i][j] = temp // unmark
	return found
}

// main exist function
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
