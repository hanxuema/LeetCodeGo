package main

import "sort"

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return (points[i][0]*points[i][0] + points[i][1]*points[i][1]) < (points[j][0]*points[j][0] + points[j][1]*points[j][1])
	})
	//fmt.Printf("point %+v\n", points)
	return points[:K]
}
