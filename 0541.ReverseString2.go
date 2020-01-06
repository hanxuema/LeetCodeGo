package main

func reverseStr(s string, k int) string { 
    sa := []byte(s)
    totalLength := len(sa)
    flag := true
    for i := 0; i < len(sa); i+=k {
        if flag {
            left := i
            right := i + k - 1
            if right > totalLength - 1{
                right = totalLength - 1
            } 
            reverseString2(sa, left, right)
        }
        flag = !flag
    } 
    return string(sa)
}

func reverseString2(s []byte, left int, right int) {
    for left < right {
        temp := s[left]
        s[left] = s[right]
        s[right] = temp
        left++
        right--
    }
}