package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	dummy := &ListNode{-1, head}
	pre := dummy
	for cur != nil {
		next := cur.Next
		if cur.Val == val {
			pre.Next = next
			cur = next
		} else {
			cur = cur.Next
			pre = pre.Next
		}

	}
	return dummy.Next
}
