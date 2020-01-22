package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			if abs(i-val) <= k {
				return true
			} else {
				m[nums[i]] = i
			}
		} else {
			m[nums[i]] = i
		}
	}
	return false
}

// func abs(val int) int {
// 	if val < 0 {
// 		val = 0 - val
// 	}
// 	return val
// }

// [1,2,3,5,6,7,1,1]
// 3
// [1,2,3,1,2,3]
// 2
