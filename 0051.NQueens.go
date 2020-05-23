package main

import "fmt"

func solveNQueens(n int) [][]string {
	res := [][]string{}
	if n == 0 {
		return res
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
	backtrack51(board, row, &res)
	return res
}

func backtrack51(board [][]string, row int, res *[][]string) {
	fmt.Printf("row %+v\n", row)
	if row == len(board) {
		fmt.Printf("row %+v\n", board)
		*res = append(*res, convert(board))
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			fmt.Printf("board col row %+v %+v\n", col, row)
			board[row][col] = "Q"
			backtrack51(board, row+1, res)
			board[row][col] = "."
		}
	}
}

func isValid(board [][]string, row int, col int) bool {
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

func convert(board [][]string) []string {
	res := []string{}

	for i := 0; i < len(board); i++ {
		row := ""
		for j := 0; j < len(board[i]); j++ {
			row += board[i][j]
		}
		res = append(res, row)
	}
	return res
}
