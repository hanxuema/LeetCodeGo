package main

import "fmt"

func partition131(s string) [][]string {
	res := [][]string{}
	if len(s) == 0 {
		return res
	}

	cur := []string{}
	backtrack131(s, 0, cur, &res)
	return res
}

func backtrack131(s string, start int, cur []string, res *[][]string) {
	if start == len(s) {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(s); i++ {
		if isPalindrome131(s[start : i+1]) {
			cur = append(cur, string(s[start:i+1]))
			fmt.Printf("add cur %+v\n", cur)
			backtrack131(s, i+1, cur, res)

			cur = cur[:len(cur)-1]
		}
	}
}

func isPalindrome131(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
