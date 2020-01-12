package main

func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }
    var pre *ListNode
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        next := slow.Next
        slow.Next = pre
        pre = slow
        slow = next
        
        fast = fast.Next.Next
    }
    var first *ListNode
    var second *ListNode
    if fast != nil { 
        //   p  s     f   
        //1<-2  3->2->1
        first = pre
        second = slow.Next
    } else {
        first = pre
        second = slow
    }  
    
    for second != nil{
        if first.Val != second.Val {
            return false
        }
        first = first.Next
        second = second.Next
    }
     
    return true
    
} 
