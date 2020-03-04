package main

func countNodes(root *TreeNode) int {
    res := 0
    if root == nil {
        return res
    }
    helper(root, &res)
   
    return res
}

func helper(root *TreeNode, res *int){
    if root == nil {
        return
    }
    *res++
    helper(root.Left, res)
    helper(root.Right, res)
}