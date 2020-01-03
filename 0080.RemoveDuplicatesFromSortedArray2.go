package main

func removeDuplicates2(nums []int) int {
    if nums == nil || len(nums) < 0 {
        return 0
    }
    res := 2
    for i := 2; i < len(nums) ; i++ {
        if nums[i] != nums[res - 2] {
            nums[res] = nums[i]
            res++
        }
    }
    return res
}