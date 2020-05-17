package main

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	if candidates == nil || len(candidates) == 0 {
		return res
	}
	cur := []int{}
	backtrack39(candidates, target, 0, cur, &res)
	return res
}

func backtrack39(candidates []int, target int, start int, cur []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		cur = append(cur, candidates[i])
		backtrack39(candidates, target-candidates[i], i, cur, res)
		cur = cur[:len(cur)-1]
	}
}
