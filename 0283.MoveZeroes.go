package main

func moveZeros(nums []int){
	//corner case
	if nums == nil || len(nums) == 0{
		return
	}
	//find the index of first zero
	left := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left = i
			break
		}
	}
	//if no zero return
	if left < 0 {
		return
	}
	//move zero to the end
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[left] == 0{
			temp := nums[i]
			nums[i] = nums[left]
			nums[left] = temp
			left++
		}
	}
}