package main

func rangeSumBST(root *TreeNode, L int, R int) int {
    res := 0
    if root == nil {
        return 0
    }
    if root.Val > L && root.Val > R{
        return rangeSumBST(root.Left, L, R)
    } else if root.Val < L && root.Val < R {
        return rangeSumBST(root.Right, L, R)
    } else if root.Val >= L && root.Val <= R{
        res += root.Val
        res += rangeSumBST(root.Right, L, R)
        res += rangeSumBST(root.Left, L, R)
    
    }
        return res
}