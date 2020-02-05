package main

func plusOne(digits []int) []int {
    res := make([]int,0)
    if digits == nil || len(digits) == 0 {
        return res
    }
    
    carry := 1
    for i := len(digits) - 1; i >= 0; i-- {
        sum := digits[i] + carry
        carry = sum / 10
        digits[i] = sum % 10
    }
    if carry == 1 {
        res = []int{1}
        digits = append(res, digits...)
    }
    return digits
}

// Your input
// [1,2,3]
// [4,3,2,1]
// [9,9,9,9]
// Output
// [1,2,4]
// [4,3,2,2]
// [1,0,0,0,0]
// Expected
// [1,2,4]
// [4,3,2,2]
// [1,0,0,0,0]