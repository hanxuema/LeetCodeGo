package main

func twoSum2(numbers []int, target int) []int {
	if numbers == nil || len(numbers) < 2 {
		return nil
	}
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
