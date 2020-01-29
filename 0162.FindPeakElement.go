package main

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if mid > 0 && mid < len(nums) && nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
