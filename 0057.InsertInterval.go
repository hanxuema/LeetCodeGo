package main

func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		return res
	}

	length := len(intervals)
	newStart := newInterval[0]
	newEnd := newInterval[1]
	for i := 0; i < length; i++ {
		if intervals[i][1] < newStart || intervals[i][0] > newEnd {
			res = append(res, intervals[i])
		} else {
			intervals[i][1] = max(intervals[i][1], newEnd)
			for k := i + 1; k < length; k++ {
				if intervals[k][0] > intervals[i][1] {
					res = append(res, intervals[i])
					i = k + 1
					break
				}
				if k == length-1 {
					intervals[i][1] = max(intervals[i][1], intervals[k][1])
					res = append(res, intervals[i])
					break
				}

			}
		}
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
