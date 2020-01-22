package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		}
		for i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j] {
			i++
		}

		for i < len(nums1) && j < len(nums2) && nums1[i] > nums2[j] {
			j++
		}

	}

	return res
}

//normal
// func intersect(nums1 []int, nums2 []int) []int {
//     res := []int{}
//     if nums1 == nil || nums2 == nil || len(nums1) == 0 || len(nums2) == 0{
//         return res
//     }
//     m := make(map[int]int)
//     for i :=0; i <len(nums1); i++{
//         m[nums1[i]]++
//     }

//     for i :=0; i <len(nums2); i++{
//         if  val, ok := m[nums2[i]]; ok {
//             if val > 0 {
//                 res = append(res, nums2[i])
//             }
//             m[nums2[i]]--
//         }
//     }

//     return res
// }
// [1,2,2,1]
// [2,2]
// [4,9,5]
// [9,4,9,8,4]
// [1,2]
// [1,1]
// [7,2,2,4,7,0,3,4,5]
// [3,9,8,6,1,9]
