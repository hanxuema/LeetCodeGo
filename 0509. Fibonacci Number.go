package main

func fib(N int) int {
    if N == 0 {
        return 0
    }
    if N == 1{
        return 1
    }
    
    temp := make([]int, N+1)
    temp[0] = 0
    temp[1] = 1
    for i := 2; i <=N ; i++{
        temp[i] = temp[i-1] + temp[i-2]
    }
    return temp[N]
}

//  0  1  2  3  4  5  6  7 index
// [0  1  1  2  3  5  8  13]

 //0 -> 1 -> 1-> 2->