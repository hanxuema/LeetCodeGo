package main

import (
	"strconv"
)

func getPermutation(n int, k int) string {
	res := make([]string, n+1)
	backtrack60(n, 1, "", &res)
	return res[k]
}

func backtrack60(n int, start int, cur string, res *[]string) {
	if len(cur) == n {
		// dc = make([]string, len(cur))
		// copy(dc, cur)
		// *res = append(*res, dc)
		// return
		*res = append(*res, cur)
	}

	for i := start; i <= n; i++ {

		cur += strconv.Itoa(i)
		backtrack60(n, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}
