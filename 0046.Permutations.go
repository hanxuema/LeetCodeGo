package main

func permute(nums []int) [][]int {
	res := [][]int{}
	if nums == nil || len(nums) == 0 {
		return res
	}

	cur := []int{}
	visited := make([]bool, len(nums))
	backtrack46(nums, cur, visited, &res)
	return res
}

func backtrack46(nums []int, cur []int, visited []bool, res *[][]int) {
	if len(cur) == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		//add
		cur = append(cur, nums[i])
		visited[i] = true
		//backtrack
		backtrack46(nums, cur, visited, res)
		//remove the last item,
		cur = cur[:len(cur)-1]
		visited[i] = false
	}

}
