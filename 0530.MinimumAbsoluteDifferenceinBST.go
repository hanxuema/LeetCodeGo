package main

import (
	"math"
)

func getMinimumDifference(root *TreeNode) int {
    if root == nil {
        return 0
    }
    arr := []int{}
    inorder530(root, &arr)
  
    res :=  math.MaxInt32
 
    for i:= 1; i < len(arr); i++{
        if arr[i] - arr[i-1] < res{
            res = arr[i] - arr[i-1]
        }
    }
    return res
}

func inorder530(root *TreeNode, arr *[]int){
    if root == nil{
        return
    }
    inorder530(root.Left, arr)
    *arr = append(*arr, root.Val)
    inorder530(root.Right, arr)
}