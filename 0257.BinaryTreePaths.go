package main

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
    res := []string{}
    if root == nil {
        return res
    } 
    binaryTreePathsHelper(root, "", &res)
    return res
}

func binaryTreePathsHelper(root *TreeNode, path string, res*[]string){
    path +=  strconv.Itoa(root.Val)
    if root.Left == nil && root.Right == nil{
        *res = append(*res, path)
        return
    }
    path += "->"
    if root.Left != nil{
        binaryTreePathsHelper(root.Left, path, res)
    }
    if root.Right != nil {
        binaryTreePathsHelper(root.Right, path, res)
    }
}


// [1,2,3,null,5]
// [null]
// [1,null]
// [1,2,3,null]