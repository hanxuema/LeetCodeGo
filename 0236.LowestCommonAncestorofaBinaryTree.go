package main

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
    if root == nil || root.Val == p.Val || root.Val == q.Val{
        return root
    }
    left := lowestCommonAncestor236(root.Left, p, q)
    right := lowestCommonAncestor236(root.Right, p, q)
    if left == nil {
        return right
    }
    if right == nil {
        return left
    }
    return root
}