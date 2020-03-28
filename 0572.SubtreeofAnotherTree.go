package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
    if s == nil {return false}
    return isSubtreeHelper(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSubtreeHelper(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil || t == nil {
        return false
    }
    if s.Val != t.Val {
        return false
    }
    return isSubtreeHelper(s.Left, t.Left) && isSubtreeHelper(s.Right, t.Right)
}