package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	cur := []int{}
	visited := make([]bool, len(nums))
	backtrack90(nums, 0, cur, visited, &res)

	return res
}

func backtrack90(nums []int, start int, cur []int, visited []bool, res *[][]int) {
	temp := make([]int, len(cur))
	copy(temp, cur)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		//add option
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == false {
			continue
		}
		cur = append(cur, nums[i])
		//backtrack
		visited[i] = true
		backtrack90(nums, i+1, cur, visited, res)
		visited[i] = false
		//remove option
		cur = cur[:len(cur)-1]
	}
}
