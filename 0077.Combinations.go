package main

func combine(n int, k int) [][]int {
    res := [][]int{}
    if k == 0 || n == 0 {
        return res
    }
    cur := []int{}
    backtrack77(n, k, 1, cur, &res)
    return res
}

func backtrack77(n int, k int, start int, cur []int, res *[][]int){
    if len(cur) == k {
        temp := make([]int, k)
        copy(temp, cur)
        *res = append(*res, temp)
        return
    } 
    
    for i := start ; i <= n; i++{
        //add option
        cur = append(cur, i)
        //backtrack
        backtrack77(n, k, i+1, cur, res)
        //remove option
        cur = cur[:len(cur)-1]
    }
}

//1,2,3,4