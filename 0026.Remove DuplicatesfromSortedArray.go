package main


func removeDuplicates(nums []int) int {
    //corner case
    if nums == nil || len(nums) == 0{
        return 0
    }
    res := 0
    for i := 1; i < len(nums); i++ {
        if nums[res] != nums[i]{
            res += 1
            nums[res] = nums[i]
        }
    }
    return res + 1
}

//0,0,1,1,1,2,2,3,3,4