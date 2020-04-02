package main

func isHappy(n int) bool {
    dic := make(map[int]int)
    var ok bool
    _, ok = dic[n]
    for n != 1 && !ok { 
        dic[n] =1
        n = nextNumber(n)
        _, ok = dic[n]
    }
    return n == 1
}

func nextNumber(n int) int{
    sum := 0
    for n >  0 {
        val := n % 10
        sum += val * val
        n = n / 10
    }
    return sum
}