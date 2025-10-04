// https://www.youtube.com/watch?v=YK78FU5Ffjw

package main

import (
	"fmt"
)

func isValid(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val || board[i][col] == val {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == val {
			return false
		}
	}
	return true
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := byte('1'); c <= byte('9'); c++ {
					if isValid(board, i, j, c) {
						board[i][j] = c
						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	solve(board)
}

func main() {
	// Example Sudoku board
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println("Original Sudoku Board:")
	printBoard(board)

	solveSudoku(board)

	fmt.Println("\nSolved Sudoku Board:")
	printBoard(board)
}

func printBoard(board [][]byte) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

/*
Time Complexity		O(9m)	Worst-case exploring 9 possibilities per empty cell
Space Complexity	O(m)	Recursion stack up to number of empty cells
*/

// func main() {
// 	// Example Sudoku board
// 	board := [][]byte{
// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	}

// 	fmt.Println("Original Sudoku Board:")
// 	printBoard(board)

// 	solveSudoku(board)

// 	fmt.Println("\nSolved Sudoku Board:")
// 	printBoard(board)
// }

// func solveSudoku(board [][]byte) {
// 	solve(board)
// }

// func solve(board [][]byte) bool {

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {

// 			// If empty
// 			if board[i][j] == '.' {

// 				// Try putting every value
// 				for c := byte('1'); c <= byte('9'); c++ {
// 					// If this satisfies , put value in board
// 					if isValid(board, i, j, c) {
// 						board[i][j] = c
// 					}

// 					// This should lead further success too as there can be many combinations
// 					if solve(board) {
// 						return true
// 					} else {
// 						board[i][j] = '.'
// 						// Put back old value and try next number. Dont return
// 					}
// 				}
// 				// If no number is fitting means older was wrong, return false
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func isValid(board [][]byte, row int, col int, c byte) bool {
// 	for i := 0; i < 9; i++ {
// 		// check if this value is in any row or col
// 		if board[row][i] == c || board[i][col] == c {
// 			// || !isValidInBox(board, row, col, c)
// 			return false

// 			// check value in same box
// 			if board[(row/3)*3+i/3][(col/3)*3+i%3] == c {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func isValidInBox(board [][]byte, row int, col int, c byte) bool {
// 	rowStart := row/3 + 1
// 	rowEnd := rowStart + 3
// 	colStart := col/3 + 1
// 	colEnd := colStart + 3

// 	for i := rowStart; i < rowEnd; i++ {
// 		for j := colStart; j < colEnd; j++ {
// 			if board[i][j] == c {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func printBoard(board [][]byte) {
// 	for _, row := range board {
// 		for _, cell := range row {
// 			fmt.Printf("%c ", cell)
// 		}
// 		fmt.Println()
// 	}
// }
