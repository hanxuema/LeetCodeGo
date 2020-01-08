package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	count := 0
	cur := head
	for cur.Next != nil {
		count++
		cur = cur.Next
	}
	count++
	cur.Next = head

	k = k % count
	newcur := head
	for i := 1; i < count-k; i++ {
		newcur = newcur.Next
	}

	newHead := newcur.Next
	newcur.Next = nil

	return newHead
}
