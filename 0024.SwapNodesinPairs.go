package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{-1, nil}
	pre.Next = head
	cur := pre
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		next := cur.Next
		nextNext := next.Next
		nextNextNext := nextNext.Next

		cur.Next = nextNext
		nextNext.Next = next
		next.Next = nextNextNext

		cur = cur.Next.Next
	}

	return pre.Next
}
