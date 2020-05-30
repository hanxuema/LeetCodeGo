package main

func reverse(x int) int {
    if x == 0 {
        return 0
    } 
    isPos := true
    if x < 0 {
        isPos = false
        x = 0 - x
    }
    var num int
    for x > 0 {
        temp :=  x % 10  
        x = x / 10
        num = num * 10 + temp
        if num < -2147483648 || num > 2147483647{
            return 0
        }
 
    }
    
    if !isPos{
        num = 0 - num
    }
    return num
}