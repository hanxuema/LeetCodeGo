package main

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}

	start, end := newInterval[0], newInterval[1]
	left, right := make([][]int, 0), make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		curStart, curEnd := intervals[i][0], intervals[i][1]
		if curEnd < start {
			left = append(left, intervals[i])
		} else if curStart > end {
			right = append(right, intervals[i])
		} else {
			start = min(start, curStart)
			end = max(end, curEnd)

		}
	}

	res = append(res, left...)

	temp := []int{start, end}
	res = append(res, temp)
	res = append(res, right...)
	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}  
// 		return b
// }

func min(a, b int) int {
	if a < b {
		return a
	}  
		return b
	 
}
