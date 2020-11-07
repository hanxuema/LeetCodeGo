package main

func findMaxConsecutiveOnes(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	//left, right := 0, 0
	count := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > res {
				res = count
			}
		} else {
			count = 0
		}
	}
	return res
}
