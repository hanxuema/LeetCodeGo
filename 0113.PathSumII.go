package main

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	pathSumHelper(root, sum, path, &res)
	return res
}

func pathSumHelper(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	pathSumHelper(root.Left, sum-root.Val, path, res)
	pathSumHelper(root.Right, sum-root.Val, path, res)
	path = path[:len(path)-1]
}
