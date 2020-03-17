package main

func kthSmallest(root *TreeNode, k int) int {
    arr := make([]int,0)
    inorderkthsmallest(root, &arr)
    return arr[k-1]
}

func inorderkthsmallest(root *TreeNode, arr *[]int){
    if root == nil {
        return
    }
    inorder(root.Left, arr)
    *arr = append(*arr, root.Val)
    inorder(root.Right,arr)
}