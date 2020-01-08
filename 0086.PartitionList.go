package main
 
func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    } 
    less := &ListNode{-1, nil} 
    lcur := less
    greater := &ListNode{-1, nil}
    gcur := greater
    for head != nil {
        if head.Val < x {
            less.Next = &ListNode{head.Val, nil} 
            less = less.Next
        } else { 
            gcur.Next = &ListNode{head.Val, nil} 
            gcur = gcur.Next
        }
        head = head.Next
    }
    less.Next = greater.Next
    
    return lcur.Next
}