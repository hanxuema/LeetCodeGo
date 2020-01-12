package main


 func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil {
        return nil
    }
    var pre *ListNode
    cur := head
    
    for m > 1{
        pre = cur
        cur = cur.Next
        m--
        n--
    } 
    p2 := pre
    c2 := cur
    for n > 0 {
            next := cur.Next
            cur.Next = pre
            pre = cur
            cur = next
        n--
    }
    if p2 == nil {
        head = pre
    }else{
        p2.Next = pre
    }
    
    c2.Next = cur
    return head
}
