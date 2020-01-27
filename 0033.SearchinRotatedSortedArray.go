package main


func search(nums []int, target int) int {
    if nums == nil || len(nums) == 0{
        return -1
    }
    
    left, right  := 0, len(nums) - 1
    
    for left <= right {
		mid := left + (right - left) / 2
        if nums[mid] > nums[right]{
            left = mid - 1
        } else {
            right = mid - 1
        }
    }
    right = len(nums) - 1
    if nums[left] == nums[0] {
        left = 0
    } 
    
    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] == target{
          return mid  
        }
        if nums[mid] > target{ 
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}