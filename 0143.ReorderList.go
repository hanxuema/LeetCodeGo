package main

 
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return 
    }  
    //find middle
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    var pre *ListNode
    cur := slow.Next 
    for cur != nil{
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    slow.Next = nil
    for pre  != nil {
        next := head.Next
        curNext := pre.Next
        
        head.Next = pre
        pre.Next = next
        
        head = next
        pre = curNext
    }
     
}

//s  s  s       pre
//f     f       f
//1->2->3 p<-4<-5 
//1 2 3 
// 5 4

//s  s  s              c
//f .   f .         f
//1->2->3   pre<-4<-5<-6
//1 2 3
// 6 5