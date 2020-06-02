package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	cur := []int{}
	visited := make([]bool, len(nums))
	backtrack47(nums, cur, visited, &res)
	return res
}

func backtrack47(nums []int, cur []int, visited []bool, res *[][]int) {
	if len(cur) == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)

		*res = append(*res, cur)
		fmt.Printf("&res %+v\n", *res)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if (i > 0 && nums[i] == nums[i-1]) && visited[i-1] == false {
			continue
		}
		//add
		cur = append(cur, nums[i])
		visited[i] = true
		//bt
		backtrack47(nums, cur, visited, res)
		//remove
		cur = cur[:len(cur)-1]
		visited[i] = false
	}
}
