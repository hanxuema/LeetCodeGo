package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left - 1
}
