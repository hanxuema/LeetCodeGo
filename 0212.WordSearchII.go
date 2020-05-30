package main

import "fmt"

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	if board == nil || len(board) == 0 || words == nil || len(words) == 0 {
		return res
	}
	dir := make([][]int, 4)
	dir[0] = []int{0, 1}
	dir[1] = []int{1, 0}
	dir[2] = []int{0, -1}
	dir[3] = []int{-1, 0}

	dic := make(map[string]int)
	for _, v := range words {
		dic[v]++
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			backtrack0212(board, i, j, dic, dir, "", &res)
		}
	}
	return res
}

func backtrack0212(board [][]byte, i int, j int, dic map[string]int, dir [][]int, cur string, res *[]string) {
	fmt.Printf("cur %+v\n", cur)
	_, ok := dic[cur]
	if ok {
		*res = append(*res, cur)
		return
	}

	for idx := 0; idx < len(dir); idx++ {
		if board[i][j] == '.' {
			continue
		}

		temp := board[i][j]
		board[i][j] = '.'

		cur = cur + string(temp)

		n := i + dir[idx][0]
		m := j + dir[idx][1]
		fmt.Printf("n %+v\n", n)
		fmt.Printf("m %+v\n", m)
		
		if n < 0 || n >= len(board) || m < 0 || m >= len(board[n]) || board[n][m] == '.' {
			continue
		}
		fmt.Printf("board[n][m] %+v\n", string(board[n][m]))
		backtrack0212(board, n, m, dic, dir, cur, res)

		board[i][j] = temp
		cur = cur[:len(cur)-1]
	}
}
