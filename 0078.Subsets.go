package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack078(nums, nil, 0, &res)
	return res
}

func backtrack078(nums []int, cur []int, start int, res *[][]int) {
	dc := make([]int, len(cur))
	copy(dc, cur)
	*res = append(*res, dc)
	//*res = append(*res, append([]int{}, cur...))
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtrack078(nums, cur, i+1, res)
		cur = cur[:len(cur)-1]
	}
}
