package main

func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {
        return 0
    }
    maxSum :=  nums[0]
    tempSum := nums[0]
    for _, num := range nums[1:]{
            tempSum = max(num, num + tempSum)
            maxSum = max(maxSum, tempSum)
    }
        return maxSum
 
}

func max(a, b int) int{
    if a > b{
        return a
    }
    return b
}