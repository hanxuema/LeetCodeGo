package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	arr := []int{4, 2, 1, 3}
	head := convertArrayToHead(arr)
	// rotateRight(head, 2)

	// partition(head, 3)
	// reorderList(head)
	// reverseBetween2(head, 1, 5)
	// isPalindrome(head)
	// oddEvenList(head)
	insertionSortList(head)
	// reorderList(head)
	// intervals := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	// newInterval := []int{4,8}
	// insert(intervals, newInterval)
	//nums := []int{1, 2, 3}
	// target := 3
	// //search(nums, target)
	// //searchRange(nums,target)
	// bsearch(nums, target)
	//majorityElement(nums)
	// subsets(nums)
	//countNumbersWithUniqueDigits(3)
	// treeArray := []int{1,2,2,3,-1,-1,3,4,-1,-1,4}
	// tree := convertArrayToTree(treeArray)
	// isBalanced(tree)
	partition131("aab")
	fmt.Printf("a %+v\n", 'a')
	fmt.Printf("z %+v\n", 'z')
	fmt.Printf("A %+v\n", 'A')
	fmt.Printf("Z %+v\n", 'Z')
	fmt.Printf("1 %+v\n", '1')
	fmt.Printf("9 %+v\n", '9')
}

func convertArrayToTree(arr []int) *TreeNode {
	// [1,2,2,3,null,null,3,4,null,null,4]
	root := TreeNode{}
	root.Val = arr[0]
	root.Left = convertToTreeHelper(arr, 1)
	root.Right = convertToTreeHelper(arr, 2)

	return &root
}

func convertToTreeHelper(arr []int, index int) *TreeNode {
	if index > len(arr)-1 || arr[index] == -1 {
		return nil
	}
	root := TreeNode{}
	root.Val = arr[index]
	root.Left = convertToTreeHelper(arr, index*2+1)
	root.Right = convertToTreeHelper(arr, index*2+2)

	return &root
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
