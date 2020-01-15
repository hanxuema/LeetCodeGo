package main

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil && fast.Next.Next == nil {
		slow = slow.Next
	}
	return slow
}

//[1,2,3,4,5]  return 3
// s s s
// f   f   f

//[1,2,3,4,5,6] return 4
// s s s
// f   f   f
