package main

import "fmt"

func getPermutation(n int, k int) string {
	per := make([]int, n)
	for i := 0; i < n; i++ {
		per[i] = i + 1
	}
	fmt.Printf("per %+v\n", per)
	res := [][]int{}
	cur := []int{}
	visited := make([]bool, n)
	backtrack60(per, visited, cur, &res)

	fmt.Printf("res %+v\n", res)
	return "A"
	//return res[k]
}

func backtrack60(per []int, visited []bool, cur []int, res *[][]int) {
	if len(cur) == len(per) {
		dc := make([]int, len(cur))
		fmt.Printf("cur %+v\n", cur)
		copy(dc, cur)
		fmt.Printf("dc %+v\n", dc)
		*res = append(*res, dc)
		return
	}

	for i := 0; i < len(per); i++ {
		if visited[i] {
			continue
		}
		cur = append(cur, per[i])
		visited[i] = true
		backtrack60(per, visited, cur, res)
		visited[i] = false
		cur = cur[:len(cur)-1]
	}
}
