package main

import "fmt"

func totalNQueens(n int) int {
	res := 0
	if n == 0 {
		return 0
	}

	board := [][]string{}
	for i := 0; i < n; i++ {
		row := []string{}
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		board = append(board, row)
	}
	row := 0
	backtrack52(board, row, &res)
	return res
}

func backtrack52(board [][]string, row int, res *int) {
	fmt.Printf("row %+v\n", row)
	if row == len(board) {
		(*res)++
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			fmt.Printf("board col row %+v %+v\n", col, row)
			board[row][col] = "Q"
			backtrack52(board, row+1, res)
			board[row][col] = "."
		}
	}
}

func isValid52(board [][]string, row int, col int) bool {
	//check the row above to the first row of same col
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == "Q" {
			return false
		}
	}
	//check 右上角
	m, n := row-1, col+1
	for m >= 0 && n < len(board) {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n++
	}
	//check 左上角
	m, n = row-1, col-1
	for m >= 0 && n >= 0 {
		if board[m][n] == "Q" {
			return false
		}
		m--
		n--
	}

	return true
}
