package main

func isSubPath(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    } 
    if root == nil {
        return false
    }
    return isSubPathHelper(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSubPathHelper(cur *ListNode, root *TreeNode) bool {
    if cur == nil  {
        return true
    }
    if root == nil && cur != nil {
        return false
    }
    if cur.Val != root.Val {
        return false
    }
    return isSubPathHelper(cur.Next, root.Left) || isSubPathHelper(cur.Next,  root.Right)
    
}