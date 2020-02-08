package main

func searchRange(nums []int, target int) []int {
    res := []int{-1,-1}
    if nums == nil || len(nums) == 0 {
        return res
    }
    //find the left
    left, right := 0, len(nums)
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] >= target {
            right = mid
        } else if nums[mid] < target {
            left = mid + 1
        }
    }
    if left > len(nums) - 1 || nums[left] > target {
        return res
    }
    
    res[0] = left
    //find the right
    
    left, right = 0, len(nums) 
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] <= target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid
        }
    }
    res[1] = right - 1
    return res
}