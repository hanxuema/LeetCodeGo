package main

func merge(intervals [][]int) [][]int {
    res := make([][]int, 0)
    if intervals == nil || len(intervals) == 0 {
        return res
    }
    
    sort.Slice(intervals, func(i, j int) bool {return intervals[i][0] < intervals[j][0]})
    res = append(res, intervals[0])
    for i := 1 ; i < len(intervals); i++ {
        curStart := intervals[i][0]
        curEnd := intervals[i][1]
         
        preEnd := intervals[i-1][1]
        
        if preEnd >= curStart {
            intervals[i-1][1] = max(preEnd, curEnd)
            intervals[i] = intervals[i-1]
        } else {
            res = append(res, intervals[i])
        }
    }
    return res
}
 
func max (a, b int) int {
                if a > b {
                    return a
                } else {
                    return b
                }
}

// Your input
// [[1,3],[2,6],[8,10],[15,18]]
// [[1,2]]
// [[1,4],[0,2],[3,5]]
// Output
// [[1,6],[8,10],[15,18]]
// [[1,2]]
// [[0,5]]
// Expected
// [[1,6],[8,10],[15,18]]
// [[1,2]]
// [[0,5]]