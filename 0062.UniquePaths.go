package main

import "fmt"

func uniquePaths(m int, n int) int {
	res := 0
	if m == 0 || n == 0 {
		return res
	}

	backtrack62(m, n, 0, 0, &res)
	return res
}

func backtrack62(m int, n int, row int, col int, res *int) {
	fmt.Println("row, col ", row, col)
	if row > n || col > m {
		return
	}
	if row == n && col == m {
		*res++
		return
	}

	//add
	row++
	//backtrack
	backtrack62(m, n, row, col, res)
	//remove
	row--

	//add
	col++
	//backtrack
	backtrack62(m, n, row, col, res)
	//remove
	col--

}
