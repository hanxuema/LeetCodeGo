package main

func isPalindrome0234(head *ListNode) bool {
	if head == nil {
		return true
	}
	//find middle
	var pre *ListNode
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next

		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	var left *ListNode
	var right *ListNode
	if fast.Next == nil {
		left = pre
		right = slow.Next
	} else { //if fast.Next.Next == nil
		right = slow.Next
		slow.Next = pre
		left = slow
	}
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
