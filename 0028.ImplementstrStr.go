package main


func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    if len(haystack) == 0 {
        return -1
    }
    if len(needle) > len(haystack) {
        return -1
    }
    
    firstChar := needle[0]
    for i := 0; i < len(haystack); i++ {
        if haystack[i] == firstChar {
            if checksubstring(haystack[i:], needle){
                return i
            }
        }
    }
    return -1
}

func checksubstring(haystack string, needle string) bool{
    if len(haystack) < len(needle){
        return false 
    }
    for i := 0; i < len(needle) ; i++{
        if haystack[i] != needle[i] {
            return false
        }
    }
    return true
}