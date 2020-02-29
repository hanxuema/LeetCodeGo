package main

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return abs(heightHelper(root.Left, 0) - heightHelper(root.Right, 0)) < 2 && isBalanced(root.Left) && isBalanced(root.Right) 
}

func heightHelper(root *TreeNode, height int) int{
    if root == nil {
        return height
    }
    return max(heightHelper(root.Left, height + 1), heightHelper(root.Right, height + 1))
}


// func abs(a int) int {
// 	if a < 0 {
// 		return 0 - a
// 	}
// 	return a
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
