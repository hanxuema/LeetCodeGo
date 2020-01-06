package main

func longestMountain(A []int) int {
	if A == nil || len(A) < 3 {
		return 0
	}
	left, right, res := 0, 0, 0 
	for left < len(A) { 
		if right < len(A)-1 && A[right+1] > A[right] {
			for right < len(A)-1 && A[right+1] > A[right] {
				right++
			}
			if right < len(A)-1 && A[right+1] < A[right] {
				for right < len(A)-1 && A[right+1] < A[right] {
					right++
				}
				if right != left && (right-left+1) > res {
					res = right - left + 1
				}
			}
		}

		if right > left {
			left = right
		} else {
			left++
			right = left
		}

	}
	return res
}
