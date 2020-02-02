package main

import (
	"fmt"
)

func bsearch(nums []int, target int) int {
    if nums == nil || len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums) - 1
    for left <= right {
        mid := left + (right - left) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    fmt.Printf("left %d\n", left)
    fmt.Printf("right %d\n", right)

    // fmt.Printf("nums left %d\n", nums[left])
    // fmt.Printf("nums right %d\n", nums[right])
    return -1
}