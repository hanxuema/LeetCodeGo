package main

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    if x == 0 {
        return true
    }
    origin := x
    var num int
    for x > 0 {
        temp := x % 10
        x = x / 10
        num = num * 10 + temp
        if num < -2147483648 || num > 2147483647{
            return false
        }
    } 
    return num == origin
}