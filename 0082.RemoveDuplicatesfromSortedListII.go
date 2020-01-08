package main

func deleteDuplicates2(head *ListNode) *ListNode {
	pre := &ListNode{-1, nil}
	cur := pre
	pre.Next = head

	for head != nil && head.Next != nil {
		if head.Val != head.Next.Val {
			head = head.Next
			cur = cur.Next
		} else {
			for head != nil && head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			if head == nil {
				cur.Next = nil
			} else {
				cur.Next = head.Next
				head = head.Next
			}
		}
	}
	return pre.Next
}
