package main

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    if(head.Next == nil) {
        root := &TreeNode{}
        root.Val = head.Val
        return root;
    }

    cur := head
    slow, fast := cur, cur
    temp := cur
    for fast != nil && fast.Next != nil {
        temp = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    next := slow.Next
    slow.Next = nil
    temp.Next = nil
    root := &TreeNode{}
    root.Val = slow.Val

    root.Left = sortedListToBST(head)
    root.Right = sortedListToBST(next)
    
    return root
}