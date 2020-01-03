package main

func reverseString(s []byte) {
	//corner case
	if s == nil || len(s) == 0 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}
