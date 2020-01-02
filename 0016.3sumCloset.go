package main

	
import (  
    "sort"
)
func threeSumClosest(nums []int, target int) int {
    //corner case
    if nums == nil || len(nums) < 3 {
        return -1
    }  
    res := nums[0] + nums[1] + nums[2]
    sort.Ints(nums) // because we need to use 2 pointers
    
    //[-1, 2, 1, -4]
    // -4, -1 ,1, 2,   -4 + -1 + 2 = -3 . target = 1
    for i := 0; i < len(nums); i++ {
        left, right := i + 1, len(nums) - 1
        for left < right{
            sum := nums[i] + nums[left] + nums[right]
            if sum == target{
                return sum
            }
            if abs(sum - target) < abs(res - target){
                res = sum
            }
            if sum < target{
                left +=1
            }
            if sum > target{
                right -= 1
            }
        }
    }
    
    return res
}

func abs(num int) int{
    if num < 0{
        return -num
    }
    return num
}