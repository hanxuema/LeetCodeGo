package main

func findPairs(nums []int, k int) int {
	if nums == nil || len(nums) < 2 || k < 0 {
		return 0
	}
	res := 0
	set := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		set[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		temp := nums[i] + k
		index, ok := set[temp]
		if ok && index != i {
			res++
			delete(set, temp)
		}
	}
	return res
}

// func findPairs(nums []int, k int) int {
// 	res := 0
// 	if nums == nil || len(nums) <= 1 {
// 		return res
// 	}

// 	set := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		set[nums[i]] = set[nums[i]] + 1
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		temp := nums[i] + k
// 		val, ok := set[temp]
// 		if ok && val > 1 {
// 			res++
// 			set[temp] = set[temp] - 1
// 		}
// 	}
// 	return res
// }

// func findPairs222(nums []int, k int) int {
//     res := make(map[int]bool)
//     if nums == nil || len(nums) <= 1{
//         return 0
//     }

//     set := make(map[int]int)
//     for i := 0; i < len(nums); i++ {
//         set[nums[i]] = set[nums[i]] + 1
//     }
//     for i := 0; i < len(nums); i++ {
//         temp := nums[i] + k
//         val, ok := set[temp]
//         if ok && val > 0 {
//             _, ok2 := res[temp]
//                 if !ok2 {
//                     res[temp] = true
//                     set[temp] = set[temp] - 1
//                 }
//         }
//     }
//     return len(res)
// }
