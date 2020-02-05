package main

func rotate(nums []int, k int)  {
    length := len(nums)
    if nums == nil || len(nums) == 0 || k % length == 0 {
        return
    }
    k = k % length
    
    for i := 0 ; i < k; i++{
        pre := nums[len(nums) - 1]
        for j := 0; j < len(nums); j++{
            temp := nums[j]
            nums[j] = pre
            pre = temp
        }
    }
}