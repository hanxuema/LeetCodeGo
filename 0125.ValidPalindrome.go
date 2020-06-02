package main

import "fmt"

func isPalindrome125(s string) bool {
	if len(s) == 0 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isLetter(s[left]) {
			left++
		}
		for left < right && !isLetter(s[right]) {
			right--
		}
		fmt.Printf("left %+v %+v\n", string(s[left]), s[left])
		fmt.Printf("right %+v %+v\n", string(s[right]), s[right])
		if toUpper(s[left]) != toUpper(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetter(char byte) bool {
	//A 65 //Z 90
	//a 97 //z 122
	return isUpper(char) || isLower(char)
}

func isUpper(char byte) bool {
	return 65 <= char && char <= 90
}

func isLower(char byte) bool {
	return 97 <= char && char <= 122
}

func toUpper(char byte) byte {
	fmt.Printf("char %+v\n", char)
	if isLower(char) {
		fmt.Printf("char 222 %+v\n", char)
		res := byte(char - 32)
		return res
	}
	return char
}
