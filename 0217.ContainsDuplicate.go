package main

func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) <= 1 {
		return false
	}
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 1 {
			return true
		}
	}
	return false
}
