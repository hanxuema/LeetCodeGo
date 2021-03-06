package main

 func preorderTraversal(root *TreeNode) []int {
    if root == nil{
        return nil
    }
    res := []int{}
    res = append(res, root.Val)
    res = append(res, preorderTraversal(root.Left)...)
    res = append(res, preorderTraversal(root.Right)...)
    
    return res
}