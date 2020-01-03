package main

func removeDuplicates2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	res := 0
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] || count == 2 {
			res += 1
			nums[res] = nums[i]
			count = 0
		} else {
			count += 1
			res += 1
		}
	}
	return res
}
