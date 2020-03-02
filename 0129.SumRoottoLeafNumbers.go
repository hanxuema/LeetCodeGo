package main

func sumNumbers(root *TreeNode) int {
    res := 0
    if root == nil {
        return res
    }
    return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, num int) int {
    if root == nil {
        return 0
    }
    num = root.Val  + num * 10
    if root.Left == nil && root.Right == nil {
        return  num
    }  
    sum := 0
    sum += sumNumbersHelper(root.Left, num)
    sum += sumNumbersHelper(root.Right,num)
    return sum
}