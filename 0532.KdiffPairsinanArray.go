package main

func findPairs(nums []int, k int) int {
	res := 0
	if nums == nil || len(nums) <= 1 {
		return res
	}

	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = set[nums[i]] + 1
	}
	for i := 0; i < len(nums); i++ {
		temp := nums[i] + k
		val, ok := set[temp]
		if ok && val > 1 {
			res++
			set[temp] = set[temp] - 1
		}
	}
	return res
}
