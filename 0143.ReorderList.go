package main

 
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return 
    } 
    pre := &ListNode{-1,head}
    cur := head
    for cur != nil && cur.Next != nil {
        cur = cur.Next 
        pre = pre.Next
    }
    
    pre.Next = nil
    cur.Next = head
    
    next := cur.Next
    nextNext := next.Next
    next.Next = cur
    cur.Next = nextNext
    
}