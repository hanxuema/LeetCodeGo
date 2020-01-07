package main

func numRescueBoats(people []int, limit int) int {
	res := 0
	if people == nil || len(people) == 0 {
		return res
	}
	left, right := 0, len(people) - 1
	for left <= right {
		sum := people[left] + people[right]
		if left == right {
			sum = people[left]
		}
		if sum <= limit {
			left+=
		}
		right--
		res++
	}

	return res
}