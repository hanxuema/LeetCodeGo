package main

func countNodes(root *TreeNode) int {
    res := 0
    if root == nil {
        return res
    }
    countNodesHelper(root, &res)
   
    return res
}

func countNodesHelper(root *TreeNode, res *int){
    if root == nil {
        return
    }
    *res++
    countNodesHelper(root.Left, res)
    countNodesHelper(root.Right, res)
}