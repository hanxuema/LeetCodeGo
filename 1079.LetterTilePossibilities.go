package main

import (
	"sort"
)

//subset with no dupliation
func numTilePossibilities(tiles string) int {
	if len(tiles) <= 1 {
		return len(tiles)
	}
	res := 0
	cur := []byte{}
	chars := []byte(tiles)
	sort.Slice(chars, func(i int, j int) bool { return chars[i] < chars[j] })
	visited := make([]bool, len(tiles))
	backtrack1079(chars, cur, visited, &res)

	return res
}

func backtrack1079(chars []byte, cur []byte, visited []bool, res *int) {
	if len(cur) > 0 {
		(*res)++
	}

	for i := 0; i < len(chars); i++ {
		//add
		if visited[i] {
			continue
		}
		if i > 0 && chars[i] == chars[i-1] && visited[i-1] == false {
			continue
		}
		cur = append(cur, chars[i])
		//fmt.Printf("cur %+v\n", cur)
		//backtrack
		visited[i] = true
		backtrack1079(chars, cur, visited, res)
		visited[i] = false
		//remove
		cur = cur[:len(cur)-1]
	}
}
