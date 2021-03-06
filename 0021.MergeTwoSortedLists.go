package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    pre := &ListNode{0, nil}
    cur := pre
    for l1 != nil && l2 != nil {
        temp := l1.Val 
        if l1.Val  > l2.Val {
            temp = l2.Val 
        }
        
        cur.Next = &ListNode{temp, nil}
        cur = cur.Next 
        if l1.Val  > l2.Val {
            l2 = l2.Next
        } else {
            l1 = l1.Next    
        }
        
    }
    if l1 != nil {
        cur.Next = l1
    }
    if l2 != nil {
        cur.Next = l2
    }
    return pre.Next
}