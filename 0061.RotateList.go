package main

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }
    cur := head
    count := 0
    //find last and connect back to the head
    for cur.Next != nil {
        cur = cur.Next
        count++
    } 
    count++
    cur.Next = head
    //find the k to break
    k = k % count
    n := count - k
    //break
    for n > 1 {
        head = head.Next
        n--
    }
    newHead := head.Next
    head.Next = nil
    //return
    return newHead
}
