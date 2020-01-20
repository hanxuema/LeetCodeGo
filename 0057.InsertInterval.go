package main

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		res = append(res, newInterval) //[] [5,7]
		return res
	}
	start, end := newInterval[0], newInterval[1]
	newIntervalUsed := false
	for i := 0; i < len(intervals); i++ {
		curStart := intervals[i][0]
		curEnd := intervals[i][1]
		if curEnd < start {
			res = append(res, intervals[i])
		} else if curStart > end {
			if !newIntervalUsed {
				temp := []int{start, end}
				res = append(res, temp)
				newIntervalUsed = true
			}
			res = append(res, intervals[i])
		} else {
			start = min(start, intervals[i][0])
			end = max(end, intervals[i][1])
			if i == len(intervals)-1 { //  [[0,2],[3,9]]     [6,8]
				temp := []int{start, end}
				res = append(res, temp)
			}
		}
	}
	if start > intervals[len(intervals)-1][1] || len(res) == 0 { //[[3,5],[17,19]]   [4,18]
		temp := []int{start, end}
		res = append(res, temp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//       [3,             10]
//[[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
