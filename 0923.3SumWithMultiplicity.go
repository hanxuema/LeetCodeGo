package main
import "sort"
func threeSumMulti(A []int, target int) int {
    if A == nil || len(A) < 3 {
        return 0
    }
    res := 0
    sort.Ints(A)
    for  i := 0; i <len(A); i++ {
        left, right := i+1, len(A) -1
        for left < right {
            sum := A[i] + A[left] + A[right]
            if sum == target {
                res++ 
                for left < right && A[right] == A[right - 1]{
                    res++
                    right--
                }
                for left < right && A[left] == A[left + 1]{
                    res++
                    left++
                }
                left++
                right--
            } else if sum > target {
                right--
            } else {
                left++
            }
        }
    }
    return res
}