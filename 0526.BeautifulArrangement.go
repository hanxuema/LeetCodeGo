package main

func countArrangement(N int) int {
    res := 0 
    if N == 0 {
        return res
    }
    
    n := []int{}
    for i := 0; i<= N; i++{
        n = append(n, i)
    }
    cur := []int{}
    visited := make([]bool, len(n))
    backtrack526(n , 1, cur, visited, &res)
    return res
}

func backtrack526(n []int, index int, cur []int, visited []bool, res *int){
    if len(cur) == len(n) -1{
        (*res)++
    }
    
    for i:= index; i < len(n); i++{
        //add
        // if visited[i] {
        //     continue
        // }
        if n[i] % i == 0 && i %  n[i] == 0 { 
            cur = append(cur,n[i])
            //backtrack
            visited[i] = true
            backtrack526(n, index+1,cur, visited, res)
            visited[i] = false
            //remove
            cur = cur[:len(cur)-1]
        }
    }
}