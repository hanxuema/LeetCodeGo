package main

func main() {
	// 	[0,1,2]
	// 0
	// nums := []int{3, 1, 4, 1, 5}
	// k := 2
	// findPairs(nums, k)
	// S := "#aaa"
	// T := "aaa"
	// res := backspaceCompare(S,T)
	// print(res)
	// A := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// longestMountain(A)
	// A := []int{1,1,2,2,2,2}
	// target := 5
	// threeSumMulti(A, target)
	
	// deleteDuplicates2(head)
	arr := []int{1,2}
	head := convertArrayToHead(arr)
	// rotateRight(head, 2)

	// partition(head, 3)
	// reorderList(head)
	// reverseBetween2(head, 1, 5)
	isPalindrome(head)
}

func convertArrayToHead(arr []int) *ListNode {
	pre := &ListNode{0, nil}
	head := &ListNode{0, nil}
	pre.Next = head
	cur := head
	for i := 0; i < len(arr); i++ {
		cur.Val = arr[i]
		cur.Next = &ListNode{0, nil}
		if i < len(arr)-1 {
			cur = cur.Next
		} else {
			cur.Next = nil
		}
	}

	return pre.Next
}
