package main

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    
    res := []int{}
    res = append(res, postorderTraversal(root.Left)...)
    res = append(res, postorderTraversal(root.Right)...)
    res = append(res, root.Val)
    return res
}