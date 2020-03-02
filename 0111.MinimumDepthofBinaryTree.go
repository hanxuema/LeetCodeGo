package main

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return minDepthHelper(root,1)
}

func minDepthHelper(root *TreeNode , height int) int {
    if root == nil {
        return height -1
    }
    if root.Left != nil && root.Right != nil{
        return min(minDepthHelper(root.Left, height + 1), minDepthHelper(root.Right , height + 1))
    }   else {
        return max(minDepthHelper(root.Left, height + 1), minDepthHelper(root.Right , height + 1))
    }
}

// func min(a, b int) int{
//     if a < b{
//         return a
//     }
//     return b
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }