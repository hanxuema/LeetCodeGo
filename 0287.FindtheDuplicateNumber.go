package main

import "sort"

func findDuplicate(nums []int) int {
	//corner case
	if nums == nil || len(nums) < 2 {
		return -1
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
