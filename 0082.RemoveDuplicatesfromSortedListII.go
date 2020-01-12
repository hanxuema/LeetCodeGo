package main

func deleteDuplicates2(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummy := &ListNode{-1, head}
    cur := head
    pre := dummy
    for cur != nil && cur.Next != nil {
        if cur.Val != cur.Next.Val {
            cur = cur.Next
            pre = pre.Next
            continue
        }
        for cur != nil && cur.Next != nil && cur.Val == cur.Next.Val {
            cur = cur.Next
        }
        cur = cur.Next
        pre.Next = cur
    }
    
    return dummy.Next
}