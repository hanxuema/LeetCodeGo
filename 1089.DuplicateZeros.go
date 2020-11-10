package main

func duplicateZeros(arr []int) {
	if arr == nil {
		return
	}
	for i := 0; i < len(arr); i++ {
		next := 0
		if arr[i] == 0 && i < len(arr)-1 {
			next = arr[i+1]
			arr[i+1] = 0
		} else if arr[i] != 0 {
			if next != 0 {
				arr[i] = next
			}
		}
	}
}

//[1,0,2,3,0,4,5,0]
//[1,0,0,2,3,0,0,4]
