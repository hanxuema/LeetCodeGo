package main

func constructFromPrePost(pre []int, post []int) *TreeNode {
	return buildTree(pre, post)
}

func buildTree(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0 {
		return nil
	}
	if len(pre) == 1 || len(post) == 1 {
		root := &TreeNode{}
		root.Val = pre[0]
		return root
	}

	left := 0
	for i, v := range post {
		if v == pre[1] {
			left = i + 1
		}
	}

	root := &TreeNode{}
	root.Val = pre[0]
	root.Left = buildTree(pre[1:left+1], post[:left])
	root.Right = buildTree(pre[left+1:], post[left:len(post)-1])

	return root
}

