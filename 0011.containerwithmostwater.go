
package main

func maxArea(height []int) int {
    //corner case
    if height == nil || len(height) < 2{
        return 0
    }
    
    maxArea := 0
    left, right := 0, len(height) -1
    for left < right {
        maxArea = maxInt(maxArea, minInt(height[left], height[right]) * (right - left))
        if height[left] >= height[right]{
            right -= 1
        }else{
            left += 1
        }
    }
    return maxArea
}

func maxInt(a,b int) int{
    if a > b{
        return a
    }
    return b
}

func minInt(a, b int) int{
    if a < b {
        return a
    }
    return b
}