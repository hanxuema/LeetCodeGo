package main

func nextGreatestLetter(letters []byte, target byte) byte {
  
    left, right := 0, len(letters)
    //find the most right
    for left < right {
        mid := left + (right - left) / 2
        if letters[mid] <= target {
            left = mid + 1
        } else if letters[mid] > target {
            right = mid
        }
    }
    if right >= len(letters) {
        return letters[0]
    }
    return letters[right]
}