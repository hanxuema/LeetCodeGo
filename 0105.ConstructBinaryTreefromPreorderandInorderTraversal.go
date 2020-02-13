package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}
	inMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inMap[inorder[i]] = i
	}
	return buildTreeHelper(preorder, inorder, inMap)
}

func buildTreeHelper(preorder []int, inorder []int, inMap map[int]int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	index := inMap[rootVal]
	root := &TreeNode{}
	root.Val = rootVal
	root.Left = buildTreeHelper(preorder[1:index+1], inorder[0:index], inMap)
	root.Right = buildTreeHelper(preorder[index+1:], inorder[index+1:], inMap)
	return root
}
