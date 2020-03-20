package main

func deleteNodeBST(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		left := root.Left
		for left.Right != nil {
			left = left.Right
		}
		left.Right = root.Right
		root = root.Left
	} else if root.Val < key {
		root.Right = deleteNodeBST(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNodeBST(root.Left, key)
	}

	return root
}
