package main

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    temp := root.Left
    root.Left = root.Right
    root.Right = temp
    
    invertTree(root.Left)
    invertTree(root.Right)
    
    return root
}


// [4,2,7,1,3,6,9]
// [1,2,null]
// [1,2,null,3,null,5]