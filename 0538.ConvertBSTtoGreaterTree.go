package main

func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    inorder538(root, &sum)
    return root
}

func inorder538(root *TreeNode, sum *int) {
    if root == nil {
        return 
    }
    inorder538(root.Right, sum)
    *sum = *sum+root.Val
    root.Val = *sum
    inorder538(root.Left, sum)
}
 