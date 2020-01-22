package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{-1, head}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val >= cur.Val {
			cur = cur.Next
		} else {
			//find the curNext, curNext is the node to be inserted somewhere between head and cur  and cur next next
			curNext := cur.Next
			curNextNext := curNext.Next
			//start point from begining
			pre := dummy
			//when point next is greater than curNext
			for pre.Next.Val <= curNext.Val {
				pre = pre.Next
			}
			//pre.next = curNext, curNext = pre.next
			curNext.Next = pre.Next
			pre.Next = curNext
			//cur.next = cur next next
			cur.Next = curNextNext
		}
	}

	return dummy.Next
}

//p      t  c
//d->-1->5->3->4->0
//
//-1->0->3->4->5
