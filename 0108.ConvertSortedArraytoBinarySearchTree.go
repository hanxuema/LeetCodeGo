package main

func sortedArrayToBST(nums []int) *TreeNode {
    if nums == nil || len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
 
        root := &TreeNode{}
        root.Val = nums[0]
        return root
    }
    
    mid := len(nums) / 2 
    root := &TreeNode{}
    root.Val = nums[mid]
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    return root
}

 