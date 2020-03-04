package main

func minDepth(root *TreeNode) int {
    return minDepthHelper(root, 0)
}

func minDepthHelper(root *TreeNode, depth int) int {
    if root == nil {
        return depth
    }
    if root.Left != nil && root.Right != nil{
        return min(minDepthHelper(root.Left, depth),minDepthHelper(root.Right, depth)) + 1
    } else if root.Left == nil {
        return minDepthHelper(root.Right, depth) + 1
    } else if root.Right == nil{
        return minDepthHelper(root.Left, depth) + 1
    } else {
        return depth
    }
}

// func min(a, b int) int{
//     if a < b{
//         return a
//     }
//     return b
// }

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