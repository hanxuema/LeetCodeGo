package main

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
