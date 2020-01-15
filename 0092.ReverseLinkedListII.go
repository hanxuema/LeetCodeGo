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

// my implementation
// func reverseBetween(head *ListNode, m int, n int) *ListNode {
//     if head == nil || m == n {
//         return head
//     }
    
//     length := n - m
//     var tail *ListNode
//     if m > 1 {
//         tail = &ListNode{-1,head}
//     }
//     cur := head

//     for m > 1 {
//         tail = tail.Next
//         cur = cur.Next
//         m--
//     }
//     newCur := cur
//     var pre *ListNode
//     next := cur.Next
//     for length >= 0 {
//         next = cur.Next
//         cur.Next = pre
//         pre = cur
//         cur = next
//         length--
//     }
//     // next := cur.Next
//     if tail != nil {
//         tail.Next = pre
//         newCur.Next = next
//     } else {
//         head = pre
//         newCur.Next = next
//     }
//     return head
// }